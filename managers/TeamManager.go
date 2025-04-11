package managers

import (
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/repository"
	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
)

func GetAllCollegeTeamsForRosterPage() models.RosterPageResponse {

	teams := GetAllCollegeTeams()
	coaches := GetAllCollegeCoaches()

	return models.RosterPageResponse{
		Teams:   teams,
		Coaches: coaches,
	}
}

func GetCollegeTeamMap() map[uint]structs.CollegeTeam {
	teams := GetAllCollegeTeams()
	teamMap := make(map[uint]structs.CollegeTeam)
	for _, t := range teams {
		teamMap[t.ID] = t
	}

	return teamMap
}

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

// GetTeamByTeamID - straightforward
func GetAllNFLTeams() []structs.NFLTeam {
	var teams []structs.NFLTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Preload("Capsheet").Order("team_name asc").Find(&teams).Error
	if err != nil {
		log.Fatal(err)
	}
	return teams
}

// Get NFL Records for Roster Page
func GetNFLRecordsForRosterPage(teamId string) models.NFLRosterPageResponse {

	team := GetNFLTeamWithCapsheetByTeamID(teamId)

	players := GetNFLPlayersForRosterPage(teamId)

	return models.NFLRosterPageResponse{
		Team:   team,
		Roster: players,
	}
}

// GetTeamByTeamID - straightforward
func GetNFLTeamByTeamID(teamId string) structs.NFLTeam {
	var team structs.NFLTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Where("id = ?", teamId).Find(&team).Error
	if err != nil {
		log.Fatal(err)
	}
	return team
}

// GetTeamByTeamID - straightforward
func GetNFLTeamWithCapsheetByTeamID(teamId string) structs.NFLTeam {
	var team structs.NFLTeam
	db := dbprovider.GetInstance().GetDB()
	err := db.Preload("Capsheet").Where("id = ?", teamId).Find(&team).Error
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

func GetNFLTeamByTeamIDForSim(id string) structs.NFLTeam {
	db := dbprovider.GetInstance().GetDB()

	var team structs.NFLTeam

	err := db.Preload("TeamGameplan").Preload("TeamDepthChart.DepthChartPlayers").Where("id = ?", id).Find(&team).Error
	if err != nil {
		log.Panicln("Could not find team by given id:"+id+"\n", err)
	}

	return team
}

func GetNFLTeamByTeamAbbr(abbr string) structs.NFLTeam {
	db := dbprovider.GetInstance().GetDB()

	var team structs.NFLTeam

	err := db.Preload("TeamGameplan").
		Preload("TeamDepthChart.DepthChartPlayers").Where("team_abbr = ?", abbr).Find(&team).Error
	if err != nil {
		log.Panicln("Could not find team by given abbreviation:"+abbr+"\n", err)
	}

	return team
}

func GetAllCollegeTeamsWithRecruitingProfileAndCoach() []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	db.Preload("CollegeCoach").Preload("RecruitingProfile.Affinities").Find(&teams)

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

func GetCollegeConferences() []structs.CollegeConference {
	db := dbprovider.GetInstance().GetDB()

	var conferences []structs.CollegeConference

	db.Preload("Divisions").Find(&conferences)

	return conferences
}

func GetCollegeTeamsByConference(conf string) []structs.CollegeTeam {
	db := dbprovider.GetInstance().GetDB()

	var teams []structs.CollegeTeam

	db.Where("conference = ?", conf).Find(&teams)

	return teams
}

func GetDashboardByTeamID(isCFB bool, teamID string) structs.DashboardResponseData {
	ts := GetTimestamp()
	_, cfbGT := ts.GetCFBCurrentGameType()
	_, nflGT := ts.GetNFLCurrentGameType()
	seasonID := strconv.Itoa(ts.CollegeSeasonID)
	collegeTeam := structs.CollegeTeam{}
	nflTeam := structs.NFLTeam{}
	if isCFB {
		collegeTeam = GetTeamByTeamID(teamID)
	} else {
		nflTeam = GetNFLTeamByTeamID(teamID)
	}
	cStandings := make(chan []structs.CollegeStandings)
	nStandings := make(chan []structs.NFLStandings)
	cGames := make(chan []structs.CollegeGame)
	nGames := make(chan []structs.NFLGame)
	newsChan := make(chan []structs.NewsLog)
	cfbPlayerChan := make(chan []structs.CollegePlayerResponse)
	nflPlayerChan := make(chan []structs.NFLPlayerResponse)
	cfbTeamStatsChan := make(chan structs.CollegeTeamSeasonStats)
	nflTeamStatsChan := make(chan structs.NFLTeamSeasonStats)
	pollChan := make(chan structs.CollegePollOfficial)

	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	go func() {
		waitGroup.Wait()
		close(cStandings)
		close(nStandings)
		close(cGames)
		close(nGames)
		close(newsChan)
		close(cfbPlayerChan)
		close(nflPlayerChan)
		close(cfbTeamStatsChan)
		close(nflTeamStatsChan)
		close(pollChan)
	}()

	go func() {
		defer waitGroup.Done()
		cSt := []structs.CollegeStandings{}
		if isCFB {
			cSt = GetStandingsByConferenceIDAndSeasonID(strconv.Itoa(collegeTeam.ConferenceID), seasonID)
		}
		cStandings <- cSt
	}()

	go func() {
		defer waitGroup.Done()
		nSt := []structs.NFLStandings{}
		if !isCFB {
			nSt = GetNFLStandingsByDivisionIDAndSeasonID(strconv.Itoa(int(nflTeam.DivisionID)), seasonID)
		}
		nStandings <- nSt
	}()

	go func() {
		defer waitGroup.Done()
		cG := []structs.CollegeGame{}
		if isCFB {
			cG = GetCollegeGamesByTeamIdAndSeasonId(teamID, seasonID, ts.CFBSpringGames)
		}
		cGames <- cG
	}()

	go func() {
		defer waitGroup.Done()
		nG := []structs.NFLGame{}
		if !isCFB {
			nG = GetNFLGamesByTeamIdAndSeasonId(teamID, seasonID)
		}
		nGames <- nG
	}()

	go func() {
		defer waitGroup.Done()
		nL := []structs.NewsLog{}
		if isCFB {
			nL = GetCFBRelatedNews(teamID)
		} else {
			nL = GetNFLRelatedNews(teamID)
		}
		newsChan <- nL
	}()

	go func() {
		defer waitGroup.Done()
		players := []structs.CollegePlayerResponse{}
		if isCFB {
			seasonKey := ts.CollegeSeasonID
			if ts.IsOffSeason {
				seasonKey -= 1
			}
			players = GetAllCollegePlayersWithSeasonStatsByTeamID(teamID, strconv.Itoa(seasonKey), cfbGT)
		}
		cfbPlayerChan <- players
	}()

	go func() {
		defer waitGroup.Done()
		players := []structs.NFLPlayerResponse{}
		if !isCFB {
			seasonKey := ts.NFLSeasonID
			if ts.IsNFLOffSeason {
				seasonKey -= 1
			}
			players = GetAllNFLPlayersWithSeasonStatsByTeamID(teamID, strconv.Itoa(seasonKey), nflGT)
		}
		nflPlayerChan <- players
	}()

	go func() {
		defer waitGroup.Done()
		stats := structs.CollegeTeamSeasonStats{}
		if isCFB {
			seasonKey := ts.CollegeSeasonID
			if ts.IsOffSeason {
				seasonKey -= 1
			}
			stats = GetCollegeTeamSeasonStatsBySeason(teamID, strconv.Itoa(seasonKey), cfbGT)
		}
		cfbTeamStatsChan <- stats
	}()

	go func() {
		defer waitGroup.Done()
		stats := structs.NFLTeamSeasonStats{}
		if !isCFB {
			seasonKey := ts.NFLSeasonID
			if ts.IsNFLOffSeason {
				seasonKey -= 1
			}
			stats = GetNFLTeamSeasonStatsByTeamANDSeason(teamID, strconv.Itoa(seasonKey), nflGT)
		}
		nflTeamStatsChan <- stats
	}()

	go func() {
		defer waitGroup.Done()
		poll := structs.CollegePollOfficial{}
		if isCFB {
			seasonKey := ts.NFLSeasonID
			if ts.IsOffSeason {
				seasonKey -= 1
			}
			polls := GetOfficialPollBySeasonID(strconv.Itoa(seasonKey))
			if len(polls) > 0 {
				poll = polls[len(polls)-1]
			}
		}
		pollChan <- poll
	}()

	collegeStandings := <-cStandings
	nflStandings := <-nStandings
	collegeGames := <-cGames
	nflGames := <-nGames
	newsLogs := <-newsChan
	collegePlayers := <-cfbPlayerChan
	nflPlayers := <-nflPlayerChan
	cfbTeamStats := <-cfbTeamStatsChan
	nflTeamStats := <-nflTeamStatsChan
	collegePoll := <-pollChan

	return structs.DashboardResponseData{
		CollegeStandings: collegeStandings,
		NFLStandings:     nflStandings,
		CollegeGames:     collegeGames,
		NFLGames:         nflGames,
		NewsLogs:         newsLogs,
		TopCFBPlayers:    collegePlayers,
		TopNFLPlayers:    nflPlayers,
		CFBTeamStats:     cfbTeamStats,
		NFLTeamStats:     nflTeamStats,
		TopTenPoll:       collegePoll,
	}
}

// Returns the CFB team's numerical value for their entire offense
func OffenseGradeCFB(depthChartPlayers structs.CollegeTeamDepthChart, gameplan structs.CollegeGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11.5 (offense weight normalization value)
	// return the resulting value
	return 0.0
}

// Returns the CFB team's numerical value for their entire offense
func DefenseGradeCFB(depthChartPlayers structs.CollegeTeamDepthChart, gameplan structs.CollegeGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11 (defense weight normalization value)
	// return the resulting value
	return 0.0
}

// Returns the CFB team's numerical value for their entire offense
func STGradeCFB(depthChartPlayers structs.CollegeTeamDepthChart) float64 {
	// Get overall values for all relevant positions
	// Weight them by position
	// Sum them all up
	// Divide by 5 (Special Teams weight normalization value)
	// return the resulting value
	return 0.0
}

// Returns the CFB team's numerical value for their entire offense
func OffenseGradeNFL(depthChartPlayers structs.NFLDepthChart, gameplan structs.NFLGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11.5 (offense weight normalization value)
	// return the resulting value
	return 0.0
}

// Returns the CFB team's numerical value for their entire offense
func DefenseGradeNFL(depthChartPlayers structs.NFLDepthChart, gameplan structs.NFLGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11 (defense weight normalization value)
	// return the resulting value
	return 0.0
}

// Returns the CFB team's numerical value for their entire offense
func STGradeNFL(depthChartPlayers structs.NFLDepthChart) float64 {
	// Get overall values for all relevant positions
	// Weight them by position
	// Sum them all up
	// Divide by 5 (Special Teams weight normalization value)
	// return the resulting value
	return 0.0
}

func GetOffensePositionGradeWeight(position string, scheme string) float64 {
	scheme = strings.ToLower(scheme)
	switch position {
	case "QB1":
		if (strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread"))
		{
			return 1.5
		} else if (strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone")) {
			return 0.5
		} else {
			return 1.0
		}
	case "RB1":
		if (strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "double"))
		{
			return 0.8
		} else {
			return 1.0
		}
	}
}

// League agnostic
func OverallGrade(offense float64, defense float64, specialTeams float64) float64 {
	var overallGrade float64 = offense * 0.45
	overallGrade = overallGrade + (defense * 0.45)
	overallGrade = overallGrade + (specialTeams * 0.1)
	return overallGrade
}

func TeamLetterGrade(value float64, mean float64, stdDev float64) string {
	// Assign a letter grade for each grade value by comparing the value to the following:
	if value > (mean + (2 * stdDev)) {
		// A+: 2.0+ std dev above the mean
		return "A+"
	} else if value > (mean + (1.75 * stdDev)) {
		// A: between 1.75-2.0 std dev above the mean
		return "A"
	} else if value > (mean + (1.5 * stdDev)) {
		// A-: between 1.5-1.75 std dev above the mean
		return "A-"
	} else if value > (mean + (1.25 * stdDev)) {
		// B+: between 1.25-1.5 std dev above the mean
		return "B+"
	} else if value > (mean + (1.0 * stdDev)) {
		// B: between 1.0-1.25 std dev above the mean
		return "B"
	} else if value > (mean + (0.75 * stdDev)) {
		// B-: between .75-1.0 std dev above the mean
		return "B-"
	} else if value > (mean + (0.5 * stdDev)) {
		// C+: between .5-.75 std dev above the mean
		return "C+"
	} else if value > (mean - (0.5 * stdDev)) {
		// C: between +/- .5 std dev from mean
		return "C"
	} else if value > (mean - (0.75 * stdDev)) {
		// C-: between .5-.75 std dev below the mean
		return "C-"
	} else if value > (mean - (1.0 * stdDev)) {
		// D+: between .75-1.0 std dev below the mean
		return "D+"
	} else if value > (mean - (1.5 * stdDev)) {
		// D: between 1.0-1.5 std dev below the mean
		return "D"
	} else if value > (mean - (2.0 * stdDev)) {
		// D-: between 1.5-2.0 std dev below the mean
		return "D-"
	} else {
		// F: 2.0+ std dev below the mean
		return "F"
	}
}

// This function should be called weekly, once 2.0 is released.
func AssignTeamGrades() {
	db := dbprovider.GetInstance().GetDB()

	// College
	collegeTeams := GetAllCollegeTeams()
	collegeDepthChartMap := GetDepthChartMap()
	collegeGameplanMap := GetCollegeGameplanMap()
	collegeTeamGrades := make(map[uint]structs.TeamGrade)

	for _, t := range collegeTeams {
		if !t.IsActive {
			continue
		}
		depthChart := collegeDepthChartMap[t.ID]
		gameplan := collegeGameplanMap[t.ID]
		offenseGrade := OffenseGradeCFB(depthChart, gameplan)
		defenseGrade := DefenseGradeCFB(depthChart, gameplan)
		STGrade := STGradeCFB(depthChart)

		collegeTeamGrades[t.ID] = structs.TeamGrade{
			OffenseGradeNumber:      offenseGrade,
			DefenseGradeNumber:      defenseGrade,
			SpecialTeamsGradeNumber: STGrade,
			OverallGradeNumber:      OverallGrade(offenseGrade, defenseGrade, STGrade),
			OffenseGradeLetter:      "",
			DefenseGradeLetter:      "",
			SpecialTeamsGradeLetter: "",
			OverallGradeLetter:      "",
		}
	}

	// Determine the mean and std dev for the data set contained in the collegeTeamGrades map
	offenseMean := 0.0
	defenseMean := 0.0
	stMean := 0.0
	overallMean := 0.0
	offenseVar := 0.0
	defenseVar := 0.0
	stVar := 0.0
	overallVar := 0.0
	dataLength := len(collegeTeamGrades)

	// Mean for all values
	for _, element := range collegeTeamGrades {
		offenseMean += element.OffenseGradeNumber
		defenseMean += element.DefenseGradeNumber
		stMean += element.SpecialTeamsGradeNumber
		overallMean += element.OverallGradeNumber
	}

	offenseMean = (offenseMean / float64(dataLength))
	defenseMean = (defenseMean / float64(dataLength))
	stMean = (stMean / float64(dataLength))
	overallMean = (overallMean / float64(dataLength))

	// Variance for all values
	for _, element := range collegeTeamGrades {
		offenseDiff := element.OffenseGradeNumber - offenseMean
		defenseDiff := element.DefenseGradeNumber - defenseMean
		stDiff := element.SpecialTeamsGradeNumber - stMean
		overallDiff := element.OffenseGradeNumber - overallMean

		offenseVar += (offenseDiff * offenseDiff)
		defenseVar += (defenseDiff * defenseDiff)
		stVar += (stDiff * stDiff)
		overallVar += (overallDiff * overallDiff)
	}

	offenseVar = (offenseVar / float64(dataLength-1))
	defenseVar = (defenseVar / float64(dataLength-1))
	stVar = (stVar / float64(dataLength-1))
	overallVar = (overallVar / float64(dataLength-1))

	// Std Dev for all values
	offenseStdDev := math.Sqrt(offenseVar)
	defenseStdDev := math.Sqrt(defenseVar)
	stStdDev := math.Sqrt(stVar)
	overallStdDev := math.Sqrt(overallVar)

	// Iterate back through the map and set the letter grades based on the number grades' relationship to the mean and std dev of the entire data set for that value
	for _, collegeTeam := range collegeTeamGrades {
		collegeTeam.SetOffenseGradeLetter(TeamLetterGrade(collegeTeam.OffenseGradeNumber, offenseMean, offenseStdDev))
		collegeTeam.SetDefenseGradeLetter(TeamLetterGrade(collegeTeam.DefenseGradeNumber, defenseMean, defenseStdDev))
		collegeTeam.SetSpecialTeamsGradeLetter(TeamLetterGrade(collegeTeam.SpecialTeamsGradeNumber, stMean, stStdDev))
		collegeTeam.SetOverallGradeLetter(TeamLetterGrade(collegeTeam.OverallGradeNumber, overallMean, overallStdDev))
	}

	// Assign those letter grades to that team's grade properties
	for _, collegeTeam := range collegeTeams {
		collegeTeam.AssignTeamGrades(collegeTeamGrades[collegeTeam.ID].OverallGradeLetter, collegeTeamGrades[collegeTeam.ID].OffenseGradeLetter,
			collegeTeamGrades[collegeTeam.ID].DefenseGradeLetter, collegeTeamGrades[collegeTeam.ID].SpecialTeamsGradeLetter)

		repository.SaveCFBTeam(collegeTeam, db)
	}

	// NFL
	nflTeams := GetAllNFLTeams()
	nflDepthChartList := GetAllNFLDepthcharts()
	nflGameplanList := GetAllNFLGameplans()
	nflDepthChartMap := make(map[uint]structs.NFLDepthChart)
	nflGameplanMap := make(map[uint]structs.NFLGameplan)

	for _, t := range nflDepthChartList {
		nflDepthChartMap[t.ID] = t
	}

	for _, t := range nflGameplanList {
		nflGameplanMap[t.ID] = t
	}

	nflTeamGrades := make(map[uint]structs.TeamGrade)

	// CHANGE ALL REFERENCES TO COLLEGE TO NFL FROM HERE ON
	for _, t := range nflTeams {
		depthChart := nflDepthChartMap[t.ID]
		gameplan := nflGameplanMap[t.ID]
		offenseGrade := OffenseGradeNFL(depthChart, gameplan)
		defenseGrade := DefenseGradeNFL(depthChart, gameplan)
		STGrade := STGradeNFL(depthChart)

		nflTeamGrades[t.ID] = structs.TeamGrade{
			OffenseGradeNumber:      offenseGrade,
			DefenseGradeNumber:      defenseGrade,
			SpecialTeamsGradeNumber: STGrade,
			OverallGradeNumber:      OverallGrade(offenseGrade, defenseGrade, STGrade),
			OffenseGradeLetter:      "",
			DefenseGradeLetter:      "",
			SpecialTeamsGradeLetter: "",
			OverallGradeLetter:      "",
		}
	}

	// Determine the mean and std dev for the data set contained in the nflTeamGrades map
	offenseMean = 0.0
	defenseMean = 0.0
	stMean = 0.0
	overallMean = 0.0
	offenseVar = 0.0
	defenseVar = 0.0
	stVar = 0.0
	overallVar = 0.0
	dataLength = len(nflTeamGrades)

	// Mean for all values
	for _, element := range nflTeamGrades {
		offenseMean += element.OffenseGradeNumber
		defenseMean += element.DefenseGradeNumber
		stMean += element.SpecialTeamsGradeNumber
		overallMean += element.OverallGradeNumber
	}

	offenseMean = (offenseMean / float64(dataLength))
	defenseMean = (defenseMean / float64(dataLength))
	stMean = (stMean / float64(dataLength))
	overallMean = (overallMean / float64(dataLength))

	// Variance for all values
	for _, element := range nflTeamGrades {
		offenseDiff := element.OffenseGradeNumber - offenseMean
		defenseDiff := element.DefenseGradeNumber - defenseMean
		stDiff := element.SpecialTeamsGradeNumber - stMean
		overallDiff := element.OffenseGradeNumber - overallMean

		offenseVar += (offenseDiff * offenseDiff)
		defenseVar += (defenseDiff * defenseDiff)
		stVar += (stDiff * stDiff)
		overallVar += (overallDiff * overallDiff)
	}

	offenseVar = (offenseVar / float64(dataLength-1))
	defenseVar = (defenseVar / float64(dataLength-1))
	stVar = (stVar / float64(dataLength-1))
	overallVar = (overallVar / float64(dataLength-1))

	// Std Dev for all values
	offenseStdDev = math.Sqrt(offenseVar)
	defenseStdDev = math.Sqrt(defenseVar)
	stStdDev = math.Sqrt(stVar)
	overallStdDev = math.Sqrt(overallVar)

	// Iterate back through the map and set the letter grades based on the number grades' relationship to the mean and std dev of the entire data set for that value
	for _, nflTeam := range nflTeamGrades {
		nflTeam.SetOffenseGradeLetter(TeamLetterGrade(nflTeam.OffenseGradeNumber, offenseMean, offenseStdDev))
		nflTeam.SetDefenseGradeLetter(TeamLetterGrade(nflTeam.DefenseGradeNumber, defenseMean, defenseStdDev))
		nflTeam.SetSpecialTeamsGradeLetter(TeamLetterGrade(nflTeam.SpecialTeamsGradeNumber, stMean, stStdDev))
		nflTeam.SetOverallGradeLetter(TeamLetterGrade(nflTeam.OverallGradeNumber, overallMean, overallStdDev))
	}

	// Assign those letter grades to that team's grade properties
	for _, nflTeam := range nflTeams {
		nflTeam.AssignTeamGrades(nflTeamGrades[nflTeam.ID].OverallGradeLetter, nflTeamGrades[nflTeam.ID].OffenseGradeLetter,
			nflTeamGrades[nflTeam.ID].DefenseGradeLetter, nflTeamGrades[nflTeam.ID].SpecialTeamsGradeLetter)

		repository.SaveNFLTeam(nflTeam, db)
	}
}
