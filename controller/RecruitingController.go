package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
)

// GetRecruitingProfileForDashboardByTeamID -- for Overall Dashboard
func GetRecruitingProfileForDashboardByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	if len(teamID) == 0 {
		panic("User did not provide teamID")
	}

	var dashboardResponse models.DashboardTeamProfileResponse

	recruitingProfile := managers.GetRecruitingProfileForDashboardByTeamID(teamID)

	dashboardResponse.SetTeamProfile(recruitingProfile)

	// Get Team Needs
	teamNeeds := managers.GetRecruitingNeeds(teamID)

	dashboardResponse.SetTeamNeedsMap(teamNeeds)

	json.NewEncoder(w).Encode(dashboardResponse)
}

// GetRecruitingProfileByTeamID -- for Overall Dashboard
func GetRecruitingProfileForTeamBoardByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	if len(teamID) == 0 {
		panic("User did not provide teamID")
	}

	var teamBoardResponse models.TeamBoardTeamProfileResponse

	recruitingProfile := managers.GetRecruitingProfileForTeamBoardByTeamID(teamID)

	teamBoardResponse.SetTeamProfile(recruitingProfile)

	// Get Team Needs
	teamNeeds := managers.GetRecruitingNeeds(teamID)

	teamBoardResponse.SetTeamNeedsMap(teamNeeds)

	json.NewEncoder(w).Encode(teamBoardResponse)
}

func GetRecruitingClassByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	if len(teamID) == 0 {
		panic("User did not provide teamID")
	}

	recruitingProfile := managers.GetRecruitingClassByTeamID(teamID)

	json.NewEncoder(w).Encode(recruitingProfile)
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

// GetAllRecruitingProfiles
func GetAllRecruitingProfiles(w http.ResponseWriter, r *http.Request) {
	recruitingProfiles := managers.GetRecruitingProfileForRecruitSync()

	json.NewEncoder(w).Encode(recruitingProfiles)
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

// SendScholarshipToRecruit
func SendScholarshipToRecruit(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
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
	enableCors(&w)
	var updateRecruitingBoardDto structs.UpdateRecruitingBoardDTO
	err := json.NewDecoder(r.Body).Decode(&updateRecruitingBoardDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := managers.GetTimestamp()

	if ts.IsRecruitingLocked {
		http.Error(w, "Recruiting is locked!", http.StatusNotAcceptable)
		return
	}

	result := make(chan structs.RecruitingTeamProfile)

	go func() {
		recruitingProfile := managers.UpdateRecruitingProfile(updateRecruitingBoardDto)
		result <- recruitingProfile
	}()

	crootProfile := <-result
	close(result)

	fmt.Println("Updated Recruiting Profile " + strconv.Itoa(crootProfile.TeamID) + " and all associated players")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(crootProfile)
}

func GenerateWalkOnRecruits(w http.ResponseWriter, r *http.Request) {
	/*
		Get all teams and team profiles
		Get all team needs for the season

		Get all signed croots per team

		Reduce needs for season

		Add remaining needs into a string array

		Shuffle array

		For loop and generate a recruit
	*/

	var teamBoardResponses []models.TeamBoardTeamProfileResponse

	teams := managers.GetAllCoachedCollegeTeams()

	for _, team := range teams {
		teamID := strconv.Itoa(int(team.ID))
		var teamBoardResponse models.TeamBoardTeamProfileResponse

		recruitingProfile := managers.GetRecruitingProfileForTeamBoardByTeamID(teamID)

		teamBoardResponse.SetTeamProfile(recruitingProfile)
		teamNeeds := managers.GetRecruitingNeeds(teamID)
		teamBoardResponse.SetTeamNeedsMap(teamNeeds)

		teamBoardResponses = append(teamBoardResponses, teamBoardResponse)
	}

	for _, board := range teamBoardResponses {
		// teamID := board.TeamProfile.TeamID
		var teamAbbr *string = &board.TeamProfile.TeamAbbreviation
		var teamNeeds map[string]int = board.TeamNeedsMap
		croots := board.TeamProfile.Recruits

		var remainingNeedsList []string

		for _, croot := range croots {
			if croot.IsSigned && croot.Recruit.College == *teamAbbr {
				teamNeeds[croot.Recruit.Position] -= 1
			}
		}

		for k, v := range teamNeeds {
			for i := 0; i < v; i++ {
				remainingNeedsList = append(remainingNeedsList, k)
			}
		}

	}

	managers.GenerateWalkOns()

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
