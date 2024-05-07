package managers

import (
	"fmt"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/repository"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
	"gorm.io/gorm"
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

func GetCollegePlayerSeasonStatsByPlayerIDAndSeason(PlayerID string, SeasonID string) structs.CollegePlayerSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats structs.CollegePlayerSeasonStats

	db.Where("college_player_id = ? and season_id = ?", PlayerID, SeasonID).Find(&playerStats)

	return playerStats
}

func GetCollegePlayerStatsByPlayerIDAndSeason(PlayerID string, SeasonID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("college_player_id = ? and season_id = ?", PlayerID, SeasonID).Find(&playerStats)

	return playerStats
}

func GetNFLPlayerStatsByPlayerIDAndSeason(PlayerID string, SeasonID string) []structs.NFLPlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.NFLPlayerStats

	db.Where("nfl_player_id = ? and season_id = ?", PlayerID, SeasonID).Find(&playerStats)

	return playerStats
}

func GetAllCollegePlayerStatsByGame(GameID string, TeamID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("game_id = ? and team_id = ?", GameID, TeamID).Find(&playerStats)

	return playerStats
}

func GetAllCollegePlayerSnapsByGame(GameID string) []structs.CollegePlayerGameSnaps {
	db := dbprovider.GetInstance().GetDB()

	var playerSnaps []structs.CollegePlayerGameSnaps

	db.Where("game_id = ?", GameID).Find(&playerSnaps)

	return playerSnaps
}

func GetAllNFLPlayerStatsByGame(GameID string, TeamID string) []structs.NFLPlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.NFLPlayerStats

	db.Where("game_id = ? and team_id = ?", GameID, TeamID).Find(&playerStats)

	return playerStats
}

func GetAllNFLPlayerSnapsByGame(GameID string) []structs.NFLPlayerGameSnaps {
	db := dbprovider.GetInstance().GetDB()

	var playerSnaps []structs.NFLPlayerGameSnaps

	db.Where("game_id = ?", GameID).Find(&playerSnaps)

	return playerSnaps
}

func GetAllPlayerStatsByWeek(WeekID string) []structs.CollegePlayerStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerStats

	db.Where("week_id = ?", WeekID).Find(&playerStats)

	return playerStats
}

func GetTeamStatsByWeekAndTeam(TeamID string, Week string) structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()

	collegeWeek := GetCollegeWeek(Week, ts)
	var collegeTeam structs.CollegeTeam

	if collegeWeek.ID == uint(ts.CollegeWeekID) {
		return structs.CollegeTeam{}
	} else {
		err := db.Preload("TeamStats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ? AND week_id = ?", collegeWeek.SeasonID, collegeWeek.ID)
		}).Where("id = ?", TeamID).Find(&collegeTeam).Error
		if err != nil {
			fmt.Println("Could not find college team and stats from week")
		}

	}
	return collegeTeam
}

// TEAM STATS
func GetSeasonalTeamStats(TeamID string, SeasonID string) models.CollegeTeamResponse {
	db := dbprovider.GetInstance().GetDB()

	var collegeTeam structs.CollegeTeam

	err := db.Preload("TeamSeasonStats", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ?", SeasonID)
	}).Where("id = ?", TeamID).Find(&collegeTeam).Error
	if err != nil {
		fmt.Println("Could not find college team and stats from week")
	}

	ct := models.CollegeTeamResponse{
		ID:           int(collegeTeam.ID),
		BaseTeam:     collegeTeam.BaseTeam,
		ConferenceID: collegeTeam.ConferenceID,
		Conference:   collegeTeam.Conference,
		DivisionID:   collegeTeam.DivisionID,
		Division:     collegeTeam.Division,
		SeasonStats:  collegeTeam.TeamSeasonStats,
	}

	return ct
}

func GetCollegeTeamSeasonStatsBySeason(TeamID string, SeasonID string) structs.CollegeTeamSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats structs.CollegeTeamSeasonStats

	err := db.Where("team_id = ? AND season_id = ?", TeamID, SeasonID).Find(&teamStats).Error
	if err != nil {
		return structs.CollegeTeamSeasonStats{}
	}

	return teamStats
}

func GetNFLTeamSeasonStatsByTeamANDSeason(TeamID string, SeasonID string) structs.NFLTeamSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats structs.NFLTeamSeasonStats

	err := db.Where("team_id = ? AND season_id = ?", TeamID, SeasonID).Find(&teamStats).Error
	if err != nil {
		return structs.NFLTeamSeasonStats{}
	}

	return teamStats
}

func GetCollegeSeasonStatsByPlayerAndSeason(PlayerID, SeasonID string) structs.CollegePlayerSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats structs.CollegePlayerSeasonStats

	err := db.Where("college_player_id = ? AND season_id = ?", PlayerID, SeasonID).Find(&playerStats).Error
	if err != nil {
		return structs.CollegePlayerSeasonStats{}
	}

	return playerStats
}

func GetCollegeSeasonSnapsByPlayerAndSeason(PlayerID, SeasonID string) structs.CollegePlayerSeasonSnaps {
	db := dbprovider.GetInstance().GetDB()

	var playerStats structs.CollegePlayerSeasonSnaps

	err := db.Where("college_player_id = ? AND season_id = ?", PlayerID, SeasonID).Find(&playerStats).Error
	if err != nil {
		return structs.CollegePlayerSeasonSnaps{}
	}

	return playerStats
}

func GetNFLSeasonStatsByPlayerAndSeason(PlayerID, SeasonID string) structs.NFLPlayerSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats structs.NFLPlayerSeasonStats

	err := db.Where("nfl_player_id = ? AND season_id = ?", PlayerID, SeasonID).Find(&playerStats).Error
	if err != nil {
		return structs.NFLPlayerSeasonStats{}
	}

	return playerStats
}

func GetNFLSeasonSnapsByPlayerAndSeason(PlayerID, SeasonID string) structs.NFLPlayerSeasonSnaps {
	db := dbprovider.GetInstance().GetDB()

	var playerStats structs.NFLPlayerSeasonSnaps

	err := db.Where("nfl_player_id = ? AND season_id = ?", PlayerID, SeasonID).Find(&playerStats).Error
	if err != nil {
		return structs.NFLPlayerSeasonSnaps{}
	}

	return playerStats
}

func GetCollegeSeasonStatsBySeason(SeasonID string) []structs.CollegeTeamSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.CollegeTeamSeasonStats

	db.Where("season_id = ?", SeasonID).Find(&teamStats)

	return teamStats
}

func GetCollegePlayerSeasonStatsBySeason(SeasonID string) []structs.CollegePlayerSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.CollegePlayerSeasonStats

	db.Where("season_id = ?", SeasonID).Find(&playerStats)

	return playerStats
}

func GetNFLTeamSeasonStatsBySeason(SeasonID string) []structs.NFLTeamSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.NFLTeamSeasonStats

	db.Where("season_id = ?", SeasonID).Find(&teamStats)

	return teamStats
}

func GetNFLPlayerSeasonStatsBySeason(SeasonID string) []structs.NFLPlayerSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.NFLPlayerSeasonStats

	db.Where("season_id = ?", SeasonID).Find(&playerStats)

	return playerStats
}

func GetAllNFLPlayerSeasonStatsByPlayerID(playerID string) []structs.NFLPlayerSeasonStats {
	db := dbprovider.GetInstance().GetDB()

	var playerStats []structs.NFLPlayerSeasonStats

	db.Where("nfl_player_id = ?", playerID).Find(&playerStats)

	return playerStats
}

func GetCollegeTeamStatsByGame(TeamID string, GameID string) structs.CollegeTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats structs.CollegeTeamStats

	db.Where("team_id = ? AND game_id = ?", TeamID, GameID).Find(&teamStats)

	return teamStats
}

func GetNFLTeamStatsByGame(TeamID string, GameID string) structs.NFLTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats structs.NFLTeamStats

	db.Where("team_id = ? AND game_id = ?", TeamID, GameID).Find(&teamStats)

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

	db.Where("team_id = ? AND season_id = ?", TeamID, SeasonID).Find(&teamStats)

	return teamStats
}

func GetNFLHistoricalTeamStats(TeamID string, SeasonID string) []structs.NFLTeamStats {
	db := dbprovider.GetInstance().GetDB()

	var teamStats []structs.NFLTeamStats

	db.Where("team_id = ? AND season_id = ?", TeamID, SeasonID).Find(&teamStats)

	return teamStats
}

func GetAllCollegeTeamsWithStatsBySeasonID(seasonID, weekID, viewType string) []structs.CollegeTeamResponse {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	if viewType == "SEASON" {
		db.Preload("TeamSeasonStats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ?", seasonID)
		}).Find(&teams)
	} else {
		db.Preload("TeamStats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ? AND week_id = ?", seasonID, weekID)
		}).Find(&teams)
	}

	var ctResponse []structs.CollegeTeamResponse

	for _, team := range teams {
		if len(team.TeamStats) == 0 && viewType == "WEEK" {
			continue
		}
		var teamstat structs.CollegeTeamStats
		if viewType == "WEEK" {
			teamstat = team.TeamStats[0]
		}
		ct := structs.CollegeTeamResponse{
			ID:           int(team.ID),
			BaseTeam:     team.BaseTeam,
			ConferenceID: team.ConferenceID,
			Conference:   team.Conference,
			DivisionID:   team.DivisionID,
			Division:     team.Division,
			SeasonStats:  team.TeamSeasonStats,
			Stats:        teamstat,
		}

		ctResponse = append(ctResponse, ct)
	}

	return ctResponse
}

func GetAllNFLTeamsWithStatsBySeasonID(seasonID, weekID, viewType string) []structs.NFLTeamResponse {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.NFLTeam

	if viewType == "SEASON" {
		db.Preload("TeamSeasonStats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ?", seasonID)
		}).Find(&teams)
	} else {
		db.Preload("TeamStats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ? AND week_id = ?", seasonID, weekID)
		}).Find(&teams)
	}

	var ctResponse []structs.NFLTeamResponse

	for _, team := range teams {
		if len(team.TeamStats) == 0 && viewType == "WEEK" {
			continue
		}
		var teamstat structs.NFLTeamStats
		var seasonstat structs.NFLTeamSeasonStats
		if viewType == "WEEK" && len(team.TeamStats) > 0 {
			teamstat = team.TeamStats[0]
		} else if len(team.TeamSeasonStats) > 0 {
			seasonstat = team.TeamSeasonStats[0]
		}
		ct := structs.NFLTeamResponse{
			ID:           int(team.ID),
			BaseTeam:     team.BaseTeam,
			ConferenceID: int(team.ConferenceID),
			Conference:   team.Conference,
			DivisionID:   int(team.DivisionID),
			Division:     team.Division,
			SeasonStats:  seasonstat,
			Stats:        teamstat,
		}

		ctResponse = append(ctResponse, ct)
	}

	return ctResponse
}

func ResetCFBSeasonalStats() {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()
	seasonID := strconv.Itoa(int(ts.CollegeSeasonID))
	teams := GetAllCollegeTeams()

	for _, team := range teams {
		teamID := strconv.Itoa(int(team.ID))
		teamStats := GetHistoricalTeamStats(teamID, seasonID)
		seasonStats := GetCollegeTeamSeasonStatsBySeason(teamID, seasonID)
		seasonStats.ResetStats()
		seasonStats.MapStats(teamStats)
		db.Save(&seasonStats)
		fmt.Println("Reset Season Stats for " + team.TeamName)
	}

	players := GetAllCollegePlayers()

	for _, player := range players {
		playerID := strconv.Itoa(int(player.ID))
		playerStats := GetCollegePlayerStatsByPlayerIDAndSeason(playerID, seasonID)
		if len(playerStats) > 0 {
			seasonStats := GetCollegeSeasonStatsByPlayerAndSeason(playerID, seasonID)
			seasonStats.ResetStats()
			seasonStats.MapStats(playerStats)
			db.Save(&seasonStats)
		}
		fmt.Println("Reset Season Stats for " + player.FirstName + " " + player.LastName + " " + player.Position)
	}

	standings := GetAllConferenceStandingsBySeasonID(seasonID)

	for _, standing := range standings {
		if standing.ID == 0 {
			continue
		}
		standing.ResetCFBStandings()

		games := GetCollegeGamesByTeamIdAndSeasonId(strconv.Itoa(int(standing.TeamID)), seasonID)

		for _, game := range games {
			if !game.GameComplete {
				continue
			}
			standing.UpdateCollegeStandings(game)
		}

		db.Save(&standing)
		fmt.Println("Updated Standings for " + standing.TeamName)
	}
}

func ResetNFLSeasonalStats() {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()
	seasonID := strconv.Itoa(int(ts.NFLSeasonID))
	teams := GetAllNFLTeams()

	for _, team := range teams {
		teamID := strconv.Itoa(int(team.ID))
		teamStats := GetNFLHistoricalTeamStats(teamID, seasonID)
		seasonStats := GetNFLTeamSeasonStatsByTeamANDSeason(teamID, seasonID)
		seasonStats.ResetStats()
		seasonStats.MapStats(teamStats)
		db.Save(&seasonStats)
		fmt.Println("Reset Season Stats for " + team.TeamName)
	}

	players := GetAllNFLPlayers()

	for _, player := range players {
		playerID := strconv.Itoa(int(player.ID))
		playerStats := GetNFLPlayerStatsByPlayerIDAndSeason(playerID, seasonID)
		if len(playerStats) > 0 {
			seasonStats := GetNFLSeasonStatsByPlayerAndSeason(playerID, seasonID)
			seasonStats.ResetStats()
			seasonStats.MapStats(playerStats, ts)
			db.Save(&seasonStats)
		}
		fmt.Println("Reset Season Stats for " + player.FirstName + " " + player.LastName + " " + player.Position)
	}

	standings := GetAllNFLStandingsBySeasonID(seasonID)

	for _, standing := range standings {
		if standing.ID == 0 {
			continue
		}
		standing.ResetNFLStandings()

		games := GetNFLGamesByTeamIdAndSeasonId(strconv.Itoa(int(standing.TeamID)), seasonID)

		for _, game := range games {
			if !game.GameComplete {
				continue
			}
			standing.UpdateNFLStandings(game)
		}

		db.Save(&standing)
		fmt.Println("Updated Standings for " + standing.TeamName)
	}
}

func ExportCFBStatisticsFromSim(gameStats []structs.GameStatDTO) {
	db := dbprovider.GetInstance().GetDB()
	fmt.Println("START")

	var teamStats []structs.CollegeTeamStats

	for _, gameDataDTO := range gameStats {
		gameID := strconv.Itoa(int(gameDataDTO.GameID))
		record := make(chan structs.CollegeGame)

		go func() {
			asynchronousGame := GetCollegeGameByGameID(gameID)
			record <- asynchronousGame
		}()

		gameRecord := <-record
		close(record)

		if gameRecord.GameComplete {
			continue
		}
		var playerStats []structs.CollegePlayerStats
		snapMap := gameDataDTO.PlayerSnapTracker.PlayerSnapCounts
		var playerSnaps []structs.CollegePlayerGameSnaps

		homeTeamID := strconv.Itoa(int(gameRecord.HomeTeamID))
		awayTeamID := strconv.Itoa(int(gameRecord.AwayTeamID))

		// Team Stats Export
		homeTeamChn := make(chan structs.CollegeTeam)

		go func() {
			homeTeam := GetTeamByTeamID(homeTeamID)
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
			awayTeam := GetTeamByTeamID(awayTeamID)
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
				OpposingTeam:    at.TeamAbbr,
				BasePlayerStats: player.MapTobasePlayerStatsObject(),
				Year:            player.Year,
				IsRedshirt:      player.IsRedshirt,
			}
			snaps := snapMap[player.GetPlayerID()]
			cpSnaps := structs.CollegePlayerGameSnaps{
				BasePlayerGameSnaps: structs.BasePlayerGameSnaps{
					SeasonID: uint(gameRecord.SeasonID),
					PlayerID: uint(player.GetPlayerID()),
					GameID:   gameRecord.ID,
					WeekID:   uint(gameRecord.WeekID),
				},
			}
			if snaps["QB"] > 0 {
				cpSnaps.MapSnapsToPosition("QB", snaps["QB"])
			}
			if snaps["RB"] > 0 {
				cpSnaps.MapSnapsToPosition("RB", snaps["RB"])
			}
			if snaps["FB"] > 0 {
				cpSnaps.MapSnapsToPosition("FB", snaps["FB"])
			}
			if snaps["WR"] > 0 {
				cpSnaps.MapSnapsToPosition("WR", snaps["WR"])
			}
			if snaps["TE"] > 0 {
				cpSnaps.MapSnapsToPosition("TE", snaps["TE"])
			}
			if snaps["OT"] > 0 {
				cpSnaps.MapSnapsToPosition("OT", snaps["OT"])
			}
			if snaps["OG"] > 0 {
				cpSnaps.MapSnapsToPosition("OG", snaps["OG"])
			}
			if snaps["C"] > 0 {
				cpSnaps.MapSnapsToPosition("C", snaps["C"])
			}
			if snaps["DT"] > 0 {
				cpSnaps.MapSnapsToPosition("DT", snaps["DT"])
			}
			if snaps["DE"] > 0 {
				cpSnaps.MapSnapsToPosition("DE", snaps["DE"])
			}
			if snaps["OLB"] > 0 {
				cpSnaps.MapSnapsToPosition("OLB", snaps["OLB"])
			}
			if snaps["ILB"] > 0 {
				cpSnaps.MapSnapsToPosition("ILB", snaps["ILB"])
			}
			if snaps["CB"] > 0 {
				cpSnaps.MapSnapsToPosition("CB", snaps["CB"])
			}
			if snaps["FS"] > 0 {
				cpSnaps.MapSnapsToPosition("FS", snaps["FS"])
			}
			if snaps["SS"] > 0 {
				cpSnaps.MapSnapsToPosition("SS", snaps["SS"])
			}
			if snaps["P"] > 0 {
				cpSnaps.MapSnapsToPosition("P", snaps["P"])
			}
			if snaps["K"] > 0 {
				cpSnaps.MapSnapsToPosition("K", snaps["K"])
			}
			if snaps["KR"] > 0 {
				cpSnaps.MapSnapsToPosition("KR", snaps["KR"])
			}
			if snaps["KOS"] > 0 {
				cpSnaps.MapSnapsToPosition("KOS", snaps["KOS"])
			}
			if snaps["PR"] > 0 {
				cpSnaps.MapSnapsToPosition("PR", snaps["PR"])
			}
			if snaps["ST"] > 0 {
				cpSnaps.MapSnapsToPosition("ST", snaps["ST"])
			}
			playerStats = append(playerStats, collegePlayerStats)
			playerSnaps = append(playerSnaps, cpSnaps)
		}

		// AwayPlayers
		for _, player := range gameDataDTO.AwayPlayers {
			collegePlayerStats := structs.CollegePlayerStats{
				CollegePlayerID: player.GetPlayerID(),
				TeamID:          awayTeam.TeamID,
				GameID:          awayTeam.GameID,
				WeekID:          gameRecord.WeekID,
				SeasonID:        gameRecord.SeasonID,
				OpposingTeam:    ht.TeamAbbr,
				BasePlayerStats: player.MapTobasePlayerStatsObject(),
				Year:            player.Year,
				IsRedshirt:      player.IsRedshirt,
			}
			snaps := snapMap[player.GetPlayerID()]
			cpSnaps := structs.CollegePlayerGameSnaps{
				BasePlayerGameSnaps: structs.BasePlayerGameSnaps{
					SeasonID: uint(gameRecord.SeasonID),
					PlayerID: uint(player.GetPlayerID()),
					GameID:   gameRecord.ID,
					WeekID:   uint(gameRecord.WeekID),
				},
			}
			if snaps["QB"] > 0 {
				cpSnaps.MapSnapsToPosition("QB", snaps["QB"])
			}
			if snaps["RB"] > 0 {
				cpSnaps.MapSnapsToPosition("RB", snaps["RB"])
			}
			if snaps["FB"] > 0 {
				cpSnaps.MapSnapsToPosition("FB", snaps["FB"])
			}
			if snaps["WR"] > 0 {
				cpSnaps.MapSnapsToPosition("WR", snaps["WR"])
			}
			if snaps["TE"] > 0 {
				cpSnaps.MapSnapsToPosition("TE", snaps["TE"])
			}
			if snaps["OT"] > 0 {
				cpSnaps.MapSnapsToPosition("OT", snaps["OT"])
			}
			if snaps["OG"] > 0 {
				cpSnaps.MapSnapsToPosition("OG", snaps["OG"])
			}
			if snaps["C"] > 0 {
				cpSnaps.MapSnapsToPosition("C", snaps["C"])
			}
			if snaps["DT"] > 0 {
				cpSnaps.MapSnapsToPosition("DT", snaps["DT"])
			}
			if snaps["DE"] > 0 {
				cpSnaps.MapSnapsToPosition("DE", snaps["DE"])
			}
			if snaps["OLB"] > 0 {
				cpSnaps.MapSnapsToPosition("OLB", snaps["OLB"])
			}
			if snaps["ILB"] > 0 {
				cpSnaps.MapSnapsToPosition("ILB", snaps["ILB"])
			}
			if snaps["CB"] > 0 {
				cpSnaps.MapSnapsToPosition("CB", snaps["CB"])
			}
			if snaps["FS"] > 0 {
				cpSnaps.MapSnapsToPosition("FS", snaps["FS"])
			}
			if snaps["SS"] > 0 {
				cpSnaps.MapSnapsToPosition("SS", snaps["SS"])
			}
			if snaps["P"] > 0 {
				cpSnaps.MapSnapsToPosition("P", snaps["P"])
			}
			if snaps["K"] > 0 {
				cpSnaps.MapSnapsToPosition("K", snaps["K"])
			}
			if snaps["KR"] > 0 {
				cpSnaps.MapSnapsToPosition("KR", snaps["KR"])
			}
			if snaps["KOS"] > 0 {
				cpSnaps.MapSnapsToPosition("KOS", snaps["KOS"])
			}
			if snaps["PR"] > 0 {
				cpSnaps.MapSnapsToPosition("PR", snaps["PR"])
			}
			if snaps["ST"] > 0 {
				cpSnaps.MapSnapsToPosition("ST", snaps["ST"])
			}
			playerSnaps = append(playerSnaps, cpSnaps)
			playerStats = append(playerStats, collegePlayerStats)
		}

		// Update Game
		gameRecord.UpdateScore(gameDataDTO.HomeScore, gameDataDTO.AwayScore)

		repository.SaveCFBGameRecord(gameRecord, db)

		repository.CreateCFBPlayerStatsInBatch(playerStats, db)

		pbp := gameDataDTO.Plays
		playByPlays := []structs.CollegePlayByPlay{}
		for _, p := range pbp {
			play := structs.CollegePlayByPlay{}
			play.Map(p)
			playByPlays = append(playByPlays, play)
		}
		repository.CreateCFBPlayByPlaysInBatch(playByPlays, db)
		repository.CreateCFBSnapsInBatch(playerSnaps, db)

		fmt.Println("Finished Game " + strconv.Itoa(int(gameRecord.ID)) + " Between " + gameRecord.HomeTeam + " and " + gameRecord.AwayTeam)
	}

	repository.CreateCFBTeamStatsInBatch(teamStats, db)
}

func ExportNFLStatisticsFromSim(gameStats []structs.GameStatDTO) {
	db := dbprovider.GetInstance().GetDB()
	fmt.Println("START")

	var teamStats []structs.NFLTeamStats

	for _, gameDataDTO := range gameStats {
		gameID := strconv.Itoa(int(gameDataDTO.GameID))
		record := make(chan structs.NFLGame)

		go func() {
			asynchronousGame := GetNFLGameByGameID(gameID)
			record <- asynchronousGame
		}()

		gameRecord := <-record
		close(record)
		if gameRecord.GameComplete {
			continue
		}
		homeTeamID := strconv.Itoa(int(gameRecord.HomeTeamID))
		awayTeamID := strconv.Itoa(int(gameRecord.AwayTeamID))

		snapMap := gameDataDTO.PlayerSnapTracker.PlayerSnapCounts
		var playerSnaps []structs.NFLPlayerGameSnaps
		var playerStats []structs.NFLPlayerStats

		// Team Stats Export
		homeTeamChn := make(chan structs.NFLTeam)

		go func() {
			homeTeam := GetNFLTeamByTeamID(homeTeamID)
			homeTeamChn <- homeTeam
		}()

		ht := <-homeTeamChn
		close(homeTeamChn)

		// Away Team
		awayTeamChn := make(chan structs.NFLTeam)

		go func() {
			awayTeam := GetNFLTeamByTeamID(awayTeamID)
			awayTeamChn <- awayTeam
		}()

		at := <-awayTeamChn
		close(awayTeamChn)

		homeTeam := structs.NFLTeamStats{
			TeamID:        ht.ID,
			GameID:        gameRecord.ID,
			WeekID:        uint(gameRecord.WeekID),
			SeasonID:      uint(gameRecord.SeasonID),
			OpposingTeam:  at.TeamAbbr,
			BaseTeamStats: gameDataDTO.HomeTeam.MapToBaseTeamStatsObject(),
		}

		teamStats = append(teamStats, homeTeam)
		awayTeam := structs.NFLTeamStats{
			TeamID:        at.ID,
			GameID:        gameRecord.ID,
			WeekID:        uint(gameRecord.WeekID),
			SeasonID:      uint(gameRecord.SeasonID),
			OpposingTeam:  ht.TeamAbbr,
			BaseTeamStats: gameDataDTO.AwayTeam.MapToBaseTeamStatsObject(),
		}

		teamStats = append(teamStats, awayTeam)

		// Player Stat Export
		// HomePlayers
		for _, player := range gameDataDTO.HomePlayers {
			nflPlayerStats := structs.NFLPlayerStats{
				NFLPlayerID:     player.GetPlayerID(),
				TeamID:          int(homeTeam.TeamID),
				GameID:          int(homeTeam.GameID),
				WeekID:          gameRecord.WeekID,
				SeasonID:        gameRecord.SeasonID,
				OpposingTeam:    gameDataDTO.AwayTeam.Abbreviation,
				BasePlayerStats: player.MapTobasePlayerStatsObject(),
				Year:            player.Year,
			}
			snaps := snapMap[player.GetPlayerID()]
			snap := structs.NFLPlayerGameSnaps{
				BasePlayerGameSnaps: structs.BasePlayerGameSnaps{
					SeasonID: uint(gameRecord.SeasonID),
					PlayerID: uint(player.GetPlayerID()),
					GameID:   gameRecord.ID,
					WeekID:   uint(gameRecord.WeekID),
				},
			}
			if snaps["QB"] > 0 {
				snap.MapSnapsToPosition("QB", snaps["QB"])
			}
			if snaps["RB"] > 0 {
				snap.MapSnapsToPosition("RB", snaps["RB"])
			}
			if snaps["FB"] > 0 {
				snap.MapSnapsToPosition("FB", snaps["FB"])
			}
			if snaps["WR"] > 0 {
				snap.MapSnapsToPosition("WR", snaps["WR"])
			}
			if snaps["TE"] > 0 {
				snap.MapSnapsToPosition("TE", snaps["TE"])
			}
			if snaps["OT"] > 0 {
				snap.MapSnapsToPosition("OT", snaps["OT"])
			}
			if snaps["OG"] > 0 {
				snap.MapSnapsToPosition("OG", snaps["OG"])
			}
			if snaps["C"] > 0 {
				snap.MapSnapsToPosition("C", snaps["C"])
			}
			if snaps["DT"] > 0 {
				snap.MapSnapsToPosition("DT", snaps["DT"])
			}
			if snaps["DE"] > 0 {
				snap.MapSnapsToPosition("DE", snaps["DE"])
			}
			if snaps["OLB"] > 0 {
				snap.MapSnapsToPosition("OLB", snaps["OLB"])
			}
			if snaps["ILB"] > 0 {
				snap.MapSnapsToPosition("ILB", snaps["ILB"])
			}
			if snaps["CB"] > 0 {
				snap.MapSnapsToPosition("CB", snaps["CB"])
			}
			if snaps["FS"] > 0 {
				snap.MapSnapsToPosition("FS", snaps["FS"])
			}
			if snaps["SS"] > 0 {
				snap.MapSnapsToPosition("SS", snaps["SS"])
			}
			if snaps["P"] > 0 {
				snap.MapSnapsToPosition("P", snaps["P"])
			}
			if snaps["K"] > 0 {
				snap.MapSnapsToPosition("K", snaps["K"])
			}
			if snaps["KR"] > 0 {
				snap.MapSnapsToPosition("KR", snaps["KR"])
			}
			if snaps["KOS"] > 0 {
				snap.MapSnapsToPosition("KOS", snaps["KOS"])
			}
			if snaps["PR"] > 0 {
				snap.MapSnapsToPosition("PR", snaps["PR"])
			}
			if snaps["ST"] > 0 {
				snap.MapSnapsToPosition("ST", snaps["ST"])
			}
			playerSnaps = append(playerSnaps, snap)
			playerStats = append(playerStats, nflPlayerStats)
		}

		// AwayPlayers
		for _, player := range gameDataDTO.AwayPlayers {
			nflPlayerStats := structs.NFLPlayerStats{
				NFLPlayerID:     player.GetPlayerID(),
				TeamID:          int(awayTeam.TeamID),
				GameID:          int(awayTeam.GameID),
				WeekID:          gameRecord.WeekID,
				SeasonID:        gameRecord.SeasonID,
				OpposingTeam:    gameDataDTO.HomeTeam.Abbreviation,
				BasePlayerStats: player.MapTobasePlayerStatsObject(),
				Year:            player.Year,
			}
			snaps := snapMap[player.GetPlayerID()]
			snap := structs.NFLPlayerGameSnaps{
				BasePlayerGameSnaps: structs.BasePlayerGameSnaps{
					SeasonID: uint(gameRecord.SeasonID),
					PlayerID: uint(player.GetPlayerID()),
					GameID:   gameRecord.ID,
					WeekID:   uint(gameRecord.WeekID),
				},
			}
			if snaps["QB"] > 0 {
				snap.MapSnapsToPosition("QB", snaps["QB"])
			}
			if snaps["RB"] > 0 {
				snap.MapSnapsToPosition("RB", snaps["RB"])
			}
			if snaps["FB"] > 0 {
				snap.MapSnapsToPosition("FB", snaps["FB"])
			}
			if snaps["WR"] > 0 {
				snap.MapSnapsToPosition("WR", snaps["WR"])
			}
			if snaps["TE"] > 0 {
				snap.MapSnapsToPosition("TE", snaps["TE"])
			}
			if snaps["OT"] > 0 {
				snap.MapSnapsToPosition("OT", snaps["OT"])
			}
			if snaps["OG"] > 0 {
				snap.MapSnapsToPosition("OG", snaps["OG"])
			}
			if snaps["C"] > 0 {
				snap.MapSnapsToPosition("C", snaps["C"])
			}
			if snaps["DT"] > 0 {
				snap.MapSnapsToPosition("DT", snaps["DT"])
			}
			if snaps["DE"] > 0 {
				snap.MapSnapsToPosition("DE", snaps["DE"])
			}
			if snaps["OLB"] > 0 {
				snap.MapSnapsToPosition("OLB", snaps["OLB"])
			}
			if snaps["ILB"] > 0 {
				snap.MapSnapsToPosition("ILB", snaps["ILB"])
			}
			if snaps["CB"] > 0 {
				snap.MapSnapsToPosition("CB", snaps["CB"])
			}
			if snaps["FS"] > 0 {
				snap.MapSnapsToPosition("FS", snaps["FS"])
			}
			if snaps["SS"] > 0 {
				snap.MapSnapsToPosition("SS", snaps["SS"])
			}
			if snaps["P"] > 0 {
				snap.MapSnapsToPosition("P", snaps["P"])
			}
			if snaps["K"] > 0 {
				snap.MapSnapsToPosition("K", snaps["K"])
			}
			if snaps["KR"] > 0 {
				snap.MapSnapsToPosition("KR", snaps["KR"])
			}
			if snaps["KOS"] > 0 {
				snap.MapSnapsToPosition("KOS", snaps["KOS"])
			}
			if snaps["PR"] > 0 {
				snap.MapSnapsToPosition("PR", snaps["PR"])
			}
			if snaps["ST"] > 0 {
				snap.MapSnapsToPosition("ST", snaps["ST"])
			}
			playerSnaps = append(playerSnaps, snap)
			playerStats = append(playerStats, nflPlayerStats)
		}

		// Update Game
		gameRecord.UpdateScore(gameDataDTO.HomeScore, gameDataDTO.AwayScore)

		repository.SaveNFLGameRecord(gameRecord, db)

		repository.CreateNFLPlayerStatsInBatch(playerStats, db)

		pbp := gameDataDTO.Plays
		playByPlays := []structs.NFLPlayByPlay{}
		for _, p := range pbp {
			play := structs.NFLPlayByPlay{}
			play.Map(p)
			playByPlays = append(playByPlays, play)
		}
		repository.CreateNFLPlayByPlaysInBatch(playByPlays, db)
		repository.CreateNFLSnapsInBatch(playerSnaps, db)

		fmt.Println("Finished Game " + strconv.Itoa(int(gameRecord.ID)) + " Between " + gameRecord.HomeTeam + " and " + gameRecord.AwayTeam)
	}

	repository.CreateNFLTeamStatsInBatch(teamStats, db)
}

func GetCFBGameResultsByGameID(gameID string) structs.GameResultsResponse {
	game := GetCollegeGameByGameID(gameID)
	htID := strconv.Itoa(game.HomeTeamID)
	atID := strconv.Itoa(game.AwayTeamID)

	homePlayers := GetAllCollegePlayersWithGameStatsByTeamID(htID, gameID)
	awayPlayers := GetAllCollegePlayersWithGameStatsByTeamID(atID, gameID)
	homeStats := GetCollegeTeamStatsByGame(htID, gameID)
	awayStats := GetCollegeTeamStatsByGame(atID, gameID)
	score := structs.ScoreBoard{
		Q1Home:  homeStats.Score1Q,
		Q2Home:  homeStats.Score2Q,
		Q3Home:  homeStats.Score3Q,
		Q4Home:  homeStats.Score4Q,
		OT1Home: homeStats.Score5Q,
		OT2Home: homeStats.Score6Q,
		OT3Home: homeStats.Score7Q,
		OT4Home: homeStats.ScoreOT,
		Q1Away:  awayStats.Score1Q,
		Q2Away:  awayStats.Score2Q,
		Q3Away:  awayStats.Score3Q,
		Q4Away:  awayStats.Score4Q,
		OT1Away: awayStats.Score5Q,
		OT2Away: awayStats.Score6Q,
		OT3Away: awayStats.Score7Q,
		OT4Away: awayStats.ScoreOT,
	}
	participantMap := getGameParticipantMap(homePlayers, awayPlayers)

	playByPlays := GetCFBPlayByPlaysByGameID(gameID)
	// Generate the Play By Play Response
	playbyPlayResponseList := GenerateCFBPlayByPlayResponse(playByPlays, participantMap, false, game.HomeTeam, game.AwayTeam)

	return structs.GameResultsResponse{
		HomePlayers: homePlayers,
		AwayPlayers: awayPlayers,
		PlayByPlays: playbyPlayResponseList,
		Score:       score,
	}
}

func GetNFLGameResultsByGameID(gameID string) structs.GameResultsResponse {
	game := GetNFLGameByGameID(gameID)
	htID := strconv.Itoa(game.HomeTeamID)
	atID := strconv.Itoa(game.AwayTeamID)
	homePlayers := GetAllNFLPlayersWithGameStatsByTeamID(strconv.Itoa(game.HomeTeamID), gameID)
	awayPlayers := GetAllNFLPlayersWithGameStatsByTeamID(strconv.Itoa(game.AwayTeamID), gameID)
	homeStats := GetNFLTeamStatsByGame(htID, gameID)
	awayStats := GetNFLTeamStatsByGame(atID, gameID)
	score := structs.ScoreBoard{
		Q1Home:  homeStats.Score1Q,
		Q2Home:  homeStats.Score2Q,
		Q3Home:  homeStats.Score3Q,
		Q4Home:  homeStats.Score4Q,
		OT1Home: homeStats.Score5Q,
		OT2Home: homeStats.Score6Q,
		OT3Home: homeStats.Score7Q,
		OT4Home: homeStats.ScoreOT,
		Q1Away:  awayStats.Score1Q,
		Q2Away:  awayStats.Score2Q,
		Q3Away:  awayStats.Score3Q,
		Q4Away:  awayStats.Score4Q,
		OT1Away: awayStats.Score5Q,
		OT2Away: awayStats.Score6Q,
		OT3Away: awayStats.Score7Q,
		OT4Away: awayStats.ScoreOT,
	}
	participantMap := getGameParticipantMap(homePlayers, awayPlayers)

	playByPlays := GetNFLPlayByPlaysByGameID(gameID)
	// Generate the Play By Play Response
	playbyPlayResponseList := GenerateNFLPlayByPlayResponse(playByPlays, participantMap, false, game.HomeTeam, game.AwayTeam)
	return structs.GameResultsResponse{
		HomePlayers: homePlayers,
		AwayPlayers: awayPlayers,
		PlayByPlays: playbyPlayResponseList,
		Score:       score,
	}
}

func GetCFBPlayByPlaysByGameID(id string) []structs.CollegePlayByPlay {
	db := dbprovider.GetInstance().GetDB()

	plays := []structs.CollegePlayByPlay{}

	db.Where("game_id = ?", id).Find(&plays)

	return plays
}

func GetNFLPlayByPlaysByGameID(id string) []structs.NFLPlayByPlay {
	db := dbprovider.GetInstance().GetDB()

	plays := []structs.NFLPlayByPlay{}

	db.Where("game_id = ?", id).Find(&plays)

	return plays
}

func GenerateCFBPlayByPlayResponse(playByPlays []structs.CollegePlayByPlay, participantMap map[uint]structs.GameResultsPlayer, isStream bool, ht, at string) []structs.PlayByPlayResponse {
	playbyPlayResponseList := []structs.PlayByPlayResponse{}
	// Get Player Information
	touchDown := false
	for idx, p := range playByPlays {
		number := idx + 1
		playType := util.GetPlayTypeByEnum(p.PlayTypeID)
		offFormation := util.GetOffensiveFormationByEnum(p.OffFormationID)
		defFormation := util.GetDefensiveFormationByEnum(p.DefensiveFormationID)
		defTendency := util.GetDefensiveTendencyByEnum(p.DefensiveTendency)
		playName := util.GetPlayNameByEnum(p.PlayNameID)
		poa := util.GetPointOfAttackByEnum(p.OffensivePoA)
		lb := util.GetCoverageStr(p.LBCoverage)
		cb := util.GetCoverageStr(p.CBCoverage)
		s := util.GetCoverageStr(p.SCoverage)
		poss := ht
		los := p.LineOfScrimmage
		losSide := ht
		if !p.HomeHasBall {
			poss = at
			losSide = at
		}
		if los > 50 {
			los = 100 - p.LineOfScrimmage
			losSide = at
			if !p.HomeHasBall {
				losSide = ht
			}
		}
		losFull := strconv.Itoa(int(los)) + " " + losSide

		play := structs.PlayByPlayResponse{
			PlayNumber:         uint(number),
			HomeTeamID:         p.HomeTeamID,
			HomeTeamScore:      p.HomeTeamScore,
			AwayTeamID:         p.AwayTeamID,
			AwayTeamScore:      p.AwayTeamScore,
			Quarter:            p.Quarter,
			TimeRemaining:      p.TimeRemaining,
			Down:               p.Down,
			Distance:           p.Distance,
			LineOfScrimmage:    losFull,
			PlayType:           playType,
			PlayName:           playName,
			PointOfAttack:      poa,
			OffensiveFormation: offFormation,
			DefensiveFormation: defFormation,
			DefensiveTendency:  defTendency,
			Possession:         poss,
			QBPlayerID:         p.QBPlayerID,
			BallCarrierID:      p.BallCarrierID,
			Tackler1ID:         p.Tackler1ID,
			Tackler2ID:         p.Tackler2ID,
			ResultYards:        p.ResultYards,
			BlitzNumber:        p.BlitzNumber,
			LBCoverage:         lb,
			CBCoverage:         cb,
			SCoverage:          s,
		}
		var result []string
		if isStream {
			result = generateStreamString(p.PlayByPlay, playType, playName, poa, participantMap, touchDown)
		} else {
			result = generateResultsString(p.PlayByPlay, playType, playName, participantMap, touchDown)
		}
		play.AddResult(result, isStream)

		if p.IsTouchdown && !touchDown {
			touchDown = true
		} else if !p.IsTouchdown && touchDown {
			touchDown = false
		}

		playbyPlayResponseList = append(playbyPlayResponseList, play)
	}

	return playbyPlayResponseList
}

func GenerateNFLPlayByPlayResponse(playByPlays []structs.NFLPlayByPlay, participantMap map[uint]structs.GameResultsPlayer, isStream bool, ht, at string) []structs.PlayByPlayResponse {
	playbyPlayResponseList := []structs.PlayByPlayResponse{}
	// Get Player Information
	touchDown := false
	for idx, p := range playByPlays {
		number := idx + 1
		playType := util.GetPlayTypeByEnum(p.PlayTypeID)
		offFormation := util.GetOffensiveFormationByEnum(p.OffFormationID)
		defFormation := util.GetDefensiveFormationByEnum(p.DefensiveFormationID)
		defTendency := util.GetDefensiveTendencyByEnum(p.DefensiveTendency)
		playName := util.GetPlayNameByEnum(p.PlayNameID)
		poa := util.GetPointOfAttackByEnum(p.OffensivePoA)
		lb := util.GetCoverageStr(p.LBCoverage)
		cb := util.GetCoverageStr(p.CBCoverage)
		s := util.GetCoverageStr(p.SCoverage)
		poss := ht

		los := p.LineOfScrimmage
		losSide := ht
		if !p.HomeHasBall {
			poss = at
			losSide = at
		}
		if los > 50 {
			los = 100 - p.LineOfScrimmage
			losSide = at
			if !p.HomeHasBall {
				losSide = ht
			}
		}
		losFull := strconv.Itoa(int(los)) + " " + losSide

		play := structs.PlayByPlayResponse{
			PlayNumber:         uint(number),
			HomeTeamID:         p.HomeTeamID,
			HomeTeamScore:      p.HomeTeamScore,
			AwayTeamID:         p.AwayTeamID,
			AwayTeamScore:      p.AwayTeamScore,
			Quarter:            p.Quarter,
			TimeRemaining:      p.TimeRemaining,
			Down:               p.Down,
			Distance:           p.Distance,
			LineOfScrimmage:    losFull,
			PlayType:           playType,
			PlayName:           playName,
			PointOfAttack:      poa,
			OffensiveFormation: offFormation,
			DefensiveFormation: defFormation,
			DefensiveTendency:  defTendency,
			Possession:         poss,
			QBPlayerID:         p.QBPlayerID,
			BallCarrierID:      p.BallCarrierID,
			Tackler1ID:         p.Tackler1ID,
			Tackler2ID:         p.Tackler2ID,
			ResultYards:        p.ResultYards,
			BlitzNumber:        p.BlitzNumber,
			LBCoverage:         lb,
			CBCoverage:         cb,
			SCoverage:          s,
		}
		var result []string
		if isStream {
			result = generateStreamString(p.PlayByPlay, playType, playName, poa, participantMap, touchDown)
		} else {
			result = generateResultsString(p.PlayByPlay, playType, playName, participantMap, touchDown)
		}
		play.AddResult(result, isStream)

		if p.IsTouchdown && !touchDown {
			touchDown = true
		} else if !p.IsTouchdown && touchDown {
			touchDown = false
		}

		playbyPlayResponseList = append(playbyPlayResponseList, play)
	}

	return playbyPlayResponseList
}

func generateResultsString(play structs.PlayByPlay, playType, playName string, participantMap map[uint]structs.GameResultsPlayer, twoPtCheck bool) []string {
	qbID := play.QBPlayerID
	bcID := play.BallCarrierID
	t1ID := play.Tackler1ID
	t2ID := play.Tackler2ID
	turnID := play.TurnoverPlayerID
	ijID := play.InjuredPlayerID
	pnID := play.PenaltyPlayerID
	yardsSTR := strconv.Itoa(int(play.ResultYards))
	firstSegment := ""
	secondSegment := ""
	thirdSegment := ""
	// First Segment
	if playType == "Pass" {
		qbLabel := getPlayerLabel(participantMap[qbID])
		yards := util.GetYardsString(play.ResultYards)
		firstSegment = qbLabel

		// Scenarios
		if play.IsSacked {
			firstSegment += util.GetScrambleText(play.IsScramble)
			tackle1Label := getPlayerLabel(participantMap[t1ID])
			if t2ID > 0 {
				tackle2Label := getPlayerLabel(participantMap[t2ID])
				tackle1Label += " and " + tackle2Label
			}
			firstSegment += " is sacked on the play by " + tackle1Label + "for a loss of " + yardsSTR + yards
		} else if play.IsComplete {
			firstSegment += util.GetScrambleText(play.IsScramble)
			bcLabel := getPlayerLabel(participantMap[bcID])
			firstSegment += " throws to " + bcLabel + " complete for " + yardsSTR + yards
		} else if play.IsINT {
			firstSegment += util.GetScrambleText(play.IsScramble)
			bcLabel := getPlayerLabel(participantMap[bcID])
			turnOverLabel := getPlayerLabel(participantMap[turnID])
			secondSegment += " throws and is intercepted! Caught by " +
				turnOverLabel + " and returned for " +
				yardsSTR + " yards from the LOS. Pass was intended for " + bcLabel + ". "
		} else if play.IsScramble {
			firstSegment += " scrambles for " + yardsSTR + yards
		} else {
			if bcID > 0 {
				firstSegment += util.GetScrambleText(play.IsScramble)
				bcLabel := getPlayerLabel(participantMap[bcID])
				firstSegment += " throws it... and it's incomplete. Pass intended for " + bcLabel + ". "
			} else {
				firstSegment += util.GetScrambleText(play.IsScramble)
				firstSegment += " can't find an open receiver and throws it away."
			}
		}

	} else if playType == "Run" {
		bcLabel := getPlayerLabel(participantMap[bcID])
		firstSegment = bcLabel + " carries for " + yardsSTR
		yards := util.GetYardsString(play.ResultYards)
		firstSegment += yards
	} else if playType == "Kickoff" {
		// Need assistance for kickign player ID and returner ID
		kickerLabel := getPlayerLabel(participantMap[qbID])
		recLabel := getPlayerLabel(participantMap[bcID])
		distanceStr := util.GetYardsString(play.KickDistance)
		verb := " kicks for "
		firstSegment = kickerLabel + verb + strconv.Itoa(int(play.KickDistance)) + distanceStr
		if play.KickDistance > 64 {
			if play.KickDistance == 65 {
				firstSegment += " Fielded at the goalline by " + recLabel + ". "
			} else {
				firstSegment += " Fielded deep in the endzone by " + recLabel + ". "
			}
		} else {
			outside := 65 - play.KickDistance
			firstSegment += " Fielded at the " + strconv.Itoa(int(outside)) + " yardline by " + recLabel + ". "
		}
		if play.IsTouchback {
			firstSegment += " Touchback. "
		} else {
			resultYdsStr := strconv.Itoa(int(100 - play.ResultYards))
			resultYards := util.GetYardsString(play.ResultYards)
			firstSegment += recLabel + " returns the ball " + resultYdsStr + resultYards
		}
	} else if playType == "Punt" {
		// Need assistance for punting player ID and returner ID
		// Need assistance for kickign player ID and returner ID
		kickerLabel := getPlayerLabel(participantMap[qbID])
		recLabel := getPlayerLabel(participantMap[bcID])
		distanceStr := util.GetYardsString(play.KickDistance)
		verb := " punts for "
		firstSegment = kickerLabel + verb + strconv.Itoa(int(play.KickDistance)) + distanceStr
		los := play.LineOfScrimmage
		if los > 50 {
			los = 100 - los
		}
		outside := los + play.KickDistance
		netReturnYards := outside - play.ResultYards
		resultYdsStr := strconv.Itoa(int(netReturnYards))
		resultYards := util.GetYardsString(netReturnYards)
		if !play.IsTouchback {
			resultYdsStr = strconv.Itoa(int(netReturnYards))
		}

		if play.IsBlocked {
			blockerLabel := getPlayerLabel(participantMap[turnID])
			firstSegment += " BLOCKED by " + blockerLabel
		} else if play.IsFairCatch {
			firstSegment += recLabel + " calls for a fair catch. "
		} else if play.IsTouchback {
			firstSegment += " Touchback. "
		} else {
			firstSegment += recLabel + " returns the ball " + resultYdsStr + resultYards
		}
	} else if playType == "XP" {
		// Need assistance for kicking player ID and outcome
		kickerLabel := getPlayerLabel(participantMap[qbID])
		firstSegment = kickerLabel + "'s extra point attempt is "
		if play.IsBlocked {
			blockerLabel := getPlayerLabel(participantMap[turnID])
			firstSegment += " BLOCKED by " + blockerLabel + ". No good. "
		} else if play.IsGood {
			firstSegment += "good. "
		} else {
			firstSegment += "no good. "
		}
	} else if playType == "FG" {
		kickerLabel := getPlayerLabel(participantMap[qbID])
		kickDistance := strconv.Itoa(int(play.KickDistance))
		firstSegment = kickerLabel + "'s " + kickDistance + " field goal attempt is "
		if play.IsBlocked {
			blockerLabel := getPlayerLabel(participantMap[turnID])
			firstSegment += " BLOCKED by " + blockerLabel + ". No good. "
		} else if play.IsGood {
			firstSegment += "good. "
		} else {
			firstSegment += "no good. "
		}
	}

	// Second Segment - Tackles and OOB
	if play.IsOutOfBounds && playType != "Kickoff" && !play.IsTouchdown {
		secondSegment = "Ran out of bounds. "
	}
	if play.IsTouchdown && !twoPtCheck {
		secondSegment += "TOUCHDOWN! "
	}
	if !play.IsSacked && t1ID > 0 && !play.IsTouchdown {
		tackle1Label := getPlayerLabel(participantMap[t1ID])
		secondSegment += "Tackled by " + tackle1Label
		if t2ID > 0 {
			tackle2Label := getPlayerLabel(participantMap[t2ID])
			secondSegment += " and " + tackle2Label
		}
		secondSegment += ". "
	}

	if play.IsTouchdown && twoPtCheck {
		secondSegment += "The 2 Point Conversion is GOOD!"
	} else if !play.IsTouchdown && twoPtCheck && playType != "XP" && playType != "FG" {
		secondSegment += "The 2 Point Conversion is NO GOOD!"
	}

	if play.IsFumble {
		turnOverLabel := getPlayerLabel(participantMap[turnID])
		secondSegment += "Fumble! Recovered by " + turnOverLabel + "."
	} else if play.IsSafety {
		secondSegment += "Safety. "
	}

	// Third Segments -- Penalties and Injuries
	if play.PenaltyID > 0 {
		penalty := util.GetPenaltyByEnum(play.PenaltyID)
		thirdSegment = "PENALTY: " + penalty + ". "
		offendingTeam := "Offense. "
		if !play.OnOffense {
			offendingTeam = "Defense. "
		}
		thirdSegment += offendingTeam
		if pnID > 0 {
			player := participantMap[pnID]
			penaltyLabel := getPlayerLabel(player)
			thirdSegment += "Player: " + penaltyLabel + ". "
		}
		penaltyYards := strconv.Itoa(int(play.PenaltyYards))
		yards := util.GetYardsString(play.PenaltyYards)
		thirdSegment += penaltyYards + yards

		if play.PenaltyID == 4 ||
			play.PenaltyID == 5 ||
			play.PenaltyID == 7 ||
			play.PenaltyID == 18 ||
			play.PenaltyID == 20 ||
			play.PenaltyID == 21 ||
			play.PenaltyID == 27 ||
			play.PenaltyID == 30 ||
			play.PenaltyID == 31 ||
			play.PenaltyID == 35 {
			firstSegment = ""
			secondSegment = ""
		}
	}

	if ijID > 0 {
		thirdSegment += "INJURY: "
		injLabel := getPlayerLabel(participantMap[ijID])
		injuryType := util.GetInjuryByEnum(play.InjuryType)
		injLength := util.GetInjuryLength(int(play.InjuryDuration))
		injSev := util.GetInjurySeverity(int(play.InjurySeverity))
		thirdSegment += injLabel + " has a " + injSev + " " + injuryType + " and will be out for " + injLength + "."
	}

	return []string{firstSegment + secondSegment + thirdSegment}
}

func generateStreamString(play structs.PlayByPlay, playType, playName, poa string, participantMap map[uint]structs.GameResultsPlayer, twoPtCheck bool) []string {
	qbID := play.QBPlayerID
	bcID := play.BallCarrierID
	t1ID := play.Tackler1ID
	t2ID := play.Tackler2ID
	turnID := play.TurnoverPlayerID
	ijID := play.InjuredPlayerID
	pnID := play.PenaltyPlayerID
	yardsSTR := strconv.Itoa(int(play.ResultYards))
	firstSegment := ""
	secondSegment := ""
	thirdSegment := ""
	list := []string{}

	if play.IsSacked || play.IsTouchdown || play.IsINT || play.IsFumble || play.IsSafety ||
		(play.ResultYards > 39 && playType != "Kickoff" && playType != "Punt" && playType != "XP" && playType != "FG") {
		firstSegment = "REDZONE ALERT: "
	}

	// First Segment
	if playType == "Pass" {
		qbLabel := getPlayerLabel(participantMap[qbID])
		recLabel := ""
		if bcID > 0 {
			recLabel = getPlayerLabel(participantMap[bcID])
		}
		turnoverLabel := ""
		if turnID > 0 {
			turnoverLabel = getPlayerLabel(participantMap[turnID])
		}
		yards := util.GetYardsString(play.ResultYards)
		firstSegment += qbLabel
		offForm := util.GetOffensiveFormationByEnum(play.OffFormationID)
		passStatement := util.GetPassStatement(int(play.ResultYards), offForm, playName, poa, recLabel, play.IsTouchdown, play.IsOutOfBounds, twoPtCheck, play.IsFumble, play.IsSafety, play.IsScramble, play.IsSacked, play.IsComplete, play.IsINT, turnoverLabel)

		firstSegment += passStatement
		// Scenarios
		if play.IsSacked {
			tackle1Label := getPlayerLabel(participantMap[t1ID])
			if t2ID > 0 {
				tackle2Label := getPlayerLabel(participantMap[t2ID])
				tackle1Label += " and " + tackle2Label
			}
			firstSegment += "Sacked on the play by " + tackle1Label + "for a loss of " + yardsSTR + yards
		} else if play.IsComplete {
			firstSegment += yardsSTR + yards
		}

	} else if playType == "Run" {
		runStatement := util.GetRunVerb(int(play.ResultYards), playName, poa, play.IsTouchdown, play.IsOutOfBounds, twoPtCheck, play.IsFumble, play.IsSafety)
		bcLabel := getPlayerLabel(participantMap[bcID])
		firstSegment = bcLabel + runStatement + yardsSTR
		yards := util.GetYardsString(play.ResultYards)
		firstSegment += yards
	} else if playType == "Kickoff" {
		// Need assistance for kickign player ID and returner ID
		kickerLabel := getPlayerLabel(participantMap[qbID])
		recLabel := getPlayerLabel(participantMap[bcID])
		distanceStr := util.GetYardsString(play.KickDistance)
		verb := util.GetKickoffVerb(1)
		firstSegment = kickerLabel + verb + strconv.Itoa(int(play.KickDistance)) + distanceStr
		outside := 0
		if play.KickDistance > 64 {
			if play.KickDistance == 65 {
				verb := util.GetKickoffVerb(2)
				firstSegment += verb + recLabel
			} else {
				verb := util.GetKickoffVerb(3)
				firstSegment += verb + recLabel
			}
		} else {
			outside = 65 - int(play.KickDistance)
			firstSegment += " Fielded at the " + strconv.Itoa(int(outside)) + " yardline by " + recLabel + ". "
		}
		if play.IsTouchback {
			firstSegment += util.GetTouchbackStatement()
		} else {
			netReturnYards := int(play.ResultYards) - outside
			resultYdsStr := strconv.Itoa(int(netReturnYards))
			resultYards := util.GetYardsString(int8(netReturnYards))
			verb := util.GetReturnVerb(netReturnYards, play.IsTouchdown, play.IsOutOfBounds)
			firstSegment += recLabel + verb + resultYdsStr + resultYards
		}
	} else if playType == "Punt" {
		// Need assistance for punting player ID and returner ID
		// Need assistance for kickign player ID and returner ID
		kickerLabel := getPlayerLabel(participantMap[qbID])
		recLabel := getPlayerLabel(participantMap[bcID])
		distanceStr := util.GetYardsString(play.KickDistance)
		verb := util.GetPuntVerb()
		firstSegment = kickerLabel + verb + strconv.Itoa(int(play.KickDistance)) + distanceStr
		los := play.LineOfScrimmage
		if los > 50 {
			los = 100 - los
		}
		outside := los + play.KickDistance
		// Line of Scrimmage + kick distance = ball spot  // Result yards - ball spot = actual yards ran // Ball spot - yards ran = next line of scrimmage
		netReturnYards := outside - play.ResultYards
		resultYards := util.GetYardsString(netReturnYards)
		resultYdsStr := strconv.Itoa(int(netReturnYards))
		if !play.IsTouchback {
			resultYdsStr = strconv.Itoa(int(netReturnYards))
		}
		if play.IsBlocked {
			blockerLabel := getPlayerLabel(participantMap[turnID])
			verb := util.GetBlockedStatement(false)
			firstSegment += verb + blockerLabel
		} else if play.IsFairCatch {
			fc := util.GetFairCatchStatement()
			firstSegment += recLabel + fc
		} else if play.IsTouchback {
			tb := util.GetTouchbackStatement()
			firstSegment += tb
		} else {
			verb := util.GetReturnVerb(int(netReturnYards), play.IsTouchdown, play.IsOutOfBounds)
			firstSegment += recLabel + verb + resultYdsStr + resultYards
		}
	} else if playType == "XP" {
		// Need assistance for kicking player ID and outcome
		kickerLabel := getPlayerLabel(participantMap[qbID])
		startingStatement := util.GetFGStartingStatement(false)
		firstSegment = kickerLabel + "'s " + startingStatement
		if play.IsBlocked {
			blockerLabel := getPlayerLabel(participantMap[turnID])
			verb := util.GetBlockedStatement(true)
			firstSegment += verb + blockerLabel + ". No good. "
		} else if play.IsGood {
			firstSegment += " the kick is good. "
		} else {
			firstSegment += "no good. "
		}
	} else if playType == "FG" {
		kickerLabel := getPlayerLabel(participantMap[qbID])
		kickDistance := strconv.Itoa(int(play.KickDistance))
		startingStatement := util.GetFGStartingStatement(true)
		firstSegment = kickerLabel + "'s " + kickDistance + " " + startingStatement
		if play.IsBlocked {
			blockerLabel := getPlayerLabel(participantMap[turnID])
			verb := util.GetBlockedStatement(true)
			firstSegment += verb + blockerLabel + ". No good. "
		} else {
			endStatement := util.GetFGEndStatement(play.IsGood, play.IsLeft, play.IsOffUpright, play.IsRight)
			firstSegment += endStatement
		}
	}

	// Second Segment - Tackles and OOB
	if !play.IsSacked && t1ID > 0 {
		tackle1Label := getPlayerLabel(participantMap[t1ID])
		firstSegment += "Tackled by " + tackle1Label
		if t2ID > 0 {
			tackle2Label := getPlayerLabel(participantMap[t2ID])
			firstSegment += " and " + tackle2Label
		}
		firstSegment += ". "
	}

	if play.IsFumble {
		turnOverLabel := getPlayerLabel(participantMap[turnID])
		secondSegment += "Fumble recovered by " + turnOverLabel + "."
	}
	list = append(list, firstSegment)
	// Second Item -- Penalties and Injuries
	if play.PenaltyID > 0 {
		penalty := util.GetPenaltyByEnum(play.PenaltyID)
		secondSegment = "PENALTY: " + penalty + ". "
		offendingTeam := "Offense. "
		if !play.OnOffense {
			offendingTeam = "Defense. "
		}
		secondSegment += offendingTeam
		if pnID > 0 {
			player := participantMap[pnID]
			penaltyLabel := getPlayerLabel(player)
			secondSegment += "Player: " + penaltyLabel + ". "
		}
		penaltyYards := strconv.Itoa(int(play.PenaltyYards))
		yards := util.GetYardsString(play.PenaltyYards)
		secondSegment += penaltyYards + yards
		if play.PenaltyID == 4 ||
			play.PenaltyID == 5 ||
			play.PenaltyID == 7 ||
			play.PenaltyID == 18 ||
			play.PenaltyID == 20 ||
			play.PenaltyID == 21 ||
			play.PenaltyID == 27 ||
			play.PenaltyID == 30 ||
			play.PenaltyID == 31 ||
			play.PenaltyID == 35 {
			list = []string{}
		}
		list = append(list, secondSegment)
	}

	if ijID > 0 {
		thirdSegment += "INJURY: "
		injLabel := getPlayerLabel(participantMap[ijID])
		injuryType := util.GetInjuryByEnum(play.InjuryType)
		injLength := util.GetInjuryLength(int(play.InjuryDuration))
		injSev := util.GetInjurySeverity(int(play.InjurySeverity))
		thirdSegment += injLabel + " has a " + injSev + " " + injuryType + " and will be out for " + injLength + "."
		list = append(list, thirdSegment)
	}
	return list
}

func getPlayerLabel(player structs.GameResultsPlayer) string {
	if player.ID == 0 {
		return ""
	}
	return player.TeamAbbr + " " + player.Position + " " + player.FirstName + " " + player.LastName
}

func getGameParticipantMap(homePlayers, awayPlayers []structs.GameResultsPlayer) map[uint]structs.GameResultsPlayer {
	playerMap := make(map[uint]structs.GameResultsPlayer)

	for _, p := range homePlayers {
		playerMap[p.ID] = p
	}

	for _, p := range awayPlayers {
		playerMap[p.ID] = p
	}
	return playerMap
}
