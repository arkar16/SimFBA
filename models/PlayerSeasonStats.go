package models

import "github.com/CalebRose/SimFBA/structs"

type PlayerSeasonStats struct {
	structs.BasePlayerStats
	QBRating     float64
	Tackles      float64
	RushingAvg   float64
	PassingAvg   float64
	ReceivingAvg float64
}
