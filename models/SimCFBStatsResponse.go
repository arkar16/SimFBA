package models

import "github.com/CalebRose/SimFBA/structs"

type SimCFBStatsResponse struct {
	CollegeConferences []structs.CollegeConference
	CollegePlayers     []CollegePlayerResponse
	CollegeTeams       []CollegeTeamResponse
}
