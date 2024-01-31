package models

import "github.com/CalebRose/SimFBA/structs"

type SimGameDataResponse struct {
	HomeTeam       SimTeamDataResponse
	HomeTeamRoster []structs.CollegePlayer
	AwayTeam       SimTeamDataResponse
	AwayTeamRoster []structs.CollegePlayer
	Stadium        structs.Stadium
	GameID         int
	WeekID         int
	SeasonID       int
	GameTemp       float64
	Cloud          string
	Precip         string
	WindSpeed      float64
	WindCategory   string
	IsPostSeason   bool
}

func (sgdr *SimGameDataResponse) AssignHomeTeam(team SimTeamDataResponse, roster []structs.CollegePlayer) {
	sgdr.HomeTeam = team
	sgdr.HomeTeamRoster = roster
}

func (sgdr *SimGameDataResponse) AssignAwayTeam(team SimTeamDataResponse, roster []structs.CollegePlayer) {
	sgdr.AwayTeam = team
	sgdr.AwayTeamRoster = roster
}

func (sgdr *SimGameDataResponse) AssignWeather(temp float64, cloud string, precip string, wind string, windspeed float64) {
	sgdr.GameTemp = temp
	sgdr.Cloud = cloud
	sgdr.Precip = precip
	sgdr.WindSpeed = windspeed
	sgdr.WindCategory = wind
}

func (sgdr *SimGameDataResponse) AssignStadium(s structs.Stadium) {
	sgdr.Stadium = s
}

func (sgdr *SimGameDataResponse) AssignPostSeasonStatus(isPostSeason bool) {
	sgdr.IsPostSeason = isPostSeason
}

type NFLSimGameDataResponse struct {
	HomeTeam       NFLSimTeamDataResponse
	HomeTeamRoster []structs.NFLPlayer
	AwayTeam       NFLSimTeamDataResponse
	AwayTeamRoster []structs.NFLPlayer
	Stadium        structs.Stadium
	GameID         int
	WeekID         int
	SeasonID       int
	GameTemp       float64
	Cloud          string
	Precip         string
	WindSpeed      float64
	WindCategory   string
	IsPostSeason   bool
	IsNeutral      bool
}

func (sgdr *NFLSimGameDataResponse) AssignHomeTeam(team NFLSimTeamDataResponse, roster []structs.NFLPlayer) {
	sgdr.HomeTeam = team
	sgdr.HomeTeamRoster = roster
}

func (sgdr *NFLSimGameDataResponse) AssignAwayTeam(team NFLSimTeamDataResponse, roster []structs.NFLPlayer) {
	sgdr.AwayTeam = team
	sgdr.AwayTeamRoster = roster
}

func (sgdr *NFLSimGameDataResponse) AssignWeather(temp float64, cloud string, precip string, wind string, windspeed float64) {
	sgdr.GameTemp = temp
	sgdr.Cloud = cloud
	sgdr.Precip = precip
	sgdr.WindSpeed = windspeed
	sgdr.WindCategory = wind
}

func (sgdr *NFLSimGameDataResponse) AssignStadium(s structs.Stadium) {
	sgdr.Stadium = s
}

func (sgdr *NFLSimGameDataResponse) AssignPostSeasonStatus(isPostSeason, isNeutral bool) {
	sgdr.IsPostSeason = isPostSeason
	sgdr.IsNeutral = isNeutral
}
