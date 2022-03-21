package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CalebRose/SimFBA/controller"
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/sendgrid"
	"github.com/rs/cors"
)

func InitialMigration() {
	initiate := dbprovider.GetInstance().InitDatabase()
	if !initiate {
		log.Println("Initiate pool failure... Ending application")
		os.Exit(1)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Health Controls
	HealthCheck := health.New(
		health.Health{
			Version:   "1",
			ReleaseID: "0.0.7-SNAPSHOT",
		},
		sendgrid.Health(),
	)
	myRouter.HandleFunc("/health", HealthCheck.Handler).Methods("GET")

	// Admin Controls

	// Game Controls
	myRouter.HandleFunc("/games/college/week/{weekID}/", controller.GetCollegeGamesByTimeslotWeekId).Methods("GET")
	myRouter.HandleFunc("/games/college/timeslot/{timeSlot}/{weekID}", controller.GetCollegeGamesByTimeslotWeekId).Methods("GET")
	myRouter.HandleFunc("/games/college/team/{teamID}/{seasonID}", controller.GetCollegeGamesByTeamIDAndSeasonID).Methods("GET")

	// Gameplan Controls
	myRouter.HandleFunc("/gameplan/college/team/{teamID}/", controller.GetTeamGameplanByTeamID).Methods("GET")
	myRouter.HandleFunc("/gameplan/college/updategameplan", controller.UpdateGameplan).Methods("PUT")
	myRouter.HandleFunc("/gameplan/college/depthchart/{teamID}/", controller.GetTeamDepthchartByTeamID).Methods("GET")
	myRouter.HandleFunc("/gameplan/college/depthchart/positions/{depthChartID}/", controller.GetDepthChartPositionsByDepthChartID).Methods("GET")
	myRouter.HandleFunc("/gameplan/college/updatedepthchart", controller.UpdateDepthChart).Methods("PUT")

	// Player Controls
	myRouter.HandleFunc("/players/all/", controller.AllPlayers).Methods("GET")
	myRouter.HandleFunc("/collegeplayers/team/{teamID}/", controller.AllCollegePlayersByTeamID).Methods("GET")
	myRouter.HandleFunc("/collegeplayers/team/nors/{teamID}/", controller.AllCollegePlayersByTeamIDWithoutRedshirts).Methods("GET")
	myRouter.HandleFunc("/collegeplayers/team/export/{teamID}/", controller.ExportRosterToCSV).Methods("GET")
	// myRouter.HandleFunc("/collegeplayers/teams/export/", controller.ExportAllRostersToCSV).Methods("GET") // DO NOT USE

	// Recruiting Controls
	myRouter.HandleFunc("/recruiting/profile/{teamID}/", controller.GetRecruitingProfileByTeamID).Methods("GET")
	myRouter.HandleFunc("/recruiting/createpointsprofile/", controller.CreateRecruitingPointsProfileForRecruit).Methods("POST")
	myRouter.HandleFunc("/recruiting/allocaterecruitpoints/", controller.AllocateRecruitingPointsForRecruit).Methods("PUT")
	myRouter.HandleFunc("/recruiting/sendscholarship/", controller.SendScholarshipToRecruit).Methods("PUT")
	myRouter.HandleFunc("/recruiting/revokescholarship/", controller.RevokeScholarshipFromRecruit).Methods("PUT")
	myRouter.HandleFunc("/recruiting/removecrootfromboard/", controller.RemoveRecruitFromBoard).Methods("PUT")
	myRouter.HandleFunc("/recruiting/saverecruitingprofile/", controller.SaveRecruitingBoard).Methods("PUT")

	// ReCroot Controls
	myRouter.HandleFunc("/recruits/all/", controller.AllRecruits).Methods("GET")
	// myRouter.HandleFunc("/recruits/juco/all/", controller.AllJUCOCollegeRecruits).Methods("GET")
	myRouter.HandleFunc("/recruits/recruit/{recruitID}/", controller.GetCollegeRecruitByRecruitID).Methods("GET")
	myRouter.HandleFunc("/recruits/profile/recruits/{recruitProfileID}/", controller.GetRecruitsByTeamProfileID).Methods("GET")
	myRouter.HandleFunc("/recruits/recruit/create/", controller.CreateCollegeRecruit).Methods("POST")
	// myRouter.HandleFunc("/recruits/recruit/update/", controller.UpdateCollegeRecruit).Methods("PUT")

	// Requests Controls
	myRouter.HandleFunc("/requests/all/", controller.GetTeamRequests).Methods("GET")
	myRouter.HandleFunc("/requests/create/", controller.CreateTeamRequest).Methods("POST")
	myRouter.HandleFunc("/requests/approve/", controller.ApproveTeamRequest).Methods("PUT")
	myRouter.HandleFunc("/requests/reject/", controller.RejectTeamRequest).Methods("DELETE")
	myRouter.HandleFunc("/requests/remove/{teamID}", controller.RemoveUserFromTeam).Methods("PUT")

	// Standings Controls

	// Stats Controls

	// Team Controls
	myRouter.HandleFunc("/teams/college/all/", controller.GetAllCollegeTeams).Methods("GET")
	myRouter.HandleFunc("/teams/college/active/", controller.GetAllActiveCollegeTeams).Methods("GET")
	myRouter.HandleFunc("/teams/college/available/", controller.GetAllAvailableCollegeTeams).Methods("GET")
	myRouter.HandleFunc("/teams/college/team/{teamID}/", controller.GetTeamByTeamID).Methods("GET")
	myRouter.HandleFunc("/teams/college/conference/{conferenceID}/", controller.GetTeamsByConferenceID).Methods("GET")
	myRouter.HandleFunc("/teams/college/division/{divisionID}/", controller.GetTeamsByDivisionID).Methods("GET")

	// Draft Controls
	myRouter.HandleFunc("/nfl/draft/draftees/export/{season}", controller.ExportDrafteesToCSV).Methods("GET")

	// Handle Controls
	handler := cors.AllowAll().Handler(myRouter)

	log.Fatal(http.ListenAndServe(":8081", handler))
}

func main() {
	InitialMigration()
	fmt.Println("Football Server Initialized.")

	handleRequests()
	fmt.Println("Hello There!")
}
