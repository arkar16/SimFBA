package managers

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetCapsheetByTeamID(TeamID string) structs.NFLCapsheet {
	db := dbprovider.GetInstance().GetDB()

	capSheet := structs.NFLCapsheet{}

	db.Where("nfl_team_id = ?", TeamID).Find(&capSheet)

	return capSheet
}
