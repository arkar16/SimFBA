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
