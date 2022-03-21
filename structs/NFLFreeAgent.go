package structs

import "github.com/jinzhu/gorm"

type NFLFreeAgent struct {
	gorm.Model
	BasePlayer
	College string
}
