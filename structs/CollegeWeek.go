package structs

import "github.com/jinzhu/gorm"

type CollegeWeek struct {
	gorm.Model
	BaseWeek
	IsRecruitingAllowed bool
	IsRegularSeason     bool
	IsBowlSeason        bool
	Games               []CollegeGame `gorm:"foreignKey:WeekID"`
}
