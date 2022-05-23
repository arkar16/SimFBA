package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

// GetCollegeStandingsByConferenceIDAndSeasonID
func GetCollegeStandingsByConferenceIDAndSeasonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conferenceID := vars["conferenceID"]
	seasonID := vars["seasonID"]
	if len(conferenceID) == 0 || len(seasonID) == 0 {
		panic("User did not provide enough information")
	}
	standings := managers.GetStandingsByConferenceIDAndSeasonID(conferenceID, seasonID)
	json.NewEncoder(w).Encode(standings)
}

// GetCollegeStandingsByDivisionIDAndSeasonID

// GetHistoricalStandingsByTeamID

// GetHistoricalStandingsByCoach
