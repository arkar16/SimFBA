package structs

import "github.com/jinzhu/gorm"

type CollegeSeason struct {
	gorm.Model
	Season
	CollegeWeeks []CollegeWeek `gorm:"foreignKey:SeasonID"`
}
