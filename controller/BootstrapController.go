package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

func FirstBootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetFirstBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}

func SecondBootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetSecondBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}

func ThirdBootstrapFootballData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collegeID := vars["collegeID"]
	proID := vars["proID"]
	data := managers.GetThirdBootstrapData(collegeID, proID)
	json.NewEncoder(w).Encode(data)
}
