package util

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/CalebRose/SimFBA/structs"
)

func GetOffensiveDefaultSchemes() map[string]structs.OffensiveFormation {
	path := filepath.Join(os.Getenv("ROOT"), "data", "defaultOffensiveSchemes.json")
	content := ReadJson(path)

	var payload map[string]structs.OffensiveFormation

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalln("Error during unmarshal: ", err)
	}

	return payload
}

func GetDefensiveDefaultSchemes() map[string]map[string]structs.DefensiveFormation {
	path := filepath.Join(os.Getenv("ROOT"), "data", "defaultDefensiveSchemes.json")
	content := ReadJson(path)

	var payload map[string]map[string]structs.DefensiveFormation

	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalln("Error during unmarshal: ", err)
	}

	return payload
}
