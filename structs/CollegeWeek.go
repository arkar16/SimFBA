package structs

import "github.com/jinzhu/gorm"

type CollegeWeek struct {
	gorm.Model
	BaseWeek
	Games []CollegeGame `gorm:"foreignKey:WeekID"`
}
