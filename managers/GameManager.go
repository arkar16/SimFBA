package managers

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetCollegeGamesByWeekId(id string) []structs.CollegeGame {
	db := dbprovider.GetInstance().GetDB()

	var games []structs.CollegeGame

	db.Order("time_slot asc").Where("WeekID = ?", id).Find(&games)

	return games
}

func GetCollegeGamesByTimeslotAndWeekId(id string, timeslot string) []structs.CollegeGame {
	db := dbprovider.GetInstance().GetDB()

	var games []structs.CollegeGame

	db.Order("time_slot asc").Where("week_id = ? AND time_slot = ?", id, timeslot).Find(&games)

	return games
}

func GetCollegeGamesByTeamIdAndSeasonId(TeamID string, SeasonID string) []structs.CollegeGame {
	db := dbprovider.GetInstance().GetDB()

	var games []structs.CollegeGame

	db.Order("week_id asc").Where("season_id = ? AND (home_team_id = ? OR away_team_id = ?)", SeasonID, TeamID, TeamID).Find(&games)

	return games
}
