package managers

import (
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
)

// For Temporary Functions that will have no general place in the simulation

func RemoveContractsFromPracticeSquadPlayers() {
	db := dbprovider.GetInstance().GetDB()

	practiceSquadPlayers := GetAllPracticeSquadPlayers()

	for _, p := range practiceSquadPlayers {
		playerID := strconv.Itoa(int(p.ID))
		teamID := strconv.Itoa(p.TeamID)
		contract := GetContractByPlayerID(playerID)
		contract.DeactivateContract()
		db.Save(&contract)

		teamCapsheet := GetCapsheetByTeamID(teamID)
		teamCapsheet.SubtractFromCapsheet(contract)
		db.Save(&teamCapsheet)
	}
}
