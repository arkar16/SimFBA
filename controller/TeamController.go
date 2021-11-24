package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
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
