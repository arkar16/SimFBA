package repository

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func FindCollegeGamesRecords(SeasonID string, isSpringGames bool) []structs.CollegeGame {
	db := dbprovider.GetInstance().GetDB()

	var games []structs.CollegeGame

	query := db.Model(&games)

	if len(SeasonID) > 0 {
		query = query.Where("season_id = ?", SeasonID)
	}

	if err := query.Order("week_id asc").Where("is_spring_game = ?", isSpringGames).Find(&games).Error; err != nil {
		return []structs.CollegeGame{}
	}
	return games
}

func FindNFLGamesRecords(SeasonID string, isSpringGames bool) []structs.NFLGame {
	db := dbprovider.GetInstance().GetDB()

	var games []structs.NFLGame

	query := db.Model(&games)

	if len(SeasonID) > 0 {
		query = query.Where("season_id = ?", SeasonID)
	}

	if err := query.Order("week_id asc").Where("is_preseason_game = ?", isSpringGames).Find(&games).Error; err != nil {
		return []structs.NFLGame{}
	}
	return games
}
