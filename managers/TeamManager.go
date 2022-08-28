package managers

import (
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
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

func GetTeamByTeamIDForDiscord(teamId string) structs.CollegeTeam {
	var team structs.CollegeTeam
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	err := db.Preload("TeamStandings", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ?", ts.CollegeSeasonID)
	}).Where("id = ?", teamId).Find(&team).Error
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

	timestamp := GetTimestamp()

	newsLog := structs.NewsLog{
		WeekID:      timestamp.CollegeWeekID,
		SeasonID:    timestamp.CollegeSeasonID,
		MessageType: "CoachJob",
		Message:     coach.CoachName + " has decided to step down as the head coach of the " + team.TeamName + " " + team.Mascot + "!",
	}

	db.Create(&newsLog)
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

func GetAllCollegeTeamsWithCurrentYearStandings() []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	var teams []structs.CollegeTeam

	db.Preload("TeamStandings", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ?", strconv.Itoa(ts.CollegeSeasonID))
	}).Find(&teams)

	return teams
}

func GetAllCollegeTeamsWithCurrentSeasonStats() []models.CollegeTeamResponse {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	var teams []structs.CollegeTeam

	db.Preload("TeamStats", func(db *gorm.DB) *gorm.DB {
		return db.Where("season_id = ? and week_id < ?", strconv.Itoa(ts.CollegeSeasonID), strconv.Itoa(ts.CollegeWeekID))
	}).Find(&teams)

	var ctResponse []models.CollegeTeamResponse

	for _, team := range teams {
		ct := models.CollegeTeamResponse{
			ID:           int(team.ID),
			BaseTeam:     team.BaseTeam,
			ConferenceID: team.ConferenceID,
			Conference:   team.Conference,
			DivisionID:   team.DivisionID,
			Division:     team.Division,
			TeamStats:    team.TeamStats,
		}

		ct.MapSeasonalStats()

		ctResponse = append(ctResponse, ct)
	}

	return ctResponse
}

func GetCollegeConferences() []structs.CollegeConference {
	db := dbprovider.GetInstance().GetDB()

	var conferences []structs.CollegeConference

	db.Preload("Divisions").Find(&conferences)

	return conferences
}
