package util

import (
	"github.com/CalebRose/SimFBA/structs"
)

func GetHeismanScore(cp structs.CollegePlayer, weightMap map[string]float64, homeTeamMapper map[int]string) float64 {
	var score float64 = 0

	homeTeam := homeTeamMapper[cp.TeamID]

	stats := cp.Stats

	for _, stat := range stats {
		var statScore float64 = 0
		opposingTeam := stat.OpposingTeam
		statScore += (float64(stat.PassingYards) * 0.075)
		statScore += (float64(stat.PassingTDs) * 4)
		statScore -= (float64(stat.Interceptions) * 2)
		statScore -= (float64(stat.Sacks) * 2)
		if cp.Position == "RB" || cp.Position == "FB" {
			statScore += (float64(stat.RushingYards) * 0.1)
			statScore += (float64(stat.RushingTDs) * 6)
		} else {
			statScore += (float64(stat.RushingYards) * 0.0775)
			statScore += (float64(stat.RushingTDs) * 4.5)
		}
		if cp.Position == "WR" || cp.Position == "TE" {
			statScore += (float64(stat.Catches) * 0.5125)
			statScore += (float64(stat.ReceivingYards) * 0.1125)
			statScore += (float64(stat.ReceivingTDs) * 6)
		} else {
			statScore += (float64(stat.Catches) * 0.25)
			statScore += (float64(stat.ReceivingYards) * 0.05)
			statScore += (float64(stat.ReceivingTDs) * 4)
		}
		statScore -= (float64(stat.Fumbles) * 2)
		statScore += (float64(stat.SoloTackles) * 1.8)
		statScore += (float64(stat.STSoloTackles) * 1.8)
		statScore += (float64(stat.AssistedTackles) * 1)
		statScore += (float64(stat.STAssistedTackles) * 1)
		statScore += (float64(stat.TacklesForLoss) * 3.75)
		statScore += (float64(stat.SacksMade) * 6)
		statScore += (float64(stat.ForcedFumbles) * 3.75)
		statScore += (float64(stat.RecoveredFumbles) * 6)
		statScore += (float64(stat.PassDeflections) * 4)
		statScore += (float64(stat.InterceptionsCaught) * 6)
		statScore += (float64(stat.Safeties) * 4)
		statScore += (float64(stat.DefensiveTDs) * 10)
		statScore += (float64(stat.FGMade) * 0.3)
		statScore += (float64(stat.ExtraPointsMade) * 0.1)
		statScore += (float64(stat.KickReturnTDs) * 6)
		statScore += (float64(stat.PuntReturnTDs) * 6)
		statScore += (float64(stat.PuntsBlocked) * 2)
		statScore += (float64(stat.FGBlocked) * 2)

		statScore *= weightMap[opposingTeam]

		statScore *= weightMap[homeTeam]

		score += statScore
	}

	score = score / float64(len(stats))

	return score
}
