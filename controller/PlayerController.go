package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
)

// AllPlayers - Get All Players Record references in table
func AllPlayers(w http.ResponseWriter, r *http.Request) {
	var players = managers.GetAllPlayers()

	json.NewEncoder(w).Encode(players)
}

// AllCollegePlayers

// AllCollegePlayersByTeamID
func AllCollegePlayersByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["teamID"]

	if len(teamId) == 0 {
		panic("User did not provide TeamID")
	}

	players := managers.GetAllCollegePlayersByTeamId(teamId)

	json.NewEncoder(w).Encode(players)
}

// AllCollegePlayersByTeamIDWithoutRedshirts
func AllCollegePlayersByTeamIDWithoutRedshirts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["teamID"]

	if len(teamId) == 0 {
		panic("User did not provide TeamID")
	}

	players := managers.GetAllCollegePlayersByTeamIdWithoutRedshirts(teamId)

	json.NewEncoder(w).Encode(players)
}

// GetCollegePlayerByID
func GetCollegePlayerByCollegePlayerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collegePlayerID := vars["playerID"]

	if len(collegePlayerID) == 0 {
		panic("User did not provide TeamID")
	}

	players := managers.GetCollegePlayerByCollegePlayerId(collegePlayerID)

	json.NewEncoder(w).Encode(players)
}

// UpdateCollegePlayer
func UpdateCollegePlayer(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "College Player successfully updated.")
}

// ToggleRedshirtStatusForPlayer
func ToggleRedshirtStatusForPlayer(w http.ResponseWriter, r *http.Request) {
	// Update DTO for College Player
	var redshirtDTO structs.RedshirtDTO
	err := json.NewDecoder(r.Body).Decode(&redshirtDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.SetRedshirtStatusForPlayer(strconv.Itoa(redshirtDTO.PlayerID))

	fmt.Println(w, "College Player successfully redshirted.")
}

func ExportRosterToCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")

	vars := mux.Vars(r)
	teamId := vars["teamID"]

	if len(teamId) == 0 {
		panic("User did not provide TeamID")
	}

	managers.ExportTeamToCSV(teamId, w)

	// ?
}

func ExportAllRostersToCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")

	CollegeTeams := managers.GetAllCollegeTeams()

	for _, team := range CollegeTeams {
		id := strconv.FormatUint(uint64(team.ID), 10)

		managers.ExportTeamToCSV(id, w)
	}

	// ?
}
