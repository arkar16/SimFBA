package managers

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func GetRecruitingModifiers() structs.AdminRecruitModifier {
	db := dbprovider.GetInstance().GetDB()

	var recruitModifiers structs.AdminRecruitModifier

	db.First(&recruitModifiers)

	return recruitModifiers
}

func SyncRecruiting(timestamp structs.Timestamp) {
	db := dbprovider.GetInstance().GetDB()
	rand.Seed(time.Now().UnixNano())
	//GetCurrentWeek

	if timestamp.RecruitingSynced {
		log.Fatalln("Recruiting already ran for this week. Please wait until next week to sync recruiting again.")
	}

	recruitModifiers := GetRecruitingModifiers()

	var recruitProfiles []structs.RecruitPlayerProfile

	// Get every recruit
	recruits := GetAllUnsignedRecruits()

	// Iterate through every recruit
	for _, recruit := range recruits {
		recruitProfiles = GetRecruitPlayerProfilesByRecruitId(strconv.Itoa(int(recruit.ID)))

		if len(recruitProfiles) == 0 {
			continue
		}

		var recruitProfilesWithScholarship []structs.RecruitPlayerProfile

		eligibleTeams := 0

		totalPointsOnRecruit := 0

		var eligiblePointThreshold float64 = 0

		var signThreshold float64

		for i := 0; i < len(recruitProfiles); i++ {

			if recruitProfiles[i].CurrentWeeksPoints == 0 {
				continue
			}

			rpa := structs.RecruitPointAllocation{
				RecruitID:        recruitProfiles[i].RecruitID,
				TeamProfileID:    recruitProfiles[i].ProfileID,
				RecruitProfileID: int(recruitProfiles[i].ID),
				WeekID:           timestamp.CollegeWeekID,
			}

			var curr float64 = 0

			res := recruitProfiles[i].RecruitingEfficiencyScore
			if recruitProfiles[i].AffinityOneEligible {
				res += .1
				rpa.ApplyAffinityOne()
			}
			if recruitProfiles[i].AffinityTwoEligible {
				res += .1
				rpa.ApplyAffinityTwo()
			}

			curr = float64(recruitProfiles[i].CurrentWeeksPoints) * res

			if recruitProfiles[i].CurrentWeeksPoints < 0 || recruitProfiles[i].CurrentWeeksPoints > 20 {
				curr = 0
				rpa.ApplyCaughtCheating()
			}

			if curr < 1 {
				curr = 1
			}

			rpa.UpdatePointsSpent(recruitProfiles[i].CurrentWeeksPoints, curr)
			recruitProfiles[i].AddCurrentWeekPointsToTotal(int(curr))
			totalPointsOnRecruit += recruitProfiles[i].TotalPoints
			if recruitProfiles[i].Scholarship {
				recruitProfilesWithScholarship = append(recruitProfilesWithScholarship, recruitProfiles[i])
			}

			if eligiblePointThreshold == 0 {
				eligiblePointThreshold = curr / 2
			}

			if recruitProfiles[i].Scholarship && recruitProfiles[i].TotalPoints > int(eligiblePointThreshold) {
				eligibleTeams += 1
			}

			// Add RPA to point allocations list
			err := db.Save(&rpa).Error
			if err != nil {
				fmt.Println(err.Error())
				log.Fatalf("Could not save Point Allocation")
			}
		}

		// Re-Sort profiles
		sort.Sort(structs.ByPoints(recruitProfilesWithScholarship))

		// Change?
		// Assign point totals
		// If there are any modifiers
		// Evaluate
		signThreshold = float64(recruitModifiers.ModifierOne-timestamp.CollegeWeek) * ((float64(eligibleTeams) / recruitModifiers.ModifierTwo) * math.Log10(float64(recruitModifiers.WeeksOfRecruiting-timestamp.CollegeWeek)))

		// Change logic to withold teams without available scholarships
		if float64(totalPointsOnRecruit) > signThreshold && eligibleTeams > 0 {

			percentageOdds := rand.Intn(totalPointsOnRecruit) + 1
			currentProbability := 0
			winningTeamID := 0

			for i := 0; i < len(recruitProfilesWithScholarship); i++ {
				// If a team has no available scholarships or if a team has 25 commitments, continue
				currentProbability += recruitProfilesWithScholarship[i].TotalPoints
				if currentProbability > percentageOdds {
					// WINNING TEAM
					winningTeamID = recruitProfilesWithScholarship[i].ProfileID
					break
				}
			}

			if winningTeamID > 0 {
				recruitTeamProfile := GetOnlyRecruitingProfileByTeamID(strconv.Itoa(winningTeamID))
				teamAbbreviation := recruitTeamProfile.TeamAbbreviation
				recruit.AssignCollege(teamAbbreviation)

				for i := 0; i < len(recruitProfiles); i++ {
					if recruitProfiles[i].ProfileID == winningTeamID {
						recruitProfiles[i].SignPlayer()
					} else {
						recruitProfiles[i].LockPlayer()
					}
				}
			}

			recruit.UpdateTeamID(winningTeamID)
		}

		// Save Player Files towards Recruit
		for _, rp := range recruitProfiles {
			// Save Team Profile
			err := db.Save(&rp).Error
			if err != nil {
				fmt.Println(err.Error())
				log.Fatalf("Could not sync recruiting profile.")
			}

			fmt.Println("Save recruit profile from " + rp.TeamAbbreviation + " towards " + recruit.FirstName + " " + recruit.LastName)
		}

		// Save Recruit
		err := db.Save(&recruit).Error
		if err != nil {
			fmt.Println(err.Error())
			log.Fatalf("Could not sync recruit")
		}
		fmt.Println("Save Recruit " + recruit.FirstName + " " + recruit.LastName)
	}

	// Update rank system for all teams
	teamRecruitingProfiles := GetRecruitingProfileForRecruitSync()

	var totalESPNScore float64 = 0
	var total247Score float64 = 0
	var totalRivalsScore float64 = 0

	for i := 0; i < len(teamRecruitingProfiles); i++ {
		signedRecruits := GetSignedRecruitsByTeamProfileID(strconv.Itoa(teamRecruitingProfiles[i].TeamID))

		teamRecruitingProfiles[i].UpdateTotalSignedRecruits(len(signedRecruits))

		team247Rank := Get247TeamRanking(teamRecruitingProfiles[i], signedRecruits)
		teamESPNRank := GetESPNTeamRanking(teamRecruitingProfiles[i], signedRecruits)
		teamRivalsRank := GetRivalsTeamRanking(teamRecruitingProfiles[i], signedRecruits)

		teamRecruitingProfiles[i].Assign247Rank(team247Rank)
		total247Score += team247Rank
		teamRecruitingProfiles[i].AssignESPNRank(teamESPNRank)
		totalESPNScore += teamESPNRank
		teamRecruitingProfiles[i].AssignRivalsRank(teamRivalsRank)
		totalRivalsScore += teamRivalsRank
	}

	averageESPNScore := totalESPNScore / 130
	average247score := total247Score / 130
	averageRivalScore := totalRivalsScore / 130

	for _, rp := range teamRecruitingProfiles {
		var avg float64 = 0
		if averageESPNScore > 0 && average247score > 0 && averageRivalScore > 0 {
			distributionESPN := rp.ESPNScore / averageESPNScore
			distribution247 := rp.Rank247Score / average247score
			distributionRivals := rp.RivalsScore / averageRivalScore

			avg = (distributionESPN + distribution247 + distributionRivals) / 3

			rp.AssignCompositeRank(avg)
		}
		rp.ResetSpentPoints()

		// Save TEAM Recruiting Profile
		err := db.Save(&rp).Error
		if err != nil {
			fmt.Println(err.Error())
			log.Fatalf("Could not save timestamp")
		}
		fmt.Println("Saved Rank Scores for Team " + rp.TeamAbbreviation)
	}
}

func SyncRecruitingEfficiency(timestamp structs.Timestamp) {
	db := dbprovider.GetInstance().GetDB()

	// Get All Team Recruiting Profiles
	teams := GetAllCollegeTeamsWithRecruitingProfileAndCoach()

	// Iterate through all profiles

	// var teamProfilesToSave []structs.RecruitingTeamProfile

	for _, team := range teams {
		// Get all games by team within a season

		teamProfile := team.RecruitingProfile

		coach := team.CollegeCoach

		// What about previous season?
		currentSeasonGames := GetCollegeGamesByTeamIdAndSeasonId(
			strconv.Itoa(teamProfile.TeamID), strconv.Itoa(timestamp.CollegeSeasonID))

		currentSeasonWins, currentSeasonLosses := util.GetWinsAndLossesForCollegeGames(currentSeasonGames, teamProfile.TeamID, false)
		currentConferenceWins, currentConferenceLosses := util.GetWinsAndLossesForCollegeGames(currentSeasonGames, teamProfile.TeamID, true)

		previousSeasonGames := GetCollegeGamesByTeamIdAndSeasonId(
			strconv.Itoa(teamProfile.TeamID), strconv.Itoa(timestamp.CollegeSeasonID-1))

		previousSeasonWins, previousSeasonLosses := util.GetWinsAndLossesForCollegeGames(previousSeasonGames, teamProfile.TeamID, false)
		previousConferenceWins, previousConferenceLosses := util.GetWinsAndLossesForCollegeGames(previousSeasonGames, teamProfile.TeamID, true)
		// Do calculation for current season losses

		// Current Season Win Percentage
		var cswp float64 = 1
		var csweight float64 = .125
		var ccwp float64 = 1
		var ccweight float64 = .125
		var pswp float64 = 1
		var psweight float64 = .05
		var pcwp float64 = 1
		var pcweight float64 = .05
		var coachwp float64 = 1
		var coachweight float64 = .1

		var postSeasonVal float64
		var conferenceChampionshipVal float64
		var postseasonweight float64 = 0.025

		if timestamp.CollegeWeek < 15 {
			postSeasonVal = util.GetPostSeasonWeight(previousSeasonGames, teamProfile.TeamID)
			conferenceChampionshipVal = util.GetConferenceChampionshipWeight(previousSeasonGames, teamProfile.TeamID)
		} else {
			postSeasonVal = util.GetPostSeasonWeight(currentSeasonGames, teamProfile.TeamID)
			conferenceChampionshipVal = util.GetConferenceChampionshipWeight(currentSeasonGames, teamProfile.TeamID)
		}

		if currentSeasonWins+currentSeasonLosses > 0 {
			cswp = float64(currentSeasonWins) / float64(currentSeasonWins+currentSeasonLosses)
			if team.ConferenceID != 13 {
				ccwp = float64(currentConferenceWins) / float64(currentConferenceWins+currentConferenceLosses)

			}
		}

		// Previous Season Win Percentage
		if previousSeasonWins+previousSeasonLosses > 0 {
			pswp = float64(previousSeasonWins) / float64(previousSeasonWins+previousSeasonLosses)
			if team.ConferenceID != 13 {
				pcwp = float64(previousConferenceWins) / float64(previousConferenceWins+previousConferenceLosses)
			}
		}

		if coach.OverallWins+coach.OverallLosses > 0 {
			coachwp = float64(coach.OverallWins) / float64(coach.OverallWins+coach.OverallLosses)
		}

		res := teamProfile.BaseEfficiencyScore

		cswsum := cswp * csweight                             // Current Season Win Percentage
		ccwsum := ccwp * ccweight                             // Current Conference Wins
		pswsum := pswp * psweight                             // Prev Season Wins
		pcwsum := pcwp * pcweight                             // Prev Conference Wins
		coachsum := coachwp * coachweight                     // Coach Overall Wins
		postseasonsum := postSeasonVal * postseasonweight     // PostSeason Check
		ccsum := conferenceChampionshipVal * postseasonweight // ConferenceChampionship Check

		totalSum := cswsum + ccwsum + pswsum + pcwsum + coachsum + postseasonsum + ccsum
		// totalWeight := csweight + ccweight + psweight + pcweight + coachweight + (postseasonweight * 2)

		// RES Calculation
		// Base of .8

		teamProfile.AssignRES(res + (totalSum * 0.4))

		err := db.Save(&teamProfile).Error
		if err != nil {
			fmt.Println(err.Error())
			log.Fatalf("Could not sync all team profiles.")
		}
		fmt.Println("Saved RES for Team: " + team.TeamAbbr)

		// teamProfilesToSave = append(teamProfilesToSave, teamProfile)
	}
	// Save the Recruiting Profiles
	// err = db.Save(&recruitProfilesToSave).Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	log.Fatalf("Could not sync res to all recruits")
	// }
}

func SyncAllMissingEfficiencies() {
	db := dbprovider.GetInstance().GetDB()
	recruitingProfiles := GetRecruitingProfileForRecruitSync()

	for _, rp := range recruitingProfiles {
		playerProfiles := GetOnlyRecruitProfilesByTeamProfileID(strconv.Itoa(int(rp.ID)))

		for _, pp := range playerProfiles {
			if pp.RecruitingEfficiencyScore != 0 {
				continue
			}
			pp.AssignRES(rp.RecruitingEfficiencyScore)

			err := db.Save(&pp).Error
			if err != nil {
				log.Panicln("COULD NOT SAVE ALL PROFILES")
			}
			fmt.Println("Saved profile from " + rp.Team + " towards " + strconv.Itoa(int(pp.ID)))
		}
	}
}
