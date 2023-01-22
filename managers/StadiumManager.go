package managers

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetAllStadiums() []structs.Stadium {
	db := dbprovider.GetInstance().GetDB()

	var stadiums []structs.Stadium

	db.Find(&stadiums)

	return stadiums
}

func GetStadiumByStadiumID(id string) structs.Stadium {
	db := dbprovider.GetInstance().GetDB()

	var stadium structs.Stadium

	db.Where("id = ?", id).Find(&stadium)

	return stadium
}
