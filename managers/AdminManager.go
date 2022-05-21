package managers

import (
	"fmt"
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// Timestamp Funcs
// GetTimestamp -- Get the Timestamp
func GetTimestamp() structs.Timestamp {
	db := dbprovider.GetInstance().GetDB()

	var timestamp structs.Timestamp

	db.First(&timestamp)

	return timestamp
}

// UpdateTimestamp - Update the timestamp
func UpdateTimestamp(updateTimestampDto structs.UpdateTimestampDto) structs.Timestamp {
	db := dbprovider.GetInstance().GetDB()

	timestamp := GetTimestamp()

	if updateTimestampDto.MoveUpCollegeWeek {
		// Update Standings based on current week's games

		// Sync to Next Week
		timestamp.SyncToNextWeek()

	} else if updateTimestampDto.ThursdayGames && !timestamp.ThursdayGames {
		timestamp.ToggleThursdayGames()
	} else if updateTimestampDto.FridayGames && !timestamp.FridayGames {
		timestamp.ToggleFridayGames()
	} else if updateTimestampDto.SaturdayMorning && !timestamp.SaturdayMorning {
		timestamp.ToggleSaturdayMorningGames()
	} else if updateTimestampDto.SaturdayNoon && !timestamp.SaturdayNoon {
		timestamp.ToggleSaturdayNoonGames()
	} else if updateTimestampDto.SaturdayEvening && !timestamp.SaturdayEvening {
		timestamp.ToggleSaturdayEveningGames()
	} else if updateTimestampDto.SaturdayNight && !timestamp.SaturdayNight {
		timestamp.ToggleSaturdayNightGames()
	}

	if updateTimestampDto.ToggleRecruitingLock {
		timestamp.ToggleLockRecruiting()
	}

	if updateTimestampDto.RESSynced && !timestamp.RecruitingEfficiencySynced {
		timestamp.ToggleRES()
		SyncRecruitingEfficiency(timestamp)
	}

	if updateTimestampDto.RecruitingSynced && !timestamp.RecruitingSynced && timestamp.RecruitingEfficiencySynced {
		SyncRecruiting(timestamp)
		timestamp.ToggleRecruiting()

	}

	err := db.Save(timestamp).Error
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Could not save timestamp")
	}

	return timestamp
}

// Week Funcs
func CreateCollegeWeek() {

}

// Season Funcs
func CreateCollegeSeason() {

}
