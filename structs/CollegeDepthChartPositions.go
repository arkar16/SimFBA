package structs

import "github.com/jinzhu/gorm"

type CollegeDepthChartPosition struct {
	gorm.Model
	DepthChartID  int
	PlayerID      int    // 123 -- CollegePlayerID
	Position      string // "QB"
	PositionLevel string // "1"
	FirstName     string // "David"
	LastName      string // "Ross"
}
