package managers

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/repository"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func MigrateFaceDataToRecruits() {
	db := dbprovider.GetInstance().GetDB()
	// Get Recruits
	recruits := GetAllRecruitRecords()
	// Get Full Name Lists
	_, lNameMap := getNameMaps()
	faceDataBlob := getFaceDataBlob()
	faceDataList := []structs.FaceData{}
	// Initialize List
	for _, r := range recruits {
		lastName := strings.ToUpper(r.LastName)
		skinColor := getSkinColor(lastName, lNameMap)
		// Store data

		face := getFace(r.ID, skinColor, faceDataBlob)

		faceDataList = append(faceDataList, face)
	}

	repository.CreateFaceRecordsBatch(db, faceDataList, 500)
}

func MigrateFaceDataToCollegePlayers() {
	db := dbprovider.GetInstance().GetDB()
	// Get Recruits
	players := GetAllCollegePlayers()
	// Get Full Name Lists
	_, lNameMap := getNameMaps()
	faceDataBlob := getFaceDataBlob()
	faceDataList := []structs.FaceData{}
	// Initialize List
	for _, p := range players {
		lastName := strings.ToUpper(p.LastName)
		skinColor := getSkinColor(lastName, lNameMap)
		// Store data

		face := getFace(p.ID, skinColor, faceDataBlob)

		faceDataList = append(faceDataList, face)
	}

	repository.CreateFaceRecordsBatch(db, faceDataList, 500)
}

func MigrateFaceDataToHistoricCollegePlayers() {
	db := dbprovider.GetInstance().GetDB()
	// Get Recruits
	players := GetAllHistoricCollegePlayers()
	// Get Full Name Lists
	_, lNameMap := getNameMaps()
	faceDataBlob := getFaceDataBlob()
	faceDataList := []structs.FaceData{}
	// Initialize List
	for _, p := range players {
		lastName := strings.ToUpper(p.LastName)
		skinColor := getSkinColor(lastName, lNameMap)
		// Store data

		face := getFace(p.ID, skinColor, faceDataBlob)

		faceDataList = append(faceDataList, face)
	}

	repository.CreateFaceRecordsBatch(db, faceDataList, 500)
}

func MigrateFaceDataToNFLPlayers() {
	db := dbprovider.GetInstance().GetDB()
	// Get Recruits
	players := GetAllNFLPlayers()
	// Get Full Name Lists
	_, lNameMap := getNameMaps()
	faceDataBlob := getFaceDataBlob()
	faceDataList := []structs.FaceData{}
	// Initialize List
	for _, p := range players {
		lastName := strings.ToUpper(p.LastName)
		skinColor := getSkinColor(lastName, lNameMap)
		// Store data

		face := getFace(p.ID, skinColor, faceDataBlob)

		faceDataList = append(faceDataList, face)
	}

	repository.CreateFaceRecordsBatch(db, faceDataList, 500)
}

func MigrateFaceDataToRetiredPlayers() {
	db := dbprovider.GetInstance().GetDB()
	// Get Recruits
	players := GetRetiredSimNFLPlayers()
	// Get Full Name Lists
	_, lNameMap := getNameMaps()
	faceDataBlob := getFaceDataBlob()
	faceDataList := []structs.FaceData{}
	// Initialize List
	for _, p := range players {
		lastName := strings.ToUpper(p.LastName)
		skinColor := getSkinColor(lastName, lNameMap)
		// Store data

		face := getFace(p.ID, skinColor, faceDataBlob)

		faceDataList = append(faceDataList, face)
	}

	repository.CreateFaceRecordsBatch(db, faceDataList, 500)
}

func getFace(id uint, ethnicity string, faceDataBlob map[string][]string) structs.FaceData {
	face := structs.FaceData{
		PlayerID:        id,
		Accessories:     uint8(util.GenerateIntFromRange(0, len(faceDataBlob["accessories"]))),
		Body:            uint8(util.GenerateIntFromRange(0, len(faceDataBlob["body"]))),
		BodySize:        float32(util.GenerateFloatFromRange(0.8, 1.2)),
		Ear:             uint8(util.GenerateIntFromRange(0, len(faceDataBlob["ear"]))),
		EarSize:         float32(util.GenerateFloatFromRange(0.5, 1.5)),
		Eye:             uint8(util.GenerateIntFromRange(0, len(faceDataBlob["eye"]))),
		EyeLine:         uint8(util.GenerateIntFromRange(0, len(faceDataBlob["eyeLine"]))),
		EyeAngle:        int8(util.GenerateIntFromRange(-10, 15)),
		Eyebrow:         uint8(util.GenerateIntFromRange(0, len(faceDataBlob["eyebrow"]))),
		EyeBrowAngle:    int8(util.GenerateIntFromRange(-15, 20)),
		FaceSize:        float32(util.GenerateFloatFromRange(0, 1)),
		FacialHair:      uint8(util.GenerateIntFromRange(0, len(faceDataBlob["facialHair"]))),
		FacialHairShave: uint8(util.GenerateIntFromRange(1, 5)),
		Glasses:         0,
		Hair:            uint8(util.GenerateIntFromRange(0, len(faceDataBlob["hair"]))),
		HairBG:          uint8(util.GenerateNormalizedIntFromRange(0, 2)),
		HairColor:       uint8(util.GenerateIntFromRange(0, len(faceDataBlob[ethnicity+"Hair"]))),
		HairFlip:        util.GenerateIntFromRange(1, 2) == 1,
		Head:            uint8(util.GenerateIntFromRange(0, len(faceDataBlob["head"]))),
		Jersey:          uint8(util.GenerateIntFromRange(0, len(faceDataBlob["jersey"]))),
		MiscLine:        uint8(util.GenerateIntFromRange(0, len(faceDataBlob["MiscLine"]))),
		Mouth:           uint8(util.GenerateIntFromRange(0, len(faceDataBlob["mouth"]))),
		MouthFlip:       util.GenerateIntFromRange(1, 2) == 1,
		Nose:            uint8(util.GenerateIntFromRange(0, len(faceDataBlob["nose"]))),
		NoseFlip:        util.GenerateIntFromRange(1, 2) == 1,
		NoseSize:        float32(util.GenerateFloatFromRange(0.5, 1.25)),
		SkinColor:       uint8(util.GenerateIntFromRange(0, len(faceDataBlob[ethnicity+"Skin"]))),
		SmileLine:       uint8(util.GenerateIntFromRange(0, len(faceDataBlob["smileLine"]))),
		SmileLineSize:   float32(util.GenerateFloatFromRange(0.25, 2.25)),
	}

	return face
}

func getSkinColor(lastName string, nameMap map[string][][]string) string {
	skinColor := "asian"
	isCaucasian := checkNameInList(lastName, nameMap["Caucasian"])
	isHispanic := checkNameInList(lastName, nameMap["Hispanic"])
	isAfrican := checkNameInList(lastName, nameMap["African"])
	isAsian := checkNameInList(lastName, nameMap["Asian"])
	isNativeAmerican := checkNameInList(lastName, nameMap["NativeAmerican"])
	if isCaucasian {
		skinColor = "white"
	} else if isHispanic {
		skinColor = "brown"
	} else if isAfrican {
		skinColor = "black"
	} else if isNativeAmerican {
		skinColor = "brown"
	}
	// Edge case for custom players
	if !isCaucasian && !isHispanic && !isAfrican && !isAsian && !isNativeAmerican {
		skinColor = util.PickFromStringList([]string{"Asian", "African", "Hispanic", "Caucasian", "Native American"})
	}
	return skinColor
}

func checkNameInList(name string, namesList [][]string) bool {
	for _, lastName := range namesList {
		if name == lastName[0] {
			return true
		}
	}
	return false
}

func getFaceDataBlob() map[string][]string {
	path := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\faceData.json"

	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}

	var payload map[string][]string
	err = json.Unmarshal(f, &payload)
	if err != nil {
		log.Fatal("Error during unmarshal: ", err)
	}

	return payload
}

func getSkinColorByEthnicity(ethn string) string {
	if ethn == "Caucasian" {
		return "white"
	}
	if ethn == "Asian" {
		return "asian"
	}
	if ethn == "African" {
		return "black"
	}
	return "brown"
}
