package structs

import "github.com/jinzhu/gorm"

type NFLCapsheet struct {
	gorm.Model
	NFLTeamID uint
	Y1Bonus   float64
	Y1Salary  float64
	Y1CapHit  float64
	Y2Bonus   float64
	Y2Salary  float64
	Y2CapHit  float64
	Y3Bonus   float64
	Y3Salary  float64
	Y3CapHit  float64
	Y4Bonus   float64
	Y4Salary  float64
	Y4CapHit  float64
	Y5Bonus   float64
	Y5Salary  float64
	Y5CapHit  float64
}
