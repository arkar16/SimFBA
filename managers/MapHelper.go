package managers

import "github.com/CalebRose/SimFBA/structs"

func MapNFLPlayers(nflPlayers []structs.NFLPlayer) map[uint]structs.NFLPlayer {
	playerMap := make(map[uint]structs.NFLPlayer)

	for _, p := range nflPlayers {
		playerMap[p.ID] = p
	}

	return playerMap
}

func MakeCollegePlayerMap(players []structs.CollegePlayer) map[uint]structs.CollegePlayer {
	playerMap := make(map[uint]structs.CollegePlayer)

	for _, p := range players {
		playerMap[p.ID] = p
	}

	return playerMap
}

func MakeCollegePlayerMapByTeamID(players []structs.CollegePlayer, excludeUnsigned bool) map[uint][]structs.CollegePlayer {
	playerMap := make(map[uint][]structs.CollegePlayer)

	for _, p := range players {
		if p.TeamID == 0 && excludeUnsigned {
			continue
		}
		if len(playerMap[uint(p.TeamID)]) > 0 {
			playerMap[uint(p.TeamID)] = append(playerMap[uint(p.TeamID)], p)
		} else {
			playerMap[uint(p.TeamID)] = []structs.CollegePlayer{p}
		}
	}

	return playerMap
}

func MakeNFLPlayerMapByTeamID(players []structs.NFLPlayer, excludeFAs bool) map[uint][]structs.NFLPlayer {
	playerMap := make(map[uint][]structs.NFLPlayer)

	for _, p := range players {
		if p.TeamID == 0 && excludeFAs {
			continue
		}
		if len(playerMap[uint(p.TeamID)]) > 0 {
			playerMap[uint(p.TeamID)] = append(playerMap[uint(p.TeamID)], p)
		} else {
			playerMap[uint(p.TeamID)] = []structs.NFLPlayer{p}
		}
	}

	return playerMap
}

func MakeCollegeDepthChartMap(dcs []structs.CollegeTeamDepthChart) map[uint]structs.CollegeTeamDepthChart {
	dcMap := make(map[uint]structs.CollegeTeamDepthChart)

	for _, dc := range dcs {
		dcMap[uint(dc.TeamID)] = dc
	}

	return dcMap
}

func MakeNFLDepthChartMap(dcs []structs.NFLDepthChart) map[uint]structs.NFLDepthChart {
	dcMap := make(map[uint]structs.NFLDepthChart)

	for _, dc := range dcs {
		dcMap[uint(dc.TeamID)] = dc
	}

	return dcMap
}
