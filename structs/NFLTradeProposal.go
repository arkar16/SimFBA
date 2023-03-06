package structs

import "github.com/jinzhu/gorm"

type NFLTradeProposal struct {
	gorm.Model
	NFLTeamID                 uint
	NFLTeam                   string
	RecepientTeamID           uint
	RecepientTeam             string
	IsTradeAccepted           bool
	IsTradeRejected           bool
	IsSynced                  bool
	NFLTeamTradeOptions       []NFLTradeOption `gorm:"foreignKey:TradeProposalID"`
	RecepientTeamTradeOptions []NFLTradeOption `gorm:"foreignKey:TradeProposalID"`
}

func (p *NFLTradeProposal) ToggleSyncStatus() {
	p.IsSynced = true
}

func (p *NFLTradeProposal) AssignID(id uint) {
	p.ID = id
}

func (p *NFLTradeProposal) AcceptTrade() {
	p.IsTradeAccepted = true
}

func (p *NFLTradeProposal) RejectTrade() {
	p.IsTradeRejected = true
}

type NFLTradeOption struct {
	gorm.Model
	TradeProposalID  uint
	NFLTeamID        uint
	NFLPlayerID      uint
	NFLDraftPickID   uint
	SalaryPercentage float64 // Will be a percentage that the recepient team (TEAM B) will pay for Y1. Will be between 0 and 100.
	// Player           NFLPlayer    // `gorm:"foreignKey:PlayerID"`       // If the NFLPlayerID is greater than 0, it will return a player.
	// Draftpick        NFLDraftPick // `gorm:"foreignKey:NFLDraftPickID"` // If the NFLDraftPickID is greater than 0, it will return a draft pick.
}

type NFLTradeOptionObj struct {
	ID               uint
	TradeProposalID  uint
	NFLTeamID        uint
	NFLPlayerID      uint
	NFLDraftPickID   uint
	SalaryPercentage float64      // Will be a percentage that the recepient team (TEAM B) will pay. Will be between 0 and 100.
	Player           NFLPlayer    // If the NFLPlayerID is greater than 0, it will return a player.
	Draftpick        NFLDraftPick // If the NFLDraftPickID is greater than 0, it will return a draft pick.
}

func (to *NFLTradeOptionObj) AssignPlayer(player NFLPlayer) {
	to.Player = player
	to.NFLPlayerID = player.ID
}

func (to *NFLTradeOptionObj) AssignPick(pick NFLDraftPick) {
	to.Draftpick = pick
	to.NFLDraftPickID = pick.ID
}

type NFLTradeProposalDTO struct {
	ID                        uint
	NFLTeamID                 uint
	NFLTeam                   string
	RecepientTeamID           uint
	RecepientTeam             string
	IsTradeAccepted           bool
	IsTradeRejected           bool
	NFLTeamTradeOptions       []NFLTradeOptionObj
	RecepientTeamTradeOptions []NFLTradeOptionObj
}

type NFLTeamProposals struct {
	SentTradeProposals     []NFLTradeProposalDTO
	ReceivedTradeProposals []NFLTradeProposalDTO
}
