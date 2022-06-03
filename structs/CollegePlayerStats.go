package structs

import "github.com/jinzhu/gorm"

type CollegePlayerStats struct {
	gorm.Model
	CollegePlayerID int
	TeamID          int
	GameID          int
	WeekID          int
	SeasonID        int
	OpposingTeam    string
	BasePlayerStats
}
