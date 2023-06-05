package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

func GetNewsLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekID := vars["weekID"]
	seasonID := vars["seasonID"]

	newsLogs := managers.GetNewsLogs(weekID, seasonID)

	json.NewEncoder(w).Encode(newsLogs)
}

func GetAllNewsLogsForASeason(w http.ResponseWriter, r *http.Request) {
	newsLogs := managers.GetAllNewsLogs()
	json.NewEncoder(w).Encode(newsLogs)
}

func GetAllNFLNewsBySeason(w http.ResponseWriter, r *http.Request) {
	newsLogs := managers.GetAllNFLNewsLogs()

	json.NewEncoder(w).Encode(newsLogs)
}

func GetNewsFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	league := vars["league"]
	teamID := vars["teamID"]

	if league == "CFB" {
		newsLogs := managers.GetCFBRelatedNews(teamID)
		json.NewEncoder(w).Encode(newsLogs)
	} else {
		newsLogs := managers.GetNFLRelatedNews(teamID)
		json.NewEncoder(w).Encode(newsLogs)
	}
}
