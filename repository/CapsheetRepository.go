package repository

import (
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func FindAllNFLCapsheets() []structs.NFLCapsheet {
	var capsheets []structs.NFLCapsheet
	db := dbprovider.GetInstance().GetDB()
	err := db.Find(&capsheets).Error
	if err != nil {
		log.Fatal(err)
	}
	return capsheets
}
