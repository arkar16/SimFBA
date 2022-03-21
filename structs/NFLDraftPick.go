package structs

import "github.com/jinzhu/gorm"

type NFLDraftPick struct {
	gorm.Model
	PickNumber     int
	OriginalTeamID int
	OriginalTeam   string
	CurrentTeamID  int
	CurrentTeam    string
	PickSelection  string
	PlayerID       int
	Round          string
	SeasonID       int
	Season         int
	TradeValue     int
}
