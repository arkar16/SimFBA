package managers

import (
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// Gets all Current Season and Beyond Draft Picks
func GetDraftPicksByTeamID(TeamID string) []structs.NFLDraftPick {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()

	seasonID := strconv.Itoa(int(ts.NFLSeasonID))
	var picks []structs.NFLDraftPick

	db.Where("team_id = ? AND season_id >= ?", TeamID, seasonID).Find(&picks)

	return picks
}

// Gets all Current Season and Beyond Draft Picks
func GetDraftPickByDraftPickID(DraftPickID string) structs.NFLDraftPick {
	db := dbprovider.GetInstance().GetDB()

	var pick structs.NFLDraftPick

	db.Where("id = ?", DraftPickID).Find(&pick)

	return pick
}
