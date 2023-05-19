package models

import "github.com/CalebRose/SimFBA/structs"

type FreeAgencyResponse struct {
	FreeAgents    []structs.NFLPlayer
	WaiverPlayers []structs.NFLPlayer
	TeamOffers    []structs.FreeAgencyOffer
	RosterCount   uint
}
