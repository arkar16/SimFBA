package structs

import "github.com/jinzhu/gorm"

type NFLStandings struct {
	gorm.Model
	TeamID           uint
	TeamName         string
	Mascot           string
	SeasonID         uint
	Season           uint
	LeagueID         uint
	LeagueName       string
	ConferenceID     uint
	ConferenceName   string
	TotalTies        uint
	ConferenceTies   uint
	DivisionID       uint
	DivisionWins     uint
	DivisionLosses   uint
	DivisionTies     uint
	PostSeasonStatus string
	BaseStandings
}
