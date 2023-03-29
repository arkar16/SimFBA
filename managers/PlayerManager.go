package managers

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
	"gorm.io/gorm"
)

// GetAllPlayers - Returns all player reference records
func GetAllPlayers() []structs.Player {
	db := dbprovider.GetInstance().GetDB()

	var players []structs.Player

	db.Find(&players)

	return players
}

func GetAllCollegePlayers() []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayers []structs.CollegePlayer

	db.Find(&CollegePlayers)

	return CollegePlayers
}

func GetAllNFLPlayers() []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var nflPlayers []structs.NFLPlayer

	db.Find(&nflPlayers)

	return nflPlayers
}

func GetAllUnsignedPlayers() []structs.UnsignedPlayer {
	db := dbprovider.GetInstance().GetDB()

	var unsignedPlayers []structs.UnsignedPlayer

	db.Find(&unsignedPlayers)

	return unsignedPlayers
}

func GetAllCollegePlayersByTeamId(TeamID string) []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayers []structs.CollegePlayer

	err := db.Order("overall desc").Where("team_id = ?", TeamID).Where("has_graduated = ?", false).Find(&CollegePlayers).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return CollegePlayers
}

func GetAllCollegePlayersByTeamIdWithoutRedshirts(TeamID string) []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayers []structs.CollegePlayer

	db.Where("team_id = ?", TeamID).Where("is_redshirting = ?", false).Where("has_graduated = ?", false).Find(&CollegePlayers)

	return CollegePlayers
}

func GetCollegePlayerByCollegePlayerId(CollegePlayerId string) structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayer structs.CollegePlayer

	db.Where("id = ?", CollegePlayerId).Find(&CollegePlayer)

	return CollegePlayer
}

func GetCollegePlayerByNameAndTeam(firstName string, lastName string, teamID string) models.CollegePlayerCSV {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayer structs.CollegePlayer

	db.Where("first_name = ? and last_name = ? and team_id = ?", firstName, lastName, teamID).Find(&CollegePlayer)

	collegePlayerResponse := models.MapPlayerToCSVModel(CollegePlayer)

	return collegePlayerResponse
}

func GetCollegePlayerByNameTeamAndWeek(firstName string, lastName string, teamID string, week string) models.CollegePlayerCSV {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	collegeWeek := GetCollegeWeek(week, ts)

	if collegeWeek.ID == uint(ts.CollegeWeekID) {
		return models.CollegePlayerCSV{}
	} else {
		var CollegePlayer structs.CollegePlayer

		db.Preload("Stats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ? AND week_id = ?", collegeWeek.SeasonID, collegeWeek.ID)
		}).Where("first_name = ? and last_name = ? and team_id = ?", firstName, lastName, teamID).Find(&CollegePlayer)

		collegePlayerResponse := models.MapPlayerForStats(CollegePlayer)

		return collegePlayerResponse
	}

}

func GetSeasonalCollegePlayerByNameTeam(firstName string, lastName string, teamID string) models.CollegePlayerResponse {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	var CollegePlayer structs.CollegePlayer

	db.Preload("Stats", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ? AND week_id < ?", strconv.Itoa(ts.CollegeSeasonID), strconv.Itoa(ts.CollegeWeekID))
	}).Where("first_name = ? and last_name = ? and team_id = ?", firstName, lastName, teamID).Find(&CollegePlayer)

	collegePlayerResponse := models.CollegePlayerResponse{
		ID:          int(CollegePlayer.ID),
		BasePlayer:  CollegePlayer.BasePlayer,
		TeamID:      CollegePlayer.TeamID,
		TeamAbbr:    CollegePlayer.TeamAbbr,
		City:        CollegePlayer.City,
		State:       CollegePlayer.State,
		Year:        CollegePlayer.Year,
		IsRedshirt:  CollegePlayer.IsRedshirt,
		PlayerStats: CollegePlayer.Stats,
	}

	collegePlayerResponse.MapSeasonalStats()

	return collegePlayerResponse
}

func UpdateCollegePlayer(cp structs.CollegePlayer) {
	db := dbprovider.GetInstance().GetDB()
	err := db.Save(&cp).Error
	if err != nil {
		log.Fatal(err)
	}
}

func SetRedshirtStatusForPlayer(playerId string) structs.CollegePlayer {
	player := GetCollegePlayerByCollegePlayerId(playerId)

	player.SetRedshirtingStatus()

	UpdateCollegePlayer(player)

	return player
}

func GetAllNFLDraftees() []structs.NFLDraftee {
	db := dbprovider.GetInstance().GetDB()

	var NFLDraftees []structs.NFLDraftee

	db.Order("overall desc").Find(&NFLDraftees)

	return NFLDraftees
}

func GetAllCollegePlayersWithStatsBySeasonID(cMap map[int]int, cNMap map[int]string, seasonID string) []models.CollegePlayerResponse {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	seasonIDVal := util.ConvertStringToInt(seasonID)

	var collegePlayers []structs.CollegePlayer

	// var distinctCollegeStats []structs.CollegePlayerStats
	var distinctCollegeStats []structs.CollegePlayerSeasonStats

	db.Distinct("college_player_id").Where("snaps > 0 AND season_id = ?", seasonID).Find(&distinctCollegeStats)

	distinctCollegePlayerIDs := util.GetCollegePlayerIDsBySeasonStats(distinctCollegeStats)

	// db.Preload("SeasonStats", func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("season_id = ? and week_id < ? and snaps > 0", strconv.Itoa(ts.CollegeSeasonID), strconv.Itoa(ts.CollegeWeekID))
	// }).Where("id in ?", distinctCollegePlayerIDs).Find(&collegePlayers)

	db.Preload("SeasonStats", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ?", seasonID)
	}).Where("id in ?", distinctCollegePlayerIDs).Find(&collegePlayers)

	var cpResponse []models.CollegePlayerResponse

	for _, player := range collegePlayers {
		cp := models.CollegePlayerResponse{
			ID:           int(player.ID),
			BasePlayer:   player.BasePlayer,
			ConferenceID: cMap[player.TeamID],
			Conference:   cNMap[player.TeamID],
			TeamID:       player.TeamID,
			TeamAbbr:     player.TeamAbbr,
			City:         player.City,
			State:        player.State,
			Year:         player.Year,
			IsRedshirt:   player.IsRedshirt,
			SeasonStats:  player.SeasonStats,
		}

		// cp.MapSeasonalStats()

		cpResponse = append(cpResponse, cp)
	}

	// If viewing a past season, get all past season players too
	if seasonIDVal < ts.CollegeSeasonID {
		var historicCollegePlayers []structs.HistoricCollegePlayer
		db.Preload("SeasonStats", func(db *gorm.DB) *gorm.DB {
			return db.Where("season_id = ?", seasonID)
		}).Where("id in ?", distinctCollegePlayerIDs).Find(&historicCollegePlayers)

		for _, player := range historicCollegePlayers {
			cp := models.CollegePlayerResponse{
				ID:           int(player.ID),
				BasePlayer:   player.BasePlayer,
				ConferenceID: cMap[player.TeamID],
				Conference:   cNMap[player.TeamID],
				TeamID:       player.TeamID,
				TeamAbbr:     player.TeamAbbr,
				City:         player.City,
				State:        player.State,
				Year:         player.Year,
				IsRedshirt:   player.IsRedshirt,
				SeasonStats:  player.SeasonStats,
			}

			cpResponse = append(cpResponse, cp)
		}
	}

	return cpResponse
}

func GetAllCollegePlayersWithStatsByTeamID(TeamID string, SeasonID string) []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var collegePlayers []structs.CollegePlayer

	db.Preload("Stats", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ? and team_id = ? and snaps > 0", SeasonID, TeamID)
	}).Where("team_id = ?", TeamID).Find(&collegePlayers)

	return collegePlayers
}

func GetHeismanList() []models.HeismanWatchModel {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	var collegePlayers []structs.CollegePlayer

	var heismanCandidates []models.HeismanWatchModel

	teamWithStandings := GetAllCollegeTeamsWithCurrentYearStandings()

	var teamWeight = make(map[string]float64)

	var homeTeamMapper = make(map[int]string)
	var teamGameMapper = make(map[int][]structs.CollegeGame)

	for _, team := range teamWithStandings {
		homeTeamMapper[int(team.ID)] = team.TeamAbbr

		games := GetCollegeGamesByTeamIdAndSeasonId(strconv.Itoa(int(team.ID)), strconv.Itoa(ts.CollegeSeasonID))

		if len(games) == 0 || len(team.TeamStandings) == 0 {
			continue
		}

		teamGameMapper[int(team.ID)] = games

		currentYearStandings := team.TeamStandings[0]

		var weight float64 = 0 // 1
		if currentYearStandings.TotalLosses+currentYearStandings.TotalWins > 0 {
			// newWeight := (float64(currentYearStandings.TotalWins) / 12) + 1

			// if newWeight > weight {
			// 	weight = newWeight
			// }
			weight = float64(currentYearStandings.TotalWins) / 100
		}

		teamWeight[team.TeamAbbr] = weight
	}

	// db.Preload("Stats", func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("snaps > 0 and season_id = ? and week_id < ?", strconv.Itoa(ts.CollegeSeasonID), strconv.Itoa(ts.CollegeWeekID))
	// }).Where("Stats.snaps > 0").Find(&collegePlayers)

	var distinctCollegeStats []structs.CollegePlayerStats

	db.Distinct("college_player_id").Where("snaps > 0").Find(&distinctCollegeStats)

	distinctCollegePlayerIDs := util.GetCollegePlayerIDs(distinctCollegeStats)

	db.Preload("Stats", func(db *gorm.DB) *gorm.DB {
		return db.Where("snaps > 0 and season_id = ? and week_id < ?", strconv.Itoa(ts.CollegeSeasonID), strconv.Itoa(ts.CollegeWeekID))
	}).Where("id IN ?", distinctCollegePlayerIDs).Find(&collegePlayers)

	for _, cp := range collegePlayers {
		if len(cp.Stats) == 0 {
			continue
		}

		score := util.GetHeismanScore(cp, teamWeight, homeTeamMapper, teamGameMapper[cp.TeamID])

		h := models.HeismanWatchModel{
			FirstName: cp.FirstName,
			LastName:  cp.LastName,
			Position:  cp.Position,
			Archetype: cp.Archetype,
			School:    cp.TeamAbbr,
			Score:     score,
			Games:     len(cp.Stats),
		}

		heismanCandidates = append(heismanCandidates, h)

	}

	sort.Sort(models.ByScore(heismanCandidates))

	return heismanCandidates
}

func GetGlobalPlayerRecord(playerID string) structs.Player {
	db := dbprovider.GetInstance().GetDB()

	var player structs.Player

	db.Where("id = ?", playerID).Find(&player)

	return player
}

func GetOnlyNFLPlayerRecord(playerID string) structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var player structs.NFLPlayer

	db.Where("id = ?", playerID).Find(&player)

	return player
}

func GetNFLPlayerRecord(playerID string) structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var player structs.NFLPlayer

	db.Preload("Contract").Where("id = ?", playerID).Find(&player)

	return player
}

func GetNFLPlayersForDCPage(TeamID string) []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var players []structs.NFLPlayer

	db.Preload("Contract").Where("team_id = ? AND is_practice_squad = ?", TeamID, false).Find(&players)

	return players
}

func GetTradableNFLPlayersByTeamID(TeamID string) []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var players []structs.NFLPlayer

	db.Preload("Contract").Where("team_id = ? AND is_on_trade_block = ?", TeamID, true).Find(&players)

	return players
}

func GetNFLPlayersWithContractsByTeamID(TeamID string) []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var players []structs.NFLPlayer

	db.Preload("Contract").Where("team_id = ?", TeamID).Find(&players)

	return players
}

func CutNFLPlayer(playerId string) {
	db := dbprovider.GetInstance().GetDB()

	player := GetOnlyNFLPlayerRecord(playerId)
	contract := GetContractByPlayerID(playerId)
	capsheet := GetCapsheetByTeamID(strconv.Itoa(int(player.TeamID)))
	ts := GetTimestamp()

	if player.Experience < 4 && !ts.IsNFLOffSeason {
		player.WaivePlayer()
	} else {
		player.ToggleIsFreeAgent()
		contract.DeactivateContract()

		// Get Minimum Value

	}

	capsheet.CutPlayerFromCapsheet(contract)
	db.Save(&contract)
	db.Save(&player)
	db.Save(&capsheet)
}

func getMinimumValue(ovr int, pos string) {

}

func GetNFLRosterForSimulation(TeamID string) []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	var players []structs.NFLPlayer

	db.Where("team_id = ? AND is_practice_squad = ?", TeamID, false).Find(&players)

	return players
}
