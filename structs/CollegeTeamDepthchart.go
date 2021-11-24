package structs

import "github.com/jinzhu/gorm"

type CollegeTeamDepthChart struct {
	gorm.Model
	TeamID            int
	DepthChartPlayers []CollegeDepthChartPosition `gorm:"foreignKey:DepthChartID"`
}
