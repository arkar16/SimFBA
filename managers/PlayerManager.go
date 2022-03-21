package managers

import (
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

// GetAllPlayers - Returns all player reference records
func GetAllPlayers() []structs.Player {
	db := dbprovider.GetInstance().GetDB()

	var players []structs.Player

	db.Find(&players)

	return players
}

func GetAllCollegePlayers() []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayers []structs.CollegePlayer

	db.Find(&CollegePlayers)

	return CollegePlayers
}

func GetAllCollegePlayersByTeamId(TeamID string) []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayers []structs.CollegePlayer

	db.Order("overall desc").Where("team_id = ?", TeamID).Find(&CollegePlayers)

	return CollegePlayers
}

func GetAllCollegePlayersByTeamIdWithoutRedshirts(TeamID string) []structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayers []structs.CollegePlayer

	db.Where("team_id = ? AND is_redshirting = false", TeamID).Find(&CollegePlayers)

	return CollegePlayers
}

func GetCollegePlayerByCollegePlayerId(CollegePlayerId string) structs.CollegePlayer {
	db := dbprovider.GetInstance().GetDB()

	var CollegePlayer structs.CollegePlayer

	db.Where("id = ?", CollegePlayerId).Find(&CollegePlayer)

	return CollegePlayer
}

func CreateRecruit(cr structs.CreateRecruitDTO) {
	db := dbprovider.GetInstance().GetDB()

	recruit := structs.Recruit{
		BasePlayer:  cr.BasePlayer,
		HighSchool:  cr.HighSchool,
		City:        cr.City,
		State:       cr.State,
		AffinityOne: cr.AffinityOne,
		AffinityTwo: cr.AffinityTwo,
	}

	err := db.Create(&recruit).Error
	if err != nil {
		log.Fatalln("Recruit already exists in DB.")
	}

	player := structs.Player{
		RecruitID: int(recruit.ID),
	}

	err = db.Create(&player).Error
	if err != nil {
		log.Fatalln("Cannot save player in DB.")
	}

	recruit.UpdatePlayerID(int(player.ID))

	err = db.Save(&recruit).Error
	if err != nil {
		log.Fatalln("Could not save recruit object")
	}
}

func UpdateRecruit(r structs.Recruit) {
	db := dbprovider.GetInstance().GetDB()
	err := db.Save(&r).Error
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateCollegePlayer(cp structs.CollegePlayer) {
	db := dbprovider.GetInstance().GetDB()
	err := db.Save(&cp).Error
	if err != nil {
		log.Fatal(err)
	}
}

func SetRedshirtStatusForPlayer(playerId string) structs.CollegePlayer {
	player := GetCollegePlayerByCollegePlayerId(playerId)

	player.SetRedshirtingStatus()

	UpdateCollegePlayer(player)

	return player
}

func GetAllNFLDraftees() []structs.NFLDraftee {
	db := dbprovider.GetInstance().GetDB()

	var NFLDraftees []structs.NFLDraftee

	db.Order("overall desc").Find(&NFLDraftees)

	return NFLDraftees
}
