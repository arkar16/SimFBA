package structs

import "github.com/jinzhu/gorm"

type RecruitRegion struct {
	gorm.Model
	StateID    int
	State      string
	RegionName string
}
