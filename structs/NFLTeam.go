package structs

import "github.com/jinzhu/gorm"

type NFLTeam struct {
	gorm.Model
	BaseTeam
	ConferenceID uint
	Conference   string
	DivisionID   uint
	Division     string
	NFLCoach     NFLCoach       `gorm:"foreignKey:TeamID"`
	Contracts    []NFLContract  `gorm:"foreignKey:TeamID"`
	DraftPicks   []NFLDraftPick `gorm:"foreignKey:TeamID"`
}
