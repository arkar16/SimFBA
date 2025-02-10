package repository

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

type StandingsQuery struct {
	SeasonID string
	TeamID   string
}

func FindAllCollegeStandingsRecords(conditions StandingsQuery) []structs.CollegeStandings {
	db := dbprovider.GetInstance().GetDB()

	var standings []structs.CollegeStandings

	query := db.Model(&standings)

	if len(conditions.SeasonID) > 0 {
		query = query.Where("season_id = ?", conditions.SeasonID)
	}

	if len(conditions.TeamID) > 0 {
		query = query.Where("team_id = ?", conditions.TeamID)
	}

	if err := query.Order("conference_id asc").Order("conference_losses asc").Order("conference_wins desc").
		Order("total_losses asc").Order("total_wins desc").Find(&standings).Error; err != nil {
		return []structs.CollegeStandings{}
	}

	return standings
}

func FindAllNFLStandingsRecords(conditions StandingsQuery) []structs.NFLStandings {
	db := dbprovider.GetInstance().GetDB()

	var standings []structs.NFLStandings

	query := db.Model(&standings)

	if len(conditions.SeasonID) > 0 {
		query = query.Where("season_id = ?", conditions.SeasonID)
	}

	if len(conditions.TeamID) > 0 {
		query = query.Where("team_id = ?", conditions.TeamID)
	}

	if err := query.Order("conference_id asc").Order("conference_losses asc").Order("conference_wins desc").
		Order("total_losses asc").Order("total_wins desc").Find(&standings).Error; err != nil {
		return []structs.NFLStandings{}
	}

	return standings
}
