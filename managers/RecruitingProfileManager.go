package managers

import (
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// GetRecruitingProfileByTeamID
func GetRecruitingProfileByTeamID(TeamID string) structs.RecruitingTeamProfile {
	db := dbprovider.GetInstance().GetDB()

	var profile structs.RecruitingTeamProfile

	err := db.Where("team_id = ?", TeamID).Find(&profile).Error
	if err != nil {
		log.Fatal(err)
	}

	return profile
}
