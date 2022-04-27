package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
)

// AllRecruits - Get all recruits in DB
func AllRecruits(w http.ResponseWriter, r *http.Request) {
	var recruits = managers.GetAllRecruits()

	json.NewEncoder(w).Encode(recruits)
}

// AllJUCOCollegeRecruits
func AllJUCOCollegeRecruits(w http.ResponseWriter, r *http.Request) {
	// Need to write manager code for getting all college recruits who age / year is greater than 18

	// json.NewEncoder(w).Encode(recruits)
}

// GetCollegeRecruitByRecruitID
func GetCollegeRecruitByRecruitID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recruitID := vars["recruitID"]

	if len(recruitID) == 0 {
		panic("User did not provide RecruitID")
	}

	recruit := managers.GetCollegeRecruitByRecruitID(recruitID)

	json.NewEncoder(w).Encode(recruit)
}

// GetRecruitsByProfileID
func GetRecruitsByTeamProfileID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recruitProfileID := vars["recruitProfileID"]

	if len(recruitProfileID) == 0 {
		panic("User did not provide RecruitProfileID")
	}

	recruits := managers.GetRecruitsByTeamProfileID(recruitProfileID)

	json.NewEncoder(w).Encode(recruits)
}

// CreateCollegeRecruit
func CreateCollegeRecruit(w http.ResponseWriter, r *http.Request) {
	// Create DTO for College Recruit
	var createRecruitDTO structs.CreateRecruitDTO
	err := json.NewDecoder(r.Body).Decode(&createRecruitDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate info from DTO
	if len(createRecruitDTO.FirstName) == 0 || len(createRecruitDTO.LastName) == 0 || createRecruitDTO.Overall == 0 {
		log.Fatalln("ERROR: Did not provide all information for recruit.")
	}

	// Send DTO to Manager Class
	managers.CreateCollegeRecruit(createRecruitDTO)

	// Send to interface?

	fmt.Println(w, "New Recruit Created")

}

// UpdateCollegeRecruit
func UpdateCollegeRecruit(w http.ResponseWriter, r *http.Request) {
	// Update DTO for College Recruit

	// validate info from DTO

	// Send DTO to Manager Class

	// Save to DB

	// Return new recruit

	// Send to interface?

	fmt.Println(w, "Recruit successfully updated.")
}

func SyncCrootRankings(w http.ResponseWriter, r *http.Request) {
	// Get all recruits, sort by overall

	// Loop through each recruit (get index?)

	// Calculate ESPN Ranking

	// Calculate Rivals Ranking

	// Third ranking?

	// Save recruit

	fmt.Println(w, "Recruit successfully updated.")
}
