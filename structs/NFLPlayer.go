package structs

import "github.com/jinzhu/gorm"

type NFLPlayer struct {
	gorm.Model
	PlayerID int
	TeamID   int
	College  string
	BasePlayer
	Contract string // Will Modify later

}
