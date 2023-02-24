package structs

import "github.com/jinzhu/gorm"

type NFLPlayer struct {
	gorm.Model
	BasePlayer
	PlayerID          int
	TeamID            int
	College           string
	TeamAbbr          string
	Experience        uint
	HighSchool        string
	Hometown          string
	State             string
	IsActive          bool
	IsPracticeSquad   bool
	IsFreeAgent       bool
	IsWaived          bool
	IsOnTradeBlock    bool
	IsAcceptingOffers bool
	IsNegotiating     bool
	NegotiationRound  uint
	SigningRound      uint
	MinimumValue      float64
	PreviousTeamID    uint
	PreviousTeam      string
	DraftedTeamID     uint
	DraftedTeam       string
	DraftedRound      uint
	DraftedPick       uint
	Contract          NFLContract       `gorm:"foreignKey:NFLPlayerID"`
	Offers            []FreeAgencyOffer `gorm:"foreignKey:NFLPlayerID"`
}

func (np *NFLPlayer) AssignMissingValues(pr int, aca string, fa string, per string, rec string, we string) {
	np.Progression = pr
	np.AcademicBias = aca
	np.FreeAgency = fa
	np.Personality = per
	np.WorkEthic = we
	np.RecruitingBias = rec
}

func (np *NFLPlayer) AssignMinimumValue(val float64) {
	np.MinimumValue = val
}

func (np *NFLPlayer) ToggleIsFreeAgent() {
	np.PreviousTeamID = uint(np.TeamID)
	np.PreviousTeam = np.TeamAbbr
	np.IsFreeAgent = true
	np.TeamID = 0
	np.TeamAbbr = ""
	np.IsAcceptingOffers = true
	np.IsNegotiating = false
}

func (np *NFLPlayer) SignPlayer(TeamID int, Abbr string) {
	np.IsFreeAgent = false
	np.TeamID = TeamID
	np.TeamAbbr = Abbr
	np.IsAcceptingOffers = false
	np.IsNegotiating = false
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

func (np *NFLPlayer) ConvertWaivedPlayerToFA() {
	np.IsWaived = false
	np.IsFreeAgent = true
	np.IsAcceptingOffers = true
}

func (np *NFLPlayer) ToggleIsNegotiating() {
	np.IsNegotiating = true
	np.IsAcceptingOffers = false
}

func (np *NFLPlayer) AssignWorkEthic(we string) {
	np.WorkEthic = we
}

func (np *NFLPlayer) AssignPersonality(we string) {
	np.Personality = we
}

func (np *NFLPlayer) AssignFreeAgency(we string) {
	np.FreeAgency = we
}

func (np *NFLPlayer) AssignFAPreferences(negotiation uint, signing uint) {
	np.NegotiationRound = negotiation
	np.SigningRound = signing
}
