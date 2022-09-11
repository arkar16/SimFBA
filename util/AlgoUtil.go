package util

import (
	"github.com/CalebRose/SimFBA/structs"
)

func GetHeismanScore(cp structs.CollegePlayer, weightMap map[string]float64, homeTeamMapper map[int]string) float64 {
	var score float64 = 0

	homeTeam := homeTeamMapper[cp.TeamID]

	stats := cp.Stats

	var opponentMod float64 = 0

	for _, stat := range stats {
		var statScore float64 = 0
		opposingTeam := stat.OpposingTeam
		statScore += (float64(stat.PassingYards) * 0.070)
		statScore += (float64(stat.PassingTDs) * 4)
		statScore -= (float64(stat.Interceptions) * 2.25)
		statScore -= (float64(stat.Sacks) * 2.25)
		if cp.Position == "RB" || cp.Position == "FB" {
			statScore += (float64(stat.RushingYards) * 0.1)
			statScore += (float64(stat.RushingTDs) * 6)
		} else {
			statScore += (float64(stat.RushingYards) * 0.0775)
			statScore += (float64(stat.RushingTDs) * 4.65)
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
		statScore -= (float64(stat.Fumbles) * 4)
		statScore += (float64(stat.SoloTackles) * 1.5)
		statScore += (float64(stat.STSoloTackles) * 1.5)
		statScore += (float64(stat.AssistedTackles) * .9)
		statScore += (float64(stat.STAssistedTackles) * .9)
		statScore += (float64(stat.TacklesForLoss) * 6)
		statScore += (float64(stat.SacksMade) * 7.5)
		statScore += (float64(stat.ForcedFumbles) * 6.5)
		statScore += (float64(stat.RecoveredFumbles) * 6.5)
		statScore += (float64(stat.PassDeflections) * 4)
		statScore += (float64(stat.InterceptionsCaught) * 10)
		statScore += (float64(stat.Safeties) * 8)
		statScore += (float64(stat.DefensiveTDs) * 20)
		statScore += (float64(stat.FGMade) * 0.3)
		statScore += (float64(stat.ExtraPointsMade) * 0.1)
		statScore += (float64(stat.KickReturnTDs) * 8)
		statScore += (float64(stat.PuntReturnTDs) * 8)
		statScore += (float64(stat.PuntsBlocked) * 6)
		statScore += (float64(stat.FGBlocked) * 6)

		opponentMod += weightMap[opposingTeam]

		// statScore *= weightMap[opposingTeam]

		// statScore *= weightMap[homeTeam]

		score += statScore
	}

	score = (score / float64(len(stats))) * (float64(len(stats)) / 12)

	score = score * (1 + opponentMod + weightMap[homeTeam])

	return score
}
