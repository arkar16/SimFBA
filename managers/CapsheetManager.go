package managers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetCapsheetByTeamID(TeamID string) structs.NFLCapsheet {
	db := dbprovider.GetInstance().GetDB()

	capSheet := structs.NFLCapsheet{}

	err := db.Where("nfl_team_id = ?", TeamID).Find(&capSheet).Error
	if err != nil {
		fmt.Println("Could not find capsheet, returning new one")
		return structs.NFLCapsheet{}
	}

	return capSheet
}

func AllocateCapsheets() {
	db := dbprovider.GetInstance().GetDB()

	teams := GetAllNFLTeams()

	for _, team := range teams {
		TeamID := strconv.Itoa(int(team.ID))

		players := GetNFLPlayersWithContractsByTeamID(TeamID)

		Capsheet := GetCapsheetByTeamID(TeamID)

		if Capsheet.ID == 0 {
			Capsheet.AssignCapsheet(team.ID)
		}

		Capsheet.ResetCapsheet()

		for _, player := range players {
			contract := player.Contract

			Capsheet.AddContractToCapsheet(contract)
		}

		db.Save(&Capsheet)
	}
}

func GetContractByPlayerID(PlayerID string) structs.NFLContract {
	db := dbprovider.GetInstance().GetDB()

	contract := structs.NFLContract{}

	err := db.Where("nfl_player_id = ? AND is_active = ?", PlayerID, true).Find(&contract).Error
	if err != nil {
		log.Fatalln("Could not find active contract for player" + PlayerID)
	}

	return contract
}

func GetAllContracts() []structs.NFLContract {
	db := dbprovider.GetInstance().GetDB()

	contracts := []structs.NFLContract{}

	err := db.Where("is_active = ?", true).Find(&contracts).Error
	if err != nil {
		log.Fatalln("Could not find all active contracts")
	}

	return contracts
}

func CalculateContractValues() {
	db := dbprovider.GetInstance().GetDB()

	contracts := GetAllContracts()

	for _, c := range contracts {
		c.CalculateContract()
		db.Save(&c)
	}
}
