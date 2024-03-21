package repository

import (
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
)

func SaveCFBGameRecord(gameRecord structs.CollegeGame, db *gorm.DB) {
	err := db.Save(&gameRecord).Error
	if err != nil {
		log.Panicln("Could not save Game " + strconv.Itoa(int(gameRecord.ID)) + "Between " + gameRecord.HomeTeam + " and " + gameRecord.AwayTeam)
	}
}

func SaveNFLGameRecord(gameRecord structs.NFLGame, db *gorm.DB) {
	err := db.Save(&gameRecord).Error
	if err != nil {
		log.Panicln("Could not save Game " + strconv.Itoa(int(gameRecord.ID)) + "Between " + gameRecord.HomeTeam + " and " + gameRecord.AwayTeam)
	}
}
