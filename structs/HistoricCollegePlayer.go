package structs

import "github.com/jinzhu/gorm"

type HistoricCollegePlayer struct {
	gorm.Model
	BasePlayer
	PlayerID      int
	TeamID        int
	TeamAbbr      string
	HighSchool    string
	City          string
	State         string
	Year          int
	IsRedshirt    bool
	IsRedshirting bool
	HasGraduated  bool
	Stats         []CollegePlayerStats `gorm:"foreignKey:CollegePlayerID"`
}
