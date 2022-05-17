package models

import "github.com/CalebRose/SimFBA/structs"

type SimGameDataResponse struct {
	HomeTeam       SimTeamDataResponse
	HomeTeamRoster []structs.CollegePlayer
	AwayTeam       SimTeamDataResponse
	AwayTeamRoster []structs.CollegePlayer
	GameID         int
	WeekID         int
	SeasonID       int
}

func (sgdr *SimGameDataResponse) AssignHomeTeam(team SimTeamDataResponse, roster []structs.CollegePlayer) {
	sgdr.HomeTeam = team
	sgdr.HomeTeamRoster = roster
}

func (sgdr *SimGameDataResponse) AssignAwayTeam(team SimTeamDataResponse, roster []structs.CollegePlayer) {
	sgdr.AwayTeam = team
	sgdr.AwayTeamRoster = roster
}
