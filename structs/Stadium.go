package structs

import "github.com/jinzhu/gorm"

type Stadium struct {
	gorm.Model
	StadiumName      string
	TeamID           uint
	TeamAbbr         string
	City             string
	State            string
	Country          string
	Region           string
	Capacity         uint
	RecordAttendance uint
	FirstSeason      uint
	LeagueID         uint
	LeagueName       string
	IsDomed          bool
}
