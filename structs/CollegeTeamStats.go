package structs

import "github.com/jinzhu/gorm"

type CollegeTeamStats struct {
	gorm.Model
	TeamID   int
	GameID   int
	WeekID   int
	SeasonID int
	BaseTeamStats
}
