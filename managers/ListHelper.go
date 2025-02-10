package managers

import "github.com/CalebRose/SimFBA/structs"

func MakeCollegeInjuryList(players []structs.CollegePlayer) []structs.CollegePlayer {
	injuryList := []structs.CollegePlayer{}

	for _, p := range players {
		if p.IsInjured {
			injuryList = append(injuryList, p)
		}
	}
	return injuryList
}

func MakeCollegePortalList(players []structs.CollegePlayer) []structs.CollegePlayer {
	portalList := []structs.CollegePlayer{}

	for _, p := range players {
		if p.TransferStatus > 0 {
			portalList = append(portalList, p)
		}
	}
	return portalList
}

func MakeProInjuryList(players []structs.NFLPlayer) []structs.NFLPlayer {
	injuryList := []structs.NFLPlayer{}

	for _, p := range players {
		if p.IsInjured {
			injuryList = append(injuryList, p)
		}
	}
	return injuryList
}
