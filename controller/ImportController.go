package controller

import (
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func ImportRecruitAICSV(w http.ResponseWriter, r *http.Request) {
	managers.ImportRecruitAICSV()
}
