package structs

import "github.com/jinzhu/gorm"

type CollegePlayerStats struct {
	gorm.Model
	CollegePlayerID int
	GameID          int
	WeekID          int
	SeasonID        int
	OpposingTeam    string
	Year            int
	IsRedshirt      bool
	BasePlayerStats
}

func (c *CollegePlayerStats) ApplyYear(year int) {
	c.Year = year
}
