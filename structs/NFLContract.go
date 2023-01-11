package structs

import "github.com/jinzhu/gorm"

type NFLContract struct {
	gorm.Model
	PlayerID       int
	NFLPlayerID    int
	TeamID         uint
	Team           string
	OriginalTeamID uint
	OriginalTeam   string
	ContractLength int
	Y1BaseSalary   float64
	Y1Bonus        float64
	Y2BaseSalary   float64
	Y2Bonus        float64
	Y3BaseSalary   float64
	Y3Bonus        float64
	Y4BaseSalary   float64
	Y4Bonus        float64
	Y5BaseSalary   float64
	Y5Bonus        float64
	ContractType   string // Pro Bowl, Starter, Veteran, New ?
	IsActive       bool
	IsComplete     bool
	IsExtended     bool
}
