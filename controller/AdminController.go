package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

// GetTimeStamp
func GetCurrentTimestamp(w http.ResponseWriter, r *http.Request) {

	timestamp := managers.GetTimestamp()

	json.NewEncoder(w).Encode(timestamp)
}

// SyncRecruiting

// SyncWeek?

// CreateCollegeRecruit?

// CreateNFLPlayer -- Create NFL Player from template, and then synthetically progress them based on the year of input

// UpdateTeamRecruitingProfile

// ApproveCoachForTeam

// RemoveCoachFromTeam

// UpdateTeam

// RunProgressionsForCollege

// RunProgressionsForNFL

// RunProgressionsForJuco?
