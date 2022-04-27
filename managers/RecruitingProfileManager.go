package managers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// GetRecruitingProfileByTeamID
func GetOnlyRecruitingProfileByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Where("team_id = ?", TeamID).Find(&profile).Error
	if err != nil {
		log.Fatal(err)
	}

	return profile
}

// GetRecruitingProfileByTeamID -- Dashboard
func GetRecruitingProfileByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Preload("Affinities").Preload("Recruits").Where("id = ?", TeamID).Find(&profile).Error
	if err != nil {
		log.Panicln(err)
	}

	// var recruitingBoard []structs.RecruitPlayerProfile

	// err = db.Where("profile_id = ?", TeamID).Find(&recruitingBoard).Error
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// profile.AddRecruitsToProfile(recruitingBoard)

	return profile
}

func AddRecruitToBoard(RecruitDTO structs.CreateRecruitPointsDto) structs.RecruitPlayerProfile {
	//
	db := dbprovider.GetInstance().GetDB()

	recruitProfile := GetRecruitProfileByPlayerId(strconv.Itoa(RecruitDTO.RecruitID), strconv.Itoa(RecruitDTO.ProfileID))

	if recruitProfile.RecruitID != 0 && recruitProfile.ProfileID != 0 {
		// Replace Recruit Onto Board
		db.Save(&recruitProfile)
		return recruitProfile
	}

	recruitingProfile := structs.RecruitPlayerProfile{
		SeasonID:            RecruitDTO.SeasonID,
		RecruitID:           RecruitDTO.RecruitID,
		ProfileID:           RecruitDTO.ProfileID,
		TotalPoints:         0,
		CurrentWeeksPoints:  0,
		Scholarship:         false,
		ScholarshipRevoked:  false,
		AffinityOneEligible: false,
		AffinityTwoEligible: false,
		TeamAbbreviation:    RecruitDTO.Team,
		RemovedFromBoard:    false,
		IsSigned:            false,
	}

	// Check for Close to Home Affinity

	// Save
	db.Create(&recruitingProfile)

	return recruitingProfile
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
