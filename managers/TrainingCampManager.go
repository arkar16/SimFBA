package managers

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/repository"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func UploadTrainingCampCSV() {
	db := dbprovider.GetInstance().GetDB()

	teamPath := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\2025\\2025_simnfl_rookie_camp_results.csv"

	resultsCSV := util.ReadCSV(teamPath)

	nflPlayerMap := GetAllNFLPlayersMap()

	for idx, row := range resultsCSV {
		if idx == 0 {
			continue
		}
		idStr := row[0]
		id := uint(util.ConvertStringToInt(idStr))

		nflPlayer := nflPlayerMap[id]

		changedAttrs := structs.CollegePlayerProgressions{
			FootballIQ:      util.ConvertStringToInt(row[2]),
			Speed:           util.ConvertStringToInt(row[3]),
			Carrying:        util.ConvertStringToInt(row[4]),
			Agility:         util.ConvertStringToInt(row[5]),
			Catching:        util.ConvertStringToInt(row[6]),
			RouteRunning:    util.ConvertStringToInt(row[7]),
			ZoneCoverage:    util.ConvertStringToInt(row[8]),
			ManCoverage:     util.ConvertStringToInt(row[9]),
			Strength:        util.ConvertStringToInt(row[10]),
			Tackle:          util.ConvertStringToInt(row[11]),
			PassBlock:       util.ConvertStringToInt(row[12]),
			RunBlock:        util.ConvertStringToInt(row[13]),
			PassRush:        util.ConvertStringToInt(row[14]),
			RunDefense:      util.ConvertStringToInt(row[15]),
			ThrowPower:      util.ConvertStringToInt(row[16]),
			ThrowAccuracy:   util.ConvertStringToInt(row[17]),
			KickAccuracy:    util.ConvertStringToInt(row[18]),
			KickPower:       util.ConvertStringToInt(row[19]),
			PuntAccuracy:    util.ConvertStringToInt(row[20]),
			PuntPower:       util.ConvertStringToInt(row[21]),
			InjuryText:      row[22],
			WeeksOfRecovery: util.ConvertStringToInt(row[23]),
		}

		nflPlayer.ApplyTrainingCampInfo(changedAttrs)
		nflPlayer.GetOverall()

		repository.SaveNFLPlayer(nflPlayer, db)
	}
}
