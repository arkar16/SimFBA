package managers

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetRecruitingModifiers() structs.AdminRecruitModifier {
	db := dbprovider.GetInstance().GetDB()

	var recruitModifiers structs.AdminRecruitModifier

	db.First(&recruitModifiers)

	return recruitModifiers
}

func SyncRecruiting() {
	db := dbprovider.GetInstance().GetDB()

	// GetCurrentWeek
	timestamp := GetTimestamp()
	if timestamp.RecruitingSynced {
		log.Fatalln("Recruiting already ran for this week. Please wait until next week to sync recruiting again.")
	}

	recruitModifiers := GetRecruitingModifiers()

	var recruitProfiles []structs.RecruitPlayerProfile

	// Get every recruit
	recruits := GetAllUnsignedRecruits()

	// Iterate through every recruit
	for _, recruit := range recruits {
		recruitProfiles = recruit.RecruitPlayerProfiles

		var recruitProfilesWithScholarship []structs.RecruitPlayerProfile

		totalTeamRecruitProfiles := len(recruitProfiles)

		totalPointsOnRecruit := 0

		var signThreshold float64

		for _, recruitProfile := range recruitProfiles {
			if recruitProfile.CurrentWeeksPoints == 0 {
				continue
			}
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
		signThreshold = float64(recruitModifiers.ModifierOne-timestamp.CollegeWeek) * (float64(totalTeamRecruitProfiles/recruitModifiers.ModifierTwo) * math.Log(float64(recruitModifiers.WeeksOfRecruiting-timestamp.CollegeWeek)))

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

	// Update rank system for all teams
	recruitingProfiles := GetRecruitingProfileForRecruitSync()

	for _, rp := range recruitingProfiles {
		signedRecruits := GetSignedRecruitsByTeamProfileID(strconv.Itoa(rp.TeamID))

		team247Rank := Get247TeamRanking(rp, signedRecruits)
		teamESPNRank := GetESPNTeamRanking(rp, signedRecruits)
		teamRivalsRank := GetRivalsTeamRanking(rp, signedRecruits)

		rp.Assign247Rank(team247Rank)
		rp.AssignESPNRank(teamESPNRank)
		rp.AssignRivalsRank(teamRivalsRank)
	}

	timestamp.ToggleRecruiting()
	// Save Recruits, Recruit Player Profiles, and Team Profiles
	// err := db.Save(&recruitProfiles).Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	log.Fatalf("Could not sync all recruiting profiles.")
	// }

	// Should save both recruits and recruit player profiles
	err := db.Save(&recruits).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not sync all recruits.")
	}

	// Save the Recruiting Profiles
	err = db.Save(&recruitingProfiles).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not save timestamp")
	}

	err = db.Save(&timestamp).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not save timestamp")
	}
}
