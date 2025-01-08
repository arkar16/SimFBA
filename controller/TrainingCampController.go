package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

// GetAllCollegeTeamsForRosterPage
func UploadTrainingCampCSVData(w http.ResponseWriter, r *http.Request) {
	managers.UploadTrainingCampCSV()

	json.NewEncoder(w).Encode("Training Camp Complete")
}
