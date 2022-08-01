package structs

import "github.com/jinzhu/gorm"

type CollegeConference struct {
	gorm.Model
	ConferenceName     string
	ConferenceAbbr     string
	ConferenceDivision string
	Divisions          []CollegeDivision `gorm:"foreignKey:ConferenceID"`
}
