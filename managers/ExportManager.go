package managers

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/CalebRose/SimFBA/models"
)

func ExportTeamToCSV(TeamID string, w http.ResponseWriter) {
	// Get Team Data
	team := GetTeamByTeamID(TeamID)
	w.Header().Set("Content-Disposition", "attachment;filename="+team.TeamName+".csv")
	w.Header().Set("Transfer-Encoding", "chunked")
	// Initialize writer
	writer := csv.NewWriter(w)

	// Get Players
	players := GetAllCollegePlayersByTeamId(TeamID)

	HeaderRow := []string{
		"Team", "First Name", "Last Name", "Position",
		"Archetype", "Year", "Age", "Stars",
		"High School", "City", "State", "Height",
		"Weight", "Overall", "Speed",
		"Football IQ", "Agility", "Carrying",
		"Catching", "Route Running", "Zone Coverage", "Man Coverage",
		"Strength", "Tackle", "Pass Block", "Run Block",
		"Pass Rush", "Run Defense", "Throw Power", "Throw Accuracy",
		"Kick Power", "Kick Accuracy", "Punt Power", "Punt Accuracy",
		"Stamina", "Injury", "Potential Grade", "Redshirt Status",
	}

	err := writer.Write(HeaderRow)
	if err != nil {
		log.Fatal("Cannot write header row", err)
	}

	for _, player := range players {
		csvModel := models.MapPlayerToCSVModel(player)
		playerRow := []string{
			team.TeamName, csvModel.FirstName, csvModel.LastName, csvModel.Position,
			csvModel.Archetype, csvModel.Year, strconv.Itoa(player.Age), strconv.Itoa(player.Stars),
			player.HighSchool, player.City, player.State, strconv.Itoa(player.Height),
			strconv.Itoa(player.Weight), csvModel.OverallGrade, csvModel.SpeedGrade,
			csvModel.FootballIQGrade, csvModel.AgilityGrade, csvModel.CarryingGrade,
			csvModel.CatchingGrade, csvModel.RouteRunningGrade, csvModel.ZoneCoverageGrade, csvModel.ManCoverageGrade,
			csvModel.StrengthGrade, csvModel.TackleGrade, csvModel.PassBlockGrade, csvModel.RunBlockGrade,
			csvModel.PassRushGrade, csvModel.RunDefenseGrade, csvModel.ThrowPowerGrade, csvModel.ThrowAccuracyGrade,
			csvModel.KickPowerGrade, csvModel.KickAccuracyGrade, csvModel.PuntPowerGrade, csvModel.PuntAccuracyGrade,
			csvModel.StaminaGrade, csvModel.InjuryGrade, csvModel.PotentialGrade, csvModel.RedshirtStatus,
		}

		err = writer.Write(playerRow)
		if err != nil {
			log.Fatal("Cannot write player row to CSV", err)
		}

		writer.Flush()
		err = writer.Error()
		if err != nil {
			log.Fatal("Error while writing to file ::", err)
		}
	}
}

func ExportCrootsToCSV(w http.ResponseWriter) {
	w.Header().Set("Content-Disposition", "attachment;filename=2022SimNFLDraftClass.csv")
	w.Header().Set("Transfer-Encoding", "chunked")
	// Initialize writer
	writer := csv.NewWriter(w)

	// Get NFL Draft Class
	croots := GetAllRecruits()

	HeaderRow := []string{
		"RecruitID", "First Name", "Last Name", "Position",
		"Archetype", "Stars", "College",
		"High School", "City", "State", "Height",
		"Weight", "Overall", "Potential Grade", "Affinity One", "Affinity Two", "Personality",
		"Recruiting Bias", "Academic Bias", "Work Ethic", "LeadingTeams",
	}

	err := writer.Write(HeaderRow)
	if err != nil {
		log.Fatal("Cannot write header row", err)
	}

	for _, croot := range croots {
		var leadingAbbr []string

		for _, lt := range croot.LeadingTeams {
			leadingAbbr = append(leadingAbbr, lt.TeamAbbr)
		}

		crootRow := []string{
			strconv.Itoa(croot.PlayerID), croot.FirstName, croot.LastName, croot.Position,
			croot.Archetype, strconv.Itoa(croot.Stars), croot.College,
			croot.HighSchool, croot.City, croot.State, strconv.Itoa(croot.Height),
			strconv.Itoa(croot.Weight), croot.OverallGrade, croot.PotentialGrade,
			croot.AffinityOne, croot.AffinityTwo, croot.Personality, croot.RecruitingBias,
			croot.AcademicBias, croot.WorkEthic, strings.Join(leadingAbbr, ", "),
		}

		err = writer.Write(crootRow)
		if err != nil {
			log.Fatal("Cannot write croot row to CSV", err)
		}

		writer.Flush()
		err = writer.Error()
		if err != nil {
			log.Fatal("Error while writing to file ::", err)
		}
	}
}

func ExportDrafteesToCSV(w http.ResponseWriter) {
	w.Header().Set("Content-Disposition", "attachment;filename=2022SimNFLDraftClass.csv")
	w.Header().Set("Transfer-Encoding", "chunked")
	// Initialize writer
	writer := csv.NewWriter(w)

	// Get NFL Draft Class
	draftees := GetAllNFLDraftees()

	HeaderRow := []string{
		"PlayerID", "First Name", "Last Name", "Position",
		"Archetype", "Age", "Stars", "College",
		"High School", "City", "State", "Height",
		"Weight", "Overall", "Speed",
		"Football IQ", "Agility", "Carrying",
		"Catching", "Route Running", "Zone Coverage", "Man Coverage",
		"Strength", "Tackle", "Pass Block", "Run Block",
		"Pass Rush", "Run Defense", "Throw Power", "Throw Accuracy",
		"Kick Power", "Kick Accuracy", "Punt Power", "Punt Accuracy",
		"Stamina", "Injury", "Potential Grade",
	}

	err := writer.Write(HeaderRow)
	if err != nil {
		log.Fatal("Cannot write header row", err)
	}

	for _, player := range draftees {
		csvModel := models.MapNFLDrafteeToModel(player)
		playerRow := []string{
			strconv.Itoa(csvModel.PlayerID), csvModel.FirstName, csvModel.LastName, csvModel.Position,
			csvModel.Archetype, strconv.Itoa(player.Age), strconv.Itoa(player.Stars), player.College,
			player.HighSchool, player.City, player.State, strconv.Itoa(player.Height),
			strconv.Itoa(player.Weight), csvModel.OverallGrade, csvModel.SpeedGrade,
			csvModel.FootballIQGrade, csvModel.AgilityGrade, csvModel.CarryingGrade,
			csvModel.CatchingGrade, csvModel.RouteRunningGrade, csvModel.ZoneCoverageGrade, csvModel.ManCoverageGrade,
			csvModel.StrengthGrade, csvModel.TackleGrade, csvModel.PassBlockGrade, csvModel.RunBlockGrade,
			csvModel.PassRushGrade, csvModel.RunDefenseGrade, csvModel.ThrowPowerGrade, csvModel.ThrowAccuracyGrade,
			csvModel.KickPowerGrade, csvModel.KickAccuracyGrade, csvModel.PuntPowerGrade, csvModel.PuntAccuracyGrade,
			csvModel.StaminaGrade, csvModel.InjuryGrade, csvModel.PotentialGrade,
		}

		err = writer.Write(playerRow)
		if err != nil {
			log.Fatal("Cannot write player row to CSV", err)
		}

		writer.Flush()
		err = writer.Error()
		if err != nil {
			log.Fatal("Error while writing to file ::", err)
		}
	}
}
