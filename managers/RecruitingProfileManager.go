package managers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/jinzhu/gorm"
)

// GetRecruitingProfileByTeamID
func GetOnlyRecruitingProfileByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Where("id = ?", TeamID).Find(&profile).Error
	if err != nil {
		log.Fatal(err)
	}

	return profile
}

// GetRecruitingProfileForDashboardByTeamID -- Dashboard
func GetRecruitingProfileForDashboardByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Preload("Affinities").Preload("Recruits.Recruit.RecruitPlayerProfiles", func(db *gorm.DB) *gorm.DB {
		return db.Order("total_points DESC").Where("total_points > 0")
	}).Where("id = ?", TeamID).Find(&profile).Error
	if err != nil {
		log.Panicln(err)
	}

	return profile
}

func GetRecruitingProfileForTeamBoardByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Preload("Affinities").Preload("Recruits.Recruit.RecruitPlayerProfiles", func(db *gorm.DB) *gorm.DB {
		return db.Order("total_points DESC").Where("total_points > 0")
	}).Where("id = ?", TeamID).Find(&profile).Error
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

func GetRecruitingNeeds(TeamID string) map[string]int {
	needsMap := make(map[string]int)

	teamRoster := GetAllCollegePlayersByTeamId(TeamID)

	for _, player := range teamRoster {
		if player.IsRedshirting {
			continue
		}
		if (player.Year == 4 && !player.IsRedshirt) || (player.Year == 5 && player.IsRedshirt) {
			needsMap[player.Position] += 1
		}
	}

	return needsMap
}

func AddRecruitToBoard(RecruitDTO structs.CreateRecruitProfileDto) structs.RecruitPlayerProfile {
	//
	db := dbprovider.GetInstance().GetDB()

	recruitProfile := GetRecruitProfileByPlayerId(strconv.Itoa(RecruitDTO.RecruitID), strconv.Itoa(RecruitDTO.ProfileID))

	if recruitProfile.RecruitID != 0 && recruitProfile.ProfileID != 0 {
		// Replace Recruit Onto Board
		recruitProfile.ToggleRemoveFromBoard()
		db.Save(&recruitProfile)
		return recruitProfile
	}

	recruitingProfile := structs.RecruitPlayerProfile{
		SeasonID:            RecruitDTO.SeasonID,
		RecruitID:           RecruitDTO.RecruitID,
		ProfileID:           RecruitDTO.ProfileID,
		TotalPoints:         0,
		CurrentWeeksPoints:  0,
		SpendingCount:       0,
		Scholarship:         false,
		ScholarshipRevoked:  false,
		AffinityOneEligible: RecruitDTO.AffinityOneEligible,
		AffinityTwoEligible: RecruitDTO.AffinityTwoEligible,
		TeamAbbreviation:    RecruitDTO.Team,
		RemovedFromBoard:    false,
		IsSigned:            false,
	}

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
		log.Panicln("Recruit has already been removed from Team Recruiting Board.")
	}

	recruitingPointsProfile.ToggleRemoveFromBoard()
	db.Save(&recruitingPointsProfile)

	return recruitingPointsProfile
}

func UpdateRecruitingProfile(updateRecruitingBoardDto structs.UpdateRecruitingBoardDTO) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var teamID = strconv.Itoa(updateRecruitingBoardDto.TeamID)

	var teamProfile = GetOnlyRecruitingProfileByTeamID(teamID)

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
