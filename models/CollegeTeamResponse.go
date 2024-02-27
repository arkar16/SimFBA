package models

import "github.com/CalebRose/SimFBA/structs"

type CollegeTeamResponse struct {
	ID int
	structs.BaseTeam
	ConferenceID int
	Conference   string
	DivisionID   int
	Division     string
	Stats        structs.CollegeTeamStats
	SeasonStats  structs.CollegeTeamSeasonStats
}

type NFLTeamResponse struct {
	ID int
	structs.BaseTeam
	ConferenceID int
	Conference   string
	DivisionID   int
	Division     string
	Stats        structs.NFLTeamStats
	SeasonStats  structs.NFLTeamSeasonStats
}
