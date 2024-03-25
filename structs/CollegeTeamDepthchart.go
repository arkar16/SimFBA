package structs

import "github.com/jinzhu/gorm"

type CollegeTeamDepthChart struct {
	gorm.Model
	TeamID            int
	DepthChartPlayers []CollegeDepthChartPosition `gorm:"foreignKey:DepthChartID"`
}

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

// Update DepthChartPosition -- Updates the Player taking the position
func (dcp *CollegeDepthChartPosition) AssignID(id uint) {
	dcp.ID = id
}

type CollegeTeamDepthChartTEST struct {
	gorm.Model
	TeamID            int
	DepthChartPlayers []CollegeDepthChartPositionTEST `gorm:"foreignKey:DepthChartID"`
}

type CollegeDepthChartPositionTEST struct {
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
func (dcp *CollegeDepthChartPositionTEST) UpdateDepthChartPosition(dto CollegeDepthChartPosition) {
	if dcp.ID != dto.ID || dcp.DepthChartID != dto.DepthChartID {
		return
	}
	dcp.PlayerID = dto.PlayerID
	dcp.FirstName = dto.FirstName
	dcp.LastName = dto.LastName
	dcp.OriginalPosition = dto.OriginalPosition
}

// Update DepthChartPosition -- Updates the Player taking the position
func (dcp *CollegeDepthChartPositionTEST) AssignID(id uint) {
	dcp.ID = id
}
