package structs

import "github.com/jinzhu/gorm"

type CollegePlayer struct {
	gorm.Model
	PlayerID int `gorm:"column:player_id"`
	TeamID   int
	TeamAbbr string
	BasePlayer
	HighSchool    string
	City          string
	State         string
	Year          int
	IsRedshirt    bool
	IsRedshirting bool
	HasGraduated  bool
	Stats         []CollegePlayerStats `gorm:"foreignKey:CollegePlayerID"`
}

func (p *CollegePlayer) SetRedshirtingStatus() {
	if !p.IsRedshirt && !p.IsRedshirting {
		p.IsRedshirting = true
	}
}

func (p *CollegePlayer) SetRedshirtStatus() {
	if p.IsRedshirting {
		p.IsRedshirting = false
		p.IsRedshirt = true
	}
}
