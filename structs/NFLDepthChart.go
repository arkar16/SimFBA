package structs

import "github.com/jinzhu/gorm"

type NFLDepthChart struct {
	gorm.Model
	TeamID            int
	DepthChartPlayers []NFLDepthChartPosition `gorm:"foreignKey:DepthChartID"`
}

type NFLDepthChartPosition struct {
	gorm.Model
	DepthChartID     uint
	PlayerID         uint      `gorm:"column:player_id"` // 123 -- CollegePlayerID
	Position         string    // "QB"
	PositionLevel    string    // "1"
	FirstName        string    // "David"
	LastName         string    // "Ross"
	OriginalPosition string    // The Original Position of the Player. Will only be used for STU position
	NFLPlayer        NFLPlayer `gorm:"foreignKey:PlayerID;references:PlayerID"`
}

// Update DepthChartPosition -- Updates the Player taking the position
func (dcp *NFLDepthChartPosition) UpdateDepthChartPosition(dto NFLDepthChartPosition) {
	if dcp.ID != dto.ID || dcp.DepthChartID != dto.DepthChartID {
		return
	}
	dcp.PlayerID = dto.PlayerID
	dcp.FirstName = dto.FirstName
	dcp.LastName = dto.LastName
	dcp.OriginalPosition = dto.OriginalPosition
}

// Update DepthChartPosition -- Updates the Player taking the position
func (dcp *NFLDepthChartPosition) AssignID(id uint) {
	dcp.ID = id
}
