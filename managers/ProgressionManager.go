package managers

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func ProgressionMain() {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()
	SeasonID := strconv.Itoa(ts.CollegeSeasonID)
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	// Get All Teams

	collegeTeams := GetAllCollegeTeams()
	// Loop
	for _, team := range collegeTeams {
		var graduatingPlayers []structs.NFLDraftee
		teamID := strconv.Itoa(int(team.ID))
		roster := GetAllCollegePlayersWithStatsByTeamID(teamID, SeasonID)
		croots := GetSignedRecruitsByTeamProfileID(teamID)

		for _, player := range roster {
			player = ProgressPlayer(player, SeasonID)
			if player.IsRedshirting {
				player.SetRedshirtStatus()
			}

			if (player.IsRedshirt && player.Year > 5) ||
				(!player.IsRedshirt && player.Year > 4) {
				player.GraduatePlayer()
				draftee := structs.NFLDraftee{}
				draftee.Map(player)
				hcp := (structs.HistoricCollegePlayer)(player)

				err := db.Save(&hcp).Error
				if err != nil {
					log.Panicln("Could not save historic player record!")
				}
				graduatingPlayers = append(graduatingPlayers, draftee)
				// CollegePlayer record will be deleted, but record will be mapped to a GraduatedCollegePlayer struct, and then saved in that table, along side with NFL Draftees table
				// GraduatedCollegePlayer will be a copy of the collegeplayers table, but only for historical players

				err = db.Delete(&player).Error
				if err != nil {
					log.Panicln("Could not delete old college player record.")
				}
				continue
			}
			err := db.Save(&player).Error
			if err != nil {
				log.Panicln("Could not save player record")
			}

		}

		for _, croot := range croots {
			// Convert to College Player Record
			cp := structs.CollegePlayer{}
			cp.MapFromRecruit(croot, team)

			// Save College Player Record
			err := db.Save(&cp).Error
			if err != nil {
				log.Panicln("Could not save new college player record")
			}

			// Delete Recruit Record
		}

		// Graduating players
		err := db.CreateInBatches(&graduatingPlayers, len(graduatingPlayers)).Error
		if err != nil {
			log.Panicln("Could not save graduating players")
		}
	}

	unsignedCroots := GetAllUnsignedRecruits()
	for _, croot := range unsignedCroots {
		croot = ProgressCroot(croot, SeasonID)

		err := db.Save(&croot).Error
		if err != nil {
			log.Panic("Recruit could not be saved!")
		}

	}
	// Get All Players from Team
	// Loop each player
	// Check for Stats
	// Progression value
	// Progress Players
	// If no stats, normal progress by 0 snaps / games played by school
	// If stats, progress player by total snaps / games played

	// Overall algorithm
	// Potential algorithm

	// Move players to NFL Draftees List

	// Import Recruits onto team

	// Get All Recruits with no team
	// Progress
	// Place onto JUCO teams?
	// Juco Table?
}

func ProgressPlayer(cp structs.CollegePlayer, SeasonID string) structs.CollegePlayer {
	stats := cp.Stats
	totalSnaps := 0

	for _, stat := range stats {
		totalSnaps += stat.Snaps
	}

	var SnapsPerGame int = 0
	if len(stats) > 0 {
		SnapsPerGame = totalSnaps / len(stats)
	}

	Agility := 0
	ThrowPower := 0
	ThrowAccuracy := 0
	Speed := 0
	FootballIQ := 0
	Strength := 0
	RunBlock := 0
	PassBlock := 0
	RunDefense := 0
	PassRush := 0
	Carrying := 0
	Tackle := 0
	RouteRunning := 0
	Catching := 0
	PuntPower := 0
	PuntAccuracy := 0
	KickPower := 0
	KickAccuracy := 0
	ManCoverage := 0
	ZoneCoverage := 0

	if cp.Position == "QB" {
		// Primary Progressions
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		ThrowPower = PrimaryProgression(cp.Progression, cp.ThrowPower, cp.Position, cp.Archetype, SnapsPerGame, "Throw Power", cp.IsRedshirting)
		ThrowAccuracy = PrimaryProgression(cp.Progression, cp.ThrowAccuracy, cp.Position, cp.Archetype, SnapsPerGame, "Throw Accuracy", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)

		// Secondary Progressions
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
		Catching = SecondaryProgression(cp.Progression, cp.Catching)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
	} else if cp.Position == "RB" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		Carrying = PrimaryProgression(cp.Progression, cp.Carrying, cp.Position, cp.Archetype, SnapsPerGame, "Carrying", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Catching = PrimaryProgression(cp.Progression, cp.Catching, cp.Position, cp.Archetype, SnapsPerGame, "Catching", cp.IsRedshirting)
		PassBlock = PrimaryProgression(cp.Progression, cp.PassBlock, cp.Position, cp.Archetype, SnapsPerGame, "Pass Blocking", cp.IsRedshirting)
		// Secondary
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
	} else if cp.Position == "FB" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		Carrying = PrimaryProgression(cp.Progression, cp.Carrying, cp.Position, cp.Archetype, SnapsPerGame, "Carrying", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Catching = PrimaryProgression(cp.Progression, cp.Catching, cp.Position, cp.Archetype, SnapsPerGame, "Catching", cp.IsRedshirting)
		PassBlock = PrimaryProgression(cp.Progression, cp.PassBlock, cp.Position, cp.Archetype, SnapsPerGame, "Pass Blocking", cp.IsRedshirting)
		RunBlock = PrimaryProgression(cp.Progression, cp.RunBlock, cp.Position, cp.Archetype, SnapsPerGame, "Run Blocking", cp.IsRedshirting)

		// Secondary
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
	} else if cp.Position == "TE" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		Carrying = PrimaryProgression(cp.Progression, cp.Carrying, cp.Position, cp.Archetype, SnapsPerGame, "Carrying", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Catching = PrimaryProgression(cp.Progression, cp.Catching, cp.Position, cp.Archetype, SnapsPerGame, "Catching", cp.IsRedshirting)
		RouteRunning = PrimaryProgression(cp.Progression, cp.RouteRunning, cp.Position, cp.Archetype, SnapsPerGame, "Route Running", cp.IsRedshirting)
		PassBlock = PrimaryProgression(cp.Progression, cp.PassBlock, cp.Position, cp.Archetype, SnapsPerGame, "Pass Blocking", cp.IsRedshirting)
		RunBlock = PrimaryProgression(cp.Progression, cp.RunBlock, cp.Position, cp.Archetype, SnapsPerGame, "Run Blocking", cp.IsRedshirting)

		// Secondary
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
	} else if cp.Position == "WR" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		Carrying = PrimaryProgression(cp.Progression, cp.Carrying, cp.Position, cp.Archetype, SnapsPerGame, "Carrying", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Catching = PrimaryProgression(cp.Progression, cp.Catching, cp.Position, cp.Archetype, SnapsPerGame, "Catching", cp.IsRedshirting)
		RouteRunning = PrimaryProgression(cp.Progression, cp.RouteRunning, cp.Position, cp.Archetype, SnapsPerGame, "Route Running", cp.IsRedshirting)

		// Secondary
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
	} else if cp.Position == "OT" || cp.Position == "OG" || cp.Position == "C" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		PassBlock = PrimaryProgression(cp.Progression, cp.PassBlock, cp.Position, cp.Archetype, SnapsPerGame, "Pass Blocking", cp.IsRedshirting)
		RunBlock = PrimaryProgression(cp.Progression, cp.RunBlock, cp.Position, cp.Archetype, SnapsPerGame, "Run Blocking", cp.IsRedshirting)

		// Secondary
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		Speed = SecondaryProgression(cp.Progression, cp.Speed)
		Catching = SecondaryProgression(cp.Progression, cp.Catching)
	} else if cp.Position == "DT" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		PassRush = PrimaryProgression(cp.Progression, cp.PassRush, cp.Position, cp.Archetype, SnapsPerGame, "Pass Rush", cp.IsRedshirting)
		RunDefense = PrimaryProgression(cp.Progression, cp.RunDefense, cp.Position, cp.Archetype, SnapsPerGame, "Run Defense", cp.IsRedshirting)
		Tackle = PrimaryProgression(cp.Progression, cp.Tackle, cp.Position, cp.Archetype, SnapsPerGame, "Tackle", cp.IsRedshirting)

		// Secondary
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		Catching = SecondaryProgression(cp.Progression, cp.Catching)
		Speed = SecondaryProgression(cp.Progression, cp.Speed)
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
	} else if cp.Position == "DE" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		PassRush = PrimaryProgression(cp.Progression, cp.PassRush, cp.Position, cp.Archetype, SnapsPerGame, "Pass Rush", cp.IsRedshirting)
		RunDefense = PrimaryProgression(cp.Progression, cp.RunDefense, cp.Position, cp.Archetype, SnapsPerGame, "Run Defense", cp.IsRedshirting)
		Tackle = PrimaryProgression(cp.Progression, cp.Tackle, cp.Position, cp.Archetype, SnapsPerGame, "Tackle", cp.IsRedshirting)

		// Secondary
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		Catching = SecondaryProgression(cp.Progression, cp.Catching)
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
	} else if cp.Position == "OLB" || cp.Position == "ILB" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		PassRush = PrimaryProgression(cp.Progression, cp.PassRush, cp.Position, cp.Archetype, SnapsPerGame, "Pass Rush", cp.IsRedshirting)
		RunDefense = PrimaryProgression(cp.Progression, cp.RunDefense, cp.Position, cp.Archetype, SnapsPerGame, "Run Defense", cp.IsRedshirting)
		Tackle = PrimaryProgression(cp.Progression, cp.Tackle, cp.Position, cp.Archetype, SnapsPerGame, "Tackle", cp.IsRedshirting)
		ManCoverage = PrimaryProgression(cp.Progression, cp.ManCoverage, cp.Position, cp.Archetype, SnapsPerGame, "Man Coverage", cp.IsRedshirting)
		ZoneCoverage = PrimaryProgression(cp.Progression, cp.ZoneCoverage, cp.Position, cp.Archetype, SnapsPerGame, "Zone Coverage", cp.IsRedshirting)

		// Secondary
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		Catching = SecondaryProgression(cp.Progression, cp.Catching)
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
	} else if cp.Position == "CB" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		Tackle = PrimaryProgression(cp.Progression, cp.Tackle, cp.Position, cp.Archetype, SnapsPerGame, "Tackle", cp.IsRedshirting)
		ManCoverage = PrimaryProgression(cp.Progression, cp.ManCoverage, cp.Position, cp.Archetype, SnapsPerGame, "Man Coverage", cp.IsRedshirting)
		ZoneCoverage = PrimaryProgression(cp.Progression, cp.ZoneCoverage, cp.Position, cp.Archetype, SnapsPerGame, "Zone Coverage", cp.IsRedshirting)
		Catching = PrimaryProgression(cp.Progression, cp.Catching, cp.Position, cp.Archetype, SnapsPerGame, "Catching", cp.IsRedshirting)

		// Secondary
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
	} else if cp.Position == "FS" || cp.Position == "SS" {
		// Primary
		Agility = PrimaryProgression(cp.Progression, cp.Agility, cp.Position, cp.Archetype, SnapsPerGame, "Agility", cp.IsRedshirting)
		FootballIQ = PrimaryProgression(cp.Progression, cp.FootballIQ, cp.Position, cp.Archetype, SnapsPerGame, "Football IQ", cp.IsRedshirting)
		Strength = PrimaryProgression(cp.Progression, cp.Strength, cp.Position, cp.Archetype, SnapsPerGame, "Strength", cp.IsRedshirting)
		Speed = PrimaryProgression(cp.Progression, cp.Speed, cp.Position, cp.Archetype, SnapsPerGame, "Speed", cp.IsRedshirting)
		RunDefense = PrimaryProgression(cp.Progression, cp.RunDefense, cp.Position, cp.Archetype, SnapsPerGame, "Run Defense", cp.IsRedshirting)
		Tackle = PrimaryProgression(cp.Progression, cp.Tackle, cp.Position, cp.Archetype, SnapsPerGame, "Tackle", cp.IsRedshirting)
		ManCoverage = PrimaryProgression(cp.Progression, cp.ManCoverage, cp.Position, cp.Archetype, SnapsPerGame, "Man Coverage", cp.IsRedshirting)
		ZoneCoverage = PrimaryProgression(cp.Progression, cp.ZoneCoverage, cp.Position, cp.Archetype, SnapsPerGame, "Zone Coverage", cp.IsRedshirting)
		Catching = PrimaryProgression(cp.Progression, cp.Catching, cp.Position, cp.Archetype, SnapsPerGame, "Catching", cp.IsRedshirting)

		// Secondary
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
	} else if cp.Position == "K" {
		// Primary
		KickPower = PrimaryProgression(cp.Progression, cp.KickPower, cp.Position, cp.Archetype, SnapsPerGame, "Kick Power", cp.IsRedshirting)
		KickAccuracy = PrimaryProgression(cp.Progression, cp.KickAccuracy, cp.Position, cp.Archetype, SnapsPerGame, "Kick Accuracy", cp.IsRedshirting)
		// Secondary
		ThrowPower = SecondaryProgression(cp.Progression, cp.ThrowPower)
		ThrowAccuracy = SecondaryProgression(cp.Progression, cp.ThrowAccuracy)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		PuntPower = SecondaryProgression(cp.Progression, cp.PuntPower)
		PuntAccuracy = SecondaryProgression(cp.Progression, cp.PuntAccuracy)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
	} else if cp.Position == "P" {
		// If David Ross
		if cp.ID == 24984 {
			ThrowPower = PrimaryProgression(cp.Progression, cp.ThrowPower, cp.Position, cp.Archetype, SnapsPerGame, "Throw Power", cp.IsRedshirting)
			ThrowAccuracy = PrimaryProgression(cp.Progression, cp.ThrowAccuracy, cp.Position, cp.Archetype, SnapsPerGame, "Throw Accuracy", cp.IsRedshirting)
		}
		// Primary Progressions
		PuntPower = PrimaryProgression(cp.Progression, cp.PuntPower, cp.Position, cp.Archetype, SnapsPerGame, "Punt Power", cp.IsRedshirting)
		PuntAccuracy = PrimaryProgression(cp.Progression, cp.PuntAccuracy, cp.Position, cp.Archetype, SnapsPerGame, "Punt Accuracy", cp.IsRedshirting)
		// Secondary Progressions
		RunBlock = SecondaryProgression(cp.Progression, cp.RunBlock)
		PassBlock = SecondaryProgression(cp.Progression, cp.PassBlock)
		RunDefense = SecondaryProgression(cp.Progression, cp.RunDefense)
		PassRush = SecondaryProgression(cp.Progression, cp.PassRush)
		Carrying = SecondaryProgression(cp.Progression, cp.Carrying)
		Tackle = SecondaryProgression(cp.Progression, cp.Tackle)
		RouteRunning = SecondaryProgression(cp.RouteRunning, cp.RouteRunning)
		Catching = SecondaryProgression(cp.Progression, cp.Catching)
		KickPower = SecondaryProgression(cp.Progression, cp.KickPower)
		KickAccuracy = SecondaryProgression(cp.Progression, cp.KickAccuracy)
		ManCoverage = SecondaryProgression(cp.Progression, cp.ManCoverage)
		ZoneCoverage = SecondaryProgression(cp.Progression, cp.ZoneCoverage)
	}

	progressions := structs.CollegePlayerProgressions{
		Agility:       Agility,
		Catching:      Catching,
		Carrying:      Carrying,
		Speed:         Speed,
		RouteRunning:  RouteRunning,
		RunBlock:      RunBlock,
		PassBlock:     PassBlock,
		RunDefense:    RunDefense,
		PassRush:      PassRush,
		Strength:      Strength,
		Tackle:        Tackle,
		ThrowPower:    ThrowPower,
		ThrowAccuracy: ThrowAccuracy,
		PuntAccuracy:  PuntAccuracy,
		PuntPower:     PuntPower,
		KickAccuracy:  KickAccuracy,
		KickPower:     KickPower,
		FootballIQ:    FootballIQ,
		ManCoverage:   ManCoverage,
		ZoneCoverage:  ZoneCoverage,
	}

	cp.Progress(progressions)

	if cp.IsRedshirting {
		cp.SetRedshirtStatus()
	}

	cp.GetOverall()

	return cp
}

func ProgressCroot(c structs.Recruit, SeasonID string) structs.Recruit {
	Agility := 0
	ThrowPower := 0
	ThrowAccuracy := 0
	Speed := 0
	FootballIQ := 0
	Strength := 0
	RunBlock := 0
	PassBlock := 0
	RunDefense := 0
	PassRush := 0
	Carrying := 0
	Tackle := 0
	RouteRunning := 0
	Catching := 0
	PuntPower := 0
	PuntAccuracy := 0
	KickPower := 0
	KickAccuracy := 0
	ManCoverage := 0
	ZoneCoverage := 0

	if c.Position == "QB" {
		// Primary Progressions
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		ThrowPower = PrimaryProgression(c.Progression, c.ThrowPower, c.Position, c.Archetype, 0, "Throw Power", true)
		ThrowAccuracy = PrimaryProgression(c.Progression, c.ThrowAccuracy, c.Position, c.Archetype, 0, "Throw Accuracy", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)

		// Secondary Progressions
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
		Catching = SecondaryProgression(c.Progression, c.Catching)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
	} else if c.Position == "RB" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		Carrying = PrimaryProgression(c.Progression, c.Carrying, c.Position, c.Archetype, 0, "Carrying", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Catching = PrimaryProgression(c.Progression, c.Catching, c.Position, c.Archetype, 0, "Catching", true)
		PassBlock = PrimaryProgression(c.Progression, c.PassBlock, c.Position, c.Archetype, 0, "Pass Blocking", true)
		// Secondary
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
	} else if c.Position == "FB" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		Carrying = PrimaryProgression(c.Progression, c.Carrying, c.Position, c.Archetype, 0, "Carrying", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Catching = PrimaryProgression(c.Progression, c.Catching, c.Position, c.Archetype, 0, "Catching", true)
		PassBlock = PrimaryProgression(c.Progression, c.PassBlock, c.Position, c.Archetype, 0, "Pass Blocking", true)
		RunBlock = PrimaryProgression(c.Progression, c.RunBlock, c.Position, c.Archetype, 0, "Run Blocking", true)

		// Secondary
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
	} else if c.Position == "TE" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		Carrying = PrimaryProgression(c.Progression, c.Carrying, c.Position, c.Archetype, 0, "Carrying", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Catching = PrimaryProgression(c.Progression, c.Catching, c.Position, c.Archetype, 0, "Catching", true)
		RouteRunning = PrimaryProgression(c.Progression, c.RouteRunning, c.Position, c.Archetype, 0, "Route Running", true)
		PassBlock = PrimaryProgression(c.Progression, c.PassBlock, c.Position, c.Archetype, 0, "Pass Blocking", true)
		RunBlock = PrimaryProgression(c.Progression, c.RunBlock, c.Position, c.Archetype, 0, "Run Blocking", true)

		// Secondary
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
	} else if c.Position == "WR" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		Carrying = PrimaryProgression(c.Progression, c.Carrying, c.Position, c.Archetype, 0, "Carrying", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Catching = PrimaryProgression(c.Progression, c.Catching, c.Position, c.Archetype, 0, "Catching", true)
		RouteRunning = PrimaryProgression(c.Progression, c.RouteRunning, c.Position, c.Archetype, 0, "Route Running", true)

		// Secondary
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
	} else if c.Position == "OT" || c.Position == "OG" || c.Position == "C" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		PassBlock = PrimaryProgression(c.Progression, c.PassBlock, c.Position, c.Archetype, 0, "Pass Blocking", true)
		RunBlock = PrimaryProgression(c.Progression, c.RunBlock, c.Position, c.Archetype, 0, "Run Blocking", true)

		// Secondary
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		Speed = SecondaryProgression(c.Progression, c.Speed)
		Catching = SecondaryProgression(c.Progression, c.Catching)
	} else if c.Position == "DT" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		PassRush = PrimaryProgression(c.Progression, c.PassRush, c.Position, c.Archetype, 0, "Pass Rush", true)
		RunDefense = PrimaryProgression(c.Progression, c.RunDefense, c.Position, c.Archetype, 0, "Run Defense", true)
		Tackle = PrimaryProgression(c.Progression, c.Tackle, c.Position, c.Archetype, 0, "Tackle", true)

		// Secondary
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		Catching = SecondaryProgression(c.Progression, c.Catching)
		Speed = SecondaryProgression(c.Progression, c.Speed)
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
	} else if c.Position == "DE" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		PassRush = PrimaryProgression(c.Progression, c.PassRush, c.Position, c.Archetype, 0, "Pass Rush", true)
		RunDefense = PrimaryProgression(c.Progression, c.RunDefense, c.Position, c.Archetype, 0, "Run Defense", true)
		Tackle = PrimaryProgression(c.Progression, c.Tackle, c.Position, c.Archetype, 0, "Tackle", true)

		// Secondary
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		Catching = SecondaryProgression(c.Progression, c.Catching)
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
	} else if c.Position == "OLB" || c.Position == "ILB" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		PassRush = PrimaryProgression(c.Progression, c.PassRush, c.Position, c.Archetype, 0, "Pass Rush", true)
		RunDefense = PrimaryProgression(c.Progression, c.RunDefense, c.Position, c.Archetype, 0, "Run Defense", true)
		Tackle = PrimaryProgression(c.Progression, c.Tackle, c.Position, c.Archetype, 0, "Tackle", true)
		ManCoverage = PrimaryProgression(c.Progression, c.ManCoverage, c.Position, c.Archetype, 0, "Man Coverage", true)
		ZoneCoverage = PrimaryProgression(c.Progression, c.ZoneCoverage, c.Position, c.Archetype, 0, "Zone Coverage", true)

		// Secondary
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		Catching = SecondaryProgression(c.Progression, c.Catching)
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
	} else if c.Position == "CB" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		Tackle = PrimaryProgression(c.Progression, c.Tackle, c.Position, c.Archetype, 0, "Tackle", true)
		ManCoverage = PrimaryProgression(c.Progression, c.ManCoverage, c.Position, c.Archetype, 0, "Man Coverage", true)
		ZoneCoverage = PrimaryProgression(c.Progression, c.ZoneCoverage, c.Position, c.Archetype, 0, "Zone Coverage", true)
		Catching = PrimaryProgression(c.Progression, c.Catching, c.Position, c.Archetype, 0, "Catching", true)

		// Secondary
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
	} else if c.Position == "FS" || c.Position == "SS" {
		// Primary
		Agility = PrimaryProgression(c.Progression, c.Agility, c.Position, c.Archetype, 0, "Agility", true)
		FootballIQ = PrimaryProgression(c.Progression, c.FootballIQ, c.Position, c.Archetype, 0, "Football IQ", true)
		Strength = PrimaryProgression(c.Progression, c.Strength, c.Position, c.Archetype, 0, "Strength", true)
		Speed = PrimaryProgression(c.Progression, c.Speed, c.Position, c.Archetype, 0, "Speed", true)
		RunDefense = PrimaryProgression(c.Progression, c.RunDefense, c.Position, c.Archetype, 0, "Run Defense", true)
		Tackle = PrimaryProgression(c.Progression, c.Tackle, c.Position, c.Archetype, 0, "Tackle", true)
		ManCoverage = PrimaryProgression(c.Progression, c.ManCoverage, c.Position, c.Archetype, 0, "Man Coverage", true)
		ZoneCoverage = PrimaryProgression(c.Progression, c.ZoneCoverage, c.Position, c.Archetype, 0, "Zone Coverage", true)
		Catching = PrimaryProgression(c.Progression, c.Catching, c.Position, c.Archetype, 0, "Catching", true)

		// Secondary
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
	} else if c.Position == "K" {
		// Primary
		KickPower = PrimaryProgression(c.Progression, c.KickPower, c.Position, c.Archetype, 0, "Kick Power", true)
		KickAccuracy = PrimaryProgression(c.Progression, c.KickAccuracy, c.Position, c.Archetype, 0, "Kick Accuracy", true)
		// Secondary
		ThrowPower = SecondaryProgression(c.Progression, c.ThrowPower)
		ThrowAccuracy = SecondaryProgression(c.Progression, c.ThrowAccuracy)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		PuntPower = SecondaryProgression(c.Progression, c.PuntPower)
		PuntAccuracy = SecondaryProgression(c.Progression, c.PuntAccuracy)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
	} else if c.Position == "P" {
		// Primary Progressions
		PuntPower = PrimaryProgression(c.Progression, c.PuntPower, c.Position, c.Archetype, 0, "Punt Power", true)
		PuntAccuracy = PrimaryProgression(c.Progression, c.PuntAccuracy, c.Position, c.Archetype, 0, "Punt Accuracy", true)
		// Secondary Progressions
		RunBlock = SecondaryProgression(c.Progression, c.RunBlock)
		PassBlock = SecondaryProgression(c.Progression, c.PassBlock)
		RunDefense = SecondaryProgression(c.Progression, c.RunDefense)
		PassRush = SecondaryProgression(c.Progression, c.PassRush)
		Carrying = SecondaryProgression(c.Progression, c.Carrying)
		Tackle = SecondaryProgression(c.Progression, c.Tackle)
		RouteRunning = SecondaryProgression(c.RouteRunning, c.RouteRunning)
		Catching = SecondaryProgression(c.Progression, c.Catching)
		KickPower = SecondaryProgression(c.Progression, c.KickPower)
		KickAccuracy = SecondaryProgression(c.Progression, c.KickAccuracy)
		ManCoverage = SecondaryProgression(c.Progression, c.ManCoverage)
		ZoneCoverage = SecondaryProgression(c.Progression, c.ZoneCoverage)
	}

	progressions := structs.CollegePlayerProgressions{
		Agility:       Agility,
		Catching:      Catching,
		Carrying:      Carrying,
		Speed:         Speed,
		RouteRunning:  RouteRunning,
		RunBlock:      RunBlock,
		PassBlock:     PassBlock,
		RunDefense:    RunDefense,
		PassRush:      PassRush,
		Strength:      Strength,
		Tackle:        Tackle,
		ThrowPower:    ThrowPower,
		ThrowAccuracy: ThrowAccuracy,
		PuntAccuracy:  PuntAccuracy,
		PuntPower:     PuntPower,
		KickAccuracy:  KickAccuracy,
		KickPower:     KickPower,
		FootballIQ:    FootballIQ,
		ManCoverage:   ManCoverage,
		ZoneCoverage:  ZoneCoverage,
	}

	c.ProgressUnsignedRecruit(progressions)

	c.GetOverall()

	return c
}

func PrimaryProgression(progression int, input int, position string, archetype string, spg int, attribute string, isRedshirting bool) int {
	if input == 0 {
		return 1
	}

	modifier := GetModifiers(position, archetype, spg, attribute)

	var progress float64 = 0

	if !isRedshirting {
		progress = ((1 - math.Pow((float64(input)/99.0), 15)) * math.Log10(float64(input)) * (0.6 + modifier)) * (1 + (float64(progression) / 70))
	} else {
		progress = ((1 - math.Pow((float64(input)/99), 15)) * math.Log10(float64(input)) * 1.225 * (1 + (float64(progression / 60))))
	}

	if progress+float64(input) > 99 {
		progress = 99
	} else {
		progress = progress + float64(input)
	}

	return int(math.Round(progress))
}

func SecondaryProgression(progression int, input int) int {
	num := rand.Intn(99)

	if num < progression && input < 99 {
		input = input + 1
		return input
	} else {
		return input
	}
}

func GetModifiers(position string, archetype string, spg int, attrib string) float64 {
	var snapMod float64 = 0.0
	if spg > 50 {
		snapMod = rand.Float64()*(1.25-1) + 1
	} else if spg > 30 {
		snapMod = rand.Float64()*(1.1-0.9) + 0.9
	} else if spg > 20 {
		snapMod = rand.Float64()*(1.0-0.75) + 0.75
	} else if spg > 10 {
		snapMod = rand.Float64()*(.9-0.6) + 0.6
	} else {
		snapMod = rand.Float64()*(0.75-0.5) + 0.5
	}
	archMod := GetArchetypeMod(position, archetype, attrib)
	return snapMod + archMod
}

func GetArchetypeMod(pos string, arch string, attribute string) float64 {
	if pos == "QB" {
		if arch == "Pocket" {
			if attribute == "Throw Power" || attribute == "Throw Accuracy" {
				return 0.05
			}
		} else if arch == "Balanced" {
			if attribute == "Agility" {
				return 0.05
			} else if attribute == "Throw Power" || attribute == "Throw Accuracy" {
				return 0.025
			}

		} else if arch == "Field General" {
			if attribute == "Football IQ" || attribute == "Throw Accuracy" {
				return 0.05
			}
		} else if arch == "Scrambler" {
			if attribute == "Agility" || attribute == "Speed" {
				return 0.05
			}
		}
		return 0
	} else if pos == "RB" {
		if arch == "Speed" {
			if attribute == "Speed" || attribute == "Carrying" {
				return 0.05
			}
			return 0.0
		} else if arch == "Power" {
			if attribute == "Carrying" || attribute == "Strength" {
				return 0.05
			}
			return 0.0
		} else if arch == "Balanced" {
			if attribute == "Agility" || attribute == "Pass Block" {
				return 0.05
			}
			return 0.0
		} else if arch == "Receiving" {
			if attribute == "Agility" || attribute == "Catching" {
				return 0.05
			}
			return 0.0
		}
		return 0
	} else if pos == "FB" {
		if arch == "Blocking" && (attribute == "Pass Blocking" || attribute == "Run Blocking") {
			return 0.05
		} else if arch == "Rushing" && (attribute == "Carrying" || attribute == "Strength") {
			return 0.05
		} else if arch == "Balanced" && (attribute == "Agility" || attribute == "Speed") {
			return 0.05
		} else if arch == "Receiving" && (attribute == "Agility" || attribute == "Catching") {
			return 0.05
		}
		return 0
	} else if pos == "WR" {
		if arch == "Possession" && (attribute == "Carrying" || attribute == "Catching") {
			return 0.05
		} else if arch == "Route Runner" && (attribute == "Agility" || attribute == "Route Running") {
			return 0.05
		} else if arch == "Speed" && attribute == "Speed" {
			return 0.1
		} else if arch == "Red Zone Threat" && (attribute == "Strength" || attribute == "Catching") {
			return 0.05
		}
		return 0
	} else if pos == "TE" {
		if arch == "Blocking" && (attribute == "Pass Blocking" || attribute == "Run Blocking") {
			return 0.05
		} else if arch == "Vertical Threat" && (attribute == "Catching" || attribute == "Speed") {
			return 0.05
		} else if arch == "Receiving" && (attribute == "Agility" || attribute == "Catching") {
			return 0.05
		}
		return 0
	} else if pos == "OT" {
		if arch == "Pass Blocking" && (attribute == "Pass Blocking" || attribute == "Strength") {
			return 0.05
		} else if arch == "Run Blocking" && (attribute == "Run Blocking" || attribute == "Strength") {
			return 0.05
		} else if arch == "Balanced" {
			if attribute == "Run Blocking" || attribute == "Pass Blocking" {
				return 0.025
			} else if attribute == "Strength" {
				return 0.05
			}
		}
	} else if pos == "OG" {
		if arch == "Pass Blocking" && (attribute == "Pass Blocking" || attribute == "Strength") {
			return 0.05
		} else if arch == "Run Blocking" && (attribute == "Run Blocking" || attribute == "Strength") {
			return 0.05
		} else if arch == "Balanced" {
			if attribute == "Run Blocking" || attribute == "Pass Blocking" {
				return 0.025
			} else if attribute == "Strength" {
				return 0.05
			}
		}
	} else if pos == "C" {
		if arch == "Pass Blocking" && (attribute == "Pass Blocking" || attribute == "Strength") {
			return 0.05
		} else if arch == "Run Blocking" && (attribute == "Run Blocking" || attribute == "Strength") {
			return 0.05
		} else if arch == "Balanced" {
			if attribute == "Run Blocking" || attribute == "Pass Blocking" {
				return 0.025
			} else if attribute == "Strength" {
				return 0.05
			}
		} else if arch == "Line Captain" {
			if attribute == "Run Blocking" || attribute == "Pass Blocking" {
				return 0.025
			} else if attribute == "Football IQ" {
				return 0.05
			}
		}
	} else if pos == "DT" {
		if arch == "Pass Rusher" && (attribute == "Pass Rush" || attribute == "Strength") {
			return 0.05
		} else if arch == "Nose Tackle" && (attribute == "Run Defense" || attribute == "Strength") {
			return 0.05
		} else if arch == "Balanced" && (attribute == "Tackle" || attribute == "Strength") {
			return 0.05
		}
	} else if pos == "DE" {
		if arch == "Speed Rusher" && attribute == "Pass Rush" {
			return 0.1
		} else if arch == "Run Stopper" && attribute == "Run Defense" {
			return 0.1
		} else if arch == "Balanced" && (attribute == "Pass Rush" || attribute == "Run Defense") {
			return 0.05
		}
	} else if pos == "ILB" {
		if arch == "Field General" && (attribute == "Football IQ" || attribute == "Tackle") {
			return 0.05
		} else if arch == "Run Stopper" && (attribute == "Run Defense" || attribute == "Strength") {
			return 0.05
		} else if arch == "Coverage" && (attribute == "Zone Coverage" || attribute == "Man Coverage") {
			return 0.05
		} else if arch == "Speed" && (attribute == "Speed" || attribute == "Agility") {
			return 0.05
		}
	} else if pos == "OLB" {
		if arch == "Pass Rush" && attribute == "Pass Rush" {
			return 0.1
		} else if arch == "Run Stopper" && (attribute == "Run Defense" || attribute == "Strength") {
			return 0.05
		} else if arch == "Coverage" && (attribute == "Zone Coverage" || attribute == "Man Coverage") {
			return 0.05
		} else if arch == "Speed" && (attribute == "Speed" || attribute == "Agility") {
			return 0.05
		}
	} else if pos == "CB" {
		if arch == "Man Coverage" && (attribute == "Man Coverage" || attribute == "Speed") {
			return 0.05
		} else if arch == "Zone Coverage" && (attribute == "Zone Coverage" || attribute == "Football IQ") {
			return 0.05
		} else if arch == "Ball Hawk" {
			if attribute == "Man Coverage" || attribute == "Zone Coverage" {
				return 0.025
			} else if attribute == "Catching" {
				return 0.05
			}
		}
	} else if pos == "FS" || pos == "SS" {
		if arch == "Man Coverage" && (attribute == "Man Coverage" || attribute == "Speed") {
			return 0.05
		} else if arch == "Zone Coverage" && (attribute == "Zone Coverage" || attribute == "Football IQ") {
			return 0.05
		} else if arch == "Ball Hawk" {
			if attribute == "Man Coverage" || attribute == "Zone Coverage" {
				return 0.025
			} else if attribute == "Catching" {
				return 0.05
			}
		} else if arch == "Run Stopper" && (attribute == "Run Defense" || attribute == "Tackle") {
			return 0.05
		}
	} else if pos == "K" {
		if arch == "Power" && attribute == "Kick Power" {
			return 0.1
		} else if arch == "Accuracy" && attribute == "Kick Accuracy" {
			return 0.1
		} else if arch == "Balanced" && (attribute == "Kick Power" || attribute == "Kick Accuracy") {
			return 0.05
		}
	} else if pos == "P" {
		if arch == "Power" && attribute == "Punt Power" {
			return 0.1
		} else if arch == "Accuracy" && attribute == "Punt Accuracy" {
			return 0.1
		} else if arch == "Balanced" && (attribute == "Punt Power" || attribute == "Punt Accuracy") {
			return 0.05
		}
	} else if pos == "ATH" {
		return 0.0075
	}
	return 0
}
