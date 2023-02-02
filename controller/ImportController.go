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

func GetMissingRecruitingClasses(w http.ResponseWriter, r *http.Request) {
	managers.GetMissingRecruitingClasses()
}
