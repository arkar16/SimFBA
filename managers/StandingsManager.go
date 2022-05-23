package managers

import (
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// GetStandingsByConferenceIDAndSeasonID
func GetStandingsByConferenceIDAndSeasonID(conferenceID string, seasonID string) []structs.CollegeStandings {
	var standings []structs.CollegeStandings
	db := dbprovider.GetInstance().GetDB()
	err := db.Where("conference_id = ? AND season_id = ?", conferenceID, seasonID).Order("total_wins desc").
		Find(&standings).Error
	if err != nil {
		log.Fatal(err)
	}
	return standings
}
