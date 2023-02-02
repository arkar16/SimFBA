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
	NFLTeamTradeOptions       []NFLTradeOption `gorm:"foreignKey:TradeProposalID"`
	RecepientTeamTradeOptions []NFLTradeOption `gorm:"foreignKey:TradeProposalID"`
}

type NFLTradeOption struct {
	gorm.Model
	TradeProposalID  uint
	NFLTeamID        int
	NFLPlayerID      int
	NFLDraftPickID   uint
	SalaryPercentage float64 // Will be a percentage that the recepient team (TEAM B) will pay. Will be between 0 and 100.
	// Player           NFLPlayer    `gorm:"foreignKey:NFLPlayerID"`    // If the NFLPlayerID is greater than 0, it will return a player.
	// Draftpick        NFLDraftPick `gorm:"foreignKey:NFLDraftPickID"` // If the NFLDraftPickID is greater than 0, it will return a draft pick.
}

type NFLTradeOptionObj struct {
	gorm.Model
	TradeProposalID  uint
	NFLTeamID        int
	NFLPlayerID      int
	NFLDraftPickID   uint
	SalaryPercentage float64      // Will be a percentage that the recepient team (TEAM B) will pay. Will be between 0 and 100.
	Player           NFLPlayer    // If the NFLPlayerID is greater than 0, it will return a player.
	Draftpick        NFLDraftPick // If the NFLDraftPickID is greater than 0, it will return a draft pick.
}
