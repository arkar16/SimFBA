package managers

import (
	"sort"
	"strconv"
	"sync"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func ProcessTransferIntention() {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()
	seasonID := strconv.Itoa(ts.CollegeSeasonID)
	allCollegePlayers := GetAllCollegePlayers()
	seasonStatMap := GetCollegePlayerSeasonStatMap(seasonID)
	fullRosterMap := GetFullTeamRosterWithCrootsMap()
	teamProfileMap := GetTeamProfileMap()

	for _, p := range allCollegePlayers {
		if p.IsRedshirting {
			continue
		}
		// Weight will be the initial barrier required for a player to consider transferring.
		// The lower the number gets, the more likely the player will transfer
		transferWeight := 100
		snapMod := 0
		ageMod := 0
		depthChartCompetitionMod := 0
		schemeMod := 0

		// Check Snaps
		seasonStats := seasonStatMap[p.ID]
		totalSnaps := seasonStats.Snaps
		snapsPerGame := totalSnaps / 12
		if snapsPerGame > 50 {
			snapMod = -25
		} else if snapsPerGame > 30 {
			snapMod = 5
		} else if snapsPerGame > 20 {
			snapMod = 15
		} else if snapsPerGame > 10 {
			snapMod = 25
		} else {
			snapMod = 33
		}

		// Check Age
		// The more experienced the player is in the league,
		// the more likely they will transfer.
		ageMod = 2 * p.Year

		// Check Team Position Rank
		teamRoster := fullRosterMap[uint(p.TeamID)]
		filteredRosterByPosition := filterRosterByPosition(teamRoster, p.Position)
		youngerPlayerAhead := false
		idFound := false
		for idx, pl := range filteredRosterByPosition {
			if pl.Age < p.Age && !idFound {
				youngerPlayerAhead = true
			}
			if pl.ID == p.ID {
				idFound = true
				// Check the index of the player.
				// If they're at the top of the list, they're considered to be starting caliber.
				if (p.Position == "QB" ||
					p.Position == "P" ||
					p.Position == "K" ||
					p.Position == "FB" ||
					p.Position == "C") && idx > 1 {
					depthChartCompetitionMod += 33
				}

				if (p.Position == "RB" ||
					p.Position == "TE" ||
					p.Position == "FS" ||
					p.Position == "SS") && idx > 2 {
					depthChartCompetitionMod += 33
				}

				if (p.Position == "WR" ||
					p.Position == "OT" ||
					p.Position == "OG" ||
					p.Position == "DE" ||
					p.Position == "DT" ||
					p.Position == "OLB" ||
					p.Position == "ILB" ||
					p.Position == "CB") && idx > 3 {
					depthChartCompetitionMod += 33
				}
			}
		}

		// If there's a modifier applied and there's a younger player ahead on the roster, double the amount on the modifier
		if depthChartCompetitionMod > 0 && youngerPlayerAhead {
			depthChartCompetitionMod += 33
		}

		// Check for scheme based on Team Recruiting Profile.
		// If it is not a good fit for the player, they will want to transfer
		// Will Need to Lock Scheme Dropdown by halfway through the season or by end of season
		teamIdStr := strconv.Itoa(p.TeamID)
		teamProfile := teamProfileMap[teamIdStr]
		goodFit := IsGoodSchemeFit(teamProfile.OffensiveScheme, teamProfile.DefensiveScheme, p.Archetype, p.Position)
		badFit := IsBadSchemeFit(teamProfile.OffensiveScheme, teamProfile.DefensiveScheme, p.Archetype, p.Position)
		if goodFit {
			schemeMod -= 33
		} else if badFit {
			schemeMod += 33
		}

		transferWeight = transferWeight - snapMod - ageMod - depthChartCompetitionMod - schemeMod
		diceRoll := util.GenerateIntFromRange(1, 100)

		// NOT INTENDING TO TRANSFER
		if diceRoll <= transferWeight {
			continue
		}

		// Is Intending to transfer
		p.DeclareTransferIntention(transferWeight)
		db.Save(&p)
	}
}

func GetCollegePlayerSeasonStatMap(seasonID string) map[uint]structs.CollegePlayerSeasonStats {
	seasonStatMap := make(map[uint]structs.CollegePlayerSeasonStats)

	seasonStats := GetCollegePlayerSeasonStatsBySeason(seasonID)
	for _, stat := range seasonStats {
		seasonStatMap[stat.CollegePlayerID] = stat
	}

	return seasonStatMap
}

func GetFullTeamRosterWithCrootsMap() map[uint][]structs.CollegePlayer {
	m := &sync.Mutex{}
	var wg sync.WaitGroup
	collegeTeams := GetAllCollegeTeams()
	fullMap := make(map[uint][]structs.CollegePlayer)
	wg.Add(len(collegeTeams))
	for _, team := range collegeTeams {
		go func(t structs.CollegeTeam) {
			defer wg.Done()
			id := strconv.Itoa(int(t.ID))
			collegePlayers := GetAllCollegePlayersByTeamId(id)
			croots := GetSignedRecruitsByTeamProfileID(id)
			fullList := collegePlayers
			for _, croot := range croots {
				p := structs.CollegePlayer{}
				p.MapFromRecruit(croot, t)

				fullList = append(fullList, p)
			}

			m.Lock()
			fullMap[t.ID] = fullList
			m.Unlock()
		}(team)
	}

	wg.Wait()
	return fullMap
}

func filterRosterByPosition(roster []structs.CollegePlayer, pos string) []structs.CollegePlayer {
	estimatedSize := len(roster) / 5 // Adjust this based on your data
	filteredList := make([]structs.CollegePlayer, 0, estimatedSize)
	for _, p := range roster {
		if p.Position != pos {
			continue
		}
		filteredList = append(filteredList, p)
	}
	sort.Slice(filteredList, func(i, j int) bool {
		return filteredList[i].Overall > filteredList[j].Overall
	})

	return filteredList
}

func GetTeamProfileMap() map[string]*structs.RecruitingTeamProfile {
	teamRecruitingProfiles := GetRecruitingProfileForRecruitSync()

	teamMap := make(map[string]*structs.RecruitingTeamProfile)
	for i := 0; i < len(teamRecruitingProfiles); i++ {
		teamMap[strconv.Itoa(int(teamRecruitingProfiles[i].ID))] = &teamRecruitingProfiles[i]
	}

	return teamMap
}

func IsGoodSchemeFit(offensiveScheme, defensiveScheme, arch, position string) bool {
	archType := arch + " " + position
	offensiveSchemeList := GetFitsByScheme(offensiveScheme, false)
	defensiveSchemeList := GetFitsByScheme(defensiveScheme, false)
	totalFitList := append(offensiveSchemeList, defensiveSchemeList...)

	return CheckPlayerFits(archType, totalFitList)
}

func IsBadSchemeFit(offensiveScheme, defensiveScheme, arch, position string) bool {
	archType := arch + " " + position
	offensiveSchemeList := GetFitsByScheme(offensiveScheme, true)
	defensiveSchemeList := GetFitsByScheme(defensiveScheme, true)
	totalFitList := append(offensiveSchemeList, defensiveSchemeList...)

	return CheckPlayerFits(archType, totalFitList)
}
