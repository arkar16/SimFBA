package structs

import "github.com/jinzhu/gorm"

type CollegeDepthChartPosition struct {
	gorm.Model
	DepthChartID     int
	PlayerID         int           `gorm:"column:player_id"` // 123 -- CollegePlayerID
	Position         string        // "QB"
	PositionLevel    string        // "1"
	FirstName        string        // "David"
	LastName         string        // "Ross"
	OriginalPosition string        // The Original Position of the Player. Will only be used for STU position
	CollegePlayer    CollegePlayer `gorm:"foreignKey:PlayerID;references:PlayerID"`
}

// Update DepthChartPosition -- Updates the Player taking the position
func (dcp *CollegeDepthChartPosition) UpdateDepthChartPosition(dto CollegeDepthChartPosition) {
	if dcp.ID != dto.ID || dcp.DepthChartID != dto.DepthChartID {
		return
	}
	dcp.PlayerID = dto.PlayerID
	dcp.FirstName = dto.FirstName
	dcp.LastName = dto.LastName
	dcp.OriginalPosition = dto.OriginalPosition
}
