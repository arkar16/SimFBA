package structs

import "github.com/jinzhu/gorm"

type CollegeGame struct {
	gorm.Model
	WeekID                   int
	SeasonID                 int
	HomeTeamID               int
	HomeTeam                 string
	AwayTeamID               int
	AwayTeam                 string
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
