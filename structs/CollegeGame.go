package structs

import "github.com/jinzhu/gorm"

type CollegeGame struct {
	gorm.Model
	WeekID                   int
	Week                     int
	SeasonID                 int
	HomeTeamID               int
	HomeTeam                 string
	HomeTeamCoach            string
	HomeTeamWin              bool
	AwayTeamID               int
	AwayTeam                 string
	AwayTeamCoach            string
	AwayTeamWin              bool
	MVP                      string
	HomeTeamScore            int
	AwayTeamScore            int
	TimeSlot                 string
	Stadium                  string
	City                     string
	State                    string
	IsNeutral                bool
	IsConference             bool
	IsDivisional             bool
	IsConferenceChampionship bool
	IsBowlGame               bool
	IsPlayoffGame            bool
	IsNationalChampionship   bool
	IsRivalryGame            bool
	GameComplete             bool
}

func (cg *CollegeGame) UpdateScore(HomeScore int, AwayScore int) {
	cg.HomeTeamScore = HomeScore
	cg.AwayTeamScore = AwayScore
	if HomeScore > AwayScore {
		cg.HomeTeamWin = true
	} else {
		cg.AwayTeamWin = true
	}
	cg.GameComplete = true
}

func (cg *CollegeGame) UpdateCoach(TeamID int, Username string) {
	if cg.HomeTeamID == TeamID {
		cg.HomeTeamCoach = Username
	} else if cg.AwayTeamID == TeamID {
		cg.AwayTeamCoach = Username
	}
}
