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
	homeTeamAbbr := vars["HomeTeamAbbr"]
	awayTeamAbbr := vars["AwayTeamAbbr"]

	var responseModel models.SimGameDataResponse

	homeTeam := managers.GetTeamByTeamAbbr(homeTeamAbbr)
	awayTeam := managers.GetTeamByTeamAbbr(awayTeamAbbr)

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

	ts := managers.GetTimestamp()

	game := managers.GetCollegeGameByAbbreviationsWeekAndSeasonID(homeTeamAbbr, strconv.Itoa(ts.CollegeWeekID), strconv.Itoa(ts.CollegeSeasonID))

	stadiumID := strconv.Itoa(int(game.StadiumID))

	stadium := managers.GetStadiumByStadiumID(stadiumID)

	responseModel.AssignWeather(game.GameTemp, game.Cloud, game.Precip, game.WindCategory, game.WindSpeed)
	responseModel.AssignStadium(stadium)

	json.NewEncoder(w).Encode(responseModel)
}
