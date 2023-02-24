package managers

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func ImportRecruitAICSV() {
	db := dbprovider.GetInstance().GetDB()
	completedCrootPath := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\FCS_Croot_Weekly_Signings.csv"
	aiPoolPath := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\2022_Croot_Class_AI.csv"
	crootMap := make(map[string][]string)
	f, err := os.Open(completedCrootPath)
	if err != nil {
		log.Fatal("Unable to read input file "+completedCrootPath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	croots, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+completedCrootPath, err)
	}

	for idx, record := range croots {
		if idx == 0 {
			continue
		}
		// Add recruit to map
		crootMap[record[0]] = record
		id := util.ConvertStringToInt(record[0])
		teamID := util.ConvertStringToInt(record[18])
		points := util.ConvertStringToInt(record[21])
		if points <= 0 {
			points = 1
		}

		if teamID > 0 {
			recruitProfile := structs.RecruitPlayerProfile{
				RecruitID:        id,
				IsSigned:         true,
				Scholarship:      false,
				TotalPoints:      float64(points),
				ProfileID:        teamID,
				TeamAbbreviation: record[19],
				SeasonID:         2,
			}

			db.Create(&recruitProfile)
			recruit := GetCollegeRecruitByRecruitID(record[0])
			// Since this croot is not in the DB yet, they will be added later
			if recruit.ID == 0 {
				continue
			}
			recruit.AssignCollege(recruitProfile.TeamAbbreviation)
			recruit.UpdateTeamID(teamID)
			recruit.UpdateSigningStatus()

			db.Save(&recruit)
		}
	}

	fmt.Println("NOW GET THE BIG ONE")

	poolFile, err := os.Open(aiPoolPath)
	if err != nil {
		log.Fatal("Unable to read input file "+aiPoolPath, err)
	}
	defer f.Close()

	csvReader = csv.NewReader(poolFile)
	croots, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+aiPoolPath, err)
	}

	for idx, croot := range croots {
		if idx == 0 {
			continue
		}
		recruit := GetCollegeRecruitByRecruitID(croot[0])
		if recruit.ID == 0 {
			// Create the record
			idStr := croot[0]
			id := util.ConvertStringToInt(croot[0])
			// If for some reason a recruit was not included in the map, then they don't have a team. Skip over for now.

			var mapRecord []string
			if len(crootMap[idStr]) > 0 {
				// Retrieve for the Team Info
				mapRecord = crootMap[idStr]
			}

			teamIDStr := ""
			abbr := ""
			if len(mapRecord) > 0 {
				teamIDStr = mapRecord[18]
				abbr = mapRecord[20]
			}

			teamID := util.ConvertStringToInt(teamIDStr)
			isSigned := false
			if teamID > 0 {
				isSigned = true
			}

			// Attributes

			base := structs.BasePlayer{
				FirstName:      croot[1],
				LastName:       croot[2],
				Stars:          util.ConvertStringToInt(croot[3]),
				Position:       croot[4],
				Archetype:      croot[5],
				Overall:        util.ConvertStringToInt(croot[6]),
				Height:         util.ConvertStringToInt(croot[7]),
				Weight:         util.ConvertStringToInt(croot[8]),
				Carrying:       util.ConvertStringToInt(croot[12]),
				Agility:        util.ConvertStringToInt(croot[13]),
				Catching:       util.ConvertStringToInt(croot[14]),
				ZoneCoverage:   util.ConvertStringToInt(croot[15]),
				ManCoverage:    util.ConvertStringToInt(croot[16]),
				FootballIQ:     util.ConvertStringToInt(croot[17]),
				KickAccuracy:   util.ConvertStringToInt(croot[18]),
				KickPower:      util.ConvertStringToInt(croot[19]),
				PassBlock:      util.ConvertStringToInt(croot[20]),
				PassRush:       util.ConvertStringToInt(croot[21]),
				PuntAccuracy:   util.ConvertStringToInt(croot[22]),
				PuntPower:      util.ConvertStringToInt(croot[23]),
				RouteRunning:   util.ConvertStringToInt(croot[24]),
				RunBlock:       util.ConvertStringToInt(croot[25]),
				RunDefense:     util.ConvertStringToInt(croot[26]),
				Speed:          util.ConvertStringToInt(croot[27]),
				Strength:       util.ConvertStringToInt(croot[28]),
				Tackle:         util.ConvertStringToInt(croot[29]),
				ThrowPower:     util.ConvertStringToInt(croot[30]),
				ThrowAccuracy:  util.ConvertStringToInt(croot[31]),
				Injury:         util.ConvertStringToInt(croot[32]),
				Stamina:        util.ConvertStringToInt(croot[33]),
				Discipline:     util.ConvertStringToInt(croot[34]),
				AcademicBias:   croot[35],
				FreeAgency:     croot[36],
				Personality:    croot[37],
				RecruitingBias: croot[38],
				WorkEthic:      croot[39],
				Progression:    util.ConvertStringToInt(croot[40]),
				PotentialGrade: croot[41],
				Age:            18,
			}

			r := structs.Recruit{
				BasePlayer: base,
				PlayerID:   id,
				TeamID:     teamID,
				HighSchool: croot[9],
				City:       croot[10],
				State:      croot[11],
				IsSigned:   isSigned,
				College:    abbr,
			}

			r.AssignID(id)

			db.Create(&r)
		} else {
			// This recruit is already in the DB
			continue
		}
	}
}

func GetMissingRecruitingClasses() {
	db := dbprovider.GetInstance().GetDB()

	teams := GetAllCollegeTeams()
	for _, team := range teams {
		if team.ID < 131 || team.ID > 134 {
			continue
		}

		count := 0
		limit := 20

		positionMap := make(map[string]int)

		unsignedPlayers := GetLeftoverRecruits()

		for _, croot := range unsignedPlayers {
			if count == limit {
				break
			}
			if (croot.Position != "K" && croot.Position != "P") && positionMap[croot.Position] > 1 || (croot.Position == "K" || croot.Position == "P" && positionMap[croot.Position] > 0) {
				continue
			}
			positionMap[croot.Position] += 1
			count++
			collegePlayer := (structs.CollegePlayer)(croot)
			collegePlayer.AssignTeamValues(team)

			db.Create(&collegePlayer)
			db.Delete(&croot)
		}
	}
}

func GetLeftoverRecruits() []structs.UnsignedPlayer {
	db := dbprovider.GetInstance().GetDB()
	var unsignedPlayers []structs.UnsignedPlayer

	db.Where("year = 1").Find(&unsignedPlayers)

	return unsignedPlayers
}

func ImportNFLPlayersCSV() {
	db := dbprovider.GetInstance().GetDB()
	playerPath := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\NFL_Progressed.csv"

	nflCSV := util.ReadCSV(playerPath)

	for idx, row := range nflCSV {
		if idx < 2 {
			continue
		}

		playerID := util.ConvertStringToInt(row[0])
		academic := row[35]
		fa := row[36]
		personality := row[37]
		recruit := row[38]
		we := row[39]
		progression := util.ConvertStringToInt(row[40])

		gp := GetGlobalPlayerRecord(row[0])
		if gp.ID == 0 {
			player := structs.Player{
				NFLPlayerID: playerID,
			}
			player.AssignID(uint(playerID))

			db.Create(&player)
		}

		NFLPlayerRecord := GetNFLPlayerRecord(row[0])
		if NFLPlayerRecord.ID == 0 {
			log.Fatalln("Something is wrong, this player was not uploaded.")
		}
		NFLPlayerRecord.AssignMissingValues(progression, academic, fa, personality, recruit, we)

		db.Save(&NFLPlayerRecord)
	}
}

func ImportMinimumFAValues() {
	db := dbprovider.GetInstance().GetDB()
	playerPath := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\2023_Free_Agency_expected_Values_CSV.csv"

	nflCSV := util.ReadCSV(playerPath)

	for idx, row := range nflCSV {
		if idx < 1 {
			continue
		}

		playerID := row[0]
		value := util.ConvertStringToFloat(row[5])

		NFLPlayerRecord := GetNFLPlayerRecord(playerID)
		if NFLPlayerRecord.ID == 0 {
			log.Fatalln("Something is wrong, this player was not uploaded.")
		}

		NFLPlayerRecord.AssignMinimumValue(value)

		db.Save(&NFLPlayerRecord)
	}
}

func ImportWorkEthic() {
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	db := dbprovider.GetInstance().GetDB()

	nflPlayers := GetAllNFLPlayers()

	for _, p := range nflPlayers {
		WorkEthic := util.GetWorkEthic()
		if p.ID == 10 {
			FreeAgency := "Highly Unlikely to Play for the Miami Dolphins."
			Personality := "Worships Himself"
			p.AssignPersonality(Personality)
			p.AssignFreeAgency(FreeAgency)
		}

		p.AssignWorkEthic(WorkEthic)

		db.Save(&p)
	}
}

func RetireAndFreeAgentPlayers() {
	db := dbprovider.GetInstance().GetDB()

	nflPlayers := GetAllNFLPlayers()

	for _, record := range nflPlayers {

		if !record.IsActive {
			retiredPlayerRecord := (structs.NFLRetiredPlayer)(record)

			db.Create(&retiredPlayerRecord)
			db.Delete(&record)
			continue
		}

		if record.TeamID == 0 {
			record.ToggleIsFreeAgent()
			db.Save(&record)
		}
	}
}
