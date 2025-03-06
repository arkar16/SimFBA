package managers

import (
	"sync"

	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
)

type BootstrapData struct {
	CollegeTeam          structs.CollegeTeam
	AllCollegeTeams      []structs.CollegeTeam
	CollegeStandings     []structs.CollegeStandings
	CollegeRosterMap     map[uint][]structs.CollegePlayer
	Recruits             []structs.Croot
	TeamProfileMap       map[string]*structs.RecruitingTeamProfile
	PortalPlayers        []structs.CollegePlayer
	CollegeInjuryReport  []structs.CollegePlayer
	CollegeNews          []structs.NewsLog
	CollegeNotifications []structs.Notification
	AllCollegeGames      []structs.CollegeGame
	CollegeGameplan      structs.CollegeGameplan
	CollegeDepthChart    structs.CollegeTeamDepthChart
	CollegeDepthChartMap map[uint]structs.CollegeTeamDepthChart

	// Player Profiles by Team?
	// Portal profiles?
	ProTeam          structs.NFLTeam
	AllProTeams      []structs.NFLTeam
	ProStandings     []structs.NFLStandings
	ProRosterMap     map[uint][]structs.NFLPlayer
	CapsheetMap      map[uint]structs.NFLCapsheet
	FreeAgency       models.FreeAgencyResponse
	ProInjuryReport  []structs.NFLPlayer
	ProNews          []structs.NewsLog
	ProNotifications []structs.Notification
	AllProGames      []structs.NFLGame
	NFLGameplan      structs.NFLGameplan
	NFLDepthChart    structs.NFLDepthChart
	NFLDepthChartMap map[uint]structs.NFLDepthChart
}

func GetFirstBootstrapData(collegeID, proID string) BootstrapData {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// College Data
	var (
		collegeTeam           structs.CollegeTeam
		allCollegeTeams       []structs.CollegeTeam
		collegeStandings      []structs.CollegeStandings
		collegePlayerMap      map[uint][]structs.CollegePlayer
		teamProfileMap        map[string]*structs.RecruitingTeamProfile
		portalPlayers         []structs.CollegePlayer
		injuredCollegePlayers []structs.CollegePlayer
		collegeNews           []structs.NewsLog
		collegeNotifications  []structs.Notification
		collegeGames          []structs.CollegeGame
		recruits              []structs.Croot
		collegeGameplan       structs.CollegeGameplan
		collegeDepthChart     structs.CollegeTeamDepthChart
		collegeDepthChartMap  map[uint]structs.CollegeTeamDepthChart
	)

	// Professional Data
	var (
		proTeam           structs.NFLTeam
		allProTeams       []structs.NFLTeam
		proStandings      []structs.NFLStandings
		proRosterMap      map[uint][]structs.NFLPlayer
		capsheetMap       map[uint]structs.NFLCapsheet
		freeAgency        models.FreeAgencyResponse
		injuredProPlayers []structs.NFLPlayer
		proNews           []structs.NewsLog
		proNotifications  []structs.Notification
		proGames          []structs.NFLGame
		proGameplan       structs.NFLGameplan
		proDepthChart     structs.NFLDepthChart
		proDepthChartMap  map[uint]structs.NFLDepthChart
	)

	// Start concurrent queries
	wg.Add(2)
	go func() {
		defer wg.Done()
		allCollegeTeams = GetAllCollegeTeams()
	}()
	go func() {
		defer wg.Done()
		allProTeams = GetAllNFLTeams()
	}()

	wg.Wait()

	if len(collegeID) > 0 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			collegeTeam = GetTeamByTeamID(collegeID)
		}()

		go func() {
			defer wg.Done()
			collegePlayers := GetAllCollegePlayers()
			mu.Lock()
			collegePlayerMap = MakeCollegePlayerMapByTeamID(collegePlayers, true)
			injuredCollegePlayers = MakeCollegeInjuryList(collegePlayers)
			portalPlayers = MakeCollegePortalList(collegePlayers)
			mu.Unlock()
		}()
		go func() {
			defer wg.Done()
			collegeGames = GetCollegeGamesBySeasonID("")
		}()
		wg.Wait()
		wg.Add(3)
		go func() {
			defer wg.Done()
			collegeNews = GetAllNewsLogs()
		}()
		go func() {
			defer wg.Done()
			collegeNotifications = GetNotificationByTeamIDAndLeague("CFB", collegeID)
		}()
		go func() {
			defer wg.Done()
			teamProfileMap = GetTeamProfileMap()
		}()

		wg.Wait()
		wg.Add(3)

		go func() {
			defer wg.Done()
			collegeStandings = GetAllCollegeStandingsBySeasonID("")
		}()
		go func() {
			defer wg.Done()
			collegeGameplan = GetGameplanByTeamID(collegeID)
		}()
		go func() {
			defer wg.Done()
			collegeDepthChart = GetDepthchartByTeamID(collegeID)

		}()

		wg.Wait()
	}
	if len(proID) > 0 {
		wg.Add(3)
		go func() {
			defer wg.Done()
			proPlayers := GetAllNFLPlayers()
			mu.Lock()
			proRosterMap = MakeNFLPlayerMapByTeamID(proPlayers, true)
			injuredProPlayers = MakeProInjuryList(proPlayers)
			mu.Unlock()
		}()
		go func() {
			defer wg.Done()
			proGames = GetNFLGamesBySeasonID("")
		}()
		go func() {
			defer wg.Done()
			proNews = GetAllNFLNewsLogs()
		}()

		wg.Wait()
		wg.Add(4)

		go func() {
			defer wg.Done()
			proNotifications = GetNotificationByTeamIDAndLeague("NFL", proID)
		}()
		go func() {
			defer wg.Done()
			capsheetMap = getCapsheetMap()
		}()
		go func() {
			defer wg.Done()
			proStandings = GetAllNFLStandingsBySeasonID("")
		}()
		go func() {
			defer wg.Done()
			proGameplan = GetNFLGameplanByTeamID(proID)
		}()
		wg.Wait()
	}
	return BootstrapData{
		CollegeTeam:          collegeTeam,
		AllCollegeTeams:      allCollegeTeams,
		CollegeStandings:     collegeStandings,
		CollegeRosterMap:     collegePlayerMap,
		CollegeInjuryReport:  injuredCollegePlayers,
		CollegeNews:          collegeNews,
		CollegeNotifications: collegeNotifications,
		CollegeGameplan:      collegeGameplan,
		CollegeDepthChart:    collegeDepthChart,
		CollegeDepthChartMap: collegeDepthChartMap,
		AllCollegeGames:      collegeGames,
		Recruits:             recruits,
		TeamProfileMap:       teamProfileMap,
		PortalPlayers:        portalPlayers,
		//
		ProTeam:          proTeam,
		AllProTeams:      allProTeams,
		ProStandings:     proStandings,
		ProRosterMap:     proRosterMap,
		CapsheetMap:      capsheetMap,
		FreeAgency:       freeAgency,
		ProInjuryReport:  injuredProPlayers,
		ProNews:          proNews,
		ProNotifications: proNotifications,
		AllProGames:      proGames,
		NFLGameplan:      proGameplan,
		NFLDepthChart:    proDepthChart,
		NFLDepthChartMap: proDepthChartMap,
	}
}

func GetSecondBootstrapData(collegeID, proID string) BootstrapData {
	var wg sync.WaitGroup

	// College Data
	var (
		collegeTeam           structs.CollegeTeam
		allCollegeTeams       []structs.CollegeTeam
		collegeStandings      []structs.CollegeStandings
		collegePlayerMap      map[uint][]structs.CollegePlayer
		teamProfileMap        map[string]*structs.RecruitingTeamProfile
		portalPlayers         []structs.CollegePlayer
		injuredCollegePlayers []structs.CollegePlayer
		collegeNews           []structs.NewsLog
		collegeNotifications  []structs.Notification
		collegeGames          []structs.CollegeGame
		recruits              []structs.Croot
		collegeGameplan       structs.CollegeGameplan
		collegeDepthChart     structs.CollegeTeamDepthChart
		collegeDepthChartMap  map[uint]structs.CollegeTeamDepthChart
	)

	// Professional Data
	var (
		proTeam           structs.NFLTeam
		allProTeams       []structs.NFLTeam
		proStandings      []structs.NFLStandings
		proRosterMap      map[uint][]structs.NFLPlayer
		capsheetMap       map[uint]structs.NFLCapsheet
		freeAgency        models.FreeAgencyResponse
		injuredProPlayers []structs.NFLPlayer
		proNews           []structs.NewsLog
		proNotifications  []structs.Notification
		proGames          []structs.NFLGame
		proGameplan       structs.NFLGameplan
		proDepthChart     structs.NFLDepthChart
		proDepthChartMap  map[uint]structs.NFLDepthChart
	)

	freeAgencyCh := make(chan models.FreeAgencyResponse, 1)

	// Start concurrent queries

	if len(collegeID) > 0 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			recruits = GetAllRecruits()
		}()

		go func() {
			defer wg.Done()
			collegeDCs := GetAllCollegeDepthcharts()
			collegeDepthChartMap = MakeCollegeDepthChartMap(collegeDCs)
		}()

		wg.Wait()
	}
	if len(proID) > 0 {
		wg.Add(2)

		go func() {
			defer wg.Done()
			dcs := GetAllNFLDepthcharts()
			proDepthChartMap = MakeNFLDepthChartMap(dcs)
		}()
		go GetAllAvailableNFLPlayersViaChan(proID, freeAgencyCh)
		freeAgency = <-freeAgencyCh
		wg.Wait()
	}
	return BootstrapData{
		CollegeTeam:          collegeTeam,
		AllCollegeTeams:      allCollegeTeams,
		CollegeStandings:     collegeStandings,
		CollegeRosterMap:     collegePlayerMap,
		CollegeInjuryReport:  injuredCollegePlayers,
		CollegeNews:          collegeNews,
		CollegeNotifications: collegeNotifications,
		CollegeGameplan:      collegeGameplan,
		CollegeDepthChart:    collegeDepthChart,
		CollegeDepthChartMap: collegeDepthChartMap,
		AllCollegeGames:      collegeGames,
		Recruits:             recruits,
		TeamProfileMap:       teamProfileMap,
		PortalPlayers:        portalPlayers,
		//
		ProTeam:          proTeam,
		AllProTeams:      allProTeams,
		ProStandings:     proStandings,
		ProRosterMap:     proRosterMap,
		CapsheetMap:      capsheetMap,
		FreeAgency:       freeAgency,
		ProInjuryReport:  injuredProPlayers,
		ProNews:          proNews,
		ProNotifications: proNotifications,
		AllProGames:      proGames,
		NFLGameplan:      proGameplan,
		NFLDepthChart:    proDepthChart,
		NFLDepthChartMap: proDepthChartMap,
	}
}
