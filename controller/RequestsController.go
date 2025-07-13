package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
)

func GetFBARequests(w http.ResponseWriter, r *http.Request) {
	requests := managers.GetAllFBARequests()
	json.NewEncoder(w).Encode(requests)
}

func GetTeamRequests(w http.ResponseWriter, r *http.Request) {
	requests := managers.GetAllTeamRequests()

	json.NewEncoder(w).Encode(requests)
}

func GetNFLTeamRequests(w http.ResponseWriter, r *http.Request) {
	requests := managers.GetAllNFLTeamRequests()

	json.NewEncoder(w).Encode(requests)
}

func CreateTeamRequest(w http.ResponseWriter, r *http.Request) {
	var request structs.TeamRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.CreateTeamRequest(request)

	fmt.Fprintf(w, "Request Successfully Created")
}

func ApproveTeamRequest(w http.ResponseWriter, r *http.Request) {
	var request structs.TeamRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.ID == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.ApproveTeamRequest(request)

	fmt.Fprintf(w, "Request: %+v", request)
}

func RejectTeamRequest(w http.ResponseWriter, r *http.Request) {
	var request structs.TeamRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.RejectTeamRequest(request)
}

func RemoveUserFromTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}

	managers.RemoveUserFromTeam(teamID)
}

func ViewCFBTeamUponRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}

	team := managers.GetCFBTeamForAvailableTeamsPage(teamID)

	json.NewEncoder(w).Encode(team)
}

func ViewNFLTeamUponRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}

	team := managers.GetNFLTeamForAvailableTeamsPage(teamID)

	json.NewEncoder(w).Encode(team)
}

func CreateNFLTeamRequest(w http.ResponseWriter, r *http.Request) {
	var request structs.NFLRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.CreateNFLTeamRequest(request)

	fmt.Fprintf(w, "Request Successfully Created")
}

func ApproveNFLTeamRequest(w http.ResponseWriter, r *http.Request) {
	var request structs.NFLRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.ID == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.ApproveNFLTeamRequest(request)

	fmt.Fprintf(w, "Request: %+v", request)
}

func RejectNFLTeamRequest(w http.ResponseWriter, r *http.Request) {
	var request structs.NFLRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.RejectNFLTeamRequest(request)
}

func RemoveNFLUserFromNFLTeam(w http.ResponseWriter, r *http.Request) {
	var request structs.NFLRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.RemoveUserFromNFLTeam(request)

	// json.NewEncoder(w).Encode(team)
}
