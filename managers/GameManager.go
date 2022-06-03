package managers

import (
	"fmt"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetCollegeGamesByWeekIdAndSeasonID(WeekID string, SeasonID string) []structs.CollegeGame {
	db := dbprovider.GetInstance().GetDB()

	var games []structs.CollegeGame

	db.Where("week_id = ? AND season_id = ?", WeekID, SeasonID).Find(&games)

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

func GetCollegeGameByAbbreviationsWeekAndSeasonID(HomeTeamAbbr string, WeekID string, SeasonID string) structs.CollegeGame {
	db := dbprovider.GetInstance().GetDB()

	var game structs.CollegeGame

	err := db.Where("season_id = ? AND week_id = ? AND (home_team = ? OR away_team = ?)", SeasonID, WeekID, HomeTeamAbbr, HomeTeamAbbr).Find(&game).Error
	if err != nil {
		fmt.Println("Could not find game!")
	}

	return game
}
