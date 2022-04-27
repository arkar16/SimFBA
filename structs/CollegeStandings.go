package structs

import "github.com/jinzhu/gorm"

type CollegeStandings struct {
	gorm.Model
	TeamID           int
	TeamName         string
	SeasonID         int
	Season           int
	ConferenceID     int
	DivisionID       int
	PostSeasonStatus string
	BaseStandings
}
