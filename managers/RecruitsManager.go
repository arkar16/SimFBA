package managers

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/jinzhu/gorm"
)

func GetAllRecruits() []models.Croot {
	db := dbprovider.GetInstance().GetDB()

	var recruits []structs.Recruit

	db.Preload("RecruitPlayerProfiles", func(db *gorm.DB) *gorm.DB {
		return db.Order("total_points DESC").Where("total_points > 0")
	}).Find(&recruits)

	var croots []models.Croot
	for _, recruit := range recruits {
		var croot models.Croot
		croot.Map(recruit)

		croots = append(croots, croot)
	}

	return croots
}

func GetAllUnsignedRecruits() []structs.Recruit {
	db := dbprovider.GetInstance().GetDB()

	var croots []structs.Recruit

	db.Where("is_signed = ?", false).Find(&croots)

	return croots
}

func GetCollegeRecruitByRecruitID(recruitID string) structs.Recruit {
	db := dbprovider.GetInstance().GetDB()

	var recruit structs.Recruit

	err := db.Where("id = ?", recruitID).Find(&recruit)
	if err != nil {
		log.Fatalln(err)
	}

	return recruit
}

func GetRecruitsByTeamProfileID(ProfileID string) []structs.RecruitPlayerProfile {
	db := dbprovider.GetInstance().GetDB()

	var croots []structs.RecruitPlayerProfile

	err := db.Preload("Recruit").Where("profile_id = ?", ProfileID).Find(&croots).Error

	if err != nil {
		log.Fatal(err)
	}

	return croots
}

func GetRecruitProfileByPlayerId(recruitID string, profileID string) structs.RecruitPlayerProfile {
	db := dbprovider.GetInstance().GetDB()

	var croot structs.RecruitPlayerProfile
	err := db.Where("recruit_id = ? and profile_id = ?", recruitID, profileID).Find(&croot).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return structs.RecruitPlayerProfile{}
		} else {
			log.Fatal(err)
		}
	}
	return croot
}

func GetRecruitPlayerProfilesByRecruitId(recruitID string) []structs.RecruitPlayerProfile {
	db := dbprovider.GetInstance().GetDB()

	var croots []structs.RecruitPlayerProfile
	err := db.Where("recruit_id = ?", recruitID).Order("total_points desc").Find(&croots).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []structs.RecruitPlayerProfile{}
		} else {
			log.Fatal(err)
		}
	}
	return croots
}

func CreateRecruitingProfileForRecruit(recruitPointsDto structs.CreateRecruitProfileDto) structs.RecruitPlayerProfile {
	db := dbprovider.GetInstance().GetDB()

	recruitEntry := GetRecruitProfileByPlayerId(strconv.Itoa(recruitPointsDto.RecruitID),
		strconv.Itoa(recruitPointsDto.ProfileID))

	if recruitEntry.RecruitID != 0 && recruitEntry.ProfileID != 0 {
		// Replace Recruit
		recruitEntry.ToggleRemoveFromBoard()
		db.Save(&recruitEntry)
		return recruitEntry
	}

	createRecruitEntry := structs.RecruitPlayerProfile{
		SeasonID:            recruitPointsDto.SeasonID,
		RecruitID:           recruitPointsDto.RecruitID,
		ProfileID:           recruitPointsDto.ProfileID,
		TotalPoints:         0,
		CurrentWeeksPoints:  0,
		SpendingCount:       0,
		Scholarship:         false,
		ScholarshipRevoked:  false,
		AffinityOneEligible: recruitPointsDto.AffinityOneEligible,
		AffinityTwoEligible: recruitPointsDto.AffinityTwoEligible,
		TeamAbbreviation:    recruitPointsDto.Team,
		RemovedFromBoard:    false,
		IsSigned:            false,
	}

	// Create
	db.Create(&createRecruitEntry)

	return createRecruitEntry
}

func AllocateRecruitPointsForRecruit(updateRecruitPointsDto structs.UpdateRecruitPointsDto) {
	db := dbprovider.GetInstance().GetDB()

	// Recruit Team Profile
	recruitingProfile := GetOnlyRecruitingProfileByTeamID(strconv.Itoa(updateRecruitPointsDto.ProfileID))

	// Recruit Player Profile
	recruitEntry := GetRecruitProfileByPlayerId(strconv.Itoa(updateRecruitPointsDto.RecruitID),
		strconv.Itoa(updateRecruitPointsDto.ProfileID))

	var pointAllocation structs.RecruitPointAllocation

	if updateRecruitPointsDto.AllocationID == 0 {
		pointAllocation = structs.RecruitPointAllocation{
			RecruitID:          updateRecruitPointsDto.RecruitID,
			TeamProfileID:      recruitingProfile.TeamID,
			RecruitProfileID:   int(recruitEntry.ID),
			WeekID:             updateRecruitPointsDto.WeekID,
			Points:             updateRecruitPointsDto.SpentPoints,
			AffinityOneApplied: recruitEntry.AffinityOneEligible,
			AffinityTwoApplied: recruitEntry.AffinityTwoEligible,
		}
	} else {
		err := db.Where("id = ?", updateRecruitPointsDto.AllocationID).Find(&pointAllocation).Error
		if err != nil {
			fmt.Println("Cannot find existing point allocation record.")
		}
	}

	// Allow the user to update points spent on a recruit during a given week.
	difference := recruitingProfile.SpentPoints - pointAllocation.Points
	if pointAllocation.Points != updateRecruitPointsDto.SpentPoints {
		pointAllocation.UpdatePointsSpent(updateRecruitPointsDto.SpentPoints)
	}
	difference += pointAllocation.Points

	recruitingProfile.AllocateSpentPoints(difference)
	if recruitingProfile.SpentPoints > recruitingProfile.WeeklyPoints {
		fmt.Printf("Recruiting Profile " + strconv.Itoa(updateRecruitPointsDto.ProfileID) + " cannot spend more points than weekly amount")
		return
	}

	recruitEntry.AllocateCurrentWeekPoints(updateRecruitPointsDto.SpentPoints)

	db.Save(&pointAllocation)
	db.Save(&recruitEntry)
	db.Save(&recruitingProfile)
}

func SendScholarshipToRecruit(updateRecruitPointsDto structs.UpdateRecruitPointsDto) (structs.RecruitPlayerProfile, structs.RecruitingTeamProfile) {
	db := dbprovider.GetInstance().GetDB()

	recruitingProfile := GetOnlyRecruitingProfileByTeamID(strconv.Itoa(updateRecruitPointsDto.ProfileID))

	if recruitingProfile.ScholarshipsAvailable == 0 {
		log.Panicln("\nTeamId: " + strconv.Itoa(updateRecruitPointsDto.ProfileID) + " does not have any availabe scholarships")
	}

	crootProfile := GetRecruitProfileByPlayerId(
		strconv.Itoa(updateRecruitPointsDto.RecruitID),
		strconv.Itoa(updateRecruitPointsDto.ProfileID),
	)

	crootProfile.ToggleScholarship(updateRecruitPointsDto.RewardScholarship, updateRecruitPointsDto.RevokeScholarship)
	if !crootProfile.ScholarshipRevoked {
		recruitingProfile.SubtractScholarshipsAvailable()
	} else {
		recruitingProfile.ReallocateScholarship()
	}

	db.Save(&crootProfile)
	db.Save(&recruitingProfile)

	return crootProfile, recruitingProfile
}

func RevokeScholarshipFromRecruit(updateRecruitPointsDto structs.UpdateRecruitPointsDto) (structs.RecruitPlayerProfile, structs.RecruitingTeamProfile) {
	db := dbprovider.GetInstance().GetDB()

	recruitingProfile := GetOnlyRecruitingProfileByTeamID(strconv.Itoa(updateRecruitPointsDto.ProfileID))

	recruitingPointsProfile := GetRecruitProfileByPlayerId(
		strconv.Itoa(updateRecruitPointsDto.RecruitID),
		strconv.Itoa(updateRecruitPointsDto.ProfileID),
	)

	if recruitingPointsProfile.Scholarship {
		fmt.Printf("\nCannot revoke an inexistant scholarship from Recruit " + strconv.Itoa(recruitingPointsProfile.RecruitID))
		return recruitingPointsProfile, recruitingProfile
	}

	// recruitingPointsProfile.ToggleScholarship()
	recruitingProfile.ReallocateScholarship()

	db.Save(&recruitingPointsProfile)
	db.Save(&recruitingProfile)

	return recruitingPointsProfile, recruitingProfile
}

func RecruitSync(CurrentWeek int) {
	db := dbprovider.GetInstance().GetDB()

	// GetCurrentWeek
	timestamp := GetTimestamp()
	if timestamp.RecruitingSynced {
		log.Fatalln("Recruiting already ran for this week. Please wait until next week to sync recruiting again.")
	}
	var recruitModifiers structs.AdminRecruitModifier
	var recruitProfiles []structs.RecruitPlayerProfile
	recruits := GetAllUnsignedRecruits()
	// Get every recruit
	for _, recruit := range recruits {
		recruitProfiles = GetRecruitPlayerProfilesByRecruitId(strconv.Itoa(int(recruit.ID)))

		var recruitProfilesWithScholarship []structs.RecruitPlayerProfile

		totalTeamRecruitProfiles := len(recruitProfiles)

		totalPointsOnRecruit := 0

		var signThreshold float64

		for _, recruitProfile := range recruitProfiles {
			// Calculate efficacy points
			// multiply points
			// Add points to total by recruit profile
			// add points to total points on recruit
			recruitProfile.AddCurrentWeekPointsToTotal()
			totalPointsOnRecruit += recruitProfile.TotalPoints
			if recruitProfile.Scholarship {
				recruitProfilesWithScholarship = append(recruitProfilesWithScholarship, recruitProfile)
			}
		}

		// Re-Sort profiles
		sort.Sort(structs.ByPoints(recruitProfilesWithScholarship))

		// Change?
		// Assign point totals
		// If there are any modifiers
		// Evaluate
		signThreshold = float64(recruitModifiers.ModifierOne-CurrentWeek) * (float64(totalTeamRecruitProfiles/recruitModifiers.ModifierTwo) * math.Log(float64(recruitModifiers.WeeksOfRecruiting-CurrentWeek)))

		// Change logic to withold teams without available scholarships
		if float64(totalPointsOnRecruit) > signThreshold {
			percentageOdds := rand.Intn(totalPointsOnRecruit) + 1
			currentProbability := 0
			winningTeamID := 0

			for _, recruitProfile := range recruitProfilesWithScholarship {
				// If a team has no available scholarships or if a team has 25 commitments, continue
				currentProbability += recruitProfile.TotalPoints
				if currentProbability > percentageOdds {
					// WINNING TEAM
					winningTeamID = recruitProfile.ProfileID
					break
				}
			}

			if winningTeamID > 0 {
				recruitTeamProfile := GetOnlyRecruitingProfileByTeamID(strconv.Itoa(winningTeamID))
				teamAbbreviation := recruitTeamProfile.TeamAbbreviation

				for _, recruitProfile := range recruitProfiles {
					recruitProfile.SetWinningTeamAbbreviation(teamAbbreviation)
					recruitProfile.SignPlayer()
				}
			}

			// Set Recruit property
			recruit.UpdateSigningStatus()
			recruit.UpdateTeamID(winningTeamID)
		}
	}
	timestamp.ToggleRecruiting()
	// Save Recruits and Recruit Player Profiles
	err := db.Save(&recruitProfiles).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not sync all recruiting profiles.")
	}

	err = db.Save(&recruits).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not sync all recruits.")
	}

	err = db.Save(&timestamp).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not save timestamp")
	}
}

func GetRecruitFromRecruitsList(id int, recruits []structs.RecruitPlayerProfile) structs.RecruitPlayerProfile {
	var recruit structs.RecruitPlayerProfile

	for i := 0; i < len(recruits); i++ {
		if recruits[i].RecruitID == id {
			recruit = recruits[i]
			break
		}
	}

	return recruit
}

func CreateCollegeRecruit(createRecruitDTO structs.CreateRecruitDTO) {
	db := dbprovider.GetInstance().GetDB()

	var lastPlayerRecord structs.Player

	err := db.Last(&lastPlayerRecord).Error
	if err != nil {
		log.Fatalln("Could not grab last player record from players table...")
	}

	newID := lastPlayerRecord.ID + 1

	collegeRecruit := &structs.Recruit{}
	collegeRecruit.Map(createRecruitDTO, newID)

	// No Player Record exists, so we shall make one.

	db.Create(&collegeRecruit)

	playerRecord := structs.Player{
		RecruitID: int(collegeRecruit.ID),
	}
	// Create Player Record
	db.Create(&playerRecord)
	// Assign PlayerID to Recruit
	collegeRecruit.AssignPlayerID(int(playerRecord.ID))
	// Save Recruit
	db.Save(&collegeRecruit)
}

func UpdateRecruit(r structs.Recruit) {
	db := dbprovider.GetInstance().GetDB()
	err := db.Save(&r).Error
	if err != nil {
		log.Fatal(err)
	}
}

func GetESPNRanking(r structs.Recruit) int {
	// ESPN Ranking = Star Rank + Archetype Modifier + weight difference + height difference
	// + potential val, and then round.

	starRank := GetESPNStarRank(r.Stars)
	archMod := GetArchetypeModifier(r.Archetype)
	potentialMod := GetPotentialModifier(r.PotentialGrade)

	espnPositionMap := config.ESPNModifiers()
	heightMod := float64(r.Height) / espnPositionMap[r.Position]["Height"]
	weightMod := float64(r.Weight) / espnPositionMap[r.Position]["Weight"]
	espnRanking := math.Round(float64(starRank) + float64(archMod) + potentialMod + heightMod + weightMod)

	return int(espnRanking)
}

func GetESPNStarRank(star int) int {
	if star == 5 {
		return 95
	} else if star == 4 {
		return 85
	} else if star == 3 {
		return 75
	} else if star == 2 {
		return 65
	}
	return 55
}

func GetArchetypeModifier(arch string) int {
	if arch == "Coverage" ||
		arch == "Run Stopper" ||
		arch == "Ball Hawk" ||
		arch == "Man Coverage" ||
		arch == "Pass Rusher" ||
		arch == "Rushing" {
		return 1
	} else if arch == "Possession" ||
		arch == "Field General" ||
		arch == "Nose Tackle" ||
		arch == "Blocking" ||
		arch == "Line Captain" {
		return -1
	} else if arch == "Speed Rusher" ||
		arch == "Pass Rush" || arch == "Scrambler" ||
		arch == "Vertical Threat" ||
		arch == "Speed" {
		return 2
	}
	return 0
}

func GetPotentialModifier(pg string) float64 {
	if pg == "A+" {
		return 1
	} else if pg == "A" {
		return 0.9
	} else if pg == "A-" {
		return 0.8
	} else if pg == "B+" {
		return 0.6
	} else if pg == "B" {
		return 0.4
	} else if pg == "B-" {
		return 0.2
	} else if pg == "C+" {
		return 0
	} else if pg == "C" {
		return -0.15
	} else if pg == "C-" {
		return -0.3
	} else if pg == "D+" {
		return -0.6
	} else if pg == "D" {
		return -0.75
	} else if pg == "D-" {
		return -0.9
	}
	return -1
}

func GetPredictiveOverall(r structs.Recruit) int {
	currentOverall := r.Overall

	var potentialProg int

	if r.PotentialGrade == "B+" ||
		r.PotentialGrade == "A-" ||
		r.PotentialGrade == "A" ||
		r.PotentialGrade == "A+" {
		potentialProg = 7
	} else if r.PotentialGrade == "B" ||
		r.PotentialGrade == "B-" ||
		r.PotentialGrade == "C+" {
		potentialProg = 5
	} else {
		potentialProg = 4
	}

	return currentOverall + (potentialProg * 3)
}
