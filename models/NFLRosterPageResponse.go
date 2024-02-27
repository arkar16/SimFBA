package models

import "github.com/CalebRose/SimFBA/structs"

type NFLRosterPageResponse struct {
	Team   structs.NFLTeam
	Roster []structs.NFLPlayer
}
