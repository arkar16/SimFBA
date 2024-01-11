package managers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func ProcessTransferIntention(w http.ResponseWriter) {
	// db := dbprovider.GetInstance().GetDB()
	w.Header().Set("Content-Disposition", "attachment;filename=transferStats.csv")
	w.Header().Set("Transfer-Encoding", "chunked")
	// Initialize writer
	writer := csv.NewWriter(w)
	ts := GetTimestamp()
	seasonID := strconv.Itoa(ts.CollegeSeasonID - 1)
	allCollegePlayers := GetAllCollegePlayers()
	seasonStatMap := GetCollegePlayerSeasonStatMap(seasonID)
	fullRosterMap := GetFullTeamRosterWithCrootsMap()
	teamProfileMap := GetTeamProfileMap()
	transferCount := 0
	freshmanCount := 0
	redshirtFreshmanCount := 0
	sophomoreCount := 0
	redshirtSophomoreCount := 0
	juniorCount := 0
	redshirtJuniorCount := 0
	seniorCount := 0
	redshirtSeniorCount := 0
	lowCount := 0
	mediumCount := 0
	highCount := 0
	bigDrop := -25.0
	smallDrop := -10.0
	tinyDrop := -5.0
	tinyGain := 5.0
	smallGain := 10.0
	mediumGain := 15.0
	bigGain := 25.0

	HeaderRow := []string{
		"Team", "First Name", "Last Name", "Stars",
		"Archetype", "Position", "Year", "Age", "Redshirt Status",
		"Overall", "Transfer Weight", "Dice Roll",
	}

	err := writer.Write(HeaderRow)
	if err != nil {
		log.Fatal("Cannot write header row", err)
	}

	for _, p := range allCollegePlayers {
		// Do not include redshirts and all graduating players
		if p.IsRedshirting || p.Year == 5 || (p.Year == 4 && !p.IsRedshirt) {
			continue
		}
		// Weight will be the initial barrier required for a player to consider transferring.
		// The lower the number gets, the more likely the player will transfer
		transferWeight := 0.0

		// Modifiers on reasons why they would transfer
		snapMod := 0.0
		ageMod := 0.0
		starMod := 0.0
		depthChartCompetitionMod := 0.0
		schemeMod := 0.0
		// closeToHomeMod := 0.0

		// Check Snaps
		seasonStats := seasonStatMap[p.ID]
		totalSnaps := seasonStats.Snaps
		snapsPerGame := totalSnaps / 12

		if p.Position == "P" || p.Position == "K" {
			if snapsPerGame > 1 {
				snapMod = bigDrop
			} else {
				snapMod = smallGain
			}
		} else if p.Position == "QB" {

		} else if p.Position == "OG" || p.Position == "OT" || p.Position == "C" {

		} else if p.Position == "SS" || p.Position == "FS" || p.Position == "RB" {
			/// next positions are usually single full time starters, but could also have backups playing ST so more tiers needed
			if snapsPerGame > 35 {
				snapMod = bigDrop
			} else if snapsPerGame > 15 {
				snapMod = tinyGain
			} else if snapsPerGame > 5 {
				snapMod = mediumGain
			} else {
				snapMod = bigGain
			}
		} else if p.Position == "FB" || p.Position == "TE" {
			if snapsPerGame > 30 {
				snapMod = bigDrop
			} else if snapsPerGame > 10 {
				snapMod = tinyDrop
			} else if snapsPerGame > 1 {
				snapMod = mediumGain
			} else {
				snapMod = bigGain
			}
		} else if p.Position == "RB" || p.Position == "DT" || p.Position == "ILB" {
			if snapsPerGame > 35 {
				snapMod = bigDrop
			} else if snapsPerGame > 20 {
				snapMod = smallDrop
			} else if snapsPerGame > 10 {
				snapMod = smallGain
			} else {
				snapMod = bigGain
			}
		} else if p.Position == "CB" || p.Position == "OLB" || p.Position == "WR" || p.Position == "DE" {
			if snapsPerGame > 35 {
				snapMod = bigDrop
			} else if snapsPerGame > 25 {
				snapMod = smallDrop
			} else if snapsPerGame > 15 {
				snapMod = smallGain
			} else if snapsPerGame > 5 {
				snapMod = mediumGain
			} else {
				snapMod = bigGain
			}
		}

		// Check Age
		// The more experienced the player is in the league,
		// the more likely they will transfer.
		/// Have this be a multiplicative factor to odds
		if p.Year == 1 {
			ageMod = .125
		} else if p.Year == 2 {
			ageMod = .33
		} else if p.Year == 3 {
			ageMod = .66
		} else if p.Year == 4 {
			ageMod = 1
		}

		/// Higher star players are more likely to transfer
		if p.Stars == 0 {
			starMod = 1
		} else if p.Stars == 1 {
			starMod = .66
		} else if p.Stars == 2 {
			starMod = .75
		} else if p.Stars == 3 {
			starMod = 1
		} else if p.Stars == 4 {
			starMod = 1.2
		} else if p.Stars == 5 {
			starMod = 1.5
		}

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
		if depthChartCompetitionMod > 0 {
			if youngerPlayerAhead {
				depthChartCompetitionMod += 33
			} else {
				depthChartCompetitionMod = .5 * depthChartCompetitionMod
			}

		}

		// Check for scheme based on Team Recruiting Profile.
		// If it is not a good fit for the player, they will want to transfer
		// Will Need to Lock Scheme Dropdown by halfway through the season or by end of season
		teamIdStr := strconv.Itoa(p.TeamID)
		teamProfile := teamProfileMap[teamIdStr]
		schemeMod = getSchemeMod(teamProfile, p, smallDrop, smallGain)

		/// Not playing = 25, low depth chart = 16 or 33, scheme = 10, if you're all 3, that's a ~60% chance of transferring pre- modifiers
		transferWeight = starMod * ageMod * (snapMod + depthChartCompetitionMod + schemeMod)
		diceRoll := util.GenerateIntFromRange(1, 100)

		// NOT INTENDING TO TRANSFER
		transferInt := int(transferWeight)
		if diceRoll > transferInt {
			continue
		}

		if p.Year == 1 {
			fmt.Println("Dice Roll: ", diceRoll)
		}

		// Is Intending to transfer
		p.DeclareTransferIntention(int(transferWeight))
		transferCount++
		if p.Year == 1 && !p.IsRedshirt {
			freshmanCount++
		} else if p.Year == 2 && p.IsRedshirt {
			redshirtFreshmanCount++
		} else if p.Year == 2 && !p.IsRedshirt {
			sophomoreCount++
		} else if p.Year == 3 && p.IsRedshirt {
			redshirtSophomoreCount++
		} else if p.Year == 3 && !p.IsRedshirt {
			juniorCount++
		} else if p.Year == 4 && p.IsRedshirt {
			redshirtJuniorCount++
		} else if p.Year == 4 && !p.IsRedshirt {
			seniorCount++
		} else if p.Year == 5 && p.IsRedshirt {
			redshirtSeniorCount++
		}

		if transferWeight < 30 {
			lowCount++
		} else if transferWeight < 70 {
			mediumCount++
		} else {
			highCount++
		}

		fmt.Println(strconv.Itoa(p.Year)+" YEAR "+p.TeamAbbr+" "+p.Position+" "+p.FirstName+" "+p.LastName+" HAS ANNOUNCED THEIR INTENTION TO TRANSFER | Weight: ", int(transferWeight))
		// db.Save(&p)
		csvModel := models.MapPlayerToCSVModel(p)
		playerRow := []string{
			p.TeamAbbr, csvModel.FirstName, csvModel.LastName, strconv.Itoa(p.Stars),
			csvModel.Archetype, csvModel.Position,
			csvModel.Year, strconv.Itoa(p.Age), csvModel.RedshirtStatus,
			csvModel.OverallGrade, strconv.Itoa(transferInt), strconv.Itoa(diceRoll),
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
	fmt.Println("Total number of players entering the transfer portal: ", transferCount)
	fmt.Println("Total number of freshmen entering the transfer portal: ", freshmanCount)
	fmt.Println("Total number of redshirt freshmen entering the transfer portal: ", redshirtFreshmanCount)
	fmt.Println("Total number of sophomores entering the transfer portal: ", sophomoreCount)
	fmt.Println("Total number of redshirt sophomores entering the transfer portal: ", redshirtSophomoreCount)
	fmt.Println("Total number of juniors entering the transfer portal: ", juniorCount)
	fmt.Println("Total number of redshirt juniors entering the transfer portal: ", redshirtJuniorCount)
	fmt.Println("Total number of seniors entering the transfer portal: ", seniorCount)
	fmt.Println("Total number of redshirt seniors entering the transfer portal: ", redshirtSeniorCount)
	fmt.Println("Total number of players with low likeliness to enter transfer portal: ", lowCount)
	fmt.Println("Total number of players with medium likeliness to enter transfer portal: ", mediumCount)
	fmt.Println("Total number of players with high likeliness to enter transfer portal: ", highCount)
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
	semaphore := make(chan struct{}, 10)
	for _, team := range collegeTeams {
		semaphore <- struct{}{}
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
			<-semaphore
		}(team)
	}

	wg.Wait()
	close(semaphore)
	return fullMap
}

func filterRosterByPosition(roster []structs.CollegePlayer, pos string) []structs.CollegePlayer {
	estimatedSize := len(roster) / 5 // Adjust this based on your data
	filteredList := make([]structs.CollegePlayer, 0, estimatedSize)
	for _, p := range roster {
		if p.Position != pos || (p.Year == 5 || (p.Year == 4 && p.IsRedshirt)) {
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

func getSchemeMod(tp *structs.RecruitingTeamProfile, p structs.CollegePlayer, drop, gain float64) float64 {
	schemeMod := 0.0

	goodFit := IsGoodSchemeFit(tp.OffensiveScheme, tp.DefensiveScheme, p.Archetype, p.Position)
	badFit := IsBadSchemeFit(tp.OffensiveScheme, tp.DefensiveScheme, p.Archetype, p.Position)
	if goodFit {
		schemeMod += drop
	} else if badFit {
		schemeMod += gain
	}

	return schemeMod
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
