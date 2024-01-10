package models

import "github.com/CalebRose/SimFBA/structs"

type FreeAgencyResponse struct {
	FreeAgents    []FreeAgentResponse
	WaiverPlayers []WaiverWirePlayerResponse
	PracticeSquad []FreeAgentResponse
	TeamOffers    []structs.FreeAgencyOffer
	RosterCount   uint
}

type FreeAgentResponse struct {
	ID       uint
	PlayerID int
	structs.BasePlayer
	TeamID          int
	College         string
	TeamAbbr        string
	Hometown        string
	State           string
	Experience      uint
	IsActive        bool
	IsFreeAgent     bool
	IsWaived        bool
	MinimumValue    float64
	PreviousTeam    string
	DraftedTeam     string
	ShowLetterGrade bool
	IsPracticeSquad bool
	Stats           []structs.NFLPlayerStats
	SeasonStats     structs.NFLPlayerSeasonStats
	Offers          []structs.FreeAgencyOffer
}

type WaiverWirePlayerResponse struct {
	ID       uint
	PlayerID int
	structs.BasePlayer
	TeamID          int
	College         string
	TeamAbbr        string
	Hometown        string
	State           string
	Experience      uint
	IsActive        bool
	IsFreeAgent     bool
	IsWaived        bool
	MinimumValue    float64
	PreviousTeam    string
	DraftedTeam     string
	ShowLetterGrade bool
	IsPracticeSquad bool
	WaiverOffers    []structs.NFLWaiverOffer
	Contract        structs.NFLContract
}
