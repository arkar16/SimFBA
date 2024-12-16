package managers

import "github.com/CalebRose/SimFBA/structs"

func MapNFLPlayers(nflPlayers []structs.NFLPlayer) map[uint]structs.NFLPlayer {
	playerMap := make(map[uint]structs.NFLPlayer)

	for _, p := range nflPlayers {
		playerMap[p.ID] = p
	}

	return playerMap
}
