package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

// GetGameplanByTeamID
func GetTeamGameplanByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide a teamID")
	}

	gamePlan := managers.GetGameplanByTeamID(teamID)

	json.NewEncoder(w).Encode(gamePlan)
}

// GetDepthChartByTeamID
func GetTeamDepthchartByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide a teamID")
	}

	depthchart := managers.GetGameplanByTeamID(teamID)

	json.NewEncoder(w).Encode(depthchart)
}

// GetDepthChartPositionsByDepthChartID
func GetDepthChartPositionsByDepthChartID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	depthChartID := vars["depthChartID"]
	if len(depthChartID) == 0 {
		panic("User did not provide a depthChartID")
	}

	positionPlayers := managers.GetDepthChartPositionPlayersByDepthchartID(depthChartID)

	json.NewEncoder(w).Encode(positionPlayers)
}

// UpdateGameplan

// UpdateDepthChart
