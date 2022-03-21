package structs

import "github.com/jinzhu/gorm"

type RecruitState struct {
	gorm.Model
	RecruitState string
}
