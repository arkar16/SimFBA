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

// Update DepthChartPosition -- Updates the Player taking the position
func (dcp *CollegeDepthChartPosition) UpdateDepthChartPosition(dto CollegeDepthChartPosition) {
	if dcp.ID != dto.ID || dcp.DepthChartID != dto.DepthChartID {
		return
	}
	dcp.PlayerID = dto.PlayerID
	dcp.FirstName = dto.FirstName
	dcp.LastName = dto.LastName
}
