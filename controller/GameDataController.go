package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/models"
	"github.com/gorilla/mux"
)

// GetHomeAndAwayTeamData
func GetHomeAndAwayTeamData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]

	game := managers.GetCollegeGameByGameID(gameID)

	var responseModel models.SimGameDataResponse

	homeTeam := managers.GetTeamByTeamAbbr(game.HomeTeam)
	awayTeam := managers.GetTeamByTeamAbbr(game.AwayTeam)

	homeTeamID := strconv.Itoa(int(homeTeam.ID))
	awayTeamID := strconv.Itoa(int(awayTeam.ID))

	var homeTeamResponse models.SimTeamDataResponse
	var homeDCResponse models.SimTeamDepthChartResponse
	var homeDCList []models.SimDepthChartPosResponse

	var awayTeamResponse models.SimTeamDataResponse
	var awayDCResponse models.SimTeamDepthChartResponse
	var awayDCList []models.SimDepthChartPosResponse

	hdc := homeTeam.TeamDepthChart

	for _, dcp := range hdc.DepthChartPlayers {
		var simDCPR models.SimDepthChartPosResponse
		simDCPR.Map(dcp)
		homeDCList = append(homeDCList, simDCPR)
	}

	adc := awayTeam.TeamDepthChart
	for _, dcp := range adc.DepthChartPlayers {
		var simDCPR models.SimDepthChartPosResponse
		simDCPR.Map(dcp)
		awayDCList = append(awayDCList, simDCPR)
	}

	homeDCResponse.Map(hdc, homeDCList)
	awayDCResponse.Map(adc, awayDCList)

	homeTeamResponse.Map(homeTeam, homeDCResponse)
	awayTeamResponse.Map(awayTeam, awayDCResponse)

	homeTeamRoster := managers.GetAllCollegePlayersByTeamIdWithoutRedshirts(homeTeamID)
	awayTeamRoster := managers.GetAllCollegePlayersByTeamIdWithoutRedshirts(awayTeamID)

	responseModel.AssignHomeTeam(homeTeamResponse, homeTeamRoster)
	responseModel.AssignAwayTeam(awayTeamResponse, awayTeamRoster)

	stadiumID := strconv.Itoa(int(game.StadiumID))

	stadium := managers.GetStadiumByStadiumID(stadiumID)

	responseModel.AssignWeather(game.GameTemp, game.Cloud, game.Precip, game.WindCategory, game.WindSpeed)
	responseModel.AssignStadium(stadium)

	json.NewEncoder(w).Encode(responseModel)
}

func GetNFLHomeAndAwayTeamData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameID"]

	game := managers.GetNFLGameByGameID(gameID)

	var responseModel models.NFLSimGameDataResponse

	homeTeam := managers.GetNFLTeamByTeamAbbr(game.HomeTeam)
	awayTeam := managers.GetNFLTeamByTeamAbbr(game.AwayTeam)

	homeTeamID := strconv.Itoa(int(homeTeam.ID))
	awayTeamID := strconv.Itoa(int(awayTeam.ID))

	var homeTeamResponse models.NFLSimTeamDataResponse
	var homeDCResponse models.NFLSimTeamDepthChartResponse
	var homeDCList []models.NFLSimDepthChartPosResponse

	var awayTeamResponse models.NFLSimTeamDataResponse
	var awayDCResponse models.NFLSimTeamDepthChartResponse
	var awayDCList []models.NFLSimDepthChartPosResponse

	hdc := homeTeam.TeamDepthChart

	for _, dcp := range hdc.DepthChartPlayers {
		var simDCPR models.NFLSimDepthChartPosResponse
		simDCPR.Map(dcp)
		homeDCList = append(homeDCList, simDCPR)
	}

	adc := awayTeam.TeamDepthChart
	for _, dcp := range adc.DepthChartPlayers {
		var simDCPR models.NFLSimDepthChartPosResponse
		simDCPR.Map(dcp)
		awayDCList = append(awayDCList, simDCPR)
	}

	homeDCResponse.Map(hdc, homeDCList)
	awayDCResponse.Map(adc, awayDCList)

	homeTeamResponse.Map(homeTeam, homeDCResponse)
	awayTeamResponse.Map(awayTeam, awayDCResponse)

	homeTeamRoster := managers.GetNFLRosterForSimulation(homeTeamID)
	awayTeamRoster := managers.GetNFLRosterForSimulation(awayTeamID)

	responseModel.AssignHomeTeam(homeTeamResponse, homeTeamRoster)
	responseModel.AssignAwayTeam(awayTeamResponse, awayTeamRoster)

	stadiumID := strconv.Itoa(int(game.StadiumID))

	stadium := managers.GetStadiumByStadiumID(stadiumID)

	responseModel.AssignWeather(game.GameTemp, game.Cloud, game.Precip, game.WindCategory, game.WindSpeed)
	responseModel.AssignStadium(stadium)

	json.NewEncoder(w).Encode(responseModel)
}
