package managers

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetCollegeCoachByCoachName(name string) structs.CollegeCoach {
	db := dbprovider.GetInstance().GetDB()

	var coach structs.CollegeCoach

	err := db.Where("coach_name = ?", name).Find(&coach).Error
	if err != nil || coach.ID == 0 {
		coach = structs.CollegeCoach{
			CoachName:                      name,
			TeamID:                         0,
			OverallWins:                    0,
			OverallLosses:                  0,
			OverallConferenceChampionships: 0,
			BowlWins:                       0,
			BowlLosses:                     0,
			PlayoffWins:                    0,
			PlayoffLosses:                  0,
			IsActive:                       false,
		}
	}

	return coach
}

func GetNFLUserByUsername(username string) structs.NFLUser {
	db := dbprovider.GetInstance().GetDB()

	var user structs.NFLUser

	err := db.Where("username = ?", username).Find(&user).Error
	if err != nil || user.ID == 0 {
		user = structs.NFLUser{
			Username:    username,
			TeamID:      0,
			TotalWins:   0,
			TotalLosses: 0,
			IsActive:    true,
		}
	}

	return user
}
