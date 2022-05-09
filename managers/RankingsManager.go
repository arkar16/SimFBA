package managers

import (
	"math"
	"math/rand"
	"time"

	"github.com/CalebRose/SimFBA/dbprovider"
	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
)

func AssignAllRecruitRanks() {
	db := dbprovider.GetInstance().GetDB()

	var recruits []structs.Recruit

	// var recruitsToSync []structs.Recruit

	db.Where("team_id = 0").Find(&recruits)

	rivalsModifiers := config.RivalsModifiers()

	for idx, croot := range recruits {
		// 247 Rankings
		rank247 := Get247Ranking(croot)
		// ESPN Rankings
		espnRank := GetESPNRanking(croot)

		// Rivals Ranking
		var rivalsRank float64 = 0
		if idx <= 249 {
			rivalsBonus := rivalsModifiers[idx]

			rivalsRank = GetRivalsRanking(croot.Stars, rivalsBonus)
		}

		croot.AssignRankValues(rank247, espnRank, rivalsRank)

		db.Save(&croot)

		// recruitsToSync = append(recruitsToSync, croot)
	}

}

func Get247Ranking(r structs.Recruit) float64 {
	ovr := r.Overall

	potentialGrade := Get247PotentialModifier(r.PotentialGrade)

	return float64(ovr) + (potentialGrade * 2)
}

func GetESPNRanking(r structs.Recruit) float64 {
	// ESPN Ranking = Star Rank + Archetype Modifier + weight difference + height difference
	// + potential val, and then round.

	starRank := GetESPNStarRank(r.Stars)
	archMod := GetArchetypeModifier(r.Archetype)
	potentialMod := GetESPNPotentialModifier(r.PotentialGrade)

	espnPositionMap := config.ESPNModifiers()
	heightMod := float64(r.Height) / espnPositionMap[r.Position]["Height"]
	weightMod := float64(r.Weight) / espnPositionMap[r.Position]["Weight"]
	espnRanking := math.Round(float64(starRank) + float64(archMod) + potentialMod + heightMod + weightMod)

	return espnRanking
}

func GetRivalsRanking(stars int, bonus int) float64 {
	return GetRivalsStarModifier(stars) + float64(bonus)
}

func GetESPNStarRank(star int) int {
	if star == 5 {
		return 95
	} else if star == 4 {
		return 85
	} else if star == 3 {
		return 75
	} else if star == 2 {
		return 65
	}
	return 55
}

func GetArchetypeModifier(arch string) int {
	if arch == "Coverage" ||
		arch == "Run Stopper" ||
		arch == "Ball Hawk" ||
		arch == "Man Coverage" ||
		arch == "Pass Rusher" ||
		arch == "Rushing" {
		return 1
	} else if arch == "Possession" ||
		arch == "Field General" ||
		arch == "Nose Tackle" ||
		arch == "Blocking" ||
		arch == "Line Captain" {
		return -1
	} else if arch == "Speed Rusher" ||
		arch == "Pass Rush" || arch == "Scrambler" ||
		arch == "Vertical Threat" ||
		arch == "Speed" {
		return 2
	}
	return 0
}

func Get247PotentialModifier(pg string) float64 {
	if pg == "A+" {
		return 7.5
	} else if pg == "A" {
		return 7
	} else if pg == "A-" {
		return 6.5
	} else if pg == "B+" {
		return 6
	} else if pg == "B" {
		return 5.5
	} else if pg == "B-" {
		return 5
	} else if pg == "C+" {
		return 4.5
	} else if pg == "C" {
		return 4
	} else if pg == "C-" {
		return 3.5
	} else if pg == "D+" {
		return 3
	} else if pg == "D" {
		return 2.5
	} else if pg == "D-" {
		return 1.75
	}
	return 1
}

func GetESPNPotentialModifier(pg string) float64 {
	if pg == "A+" {
		return 1
	} else if pg == "A" {
		return 0.9
	} else if pg == "A-" {
		return 0.8
	} else if pg == "B+" {
		return 0.6
	} else if pg == "B" {
		return 0.4
	} else if pg == "B-" {
		return 0.2
	} else if pg == "C+" {
		return 0
	} else if pg == "C" {
		return -0.15
	} else if pg == "C-" {
		return -0.3
	} else if pg == "D+" {
		return -0.6
	} else if pg == "D" {
		return -0.75
	} else if pg == "D-" {
		return -0.9
	}
	return -1
}

func GetPredictiveOverall(r structs.Recruit) int {
	currentOverall := r.Overall

	var potentialProg int

	if r.PotentialGrade == "B+" ||
		r.PotentialGrade == "A-" ||
		r.PotentialGrade == "A" ||
		r.PotentialGrade == "A+" {
		potentialProg = 7
	} else if r.PotentialGrade == "B" ||
		r.PotentialGrade == "B-" ||
		r.PotentialGrade == "C+" {
		potentialProg = 5
	} else {
		potentialProg = 4
	}

	return currentOverall + (potentialProg * 3)
}

func GetRivalsStarModifier(stars int) float64 {
	rand.Seed(time.Now().UnixNano())
	if stars == 5 {
		return 6.1
	} else if stars == 4 {
		return RoundToFixedDecimalPlace(rand.Float64()*((6.0-5.8)+5.8), 1)
	} else if stars == 3 {
		return RoundToFixedDecimalPlace(rand.Float64()*((5.7-5.5)+5.5), 1)
	} else if stars == 2 {
		return RoundToFixedDecimalPlace(rand.Float64()*((5.4-5.2)+5.2), 1)
	} else {
		return 5
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundToFixedDecimalPlace(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func Get247TeamRanking(rp structs.RecruitingTeamProfile, signedCroots []structs.Recruit) float64 {
	stddev := 10

	var Rank247 float64 = 0

	for idx, croot := range signedCroots {

		rank := float64((idx - 1) / stddev)

		expo := (-0.5 * (math.Pow(rank, 2)))

		weightedScore := (croot.RivalsRank - 20) * math.Pow(math.E, expo)

		Rank247 += (weightedScore)
	}

	return Rank247
}

func GetESPNTeamRanking(rp structs.RecruitingTeamProfile, signedCroots []structs.Recruit) float64 {

	var espnRank float64 = 0

	for _, croot := range signedCroots {
		espnRank += croot.ESPNRank
	}

	return espnRank
}

func GetRivalsTeamRanking(rp structs.RecruitingTeamProfile, signedCroots []structs.Recruit) float64 {

	var rivalsRank float64 = 0

	for _, croot := range signedCroots {
		rivalsRank += croot.RivalsRank
	}

	return rivalsRank
}
