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
	"github.com/CalebRose/SimFBA/structs"
	"github.com/jinzhu/gorm"
)

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

func CreateRecruitingProfileForRecruit(recruitPointsDto structs.CreateRecruitPointsDto) structs.RecruitPlayerProfile {
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
		TeamAbbreviation:    recruitPointsDto.Team,
		TotalPoints:         0,
		CurrentWeeksPoints:  0,
		SpendingCount:       0,
		Scholarship:         false,
		AffinityOneEligible: recruitPointsDto.AffinityOneEligible,
		AffinityTwoEligible: recruitPointsDto.AffinityTwoEligible,
		RemovedFromBoard:    false,
		IsSigned:            false,
	}
	// Do a check on affinities here

	// Create
	db.Create(&createRecruitEntry)

	return createRecruitEntry
}

func AllocateRecruitPointsForRecruit(updateRecruitPointsDto structs.UpdateRecruitPointsDto) {
	db := dbprovider.GetInstance().GetDB()

	// Recruit Team Profile
	recruitingProfile := GetRecruitingProfileByTeamID(strconv.Itoa(updateRecruitPointsDto.ProfileID))

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

	recruitingProfile := GetRecruitingProfileByTeamID(strconv.Itoa(updateRecruitPointsDto.ProfileID))

	if recruitingProfile.ScholarshipsAvailable == 0 {
		log.Fatalf("\nTeamId: " + strconv.Itoa(updateRecruitPointsDto.ProfileID) + " does not have any availabe scholarships")
	}

	recruitingPointsEntry := GetRecruitProfileByPlayerId(
		strconv.Itoa(updateRecruitPointsDto.RecruitID),
		strconv.Itoa(updateRecruitPointsDto.ProfileID),
	)

	if recruitingPointsEntry.Scholarship {
		log.Fatalf("\nRecruit " + strconv.Itoa(recruitingPointsEntry.RecruitID) + "already has a scholarship")
	}

	recruitingPointsEntry.ToggleScholarship()
	recruitingProfile.SubtractScholarshipsAvailable()

	db.Save(&recruitingPointsEntry)
	db.Save(&recruitingProfile)

	return recruitingPointsEntry, recruitingProfile
}

func RevokeScholarshipFromRecruit(updateRecruitPointsDto structs.UpdateRecruitPointsDto) (structs.RecruitPlayerProfile, structs.RecruitingTeamProfile) {
	db := dbprovider.GetInstance().GetDB()

	recruitingProfile := GetRecruitingProfileByTeamID(strconv.Itoa(updateRecruitPointsDto.ProfileID))

	recruitingPointsProfile := GetRecruitProfileByPlayerId(
		strconv.Itoa(updateRecruitPointsDto.RecruitID),
		strconv.Itoa(updateRecruitPointsDto.ProfileID),
	)

	if recruitingPointsProfile.Scholarship {
		fmt.Printf("\nCannot revoke an inexistant scholarship from Recruit " + strconv.Itoa(recruitingPointsProfile.RecruitID))
		return recruitingPointsProfile, recruitingProfile
	}

	recruitingPointsProfile.ToggleScholarship()
	recruitingProfile.ReallocateScholarship()

	db.Save(&recruitingPointsProfile)
	db.Save(&recruitingProfile)

	return recruitingPointsProfile, recruitingProfile
}

func RecruitSync(CurrentWeek int) {
	db := dbprovider.GetInstance().GetDB()

	// GetCurrentWeek
	var recruitModifiers structs.AdminRecruitModifier
	recruits := GetAllRecruits()
	// Get every recruit
	for _, recruit := range recruits {
		recruitProfiles := GetRecruitPlayerProfilesByRecruitId(strconv.Itoa(int(recruit.ID)))
		var recruitProfilesWithScholarship []structs.RecruitPlayerProfile
		totalTeamRecruitProfiles := len(recruitProfiles)
		totalPointsOnRecruit := 0
		var signThreshold float64
		for _, recruitProfile := range recruitProfiles {
			recruitProfile.AddCurrentWeekPointsToTotal()
			totalPointsOnRecruit += recruitProfile.TotalPoints
			if recruitProfile.Scholarship {
				recruitProfilesWithScholarship = append(recruitProfilesWithScholarship, recruitProfile)
			}
		}

		// Re-Sort profiles
		sort.Sort(structs.ByPoints(recruitProfilesWithScholarship))

		signThreshold = float64(recruitModifiers.ModifierOne-CurrentWeek) * (float64(totalTeamRecruitProfiles/recruitModifiers.ModifierTwo) * math.Log(float64(recruitModifiers.WeeksOfRecruiting-CurrentWeek)))
		if float64(totalPointsOnRecruit) > signThreshold {
			percentageOdds := rand.Intn(totalPointsOnRecruit) + 1
			currentProbability := 0
			winningTeamID := 0

			for _, recruitProfile := range recruitProfilesWithScholarship {
				currentProbability += recruitProfile.TotalPoints
				if currentProbability > percentageOdds {
					// WINNING TEAM
					winningTeamID = recruitProfile.ProfileID
					break
				}
			}

			if winningTeamID > 0 {
				recruitTeamProfile := GetRecruitingProfileByTeamID(strconv.Itoa(winningTeamID))
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
		// Save all recruit profiles after iterating recruit
		db.Save(&recruitProfiles)
	}
	// Save Recruits
	db.Save(&recruits)
}
