package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/gorilla/mux"
)

// Flex: Compare Two Program's history against one another
func GetCollegeConferenceStandings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	confID := vars["conference"]
	if len(confID) == 0 {
		panic("User did not provide teamID")
	}

	ts := managers.GetTimestamp()
	seasonID := strconv.Itoa(ts.CollegeSeasonID)

	res := managers.GetStandingsByConferenceIDAndSeasonID(confID, seasonID)

	json.NewEncoder(w).Encode(res)
}

// Flex: Compare Two Program's history against one another
func CompareTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamOneID := vars["teamOneID"]
	if len(teamOneID) == 0 {
		panic("User did not provide teamID")
	}

	teamTwoID := vars["teamTwoID"]
	if len(teamTwoID) == 0 {
		panic("User did not provide teamID")
	}

	res := managers.CompareTwoTeams(teamOneID, teamTwoID)

	json.NewEncoder(w).Encode(res)
}

func GetTeamByTeamIDForDiscord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]
	if len(teamID) == 0 {
		panic("User did not provide TeamID")
	}
	team := managers.GetCFBTeamDataForDiscord(teamID)
	json.NewEncoder(w).Encode(team)
}

func GetCollegePlayerStatsByNameTeamAndWeek(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	week := vars["week"]

	if len(id) == 0 {
		panic("User did not provide a first name")
	}

	player := managers.GetCollegePlayerByIdAndWeek(id, week)

	json.NewEncoder(w).Encode(player)
}

func GetCurrentSeasonCollegePlayerStatsByNameTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) == 0 {
		panic("User did not provide a first name")
	}

	player := managers.GetSeasonalCollegePlayerByNameTeam(id)

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

// GetCollegePlayer
func GetCollegePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) == 0 {
		panic("User did not provide a first name")
	}

	player := managers.GetCollegePlayerViaDiscord(id)

	json.NewEncoder(w).Encode(player)
}

func GetRecruitingClassByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	if len(teamID) == 0 {
		panic("User did not provide teamID")
	}

	recruitingProfile := managers.GetRecruitingClassByTeamID(teamID)

	json.NewEncoder(w).Encode(recruitingProfile)
}

func GetRecruitViaDiscord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) == 0 {
		panic("User did not provide a first name")
	}

	recruit := managers.GetCollegeRecruitViaDiscord(id)

	json.NewEncoder(w).Encode(recruit)
}

// GetCollegeGamesByTeamIDAndSeasonID
func GetCurrentWeekGamesByLeague(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	league := vars["league"]
	if len(league) == 0 {
		panic("User did not provide teamID")
	}

	if league == "simcfb" {
		collegeGames := managers.GetCFBCurrentWeekSchedule()
		json.NewEncoder(w).Encode(collegeGames)
	} else {
		nflgames := managers.GetNFLCurrentWeekSchedule()
		json.NewEncoder(w).Encode(nflgames)
	}
}

// GetCollegeGamesByConference
func GetCollegeGamesByConference(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conference := vars["conference"]
	if len(conference) == 0 {
		panic("User did not provide conference name")
	}

	collegeGames := managers.GetCFBScheduleByConference(conference)
	json.NewEncoder(w).Encode(collegeGames)
}

func GetFBSGameStreams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timeslot := vars["timeslot"]
	if len(timeslot) == 0 {
		panic("User did not provide timeslot")
	}
	week := vars["week"]
	if len(timeslot) == 0 {
		panic("User did not provide week")
	}
	streams := managers.GetCFBPlayByPlayStreamData(timeslot, week, true)
	json.NewEncoder(w).Encode(streams)
}

func GetFCSGameStreams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timeslot := vars["timeslot"]
	if len(timeslot) == 0 {
		panic("User did not provide timeslot")
	}
	week := vars["week"]
	if len(timeslot) == 0 {
		panic("User did not provide week")
	}
	streams := managers.GetCFBPlayByPlayStreamData(timeslot, week, false)
	json.NewEncoder(w).Encode(streams)
}

func GetNFLGameStreams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timeslot := vars["timeslot"]
	if len(timeslot) == 0 {
		panic("User did not provide timeslot")
	}
	week := vars["week"]
	if len(timeslot) == 0 {
		panic("User did not provide week")
	}
	streams := managers.GetNFLPlayByPlayStreamData(timeslot, week, false)
	json.NewEncoder(w).Encode(streams)
}
