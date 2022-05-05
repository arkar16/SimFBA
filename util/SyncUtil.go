package util

import "github.com/CalebRose/SimFBA/structs"

func GetWinsAndLossesForCollegeGames(games []structs.CollegeGame, TeamID int, ConferenceCheck bool) (int, int) {
	wins := 0
	losses := 0

	for _, game := range games {
		if ConferenceCheck && !game.IsConference {
			continue
		}
		if (game.HomeTeamID == TeamID && game.HomeTeamScore > game.AwayTeamScore) ||
			(game.AwayTeamID == TeamID && game.AwayTeamScore > game.HomeTeamScore) {
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
