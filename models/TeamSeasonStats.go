package models

import "github.com/CalebRose/SimFBA/structs"

type TeamSeasonStats struct {
	structs.BaseTeamStats
	TotalOffensiveYards int
	TotalYardsAllowed   int
	Fumbles             int
	QBRating            float64
}
