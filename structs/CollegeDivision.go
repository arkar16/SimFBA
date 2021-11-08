package structs

import "github.com/jinzhu/gorm"

type CollegeDivision struct {
	gorm.Model
	DivisionName string
	ConferenceID int
}
