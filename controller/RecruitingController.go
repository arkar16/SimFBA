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

// GetRecruitingProfileByTeamID -- for Overall Dashboard
func GetRecruitingProfileByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	if len(teamID) == 0 {
		panic("User did not provide teamID")
	}

	var dashboardResponse structs.DashboardTeamProfileResponse

	recruitingProfile := managers.GetRecruitingProfileByTeamID(teamID)

	dashboardResponse.SetTeamProfile(recruitingProfile)

	// Get Team Needs
	teamNeeds := managers.GetRecruitingNeeds(teamID)

	dashboardResponse.SetTeamNeedsMap(teamNeeds)

	json.NewEncoder(w).Encode(dashboardResponse)
}

// GetOnlyRecruitingProfileByTeamID
func GetOnlyRecruitingProfileByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	if len(teamID) == 0 {
		panic("User did not provide teamID")
	}

	recruitingProfile := managers.GetOnlyRecruitingProfileByTeamID(teamID)

	json.NewEncoder(w).Encode(recruitingProfile)
}

// CreateRecruitPlayerProfile
func CreateRecruitPlayerProfile(w http.ResponseWriter, r *http.Request) {

	var recruitPointsDto structs.CreateRecruitProfileDto
	err := json.NewDecoder(r.Body).Decode(&recruitPointsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruitingProfile := managers.AddRecruitToBoard(recruitPointsDto)

	json.NewEncoder(w).Encode(recruitingProfile)

	fmt.Fprintf(w, "New Recruiting Profile Created")
}

// CreateRecruitingPointsProfileForRecruit
func CreateRecruitingPointsProfileForRecruit(w http.ResponseWriter, r *http.Request) {

	var recruitPointsDto structs.CreateRecruitProfileDto
	err := json.NewDecoder(r.Body).Decode(&recruitPointsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruitingProfile := managers.CreateRecruitingProfileForRecruit(recruitPointsDto)

	json.NewEncoder(w).Encode(recruitingProfile)

	fmt.Fprintf(w, "New Recruiting Profile Created")
}

// AllocateRecruitingPointsForRecruit
func AllocateRecruitingPointsForRecruit(w http.ResponseWriter, r *http.Request) {
	var updateRecruitPointsDto structs.UpdateRecruitPointsDto
	err := json.NewDecoder(r.Body).Decode(&updateRecruitPointsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	managers.AllocateRecruitPointsForRecruit(updateRecruitPointsDto)

	fmt.Printf("Updated Recruiting Points Profile")
}

// SendScholarshipToRecruit
func SendScholarshipToRecruit(w http.ResponseWriter, r *http.Request) {
	var updateRecruitPointsDto structs.UpdateRecruitPointsDto
	err := json.NewDecoder(r.Body).Decode(&updateRecruitPointsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruitingPointsProfile, recruitingProfile := managers.SendScholarshipToRecruit(updateRecruitPointsDto)
	fmt.Printf("\nScholarship allocated to player " + strconv.Itoa(recruitingPointsProfile.RecruitID) + ". Record saved")
	fmt.Printf("\nProfile: " + strconv.Itoa(recruitingProfile.TeamID) + " Saved")
}

// RevokeScholarshipToRecruit
func RevokeScholarshipFromRecruit(w http.ResponseWriter, r *http.Request) {
	var updateRecruitPointsDto structs.UpdateRecruitPointsDto
	err := json.NewDecoder(r.Body).Decode(&updateRecruitPointsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruitingPointsProfile, recruitingProfile := managers.RevokeScholarshipFromRecruit(updateRecruitPointsDto)

	fmt.Printf("\nScholarship revoked from player " + strconv.Itoa(recruitingPointsProfile.RecruitID) + ". Record saved")
	fmt.Printf("\nProfile: " + strconv.Itoa(recruitingProfile.TeamID) + " Saved")
}

// RemoveRecruitFromBoard
func RemoveRecruitFromBoard(w http.ResponseWriter, r *http.Request) {
	var updateRecruitPointsDto structs.UpdateRecruitPointsDto
	err := json.NewDecoder(r.Body).Decode(&updateRecruitPointsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruitingPointsProfile := managers.RemoveRecruitFromBoard(updateRecruitPointsDto)

	fmt.Printf("\nPlayer " + strconv.Itoa(recruitingPointsProfile.RecruitID) + " removed from board.")
}

// SaveRecruitingBoard
func SaveRecruitingBoard(w http.ResponseWriter, r *http.Request) {
	var updateRecruitingBoardDto structs.UpdateRecruitingBoardDTO
	err := json.NewDecoder(r.Body).Decode(&updateRecruitingBoardDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recruitingProfile := managers.UpdateRecruitingProfile(updateRecruitingBoardDto)

	fmt.Println("Updated Recruiting Profile " + strconv.Itoa(recruitingProfile.TeamID) + " and all associated players")
	w.WriteHeader(http.StatusOK)
}
