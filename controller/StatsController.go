package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
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

func ExportPlayerStatsToCSV(w http.ResponseWriter, r *http.Request) {
	teamsChan := make(chan []structs.CollegeTeam)

	go func() {
		ct := managers.GetAllCollegeTeams()
		teamsChan <- ct
	}()

	collegeTeams := <-teamsChan
	close(teamsChan)

	var conferenceMap = make(map[int]int)
	var conferenceNameMap = make(map[int]string)

	for _, team := range collegeTeams {
		conferenceMap[int(team.ID)] = team.ConferenceID
		conferenceNameMap[int(team.ID)] = team.Conference
	}

	playersChan := make(chan []models.CollegePlayerResponse)
	go func() {
		cp := managers.GetAllCollegePlayersWithCurrentYearStatistics(conferenceMap, conferenceNameMap)
		playersChan <- cp
	}()

	collegePlayers := <-playersChan
	close(playersChan)

	managers.ExportPlayerStatsToCSV(collegePlayers, w)
}

func GetStatsPageContentForCurrentSeason(w http.ResponseWriter, r *http.Request) {

	teamsChan := make(chan []models.CollegeTeamResponse)

	go func() {
		ct := managers.GetAllCollegeTeamsWithCurrentSeasonStats()
		teamsChan <- ct
	}()

	collegeTeams := <-teamsChan
	close(teamsChan)

	var conferenceMap = make(map[int]int)
	var conferenceNameMap = make(map[int]string)

	for _, team := range collegeTeams {
		conferenceMap[int(team.ID)] = team.ConferenceID
		conferenceNameMap[int(team.ID)] = team.Conference
	}

	playersChan := make(chan []models.CollegePlayerResponse)
	go func() {
		cp := managers.GetAllCollegePlayersWithCurrentYearStatistics(conferenceMap, conferenceNameMap)
		playersChan <- cp
	}()

	collegePlayers := <-playersChan
	close(playersChan)

	confChan := make(chan []structs.CollegeConference)
	go func() {
		cf := managers.GetCollegeConferences()
		confChan <- cf
	}()

	collegeConferences := <-confChan
	close(confChan)

	response := models.SimCFBStatsResponse{
		CollegePlayers:     collegePlayers,
		CollegeTeams:       collegeTeams,
		CollegeConferences: collegeConferences,
	}

	json.NewEncoder(w).Encode(response)
}

func GetCollegePlayerStatsByNameTeamAndWeek(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstName := vars["firstName"]
	lastName := vars["lastName"]
	teamID := vars["team"]
	week := vars["week"]

	if len(firstName) == 0 {
		panic("User did not provide a first name")
	}

	player := managers.GetCollegePlayerByNameTeamAndWeek(firstName, lastName, teamID, week)

	json.NewEncoder(w).Encode(player)
}

func GetCurrentSeasonCollegePlayerStatsByNameTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstName := vars["firstName"]
	lastName := vars["lastName"]
	teamID := vars["team"]

	if len(firstName) == 0 {
		panic("User did not provide a first name")
	}

	player := managers.GetSeasonalCollegePlayerByNameTeam(firstName, lastName, teamID)

	json.NewEncoder(w).Encode(player)
}

func GetWeeklyTeamStatsByTeamAbbrAndWeek(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["team"]
	week := vars["week"]

	if len(teamID) == 0 {
		panic("User did not provide a first name")
	}

	team := managers.GetTeamStatsByWeekAndTeam(teamID, week)

	json.NewEncoder(w).Encode(team)
}

func GetSeasonTeamStatsByTeamAbbrAndSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["team"]
	season := vars["season"]

	if len(teamID) == 0 {
		panic("User did not provide a first name")
	}

	team := managers.GetSeasonalTeamStats(teamID, season)

	json.NewEncoder(w).Encode(team)
}

func MapAllStatsToSeason(w http.ResponseWriter, r *http.Request) {
	managers.MapAllStatsToSeason()

}
