package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

func BootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}
