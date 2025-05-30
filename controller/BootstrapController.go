package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

func BootstrapTeamData(w http.ResponseWriter, r *http.Request) {
	data := managers.GetTeamsBootstrap()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func FirstBootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetFirstBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}

func SecondBootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetSecondBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}

func ThirdBootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetThirdBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}

func GetCollegeHistoryProfile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	data := managers.GetCollegeTeamProfilePageData()
	json.NewEncoder(w).Encode(data)
}
