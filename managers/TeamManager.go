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
	"github.com/CalebRose/SimFBA/util"
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

func GetCollegePlayer(depthChartPlayers structs.CollegeTeamDepthChart, position string, level int) structs.CollegePlayer {
	for _, player := range depthChartPlayers.DepthChartPlayers {
		if player.Position == position && util.ConvertStringToInt(player.PositionLevel) == level {
			return player.CollegePlayer
		}
	}

	return structs.CollegePlayer{}
}

func GetNFLPlayer(depthChartPlayers structs.NFLDepthChart, position string, level int) structs.NFLPlayer {
	for _, player := range depthChartPlayers.DepthChartPlayers {
		if player.Position == position && util.ConvertStringToInt(player.PositionLevel) == level {
			return player.NFLPlayer
		}
	}

	return structs.NFLPlayer{}
}

// Returns the CFB team's numerical value for their entire offense
func OffenseGradeCFB(depthChartPlayers structs.CollegeTeamDepthChart, gameplan structs.CollegeGameplan) float64 {
	// Get overall values for all relevant positions
	qb1 := GetCollegePlayer(depthChartPlayers, "QB", 1)
	log.Println("ERROR DURING TEAM GRADING: QB1 was not found!!! Player ID: " + strconv.Itoa(int(qb1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	rb1 := GetCollegePlayer(depthChartPlayers, "RB", 1)
	log.Println("ERROR DURING TEAM GRADING: RB1 was not found!!! Player ID: " + strconv.Itoa(int(rb1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	rb2 := GetCollegePlayer(depthChartPlayers, "RB", 2)
	log.Println("ERROR DURING TEAM GRADING: RB2 was not found!!! Player ID: " + strconv.Itoa(int(rb2.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	rb3 := GetCollegePlayer(depthChartPlayers, "RB", 3)
	log.Println("ERROR DURING TEAM GRADING: RB3 was not found!!! Player ID: " + strconv.Itoa(int(rb3.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	fb1 := GetCollegePlayer(depthChartPlayers, "FB", 1)
	log.Println("ERROR DURING TEAM GRADING: FB1 was not found!!! Player ID: " + strconv.Itoa(int(fb1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	te1 := GetCollegePlayer(depthChartPlayers, "TE", 1)
	log.Println("ERROR DURING TEAM GRADING: TE1 was not found!!! Player ID: " + strconv.Itoa(int(te1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	te2 := GetCollegePlayer(depthChartPlayers, "TE", 2)
	log.Println("ERROR DURING TEAM GRADING: TE2 was not found!!! Player ID: " + strconv.Itoa(int(te2.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	wr1 := GetCollegePlayer(depthChartPlayers, "WR", 1)
	log.Println("ERROR DURING TEAM GRADING: WR1 was not found!!! Player ID: " + strconv.Itoa(int(wr1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	wr2 := GetCollegePlayer(depthChartPlayers, "WR", 2)
	log.Println("ERROR DURING TEAM GRADING: WR2 was not found!!! Player ID: " + strconv.Itoa(int(wr2.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	wr3 := GetCollegePlayer(depthChartPlayers, "WR", 3)
	log.Println("ERROR DURING TEAM GRADING: WR3 was not found!!! Player ID: " + strconv.Itoa(int(wr3.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	wr4 := GetCollegePlayer(depthChartPlayers, "WR", 4)
	log.Println("ERROR DURING TEAM GRADING: WR4 was not found!!! Player ID: " + strconv.Itoa(int(wr4.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	wr5 := GetCollegePlayer(depthChartPlayers, "WR", 5)
	log.Println("ERROR DURING TEAM GRADING: WR5 was not found!!! Player ID: " + strconv.Itoa(int(wr5.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	lt1 := GetCollegePlayer(depthChartPlayers, "LT", 1)
	log.Println("ERROR DURING TEAM GRADING: LT1 was not found!!! Player ID: " + strconv.Itoa(int(lt1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	lg1 := GetCollegePlayer(depthChartPlayers, "LG", 1)
	log.Println("ERROR DURING TEAM GRADING: LG1 was not found!!! Player ID: " + strconv.Itoa(int(lg1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	c1 := GetCollegePlayer(depthChartPlayers, "C", 1)
	log.Println("ERROR DURING TEAM GRADING: C1 was not found!!! Player ID: " + strconv.Itoa(int(c1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	rg1 := GetCollegePlayer(depthChartPlayers, "RG", 1)
	log.Println("ERROR DURING TEAM GRADING: RG1 was not found!!! Player ID: " + strconv.Itoa(int(rg1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))
	rt1 := GetCollegePlayer(depthChartPlayers, "RT", 1)
	log.Println("ERROR DURING TEAM GRADING: RT1 was not found!!! Player ID: " + strconv.Itoa(int(rt1.ID)) + " Team ID: " + strconv.Itoa(int(depthChartPlayers.TeamID)))

	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// ENSURE TO TRANSLATE THEIR POSITION PROPERLY FOR THE SCHEME FIT!!!
	qb1Overall := ApplySchemeModifiers(float64(qb1.Overall), "QB", qb1.Archetype, gameplan.OffensiveScheme)
	rb1Overall := ApplySchemeModifiers(float64(rb1.Overall), "RB", rb1.Archetype, gameplan.OffensiveScheme)
	rb2Overall := ApplySchemeModifiers(float64(rb2.Overall), "RB", rb2.Archetype, gameplan.OffensiveScheme)
	rb3Overall := ApplySchemeModifiers(float64(rb3.Overall), "RB", rb3.Archetype, gameplan.OffensiveScheme)
	fb1Overall := ApplySchemeModifiers(float64(fb1.Overall), "FB", fb1.Archetype, gameplan.OffensiveScheme)
	te1Overall := ApplySchemeModifiers(float64(te1.Overall), "TE", te1.Archetype, gameplan.OffensiveScheme)
	te2Overall := ApplySchemeModifiers(float64(te2.Overall), "TE", te2.Archetype, gameplan.OffensiveScheme)
	wr1Overall := ApplySchemeModifiers(float64(wr1.Overall), "WR", wr1.Archetype, gameplan.OffensiveScheme)
	wr2Overall := ApplySchemeModifiers(float64(wr2.Overall), "WR", wr2.Archetype, gameplan.OffensiveScheme)
	wr3Overall := ApplySchemeModifiers(float64(wr3.Overall), "WR", wr3.Archetype, gameplan.OffensiveScheme)
	wr4Overall := ApplySchemeModifiers(float64(wr4.Overall), "WR", wr4.Archetype, gameplan.OffensiveScheme)
	wr5Overall := ApplySchemeModifiers(float64(wr5.Overall), "WR", wr5.Archetype, gameplan.OffensiveScheme)
	lt1Overall := ApplySchemeModifiers(float64(lt1.Overall), "OT", lt1.Archetype, gameplan.OffensiveScheme)
	lg1Overall := ApplySchemeModifiers(float64(lg1.Overall), "OG", lg1.Archetype, gameplan.OffensiveScheme)
	c1Overall := ApplySchemeModifiers(float64(c1.Overall), "C", c1.Archetype, gameplan.OffensiveScheme)
	rg1Overall := ApplySchemeModifiers(float64(rg1.Overall), "RG", rg1.Archetype, gameplan.OffensiveScheme)
	rt1Overall := ApplySchemeModifiers(float64(rt1.Overall), "RT", rt1.Archetype, gameplan.OffensiveScheme)

	// Depending on scheme, weight them
	qb1Overall = qb1Overall * GetOffensePositionGradeWeight("QB1", gameplan.OffensiveScheme)
	rb1Overall = rb1Overall * GetOffensePositionGradeWeight("RB1", gameplan.OffensiveScheme)
	rb2Overall = rb2Overall * GetOffensePositionGradeWeight("RB2", gameplan.OffensiveScheme)
	rb3Overall = rb3Overall * GetOffensePositionGradeWeight("RB3", gameplan.OffensiveScheme)
	fb1Overall = fb1Overall * GetOffensePositionGradeWeight("FB1", gameplan.OffensiveScheme)
	te1Overall = te1Overall * GetOffensePositionGradeWeight("TE1", gameplan.OffensiveScheme)
	te2Overall = te2Overall * GetOffensePositionGradeWeight("TE2", gameplan.OffensiveScheme)
	wr1Overall = wr1Overall * GetOffensePositionGradeWeight("WR1", gameplan.OffensiveScheme)
	wr2Overall = wr2Overall * GetOffensePositionGradeWeight("WR2", gameplan.OffensiveScheme)
	wr3Overall = wr3Overall * GetOffensePositionGradeWeight("WR3", gameplan.OffensiveScheme)
	wr4Overall = wr4Overall * GetOffensePositionGradeWeight("WR4", gameplan.OffensiveScheme)
	wr5Overall = wr5Overall * GetOffensePositionGradeWeight("WR5", gameplan.OffensiveScheme)
	lt1Overall = lt1Overall * GetOffensePositionGradeWeight("LT1", gameplan.OffensiveScheme)
	lg1Overall = lg1Overall * GetOffensePositionGradeWeight("LG1", gameplan.OffensiveScheme)
	c1Overall = c1Overall * GetOffensePositionGradeWeight("C1", gameplan.OffensiveScheme)
	rg1Overall = rg1Overall * GetOffensePositionGradeWeight("RG1", gameplan.OffensiveScheme)
	rt1Overall = rt1Overall * GetOffensePositionGradeWeight("RT1", gameplan.OffensiveScheme)

	// Sum them all up
	grade := qb1Overall + rb1Overall + rb2Overall + rb3Overall + fb1Overall + te1Overall + te2Overall + wr1Overall + wr2Overall + wr3Overall + wr4Overall + wr5Overall + lt1Overall + lg1Overall + c1Overall + rg1Overall + rt1Overall
	// Divide by 11.5 (offense weight normalization value)
	grade = grade / 11.5
	// return the resulting value
	return grade
}

// Returns the CFB team's numerical value for their entire offense
func DefenseGradeCFB(depthChartPlayers structs.CollegeTeamDepthChart, gameplan structs.CollegeGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11 (defense weight normalization value)
	// return the resulting value
	return
}

// Returns the CFB team's numerical value for their entire offense
func STGradeCFB(depthChartPlayers structs.CollegeTeamDepthChart) float64 {
	// Get overall values for all relevant positions
	// Weight them by position
	// Sum them all up
	// Divide by 5 (Special Teams weight normalization value)
	// return the resulting value
	return
}

// Returns the CFB team's numerical value for their entire offense
func OffenseGradeNFL(depthChartPlayers structs.NFLDepthChart, gameplan structs.NFLGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11.5 (offense weight normalization value)
	// return the resulting value
	return
}

// Returns the CFB team's numerical value for their entire offense
func DefenseGradeNFL(depthChartPlayers structs.NFLDepthChart, gameplan structs.NFLGameplan) float64 {
	// Get overall values for all relevant positions
	// If the player is a scheme fit, give them a bonus, if they are a bad fit, give them a malus
	// Depending on scheme, weight them
	// Sum them all up
	// Divide by 11 (defense weight normalization value)
	// return the resulting value
	return
}

// Returns the CFB team's numerical value for their entire offense
func STGradeNFL(depthChartPlayers structs.NFLDepthChart) float64 {
	// Get overall values for all relevant positions
	// Weight them by position
	// Sum them all up
	// Divide by 5 (Special Teams weight normalization value)
	// return the resulting value
	return
}

// League agnostic
func OverallGrade(offense float64, defense float64, specialTeams float64) float64 {
	var overallGrade float64 = offense * 0.45
	overallGrade = overallGrade + (defense * 0.45)
	overallGrade = overallGrade + (specialTeams * 0.1)
	return overallGrade
}

func ApplySchemeModifiers(overall float64, position string, archetype string, scheme string) float64 {
	if IsSchemeFit(position, archetype, scheme) {
		return (overall * 1.05)
	} else if IsBadFit(position, archetype, scheme) {
		return (overall * 0.95)
	} else {
		return overall
	}
}

func IsSchemeFit(position string, archetype string, scheme string) bool {
	scheme = strings.ToLower(scheme)
	archetype = strings.ToLower(archetype)
	switch position {
	case "QB":
		switch archetype {
		case "scrambler":
			if strings.Contains(scheme, "option") || strings.Contains(scheme, "flexbone") {
				return true
			} else {
				return false
			}
		case "balanced":
			if strings.Contains(scheme, "wing-t") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "wishbone") {
				return true
			} else {
				return false
			}
		case "pocket":
			if strings.Contains(scheme, "raid") || strings.Contains(scheme, "vert") || strings.Contains(scheme, "pistol") {
				return true
			} else {
				return false
			}
		case "field":
			if strings.Contains(scheme, "west") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "wish") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "RB":
		switch archetype {
		case "balanced":
			if strings.Contains(scheme, "pistol") || strings.Contains(scheme, "wing-t") || strings.Contains(scheme, "wish") {
				return true
			} else {
				return false
			}
		case "receiving":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "spread") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "flex") {
				return true
			} else {
				return false
			}
		case "power":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "i option") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "FB":
		switch archetype {
		case "balanced":
			if strings.Contains(scheme, "west") || strings.Contains(scheme, "wing-t") || strings.Contains(scheme, "flex") {
				return true
			} else {
				return false
			}
		case "receiving":
			if strings.Contains(scheme, "west") || strings.Contains(scheme, "spread") || strings.Contains(scheme, "vert") {
				return true
			} else {
				return false
			}
		case "rushing":
			if strings.Contains(scheme, "i option") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "WR":
		switch archetype {
		case "route runner":
			if strings.Contains(scheme, "west") || strings.Contains(scheme, "vert") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
				return true
			} else {
				return false
			}
		case "red zone threat":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") || strings.Contains(scheme, "bone") {
				return true
			} else {
				return false
			}
		case "possession":
			if strings.Contains(scheme, "option") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "west") {
				return true
			} else {
				return false
			}
		case "possesion":
			if strings.Contains(scheme, "option") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "west") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "wing-t") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "TE":
		switch archetype {
		case "vertical threat":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") {
				return true
			} else {
				return false
			}
		case "receiving":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "west") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") || strings.Contains(scheme, "i option") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "OT":
		switch archetype {
		case "run blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "pass blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "OG":
		switch archetype {
		case "run blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "pass blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "C":
		switch archetype {
		case "run blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "pass blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "line captain":
			if strings.Contains(scheme, "west") || strings.Contains(scheme, "shoot") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "DE":
		switch archetype {
		case "speed rusher":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		case "run stopper":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "DT":
		switch archetype {
		case "pass rusher":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		case "nose tackle":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "OLB":
		switch archetype {
		case "pass rush":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "speed") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "coverage":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		case "run stopper":
			if strings.Contains(scheme, "old") || strings.Contains(scheme, "2") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "ILB":
		switch archetype {
		case "field general":
			if strings.Contains(scheme, "old") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "speed") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "coverage":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		case "run stopper":
			if strings.Contains(scheme, "old") || strings.Contains(scheme, "2") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	default:
		return false
	}
}

func IsBadFit(position string, archetype string, scheme string) bool {
	scheme = strings.ToLower(scheme)
	archetype = strings.ToLower(archetype)
	switch position {
	case "QB":
		switch archetype {
		case "scrambler":
			if strings.Contains(scheme, "raid") || strings.Contains(scheme, "vert") {
				return true
			} else {
				return false
			}
		case "balanced":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "pocket":
			if strings.Contains(scheme, "i option") || strings.Contains(scheme, "flex") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "field":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "RB":
		switch archetype {
		case "balanced":
			if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "spread") || strings.Contains(scheme, "flex") {
				return true
			} else {
				return false
			}
		case "receiving":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "i option") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "i option") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "power":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "FB":
		switch archetype {
		case "balanced":
			if strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") || strings.Contains(scheme, "wish") {
				return true
			} else {
				return false
			}
		case "receiving":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "i option") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "rushing":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "west") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "WR":
		switch archetype {
		case "route runner":
			if strings.Contains(scheme, "wish") {
				return true
			} else {
				return false
			}
		case "red zone threat":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "west") {
				return true
			} else {
				return false
			}
		case "possession":
			if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "flex") {
				return true
			} else {
				return false
			}
		case "possesion":
			if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "flex") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "flex") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "TE":
		switch archetype {
		case "vertical threat":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "i option") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "receiving":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "i option") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "OT":
		switch archetype {
		case "run blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "pass blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "OG":
		switch archetype {
		case "run blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "pass blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "C":
		switch archetype {
		case "run blocking":
			if strings.Contains(scheme, "vert") || strings.Contains(scheme, "raid") {
				return true
			} else {
				return false
			}
		case "pass blocking":
			if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
				return true
			} else {
				return false
			}
		case "line captain":
			if strings.Contains(scheme, "pistol") || strings.Contains(scheme, "wish") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "DE":
		switch archetype {
		case "speed rusher":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "run stopper":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "DT":
		switch archetype {
		case "pass rusher":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "nose tackle":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "OLB":
		switch archetype {
		case "pass rush":
			if strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		case "coverage":
			if strings.Contains(scheme, "old") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "run stopper":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	case "ILB":
		switch archetype {
		case "field general":
			if strings.Contains(scheme, "3") || strings.Contains(scheme, "speed") {
				return true
			} else {
				return false
			}
		case "speed":
			if strings.Contains(scheme, "2") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		case "coverage":
			if strings.Contains(scheme, "old") || strings.Contains(scheme, "multiple") {
				return true
			} else {
				return false
			}
		case "run stopper":
			if strings.Contains(scheme, "4") || strings.Contains(scheme, "3") {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	default:
		return false
	}
}

func GetOffensePositionGradeWeight(position string, scheme string) float64 {
	scheme = strings.ToLower(scheme)
	switch position {
	case "QB1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
			return 1.5
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 0.5
		} else {
			return 1.0
		}
	case "RB1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "double") {
			return 0.8
		} else {
			return 1.0
		}
	case "RB2":
		if strings.Contains(scheme, "wing-t") || strings.Contains(scheme, "bone") {
			return 1.0
		} else if strings.Contains(scheme, "double") {
			return 0.8
		} else if strings.Contains(scheme, "spread") {
			return 0.7
		} else if strings.Contains(scheme, "raid") {
			return 0.6
		} else if strings.Contains(scheme, "shoot") {
			return 0.4
		} else {
			return 0.5
		}
	case "RB3":
		if strings.Contains(scheme, "wing-t") || strings.Contains(scheme, "wishbone") {
			return 0.1
		} else if strings.Contains(scheme, "flexbone") {
			return 0.5
		} else if strings.Contains(scheme, "double") {
			return 0.3
		} else {
			return 0.0
		}
	case "FB1":
		if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 1.0
		} else if strings.Contains(scheme, "power") {
			return 0.6
		} else if strings.Contains(scheme, "vert") || strings.Contains(scheme, "spread") {
			return 0.2
		} else if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
			return 0.0
		} else {
			return 0.4
		}
	case "TE1":
		if strings.Contains(scheme, "vert") || strings.Contains(scheme, "west") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "wishbone") {
			return 0.8
		} else if strings.Contains(scheme, "spread") || strings.Contains(scheme, "double") || strings.Contains(scheme, "flex") {
			return 0.6
		} else if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
			return 0.4
		} else {
			return 1.0
		}
	case "TE2":
		if strings.Contains(scheme, "power") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "double") || strings.Contains(scheme, "bone") {
			return 0.4
		} else if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "west") {
			return 0.0
		} else {
			return 0.2
		}
	case "WR1":
		if strings.Contains(scheme, "wing") || strings.Contains(scheme, "wishbone") {
			return 0.8
		} else if strings.Contains(scheme, "flexbone") {
			return 0.6
		} else {
			return 1.0
		}
	case "WR2":
		if strings.Contains(scheme, "power") || strings.Contains(scheme, "pistol") {
			return 0.8
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 0.4
		} else {
			return 1.0
		}
	case "WR3":
		if strings.Contains(scheme, "raid") {
			return 0.8
		} else if strings.Contains(scheme, "vert") || strings.Contains(scheme, "west") || strings.Contains(scheme, "spread") {
			return 0.6
		} else if strings.Contains(scheme, "power") || strings.Contains(scheme, "double") {
			return 0.2
		} else if strings.Contains(scheme, "raid") {
			return 1.0
		} else if strings.Contains(scheme, "bone") || strings.Contains(scheme, "wing-t") {
			return 0.0
		} else {
			return 0.4
		}
	case "WR4":
		if strings.Contains(scheme, "shoot") {
			return 0.6
		} else if strings.Contains(scheme, "vert") || strings.Contains(scheme, "west") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") || strings.Contains(scheme, "double") {
			return 0.2
		} else if strings.Contains(scheme, "raid") {
			return 0.4
		} else {
			return 0.0
		}
	case "WR5":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") {
			return 0.3
		} else {
			return 0.0
		}
	case "LT1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
			return 0.9
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 1.1
		} else {
			return 1.0
		}
	case "LG1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
			return 0.9
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 1.1
		} else {
			return 1.0
		}
	case "C1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
			return 0.9
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 1.1
		} else {
			return 1.0
		}
	case "RG1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
			return 0.9
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 1.1
		} else {
			return 1.0
		}
	case "RT1":
		if strings.Contains(scheme, "shoot") || strings.Contains(scheme, "raid") || strings.Contains(scheme, "pistol") || strings.Contains(scheme, "spread") {
			return 0.9
		} else if strings.Contains(scheme, "wing") || strings.Contains(scheme, "bone") {
			return 1.1
		} else {
			return 1.0
		}
	default:
		return 0.0
	}
}

func GetDefensePositionGradeWeight(position string, scheme string) float64 {
	scheme = strings.ToLower(scheme)
	switch position {
	case "LE1":
		return 1.0
	case "DT1":
		if strings.Contains(scheme, "2") {
			return 0.6
		} else {
			return 1.0
		}
	case "DT2":
		if strings.Contains(scheme, "2") || strings.Contains(scheme, "3") {
			return 0.0
		} else if strings.Contains(scheme, "multiple") {
			return 0.6
		} else {
			return 1.0
		}
	case "RE1":
		return 1.0
	case "LOLB1":
		if strings.Contains(scheme, "4") {
			return 0.0
		} else if strings.Contains(scheme, "old") || strings.Contains(scheme, "2") {
			return 0.8
		} else {
			return 0.6
		}
	case "MLB1":
		return 1.0
	case "MLB2":
		if strings.Contains(scheme, "2") {
			return 1.0
		} else if strings.Contains(scheme, "old") {
			return 0.6
		} else if strings.Contains(scheme, "multiple") {
			return 0.4
		} else {
			return 0.0
		}
	case "ROLB1":
		if strings.Contains(scheme, "speed") {
			return 0.8
		} else if strings.Contains(scheme, "4") {
			return 0.6
		} else {
			return 1.0
		}
	case "CB1":
		return 1.0
	case "CB2":
		return 1.0
	case "CB3":
		if strings.Contains(scheme, "3") {
			return 1.0
		} else if strings.Contains(scheme, "4") {
			return 0.6
		} else if strings.Contains(scheme, "old") {
			return 0.2
		} else {
			return 0.4
		}
	case "CB4":
		if strings.Contains(scheme, "old") || strings.Contains(scheme, "multiple") {
			return 1.0
		} else {
			return 0.2
		}
	case "FS1":
		return 1.0
	case "SS1":
		if strings.Contains(scheme, "old") {
			return 0.4
		} else {
			return 1.0
		}
	case "SS2":
		if strings.Contains(scheme, "4") {
			return 0.6
		} else if strings.Contains(scheme, "3") {
			return 0.2
		} else {
			return 0.0
		}
	default:
		return 0.0
	}
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
