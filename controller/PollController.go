package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
)

func CreatePollSubmission(w http.ResponseWriter, r *http.Request) {
	var dto structs.CollegePollSubmission
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate info from DTO
	if len(dto.Username) == 0 {
		log.Fatalln("ERROR: Cannot submit poll.")
	}

	poll := managers.CreatePoll(dto)
	json.NewEncoder(w).Encode(poll)
}

func GetPollSubmission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	ts := managers.GetTimestamp()
	seasonID := strconv.Itoa(int(ts.CollegeSeasonID))
	weekID := strconv.Itoa(int(ts.CollegeWeekID))
	poll := managers.GetPollSubmissionByUsernameWeekAndSeason(username)
	conferenceStandings := managers.GetAllConferenceStandingsBySeasonID(seasonID)
	collegeGames := managers.GetCollegeGamesByWeekIdAndSeasonID(weekID, seasonID)

	res := structs.PollDataResponse{
		Poll:      poll,
		Matches:   collegeGames,
		Standings: conferenceStandings,
	}

	json.NewEncoder(w).Encode(res)
}

func SyncCollegePoll(w http.ResponseWriter, r *http.Request) {
	managers.SyncCollegePollSubmissionForCurrentWeek()
}

func GetOfficialPollsBySeasonID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonID := vars["seasonID"]
	if len(seasonID) == 0 {
		panic("User did not provide seasonID")
	}
	polls := managers.GetOfficialPollBySeasonID(seasonID)
	conferenceStandings := managers.GetAllConferenceStandingsBySeasonID(seasonID)
	// collegeGames := managers.GetCBBMatchesBySeasonID(seasonID)

	res := structs.PollDataResponse{
		OfficialPolls: polls,
		// Matches:       collegeGames,
		Standings: conferenceStandings,
	}

	json.NewEncoder(w).Encode(res)
}
