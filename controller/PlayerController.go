package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

// AllPlayers - Get All Players Record references in table
func AllPlayers(w http.ResponseWriter, r *http.Request) {
	var players = managers.GetAllPlayers()

	json.NewEncoder(w).Encode(players)
}

func AllRecruits(w http.ResponseWriter, r *http.Request) {
	var recruits = managers.GetAllRecruits()

	json.NewEncoder(w).Encode(recruits)
}
