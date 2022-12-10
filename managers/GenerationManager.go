package managers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func GenerateWalkOns() {
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	db := dbprovider.GetInstance().GetDB()
	AllTeams := GetRecruitingProfileForRecruitSync()
	count := 0
	attributeBlob := getAttributeBlob()

	firstNameMap, lastNameMap := getNameMaps()

	var lastPlayerRecord structs.Player

	err := db.Last(&lastPlayerRecord).Error
	if err != nil {
		log.Fatalln("Could not grab last player record from players table...")
	}

	newID := lastPlayerRecord.ID + 1

	for _, team := range AllTeams {
		if team.TotalCommitments == team.RecruitClassSize {
			continue
		}

		limit := team.RecruitClassSize - team.TotalCommitments
		positionList := []string{}
		id := strconv.Itoa(int(team.ID))

		// Get Team Needs
		teamNeeds := GetRecruitingNeeds(id)
		signedRecruits := GetSignedRecruitsByTeamProfileID(id)

		for _, recruit := range signedRecruits {
			if teamNeeds[recruit.Position] > 0 {
				teamNeeds[recruit.Position] -= 1
			}
		}

		// Get All Needed Positions into a list
		for k, v := range teamNeeds {
			i := v
			for i > 0 {
				positionList = append(positionList, k)
				i--
			}
		}

		// Randomize List
		rand.Shuffle(len(positionList), func(i, j int) {
			positionList[i], positionList[j] = positionList[j], positionList[i]
		})

		// Recruit Generation
		for _, pos := range positionList {
			if count >= limit {
				break
			}

			year := 1
			ethnicity := pickEthnicity()

			recruit := createRecruit(ethnicity, pos, year, firstNameMap[ethnicity], lastNameMap[ethnicity], newID, attributeBlob)

			recruit.AssignWalkon(team.TeamAbbreviation, int(team.ID), newID)

			recruitPlayerRecord := structs.RecruitPlayerProfile{
				ProfileID:        int(team.ID),
				RecruitID:        int(newID),
				IsSigned:         true,
				IsLocked:         true,
				TeamAbbreviation: team.TeamAbbreviation,
				SeasonID:         2,
				TotalPoints:      1,
			}

			playerRecord := structs.Player{
				RecruitID:       int(newID),
				CollegePlayerID: int(newID),
				NFLPlayerID:     int(newID),
			}
			playerRecord.AssignID(newID)
			count++

			db.Save(&recruit)
			db.Create(&recruit)
			db.Create(&recruitPlayerRecord)
		}
	}
}

func createRecruit(ethnicity string, position string, year int, firstNameList [][]string, lastNameList [][]string, id uint, blob map[string]map[string]map[string]map[string]interface{}) structs.Recruit {
	fName := getName(firstNameList)
	lName := getName(lastNameList)
	firstName := strings.Title(strings.ToLower(fName))
	lastName := strings.Title(strings.ToLower(lName))
	age := 18
	state := ""
	city := ""
	highSchool := ""

	archetype := getArchetype(position)
	stars := getStarRating()
	height := getAttributeValue(position, archetype, stars, "Height", blob)
	weight := getAttributeValue(position, archetype, stars, "Weight", blob)
	footballIQ := getAttributeValue(position, archetype, stars, "Football IQ", blob)
	speed := getAttributeValue(position, archetype, stars, "Speed", blob)
	agility := getAttributeValue(position, archetype, stars, "Agility", blob)
	carrying := getAttributeValue(position, archetype, stars, "Carrying", blob)
	catching := getAttributeValue(position, archetype, stars, "Catching", blob)
	routeRunning := getAttributeValue(position, archetype, stars, "Route Running", blob)
	zoneCoverage := getAttributeValue(position, archetype, stars, "Zone Coverage", blob)
	manCoverage := getAttributeValue(position, archetype, stars, "Man Coverage", blob)
	strength := getAttributeValue(position, archetype, stars, "Strength", blob)
	tackle := getAttributeValue(position, archetype, stars, "Tackle", blob)
	passBlock := getAttributeValue(position, archetype, stars, "Pass Block", blob)
	runBlock := getAttributeValue(position, archetype, stars, "Run Block", blob)
	passRush := getAttributeValue(position, archetype, stars, "Pass Rush", blob)
	runDefense := getAttributeValue(position, archetype, stars, "Run Defense", blob)
	throwPower := getAttributeValue(position, archetype, stars, "Throw Power", blob)
	throwAccuracy := getAttributeValue(position, archetype, stars, "Throw Accuracy", blob)
	kickAccuracy := getAttributeValue(position, archetype, stars, "Kick Accuracy", blob)
	kickPower := getAttributeValue(position, archetype, stars, "Kick Power", blob)
	puntAccuracy := getAttributeValue(position, archetype, stars, "Punt Accuracy", blob)
	puntPower := getAttributeValue(position, archetype, stars, "Punt Power", blob)
	injury := util.GenerateIntFromRange(10, 100)
	stamina := util.GenerateIntFromRange(10, 100)
	discipline := util.GenerateIntFromRange(10, 100)
	progression := util.GetProgressionRating()

	freeAgency := util.GetFreeAgencyBias()
	personality := util.GetPersonality()
	recruitingBias := util.GetRecruitingBias()
	workEthic := util.GetWorkEthic()
	academicBias := util.GetAcademicBias()
	potentialGrade := util.GetWeightedPotentialGrade(progression)

	basePlayer := structs.BasePlayer{
		FirstName:      firstName,
		LastName:       lastName,
		Age:            age,
		Stars:          stars,
		Height:         height,
		Weight:         weight,
		Stamina:        stamina,
		Injury:         injury,
		FootballIQ:     footballIQ,
		Speed:          speed,
		Carrying:       carrying,
		Agility:        agility,
		Catching:       catching,
		RouteRunning:   routeRunning,
		ZoneCoverage:   zoneCoverage,
		ManCoverage:    manCoverage,
		Strength:       strength,
		Tackle:         tackle,
		PassBlock:      passBlock,
		RunBlock:       runBlock,
		PassRush:       passRush,
		RunDefense:     runDefense,
		ThrowPower:     throwPower,
		ThrowAccuracy:  throwAccuracy,
		KickAccuracy:   kickAccuracy,
		KickPower:      kickPower,
		PuntAccuracy:   puntAccuracy,
		PuntPower:      puntPower,
		Progression:    progression,
		Discipline:     discipline,
		PotentialGrade: potentialGrade,
		FreeAgency:     freeAgency,
		Personality:    personality,
		RecruitingBias: recruitingBias,
		WorkEthic:      workEthic,
		AcademicBias:   academicBias,
	}

	return structs.Recruit{
		BasePlayer: basePlayer,
		City:       city,
		HighSchool: highSchool,
		State:      state,
		IsSigned:   true,
	}
}

func pickEthnicity() string {
	min := 0
	max := 10000
	num := util.GenerateIntFromRange(min, max)

	if num < 6000 {
		return "Caucasian"
	} else if num < 7800 {
		return "African"
	} else if num < 8900 {
		return "Hispanic"
	} else if num < 9975 {
		return "Asian"
	}
	return "NativeAmerican"
}

func getAttributeBlob() map[string]map[string]map[string]map[string]interface{} {
	path := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\attributeBlob.json"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload map[string]map[string]map[string]map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during unmarshal: ", err)
	}

	return payload
}

func getNameList(ethnicity string, isFirstName bool) [][]string {
	path := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimNBA\\data"
	var fileName string
	if ethnicity == "Caucasian" {
		if isFirstName {
			fileName = "FNameW.csv"
		} else {
			fileName = "LNameW.csv"
		}
	} else if ethnicity == "African" {
		if isFirstName {
			fileName = "FNameB.csv"
		} else {
			fileName = "LNameB.csv"
		}
	} else if ethnicity == "Asian" {
		if isFirstName {
			fileName = "FNameA.csv"
		} else {
			fileName = "LNameA.csv"
		}
	} else if ethnicity == "NativeAmerican" {
		if isFirstName {
			fileName = "FNameN.csv"
		} else {
			fileName = "LNameN.csv"
		}
	} else {
		if isFirstName {
			fileName = "FNameH.csv"
		} else {
			fileName = "LNameH.csv"
		}
	}
	path = path + "\\" + fileName
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}

	return records
}

func getNameMaps() (map[string][][]string, map[string][][]string) {
	CaucasianFirstNameList := getNameList("Caucasian", true)
	CaucasianLastNameList := getNameList("Caucasian", false)
	AfricanFirstNameList := getNameList("African", true)
	AfricanLastNameList := getNameList("African", false)
	HispanicFirstNameList := getNameList("Hispanic", true)
	HispanicLastNameList := getNameList("Hispanic", false)
	NativeFirstNameList := getNameList("NativeAmerican", true)
	NativeLastNameList := getNameList("NativeAmerican", false)
	AsianFirstNameList := getNameList("Asian", true)
	AsianLastNameList := getNameList("Asian", false)

	firstNameMap := make(map[string][][]string)
	firstNameMap["Caucasian"] = CaucasianFirstNameList
	firstNameMap["African"] = AfricanFirstNameList
	firstNameMap["Hispanic"] = HispanicFirstNameList
	firstNameMap["NativeAmerican"] = NativeFirstNameList
	firstNameMap["Asian"] = AsianFirstNameList

	lastNameMap := make(map[string][][]string)
	lastNameMap["Caucasian"] = CaucasianLastNameList
	lastNameMap["African"] = AfricanLastNameList
	lastNameMap["Hispanic"] = HispanicLastNameList
	lastNameMap["NativeAmerican"] = NativeLastNameList
	lastNameMap["Asian"] = AsianLastNameList

	return (firstNameMap), (lastNameMap)
}

func getName(list [][]string) string {
	endOfListWeight, err := strconv.Atoi(list[len(list)-1][1])
	if err != nil {
		log.Fatalln("Could not convert number from string")
	}
	name := ""
	num := util.GenerateIntFromRange(1, endOfListWeight)
	for i := 1; i < len(list); i++ {
		weight, err := strconv.Atoi(list[i][1])
		if err != nil {
			log.Fatalln("Could not convert number from string in name generator")
		}
		if num < weight {
			name = list[i][0]
			break
		}
	}
	return name
}

func getArchetype(pos string) string {
	if pos == "QB" {
		return util.PickFromStringList([]string{"Balanced", "Pocket", "Scrambler", "Field General"})
	} else if pos == "RB" {
		return util.PickFromStringList([]string{"Balanced", "Power", "Speed", "Receiving"})
	} else if pos == "FB" {
		return util.PickFromStringList([]string{"Balanced", "Blocking", "Receiving", "Rushing"})
	} else if pos == "WR" {
		return util.PickFromStringList([]string{"Speed", "Possession", "Route Runner", "Red Zone Threat"})
	} else if pos == "TE" {
		return util.PickFromStringList([]string{"Blocking", "Receiving", "Vertical Threat"})
	} else if pos == "OT" || pos == "OG" {
		return util.PickFromStringList([]string{"Balanced", "Pass Blocking", "Run Blocking"})
	} else if pos == "C" {
		return util.PickFromStringList([]string{"Balanced", "Pass Blocking", "Run Blocking", "Line Captain"})
	} else if pos == "DT" {
		return util.PickFromStringList([]string{"Balanced", "Nose Tackle", "Pass Rusher"})
	} else if pos == "DE" {
		return util.PickFromStringList([]string{"Balanced", "Run Stopper", "Speed Rusher"})
	} else if pos == "ILB" {
		return util.PickFromStringList([]string{"Coverage", "Field General", "Run Stopper", "Speed"})
	} else if pos == "OLB" {
		return util.PickFromStringList([]string{"Coverage", "Pass Rush", "Run Stopper", "Speed"})
	} else if pos == "CB" {
		return util.PickFromStringList([]string{"Balanced", "Ball Hawk", "Man Coverage", "Zone Coverage"})
	} else if pos == "FS" || pos == "SS" {
		return util.PickFromStringList([]string{"Run Stopper", "Ball Hawk", "Man Coverage", "Zone Coverage"})
	} else if pos == "K" || pos == "P" {
		return util.PickFromStringList([]string{"Balanced", "Accuracy", "Power"})
	}
	return ""

}

// Going to be honest, this should be a JSON file. This would be a huge blob of a map.
func getAttributeValue(pos string, arch string, star int, attr string, blob map[string]map[string]map[string]map[string]interface{}) int {
	starStr := strconv.Itoa(star)
	if pos == "QB" {
		if attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" || attr == "Run Defense" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])

	} else if pos == "RB" {
		if attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" || attr == "Run Defense" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "FB" {
		if attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Rush" || attr == "Run Defense" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "WR" {
		if attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Pass Rush" || attr == "Run Defense" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "TE" {
		if attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" || attr == "Pass Rush" || attr == "Run Defense" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "OT" || pos == "OG" {
		if attr == "Carrying" || attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Rush" || attr == "Run Defense" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "C" {
		if attr == "Carrying" || attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Tackle" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Rush" || attr == "Run Defense" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "DT" {
		if attr == "Carrying" || attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])

	} else if pos == "DE" {
		if attr == "Carrying" || attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "ILB" {
		if attr == "Carrying" || attr == "Catching" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "OLB" {
		if attr == "Carrying" || attr == "Catching" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Route Running" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "CB" {
		if attr == "Carrying" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" || attr == "Run Defense" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "FS" {
		if attr == "Carrying" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "SS" {
		if attr == "Carrying" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "K" {
		if attr == "Carrying" || attr == "Agility" || attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Punt Accuracy" || attr == "Punt Power" || attr == "Speed" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" || attr == "Run Defense" || attr == "Route Running" || attr == "Tackle" || attr == "Strength" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	} else if pos == "P" {
		if attr == "Carrying" || attr == "Agility" || attr == "Catching" || attr == "Zone Coverage" || attr == "Man Coverage" || attr == "Kick Accuracy" || attr == "Kick Power" || attr == "Speed" || attr == "Throw Power" || attr == "Throw Accuracy" {
			return getValueFromInterfaceRange(starStr, blob["Default"]["Default"]["Default"])
		} else if attr == "Pass Block" || attr == "Run Block" || attr == "Pass Rush" || attr == "Run Defense" || attr == "Route Running" || attr == "Tackle" || attr == "Strength" {
			return getValueFromInterfaceRange(starStr, blob["Under"]["Under"]["Under"])
		}
		return getValueFromInterfaceRange(starStr, blob[pos][arch][attr])
	}
	return util.GenerateIntFromRange(5, 15)
}

func getValueFromInterfaceRange(star string, starMap map[string]interface{}) int {
	u, ok := starMap[star].([]int)
	if ok {
		fmt.Println("Was able to get value)")
	}

	return util.GenerateIntFromRange(u[0], u[1])
}

func getStarRating() int {
	weightedRoll := util.GenerateIntFromRange(0, 10000)

	if weightedRoll < 9001 {
		return 1
	} else if weightedRoll < 9601 {
		return 2
	} else if weightedRoll < 9901 {
		return 3
	} else if weightedRoll < 9976 {
		return 4
	}
	return 5
}
