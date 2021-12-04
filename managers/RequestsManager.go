package managers

import (
	"fmt"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetAllTeamRequests() []structs.CreateRequestDTO {
	db := dbprovider.GetInstance().GetDB()
	var CollegeTeamRequests []structs.CreateRequestDTO
	var NFLTeamRequests []structs.CreateRequestDTO
	var AllRequests []structs.CreateRequestDTO

	// College Team Requests
	db.Raw("SELECT requests.id, requests.team_id, college_teams.team_name, college_teams.team_abbr, requests.username, college_teams.conference, requests.is_approved FROM simfbaah_interface_3.requests INNER JOIN simfbaah_interface_3.college_teams on college_teams.id = requests.team_id WHERE requests.deleted_at is null AND requests.is_approved = 0").
		Scan(&CollegeTeamRequests)

	// NFL Team Requests
	db.Raw("SELECT requests.id, requests.team_id, nfl_teams.team_name, nfl_teams.team_abbr, requests.username, nfl_teams.conference, requests.is_approved FROM simfbaah_interface_3.requests INNER JOIN simfbaah_interface_3.nfl_teams on nfl_teams.id = requests.team_id WHERE requests.deleted_at is null AND requests.is_approved = 0").
		Scan(&NFLTeamRequests)

	// Append
	AllRequests = append(AllRequests, CollegeTeamRequests...)
	AllRequests = append(AllRequests, NFLTeamRequests...)

	return AllRequests
}

func CreateTeamRequest(request structs.TeamRequest) {
	db := dbprovider.GetInstance().GetDB()

	var ExistingTeamRequest structs.TeamRequest
	err := db.Where("username = ? AND team_id = ?", request.Username, request.TeamID).Find(&ExistingTeamRequest).Error
	if err != nil {
		// Then there's no existing record, I guess? Which is fine.
		fmt.Println("Creating Team Request for TEAM " + strconv.Itoa(request.TeamID))
	}
	if ExistingTeamRequest.ID != 0 {
		// There is already an existing record.
		panic("There is already an existing request in place for the user. Please be patient while admin approves your formal request. If there is an issue, please reach out to TuscanSota.")
	}

	db.Create(&request)
}

func ApproveTeamRequest(request structs.TeamRequest) {
	db := dbprovider.GetInstance().GetDB()

	// Approve Request
	request.ApproveTeamRequest()

	fmt.Println("Team Approved...")

	db.Save(&request)

	// Assign Team
	fmt.Println("Assigning team...")

	team := GetTeamByTeamID(strconv.Itoa(request.TeamID))

	team.AssignUserToTeam(request.Username)

	db.Save(&team)
}

func RejectTeamRequest(request structs.TeamRequest) {
	db := dbprovider.GetInstance().GetDB()

	request.RejectTeamRequest()

	db.Delete(&request)
}
