package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
)

func ExportStatisticsFromSim(w http.ResponseWriter, r *http.Request) {
	// Create DTO for College Recruit
	var exportStatsDTO structs.ExportStatsDTO
	err := json.NewDecoder(r.Body).Decode(&exportStatsDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate info from DTO
	if len(exportStatsDTO.GameStatDTOs) == 0 {
		log.Fatalln("ERROR: Could not acquire all data for export")
	}

	// Send DTO to Manager Class
	managers.ExportStatisticsFromSim(exportStatsDTO)

	fmt.Println(w, "Game Data Exported")

}
