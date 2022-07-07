package managers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetCollegePlayerStatsByGame(PlayerID string, GameID string) structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats structs.CollegePlayerStats

	db.Where("college_player_id = ? and game_id = ?", PlayerID, GameID).Find(&playerStats)

	return playerStats
}

func GetCareerCollegePlayerStatsByPlayerID(PlayerID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("college_player_id = ?", PlayerID).Find(&playerStats)

	return playerStats
}

func GetCollegePlayerStatsByPlayerIDAndSeason(PlayerID string, SeasonID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("college_player_id = ? and season_id = ?", PlayerID, SeasonID).Find(&playerStats)

	return playerStats
}

func GetAllPlayerStatsByGame(GameID string, TeamID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("game_id = ? and season_id = ?", GameID, TeamID).Find(&playerStats)

	return playerStats
}

func GetAllPlayerStatsByWeek(WeekID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("week_id = ?", WeekID).Find(&playerStats)

	return playerStats
}

// TEAM STATS
func GetSeasonalTeamStats(TeamID string, SeasonID string) []structs.CollegeTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.CollegeTeamStats

	db.Where("team_id = ? AND season_id = ?", TeamID, SeasonID).Find(&teamStats)

	return teamStats
}

func GetCollegeTeamStatsByGame(GameID string) []structs.CollegeTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.CollegeTeamStats

	db.Where("game_id = ?", GameID).Find(&teamStats)

	return teamStats
}

func GetCollegeTeamStatsByWeek(WeekID string) []structs.CollegeTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.CollegeTeamStats

	db.Where("week_id = ?", WeekID).Find(&teamStats)

	return teamStats
}

func GetCollegeTeamStatsBySeason(SeasonID string) []structs.CollegeTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.CollegeTeamStats

	db.Where("season_id = ?", SeasonID).Find(&teamStats)

	return teamStats
}

func GetHistoricalTeamStats(TeamID string, SeasonID string) []structs.CollegeTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.CollegeTeamStats

	db.Where("team_id = ? AND season_id != ?", TeamID, SeasonID).Find(&teamStats)

	return teamStats
}

func ExportStatisticsFromSim(exportStatsDTO structs.ExportStatsDTO) {
	db := dbprovider.GetInstance().GetDB()
	fmt.Println("START")

	tsChn := make(chan structs.Timestamp)

	go func() {
		ts := GetTimestamp()
		tsChn <- ts
	}()

	timestamp := <-tsChn
	close(tsChn)

	var teamStats []structs.CollegeTeamStats

	for _, gameDataDTO := range exportStatsDTO.GameStatDTOs {

		record := make(chan structs.CollegeGame)

		go func() {
			asynchronousGame := GetCollegeGameByAbbreviationsWeekAndSeasonID(gameDataDTO.HomeTeam.Abbreviation, strconv.Itoa(timestamp.CollegeWeekID), strconv.Itoa(timestamp.CollegeSeasonID))
			record <- asynchronousGame
		}()

		gameRecord := <-record
		close(record)
		var playerStats []structs.CollegePlayerStats

		// Team Stats Export
		homeTeamChn := make(chan structs.CollegeTeam)

		go func() {
			homeTeam := GetTeamByTeamAbbr(gameDataDTO.HomeTeam.Abbreviation)
			homeTeamChn <- homeTeam
		}()

		ht := <-homeTeamChn
		close(homeTeamChn)

		homeTeam := structs.CollegeTeamStats{
			TeamID:        int(ht.ID),
			GameID:        int(gameRecord.ID),
			WeekID:        gameRecord.WeekID,
			SeasonID:      gameRecord.SeasonID,
			OpposingTeam:  gameDataDTO.AwayTeam.Abbreviation,
			BaseTeamStats: gameDataDTO.HomeTeam.MapToBaseTeamStatsObject(),
		}

		teamStats = append(teamStats, homeTeam)

		// Away Team
		awayTeamChn := make(chan structs.CollegeTeam)

		go func() {
			awayTeam := GetTeamByTeamAbbr(gameDataDTO.AwayTeam.Abbreviation)
			awayTeamChn <- awayTeam
		}()

		at := <-awayTeamChn
		close(awayTeamChn)

		awayTeam := structs.CollegeTeamStats{
			TeamID:        int(at.ID),
			GameID:        int(gameRecord.ID),
			WeekID:        gameRecord.WeekID,
			SeasonID:      gameRecord.SeasonID,
			OpposingTeam:  gameDataDTO.HomeTeam.Abbreviation,
			BaseTeamStats: gameDataDTO.AwayTeam.MapToBaseTeamStatsObject(),
		}

		teamStats = append(teamStats, awayTeam)

		// Player Stat Export
		// HomePlayers
		for _, player := range gameDataDTO.HomePlayers {
			collegePlayerStats := structs.CollegePlayerStats{
				CollegePlayerID: player.GetPlayerID(),
				TeamID:          homeTeam.TeamID,
				GameID:          homeTeam.GameID,
				WeekID:          gameRecord.WeekID,
				SeasonID:        gameRecord.SeasonID,
				OpposingTeam:    gameDataDTO.AwayTeam.Abbreviation,
				BasePlayerStats: player.MapTobasePlayerStatsObject(),
			}
			playerStats = append(playerStats, collegePlayerStats)
		}

		// AwayPlayers
		for _, player := range gameDataDTO.AwayPlayers {
			collegePlayerStats := structs.CollegePlayerStats{
				CollegePlayerID: player.GetPlayerID(),
				TeamID:          awayTeam.TeamID,
				GameID:          awayTeam.GameID,
				WeekID:          gameRecord.WeekID,
				SeasonID:        gameRecord.SeasonID,
				OpposingTeam:    gameDataDTO.HomeTeam.Abbreviation,
				BasePlayerStats: player.MapTobasePlayerStatsObject(),
			}
			playerStats = append(playerStats, collegePlayerStats)
		}

		// Update Game
		gameRecord.UpdateScore(gameDataDTO.HomeScore, gameDataDTO.AwayScore)

		err := db.Save(&gameRecord).Error
		if err != nil {
			log.Panicln("Could not save Game " + strconv.Itoa(int(gameRecord.ID)) + "Between " + gameRecord.HomeTeam + " and " + gameRecord.AwayTeam)
		}

		err = db.CreateInBatches(&playerStats, len(playerStats)).Error
		if err != nil {
			log.Panicln("Could not save player stats from week " + strconv.Itoa(timestamp.CollegeWeek))
		}

		fmt.Println("Finished Game " + strconv.Itoa(int(gameRecord.ID)) + " Between " + gameRecord.HomeTeam + " and " + gameRecord.AwayTeam)
	}

	err := db.CreateInBatches(&teamStats, len(teamStats)).Error
	if err != nil {
		log.Panicln("Could not save team stats!")
	}

}
