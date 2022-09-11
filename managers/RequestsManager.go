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
	// var NFLTeamRequests []structs.CreateRequestDTO
	var AllRequests []structs.CreateRequestDTO

	// College Team Requests
	db.Raw("SELECT team_requests.id, team_requests.team_id, college_teams.team_name, college_teams.team_abbr, team_requests.username, college_teams.conference, team_requests.is_approved FROM simfbaah_interface_3.team_requests INNER JOIN simfbaah_interface_3.college_teams on college_teams.id = team_requests.team_id WHERE team_requests.deleted_at is null AND team_requests.is_approved = 0").
		Scan(&CollegeTeamRequests)

	// NFL Team Requests
	// db.Raw("SELECT team_requests.id, team_requests.team_id, nfl_teams.team_name, nfl_teams.team_abbr, team_requests.username, nfl_teams.conference, team_requests.is_approved FROM simfbaah_interface_3.team_requests INNER JOIN simfbaah_interface_3.nfl_teams on nfl_teams.id = team_requests.team_id WHERE team_requests.deleted_at is null AND requests.is_approved = 0").
	// 	Scan(&NFLTeamRequests)

	// Append
	AllRequests = append(AllRequests, CollegeTeamRequests...)
	// AllRequests = append(AllRequests, NFLTeamRequests...)

	return AllRequests
}

func CreateTeamRequest(request structs.TeamRequest) {
	db := dbprovider.GetInstance().GetDB()

	var ExistingTeamRequest structs.TeamRequest
	err := db.Where("username = ? AND team_id = ? AND is_approved = false AND deleted_at is null", request.Username, request.TeamID).Find(&ExistingTeamRequest).Error
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

func ApproveTeamRequest(request structs.TeamRequest) structs.TeamRequest {
	db := dbprovider.GetInstance().GetDB()

	timestamp := GetTimestamp()

	// Approve Request
	request.ApproveTeamRequest()

	fmt.Println("Team Approved...")

	db.Save(&request)

	// Assign Team
	fmt.Println("Assigning team...")

	team := GetTeamByTeamID(strconv.Itoa(request.TeamID))

	coach := GetCollegeCoachByCoachName(request.Username)

	coach.SetTeam(request.TeamID)

	team.AssignUserToTeam(request.Username)

	seasonalGames := GetCollegeGamesByTeamIdAndSeasonId(strconv.Itoa(request.TeamID), strconv.Itoa(timestamp.CollegeSeasonID))

	for _, game := range seasonalGames {
		if game.Week >= timestamp.CollegeWeek {
			game.UpdateCoach(request.TeamID, request.Username)
			db.Save(&game)
		}

	}

	db.Save(&team)

	db.Save(&coach)

	newsLog := structs.NewsLog{
		WeekID:      timestamp.CollegeWeekID,
		SeasonID:    timestamp.CollegeSeasonID,
		MessageType: "CoachJob",
		Message:     "Breaking News! The " + team.TeamName + " " + team.Mascot + " have hired " + coach.CoachName + " as their new coach for the " + strconv.Itoa(timestamp.Season) + " season!",
	}

	db.Create(&newsLog)

	return request
}

func RejectTeamRequest(request structs.TeamRequest) {
	db := dbprovider.GetInstance().GetDB()

	request.RejectTeamRequest()

	db.Delete(&request)
}
