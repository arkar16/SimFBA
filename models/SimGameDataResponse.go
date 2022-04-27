package models

import "github.com/CalebRose/SimFBA/structs"

type SimGameDataResponse struct {
	HomeTeam       structs.CollegeTeam
	HomeTeamRoster []structs.CollegePlayer
	AwayTeam       structs.CollegeTeam
	AwayTeamRoster []structs.CollegePlayer
}

func (sgdr *SimGameDataResponse) AssignHomeTeam(team structs.CollegeTeam, roster []structs.CollegePlayer) {
	sgdr.HomeTeam = team
	sgdr.HomeTeamRoster = roster
}

func (sgdr *SimGameDataResponse) AssignAwayTeam(team structs.CollegeTeam, roster []structs.CollegePlayer) {
	sgdr.AwayTeam = team
	sgdr.AwayTeamRoster = roster
}
