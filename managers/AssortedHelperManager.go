package managers

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func GetHeismanScore(cp structs.CollegePlayer, weightMap map[string]float64, homeTeamMapper map[int]string, games []structs.CollegeGame) float64 {
	var score float64 = 0

	homeTeam := homeTeamMapper[cp.TeamID]

	stats := cp.Stats

	var totalMod float64 = 0

	for idx, stat := range stats {
		if idx > 12 {
			continue
		}
		var statScore float64 = 0

		opposingTeam := stat.OpposingTeam
		statScore += (float64(stat.PassingYards) * 0.069)
		statScore += (float64(stat.PassingTDs) * 4)
		statScore -= (float64(stat.Interceptions) * 2.25)
		statScore -= (float64(stat.Sacks) * 2.25)
		if cp.Position == "RB" || cp.Position == "FB" {
			statScore += (float64(stat.RushingYards) * 0.1)
			statScore += (float64(stat.RushingTDs) * 6)
		} else {
			statScore += (float64(stat.RushingYards) * 0.1)
			statScore += (float64(stat.RushingTDs) * 5)
		}

		drops := stat.Targets - stat.Catches
		if cp.Position == "WR" || cp.Position == "TE" {
			statScore += (float64(stat.Catches) * 0.525)
			statScore += (float64(stat.ReceivingYards) * 0.1125)
			statScore += (float64(stat.ReceivingTDs) * 6)
			statScore -= float64(drops) * 0.75
		} else {
			statScore += (float64(stat.Catches) * 0.25)
			statScore += (float64(stat.ReceivingYards) * 0.05)
			statScore += (float64(stat.ReceivingTDs) * 4)
			statScore -= float64(drops) * 0.75
		}
		statScore -= (float64(stat.Fumbles) * 6)

		statScore += (float64(stat.SoloTackles) * 1)
		statScore += (float64(stat.STSoloTackles) * 1)
		statScore += (float64(stat.AssistedTackles) * .9)
		statScore += (float64(stat.STAssistedTackles) * .9)

		statScore += (float64(stat.TacklesForLoss) * 6.25)
		statScore += (float64(stat.SacksMade) * 7.125)
		statScore += (float64(stat.PassDeflections) * 6.25)

		statScore += (float64(stat.ForcedFumbles) * 8)
		statScore += (float64(stat.RecoveredFumbles) * 6)
		statScore += (float64(stat.InterceptionsCaught) * 15)
		statScore += (float64(stat.PuntsBlocked) * 10)

		statScore += (float64(stat.Safeties) * 10)
		statScore += (float64(stat.DefensiveTDs) * 20)
		statScore += (float64(stat.KickReturnTDs) * 14)
		statScore += (float64(stat.PuntReturnTDs) * 14)
		statScore += (float64(stat.FGBlocked) * 12)
		statScore += (float64(stat.FGMade) * 0.3)
		statScore += (float64(stat.ExtraPointsMade) * 0.1)

		game := GetCollegeGameStructByGameID(games, stat.GameID)

		opposingTeamWeight := weightMap[opposingTeam]

		if (game.HomeTeamWin && cp.TeamID != game.HomeTeamID) || (game.AwayTeamWin && cp.TeamID != game.AwayTeamID) {
			opposingTeamWeight *= -.4125
		}

		totalMod += opposingTeamWeight

		// statScore *= weightMap[opposingTeam]

		// statScore *= weightMap[homeTeam]

		score += statScore
	}

	score = (score / float64(len(stats))) * (float64(len(stats)) / 12)

	score = score * (1 + totalMod + weightMap[homeTeam])

	return score
}

func GetCollegeGameStructByGameID(games []structs.CollegeGame, gameID int) structs.CollegeGame {
	for _, game := range games {
		if int(game.ID) == gameID {
			return game
		}
	}

	return structs.CollegeGame{}
}

func GetCollegePlayerIDsBySeasonStats(cps []structs.CollegePlayerSeasonStats) []string {
	var list []string

	for _, cp := range cps {
		list = append(list, strconv.Itoa(int(cp.CollegePlayerID)))
	}

	return list
}

func GetNFLPlayerIDsBySeasonStats(cps []structs.NFLPlayerSeasonStats) []string {
	var list []string

	for _, cp := range cps {
		list = append(list, strconv.Itoa(int(cp.NFLPlayerID)))
	}

	return list
}

func GetCollegePlayerIDs(cps []structs.CollegePlayerStats) []string {
	var list []string

	for _, cp := range cps {
		list = append(list, strconv.Itoa(int(cp.CollegePlayerID)))
	}

	return list
}

func GetOffensiveDefaultSchemes() map[string]structs.OffensiveFormation {
	path := filepath.Join(os.Getenv("ROOT"), "data", "defaultOffensiveSchemes.json")
	content := util.ReadJson(path)

	var payload map[string]structs.OffensiveFormation

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalln("Error during unmarshal: ", err)
	}

	return payload
}

func GetDefensiveDefaultSchemes() map[string]map[string]structs.DefensiveFormation {
	path := filepath.Join(os.Getenv("ROOT"), "data", "defaultDefensiveSchemes.json")
	content := util.ReadJson(path)

	var payload map[string]map[string]structs.DefensiveFormation

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalln("Error during unmarshal: ", err)
	}

	return payload
}

func IsAITeamContendingForCroot(profiles []structs.RecruitPlayerProfile) float64 {
	if len(profiles) == 0 {
		return 0
	}
	var leadingVal float64 = 0
	for _, profile := range profiles {
		if profile.TotalPoints != 0 && profile.TotalPoints > float64(leadingVal) {
			leadingVal = profile.TotalPoints
		}
	}

	return leadingVal
}

func GetWinsAndLossesForCollegeGames(games []structs.CollegeGame, TeamID int, ConferenceCheck bool) (int, int) {
	wins := 0
	losses := 0

	for _, game := range games {
		if !game.GameComplete {
			continue
		}
		if ConferenceCheck && !game.IsConference {
			continue
		}
		if (game.HomeTeamID == TeamID && game.HomeTeamWin) ||
			(game.AwayTeamID == TeamID && game.AwayTeamWin) {
			wins += 1
		} else {
			losses += 1
		}
	}

	return wins, losses
}

func GetConferenceChampionshipWeight(games []structs.CollegeGame, TeamID int) float64 {
	var weight float64 = 0

	for _, game := range games {
		if !game.IsConference {
			continue
		}
		if (game.HomeTeamID == TeamID && game.HomeTeamScore > game.AwayTeamScore) ||
			(game.AwayTeamID == TeamID && game.AwayTeamScore > game.HomeTeamScore) {
			weight = 1
		} else {
			weight = 0.5
		}
	}

	return weight
}

func GetPostSeasonWeight(games []structs.CollegeGame, TeamID int) float64 {
	for _, game := range games {
		if !game.IsPlayoffGame || !game.IsBowlGame {
			continue
		}
		return 1
	}
	return 0
}

func FilterOutRecruitingProfile(profiles []structs.RecruitPlayerProfile, ID int) []structs.RecruitPlayerProfile {
	var rp []structs.RecruitPlayerProfile

	for _, profile := range profiles {
		if profile.ProfileID != ID {
			rp = append(rp, profile)
		}
	}

	return rp
}

func IsAITeamContendingForPortalPlayer(profiles []structs.TransferPortalProfile) int {
	if len(profiles) == 0 {
		return 0
	}
	leadingVal := 0
	for _, profile := range profiles {
		if profile.TotalPoints != 0 && profile.TotalPoints > float64(leadingVal) {
			leadingVal = int(profile.TotalPoints)
		}
	}

	return leadingVal
}

func FilterOutPortalProfile(profiles []structs.TransferPortalProfile, ID uint) []structs.TransferPortalProfile {
	var rp []structs.TransferPortalProfile

	for _, profile := range profiles {
		if profile.ID != ID {
			rp = append(rp, profile)
		}
	}

	return rp
}
