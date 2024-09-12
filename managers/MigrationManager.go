package managers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/repository"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func MigrateHistoricPlayersToNFLDraftees() {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()
	SeasonID := strconv.Itoa(ts.CollegeSeasonID)
	historicPlayers := []structs.HistoricCollegePlayer{}
	draftees := []models.NFLDraftee{}
	targetDate, err := time.Parse("2006-01-02", "2024-09-10")
	if err != nil {
		log.Panic(err)
	}
	err = db.Where("created_at > ?", targetDate).Find(&historicPlayers).Error
	if err != nil {
		log.Panic(err)
	}

	for _, p := range historicPlayers {
		if p.CreatedAt.Before(targetDate) {
			continue
		}

		grad := (structs.CollegePlayer)(p)

		draftee := models.NFLDraftee{}
		draftee.Map(grad)
		// Map New Progression value for NFL
		newProgression := util.GenerateNFLPotential(grad.Progression)
		newPotentialGrade := util.GetWeightedPotentialGrade(newProgression)
		draftee.MapProgression(newProgression, newPotentialGrade)

		if draftee.Position == "RB" {
			draftee = BoomBustDraftee(draftee, SeasonID, 31, true)
		}

		draftee.GetLetterGrades()

		/*
			Boom/Bust Function
		*/
		tier := 1
		isBoom := false
		enableBoomBust := false
		boomBustStatus := "None"
		tierRoll := util.GenerateIntFromRange(1, 10)
		diceRoll := util.GenerateIntFromRange(1, 20)

		if tierRoll > 7 && tierRoll < 10 {
			tier = 2
		} else if tierRoll > 9 {
			tier = 3
		}

		// Generate Tier
		if diceRoll == 1 {
			boomBustStatus = "Bust"
			enableBoomBust = true
			// Bust
			fmt.Println(draftee.FirstName + " " + draftee.LastName + " has BUSTED!")
			draftee.AssignBoomBustStatus(boomBustStatus)

		} else if diceRoll == 20 {
			enableBoomBust = true
			// Boom
			fmt.Println(draftee.FirstName + " " + draftee.LastName + " has BOOMED!")
			boomBustStatus = "Boom"
			isBoom = true
			draftee.AssignBoomBustStatus(boomBustStatus)
		} else {
			tier = 0
		}
		if enableBoomBust {
			for i := 0; i < tier; i++ {
				draftee = BoomBustDraftee(draftee, SeasonID, 51, isBoom)
			}
		}

		draftees = append(draftees, draftee)
	}
	repository.CreateNFLDrafteesSafely(db, draftees, 500)
}
