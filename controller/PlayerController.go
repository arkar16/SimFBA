package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
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
	// Update DTO for College Player

	// validate info from DTO

	// Send DTO to Manager Class

	// Save to DB

	// Return updated player

	// Send to interface?

	fmt.Println(w, "College Player successfully updated.")
}

// ToggleRedshirtStatusForPlayer
func ToggleRedshirtStatusForPlayer(w http.ResponseWriter, r *http.Request) {
	// Use DTO for redshirts?

	// validate info from DTO

	// Send DTO to Manager Class

	// Update redshirting status for players

	// Save to DB

	// Return updated info

	// Send to interface?

	fmt.Println(w, "College Player successfully redshirted.")
}
