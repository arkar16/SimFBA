package structs

import "github.com/jinzhu/gorm"

type NFLDraftPick struct {
	gorm.Model
	OriginalTeamID uint
	OriginalTeam   string
	TeamID         uint
	Team           string
	PickSelection  string
	PlayerID       uint
	Round          uint
	PickNumber     uint
	SeasonID       uint
	Season         uint
	TradeValue     uint
}
