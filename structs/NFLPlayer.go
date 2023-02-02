package structs

import "github.com/jinzhu/gorm"

type NFLPlayer struct {
	gorm.Model
	BasePlayer
	PlayerID        int
	TeamID          int
	College         string
	TeamAbbr        string
	Experience      uint
	HighSchool      string
	Hometown        string
	State           string
	IsActive        bool
	IsPracticeSquad bool
	IsFreeAgent     bool
	IsWaived        bool
	IsOnTradeBlock  bool
	Contract        NFLContract `gorm:"foreignKey:NFLPlayerID"`
}

func (np *NFLPlayer) AssignMissingValues(pr int, aca string, fa string, per string, rec string, we string) {
	np.Progression = pr
	np.AcademicBias = aca
	np.FreeAgency = fa
	np.Personality = per
	np.WorkEthic = we
	np.RecruitingBias = rec
}

func (np *NFLPlayer) ToggleIsFreeAgent() {
	np.IsFreeAgent = true
}

func (np *NFLPlayer) ToggleIsPracticeSquad() {
	np.IsPracticeSquad = true
}

func (np *NFLPlayer) PlaceOnTradeBlock() {
	np.IsOnTradeBlock = true
}

func (np *NFLPlayer) RemoveFromTradeBlock() {
	np.IsOnTradeBlock = false
}

func (np *NFLPlayer) WaivePlayer() {
	np.TeamID = 0
	np.TeamAbbr = ""
	np.IsWaived = true
}
