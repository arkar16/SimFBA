package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

func GenerateCapsheets(w http.ResponseWriter, r *http.Request) {
	managers.AllocateCapsheets()
	fmt.Println(w, "Congrats, you generated the Capsheets!")
}

// GetTeamByTeamID
func GetNFLCapsheetByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}
	team := managers.GetCapsheetByTeamID(teamID)
	json.NewEncoder(w).Encode(team)
}
