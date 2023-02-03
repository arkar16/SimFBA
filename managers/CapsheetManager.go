package managers

import (
	"fmt"
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

		for _, player := range players {
			contract := player.Contract

			Capsheet.AddContractToCapsheet(contract)
		}

		db.Save(&Capsheet)
	}
}
