package structs

import "github.com/jinzhu/gorm"

type CollegeCoach struct {
	gorm.Model
	CoachName                      string
	TeamID                         int
	OverallWins                    int
	OverallLosses                  int
	OverallConferenceChampionships int
	BowlWins                       int
	BowlLosses                     int
	PlayoffWins                    int
	PlayoffLosses                  int
	NationalChampionships          int
	IsActive                       bool
}

func (cc *CollegeCoach) SetTeam(TeamID int) {
	cc.TeamID = TeamID
	cc.IsActive = true
}

func (cc *CollegeCoach) SetAsInactive() {
	cc.IsActive = false
	cc.TeamID = 0
}
