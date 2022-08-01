package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/models"
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
