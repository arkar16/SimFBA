package managers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// GetRecruitingProfileByTeamID
func GetRecruitingProfileByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Where("team_id = ?", TeamID).Find(&profile).Error
	if err != nil {
		log.Fatal(err)
	}

	return profile
}

func RemoveRecruitFromBoard(updateRecruitPointsDto structs.UpdateRecruitPointsDto) structs.RecruitPlayerProfile {
	db := dbprovider.GetInstance().GetDB()

	recruitingPointsProfile := GetRecruitProfileByPlayerId(
		strconv.Itoa(updateRecruitPointsDto.RecruitID),
		strconv.Itoa(updateRecruitPointsDto.ProfileID))

	if recruitingPointsProfile.RemovedFromBoard {
		panic("Recruit has already been removed from Team Recruiting Board.")
	}

	recruitingPointsProfile.ToggleRemoveFromBoard()
	db.Save(&recruitingPointsProfile)

	return recruitingPointsProfile
}

func UpdateRecruitingProfile(updateRecruitingBoardDto structs.UpdateRecruitingBoardDTO) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var teamID = strconv.Itoa(updateRecruitingBoardDto.TeamID)

	var teamProfile = GetRecruitingProfileByTeamID(teamID)

	var recruitProfiles = GetRecruitsByTeamProfileID(teamID)

	var updatedRecruits = updateRecruitingBoardDto.Recruits

	currentPoints := 0

	for i := 0; i < len(recruitProfiles); i++ {
		updatedRecruit := GetRecruitFromRecruitsList(recruitProfiles[i].RecruitID, updatedRecruits)

		if updatedRecruit.CurrentWeeksPoints > 0 &&
			recruitProfiles[i].CurrentWeeksPoints != updatedRecruit.CurrentWeeksPoints {

			// Allocate Points to Profile
			currentPoints += updatedRecruit.CurrentWeeksPoints
			teamProfile.AllocateSpentPoints(currentPoints)
			// If total not surpassed, allocate to the recruit and continue
			if teamProfile.SpentPoints <= teamProfile.WeeklyPoints {
				recruitProfiles[i].AllocateCurrentWeekPoints(updatedRecruit.CurrentWeeksPoints)
				fmt.Println("Saving recruit " + strconv.Itoa(recruitProfiles[i].RecruitID))
				db.Save(&recruitProfiles[i])
			} else {
				panic("Error: Allocated more points for Profile " + strconv.Itoa(teamProfile.TeamID) + " than what is allowed.")
			}
		}
	}

	// Save profile
	db.Save(&teamProfile)

	return teamProfile
}
