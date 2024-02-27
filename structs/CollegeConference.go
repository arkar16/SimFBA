package structs

import "github.com/jinzhu/gorm"

type CollegeConference struct {
	gorm.Model

	ConferenceName     string
	ConferenceAbbr     string
	LeagueID           uint
	LeagueName         string
	ConferenceDivision string
	Divisions          []CollegeDivision `gorm:"foreignKey:ConferenceID"`
}
