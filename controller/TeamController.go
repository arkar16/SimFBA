package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/models"
	"github.com/gorilla/mux"
)

// GetAllCollegeTeams
func GetAllCollegeTeams(w http.ResponseWriter, r *http.Request) {
	collegeTeams := managers.GetAllCollegeTeams()

	json.NewEncoder(w).Encode(collegeTeams)
}

// GetAllNFLTeams

// GetAllActiveCollegeTeams
func GetAllActiveCollegeTeams(w http.ResponseWriter, r *http.Request) {
	collegeTeams := managers.GetAllCoachedCollegeTeams()

	json.NewEncoder(w).Encode(collegeTeams)
}

// GetAllAvailableCollegeTeams
func GetAllAvailableCollegeTeams(w http.ResponseWriter, r *http.Request) {
	collegeTeams := managers.GetAllAvailableCollegeTeams()

	json.NewEncoder(w).Encode(collegeTeams)
}

// GetAllAvailableNFLTeams

// GetAllCoachedNFLTeams

// GetTeamByTeamID
func GetTeamByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}
	team := managers.GetTeamByTeamID(teamID)
	json.NewEncoder(w).Encode(team)
}

func GetTeamByTeamIDForDiscord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}
	team := managers.GetTeamByTeamIDForDiscord(teamID)
	json.NewEncoder(w).Encode(team)
}

// GetTeamsByConferenceID
func GetTeamsByConferenceID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conferenceID := vars["conferenceID"]
	if len(conferenceID) == 0 {
		panic("User did not provide conferenceID")
	}
	team := managers.GetTeamByTeamID(conferenceID)
	json.NewEncoder(w).Encode(team)
}

// GetTeamsByDivisionID
func GetTeamsByDivisionID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	divisionID := vars["divisionID"]
	if len(divisionID) == 0 {
		panic("User did not provide divisionID")
	}
	team := managers.GetTeamByTeamID(divisionID)
	json.NewEncoder(w).Encode(team)
}

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

	json.NewEncoder(w).Encode(responseModel)
}
