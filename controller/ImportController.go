package controller

import (
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func ImportRecruitAICSV(w http.ResponseWriter, r *http.Request) {
	managers.ImportRecruitAICSV()
}

func ImportNFLRecords(w http.ResponseWriter, r *http.Request) {
	managers.RetireAndFreeAgentPlayers()
}

func ImportWorkEthic(w http.ResponseWriter, r *http.Request) {
	managers.ImportWorkEthic()
}

func ImportFAPreferences(w http.ResponseWriter, r *http.Request) {
	managers.ImportFAPreferences()
}

func GetMissingRecruitingClasses(w http.ResponseWriter, r *http.Request) {
	managers.GetMissingRecruitingClasses()
}
