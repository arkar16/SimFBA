package structs

import "github.com/jinzhu/gorm"

type CollegeRival struct {
	gorm.Model
	TeamID        int
	Team          string
	TeamWins      int
	RivalID       int
	RivalTeam     string
	RivalWins     int
	TeamStreak    int
	RivalStreak   int
	CurrentStreak int
	LatestVictor  string
}
