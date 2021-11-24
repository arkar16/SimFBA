package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

// GetCollegeGamesByWeekId
func GetCollegeGamesByWeekId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekID := vars["weekID"]

	if len(weekID) == 0 {
		panic("User did not provide weekID")
	}

	collegeGames := managers.GetCollegeGamesByWeekId(weekID)

	json.NewEncoder(w).Encode(collegeGames)
}

// GetCollegeGamesByTimeslotAndWeekID
func GetCollegeGamesByTimeslotWeekId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekID := vars["weekID"]
	timeSlot := vars["timeSlot"]

	if len(weekID) == 0 {
		panic("User did not provide weekID")
	}

	if len(timeSlot) == 0 {
		panic("No time slot selected!")
	}

	collegeGames := managers.GetCollegeGamesByTimeslotAndWeekId(weekID, timeSlot)

	json.NewEncoder(w).Encode(collegeGames)
}

// GetCollegeGamesByTeamIDAndSeasonID
func GetCollegeGamesByTeamIDAndSeasonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	seasonID := vars["seasonID"]

	if len(teamID) == 0 {
		panic("User did not provide weekID")
	}

	if len(seasonID) == 0 {
		panic("No time slot selected!")
	}

	collegeGames := managers.GetCollegeGamesByTeamIdAndSeasonId(teamID, seasonID)

	json.NewEncoder(w).Encode(collegeGames)
}

// GetNFLGamesByWeekID

// GetNFLGamesByTimeslotAndWeekID

// GetNFLGamesByTeamIDAndSeasonID
