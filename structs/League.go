package structs

import "github.com/jinzhu/gorm"

type League struct {
	gorm.Model
	LeagueName     string
	IsProfessional bool
}
