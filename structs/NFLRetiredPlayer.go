package structs

import "github.com/jinzhu/gorm"

type NFLRetiredPlayer struct {
	gorm.Model
	BasePlayer
	PlayerID        int
	TeamID          int
	College         string
	TeamAbbr        string
	Experience      uint
	HighSchool      string
	Hometown        string
	State           string
	IsActive        bool
	IsPracticeSquad bool
	IsFreeAgent     bool
	IsWaived        bool
	IsOnTradeBlock  bool
	Contract        NFLContract `gorm:"foreignKey:NFLPlayerID"`
}
