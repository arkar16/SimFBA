package structs

import "github.com/jinzhu/gorm"

type CollegeStandings struct {
	gorm.Model
	TeamID       int
	SeasonID     int
	ConferenceID int
	DivisionID   int
	BaseStandings
}
