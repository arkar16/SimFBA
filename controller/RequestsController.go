package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
)

func GetTeamRequests(w http.ResponseWriter, r *http.Request) {
	requests := managers.GetAllTeamRequests()

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

	fmt.Fprintf(w, "Request: %+v", request)
}
