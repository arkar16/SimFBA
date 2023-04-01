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

func GetCollegeWeek(weekID string, ts structs.Timestamp) structs.CollegeWeek {
	db := dbprovider.GetInstance().GetDB()

	var week structs.CollegeWeek

	db.Where("week = ? AND season_id = ?", weekID, ts.CollegeSeasonID).Find(&week)

	return week
}

func MoveUpWeek() structs.Timestamp {
	db := dbprovider.GetInstance().GetDB()
	timestamp := GetTimestamp()
	if timestamp.RecruitingSynced {
		// Sync to Next Week
		UpdateStandings(timestamp)
		UpdateGameplanPenalties()
		timestamp.SyncToNextWeek()
		db.Save(&timestamp)
	}

	return timestamp
}

// UpdateTimestamp - Update the timestamp
func UpdateTimestamp(updateTimestampDto structs.UpdateTimestampDto) structs.Timestamp {
	db := dbprovider.GetInstance().GetDB()

	timestamp := GetTimestamp()

	if updateTimestampDto.MoveUpCollegeWeek {
		// Update Standings based on current week's games

		// Sync to Next Week
		UpdateStandings(timestamp)
		UpdateGameplanPenalties()
		timestamp.SyncToNextWeek()
	}
	// else if updateTimestampDto.ThursdayGames && !timestamp.ThursdayGames {
	// 	timestamp.ToggleThursdayGames()
	// } else if updateTimestampDto.FridayGames && !timestamp.FridayGames {
	// 	timestamp.ToggleFridayGames()
	// } else if updateTimestampDto.SaturdayMorning && !timestamp.SaturdayMorning {
	// 	timestamp.ToggleSaturdayMorningGames()
	// } else if updateTimestampDto.SaturdayNoon && !timestamp.SaturdayNoon {
	// 	timestamp.ToggleSaturdayNoonGames()
	// } else if updateTimestampDto.SaturdayEvening && !timestamp.SaturdayEvening {
	// 	timestamp.ToggleSaturdayEveningGames()
	// } else if updateTimestampDto.SaturdayNight && !timestamp.SaturdayNight {
	// 	timestamp.ToggleSaturdayNightGames()
	// }

	if updateTimestampDto.ToggleRecruitingLock {
		timestamp.ToggleLockRecruiting()
	}

	// if updateTimestampDto.RESSynced && !timestamp.RecruitingEfficiencySynced {
	// 	timestamp.ToggleRES()
	// 	SyncRecruitingEfficiency(timestamp)
	// }

	if updateTimestampDto.RecruitingSynced && !timestamp.RecruitingSynced && timestamp.IsRecruitingLocked {
		SyncRecruiting(timestamp)
		timestamp.ToggleRecruiting()
	}

	err := db.Save(&timestamp).Error
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

// Season Funcs
func MoveUpInOffseasonFreeAgency() {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()
	ts.MoveUpFreeAgencyRound()
	db.Save(&ts)
}

func GetNewsLogs(weekID string, seasonID string) []structs.NewsLog {
	db := dbprovider.GetInstance().GetDB()

	var logs []structs.NewsLog

	err := db.Where("week_id = ? AND season_id = ?", weekID, seasonID).Find(&logs).Error
	if err != nil {
		fmt.Println(err)
	}

	return logs
}

func GetAllNewsLogs() []structs.NewsLog {
	db := dbprovider.GetInstance().GetDB()

	var logs []structs.NewsLog

	err := db.Where("league = ?", "CFB").Find(&logs).Error
	if err != nil {
		fmt.Println(err)
	}

	return logs
}

func GetAllNFLNewsLogs() []structs.NewsLog {
	db := dbprovider.GetInstance().GetDB()

	var logs []structs.NewsLog

	err := db.Where("league = ?", "NFL").Find(&logs).Error
	if err != nil {
		fmt.Println(err)
	}

	return logs
}

func GetWeeksInASeason(seasonID string, weekID string) []structs.CollegeWeek {
	db := dbprovider.GetInstance().GetDB()

	var weeks []structs.CollegeWeek

	err := db.Where("season_id = ?", seasonID).Find(&weeks).Error
	if err != nil {
		fmt.Println(err)
	}

	return weeks
}
