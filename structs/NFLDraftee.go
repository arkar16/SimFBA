package structs

import "github.com/jinzhu/gorm"

type NFLDraftee struct {
	gorm.Model
	BasePlayer
	PlayerID   int
	HighSchool string
	College    string
	City       string
	State      string
}
