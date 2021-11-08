package structs

import "github.com/jinzhu/gorm"

type CollegeConference struct {
	gorm.Model
	ConferenceName     int
	ConferenceDivision string
	Divisions          []CollegeDivision
}
