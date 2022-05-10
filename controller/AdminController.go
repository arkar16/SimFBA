package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
)

// GetTimeStamp
func GetCurrentTimestamp(w http.ResponseWriter, r *http.Request) {

	timestamp := managers.GetTimestamp()

	json.NewEncoder(w).Encode(timestamp)
}

// SyncWeek?
func SyncTimestamp(w http.ResponseWriter, r *http.Request) {
	var updateTimestampDto structs.UpdateTimestampDto
	err := json.NewDecoder(r.Body).Decode(&updateTimestampDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTimestamp := managers.UpdateTimestamp(updateTimestampDto)

	json.NewEncoder(w).Encode(newTimestamp)
}

func SyncMissingRES(w http.ResponseWriter, r *http.Request) {
	managers.SyncAllMissingEfficiencies()

}

// CreateCollegeRecruit?

// CreateNFLPlayer -- Create NFL Player from template, and then synthetically progress them based on the year of input

// UpdateTeamRecruitingProfile

// ApproveCoachForTeam

// RemoveCoachFromTeam

// UpdateTeam

// RunProgressionsForCollege

// RunProgressionsForNFL

// RunProgressionsForJuco?
