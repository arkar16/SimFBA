package structs

import "github.com/jinzhu/gorm"

type NFLContract struct {
	gorm.Model
	PlayerID       int
	NFLPlayerID    int
	ContractLength int
	BaseSalary     float32
	Bonus          float32
	ContractType   string // Pro Bowl, Starter, Veteran, New ?
	IsExtended     bool
}
