package managers

import (
	"log"
	"strconv"

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

// GetStandingsByConferenceIDAndSeasonID
func GetStandingsByTeamIDAndSeasonID(TeamID string, seasonID string) structs.CollegeStandings {
	var standings structs.CollegeStandings
	db := dbprovider.GetInstance().GetDB()
	err := db.Where("team_id = ? AND season_id = ?", TeamID, seasonID).
		Find(&standings).Error
	if err != nil {
		log.Fatal(err)
	}
	return standings
}

func UpdateStandings(ts structs.Timestamp) {
	db := dbprovider.GetInstance().GetDB()

	games := GetCollegeGamesByWeekIdAndSeasonID(strconv.Itoa(ts.CollegeWeekID), strconv.Itoa(ts.CollegeSeasonID))

	for i := 0; i < len(games); i++ {
		HomeID := games[i].HomeTeamID
		AwayID := games[i].AwayTeamID

		homeStandings := GetStandingsByTeamIDAndSeasonID(strconv.Itoa(HomeID), strconv.Itoa(ts.CollegeSeasonID))
		awayStandings := GetStandingsByTeamIDAndSeasonID(strconv.Itoa(AwayID), strconv.Itoa(ts.CollegeSeasonID))

		homeStandings.UpdateCollegeStandings(games[i])
		awayStandings.UpdateCollegeStandings(games[i])

		err := db.Save(&homeStandings).Error
		if err != nil {
			log.Panicln("Could not save standings for team " + strconv.Itoa(HomeID))
		}

		err = db.Save(&awayStandings).Error
		if err != nil {
			log.Panicln("Could not save standings for team " + strconv.Itoa(AwayID))
		}
	}
}
