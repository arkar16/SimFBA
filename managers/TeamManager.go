package managers

import (
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetAllCollegeTeams() []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	db.Find(&teams)

	return teams
}

func GetAllAvailableCollegeTeams() []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	db.Where("coach IN (?,?)", "", "AI").Find(&teams)

	return teams
}

func GetAllCoachedCollegeTeams() []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	db.Where("coach is not null AND coach NOT IN (?,?)", "", "AI").Find(&teams)

	return teams
}

// GetTeamByTeamID - straightforward
func GetTeamByTeamID(teamId string) structs.CollegeTeam {
	var team structs.CollegeTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Where("id = ?", teamId).Find(&team).Error
	if err != nil {
		log.Fatal(err)
	}
	return team
}

// GetTeamsByConferenceID
func GetTeamsByConferenceID(conferenceID string) []structs.CollegeTeam {
	var teams []structs.CollegeTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Where("conference_id = ?", conferenceID).Find(&teams).Error
	if err != nil {
		log.Fatal(err)
	}
	return teams
}

// GetTeamsByConferenceIDWithStandings
func GetTeamsByConferenceIDWithStandings(conferenceID string, seasonID string) []structs.CollegeTeam {
	var teams []structs.CollegeTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Preload("TeamStandings").
		Where("conference_id = ? AND TeamStandings.season_id = ?", conferenceID, seasonID).
		Find(&teams).Error
	if err != nil {
		log.Fatal(err)
	}
	return teams
}

// GetTeamsByDivisionID
func GetTeamsByDivisionID(conferenceID string) []structs.CollegeTeam {
	var teams []structs.CollegeTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Where("division_id = ?", conferenceID).Find(&teams).Error
	if err != nil {
		log.Fatal(err)
	}
	return teams
}

func RemoveUserFromTeam(teamId string) {
	db := dbprovider.GetInstance().GetDB()

	team := GetTeamByTeamID(teamId)

	coach := GetCollegeCoachByCoachName(team.Coach)

	coach.SetAsInactive()

	team.RemoveUserFromTeam()

	db.Save(&team)

	db.Save(&coach)
}

func GetTeamsInConference(conference string) []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	err := db.Where("conference = ?", conference).Find(&teams).Error
	if err != nil {
		log.Fatal(err)
	}
	return teams
}

func GetTeamByTeamAbbr(abbr string) structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var team structs.CollegeTeam

	err := db.Preload("TeamGameplan").Preload("TeamDepthChart.DepthChartPlayers").Where("team_abbr = ?", abbr).Find(&team).Error
	if err != nil {
		log.Panicln("Could not find team by given abbreviation:"+abbr+"\n", err)
	}

	return team
}

func GetAllCollegeTeamsWithRecruitingProfileAndCoach() []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	db.Preload("CollegeCoach").Preload("RecruitingProfile").Find(&teams)

	return teams
}
