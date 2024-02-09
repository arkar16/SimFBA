package models

import (
	"strconv"

	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

type CollegePlayerCSV struct {
	FirstName          string
	LastName           string
	Position           string
	Archetype          string
	Year               string
	Team               string
	Age                int
	Stars              int
	HighSchool         string
	City               string
	State              string
	Height             int
	Weight             int
	Shotgun            int
	OverallGrade       string
	StaminaGrade       string
	InjuryGrade        string
	FootballIQGrade    string
	SpeedGrade         string
	CarryingGrade      string
	AgilityGrade       string
	CatchingGrade      string
	RouteRunningGrade  string
	ZoneCoverageGrade  string
	ManCoverageGrade   string
	StrengthGrade      string
	TackleGrade        string
	PassBlockGrade     string
	RunBlockGrade      string
	PassRushGrade      string
	RunDefenseGrade    string
	ThrowPowerGrade    string
	ThrowAccuracyGrade string
	KickAccuracyGrade  string
	KickPowerGrade     string
	PuntAccuracyGrade  string
	PuntPowerGrade     string
	PotentialGrade     string
	RedshirtStatus     string
	Stats              []structs.CollegePlayerStats
}

func MapPlayerForStats(player structs.CollegePlayer) CollegePlayerCSV {
	attributeMeans := config.AttributeMeans()
	Year, RedShirtStatus := util.GetYearAndRedshirtStatus(player.Year, player.IsRedshirt)
	OverallGrade := util.GetOverallGrade(player.Overall, player.Year)
	StaminaGrade := util.GetLetterGrade(player.Stamina, attributeMeans["Stamina"][player.Position]["mean"], attributeMeans["Stamina"][player.Position]["stddev"], player.Year)
	InjuryGrade := util.GetLetterGrade(player.Injury, attributeMeans["Injury"][player.Position]["mean"], attributeMeans["Injury"][player.Position]["stddev"], player.Year)
	SpeedGrade := util.GetLetterGrade(player.Speed, attributeMeans["Speed"][player.Position]["mean"], attributeMeans["Speed"][player.Position]["stddev"], player.Year)
	FootballIQGrade := util.GetLetterGrade(player.FootballIQ, attributeMeans["FootballIQ"][player.Position]["mean"], attributeMeans["FootballIQ"][player.Position]["stddev"], player.Year)
	AgilityGrade := util.GetLetterGrade(player.Agility, attributeMeans["Agility"][player.Position]["mean"], attributeMeans["Agility"][player.Position]["stddev"], player.Year)
	CarryingGrade := util.GetLetterGrade(player.Carrying, attributeMeans["Carrying"][player.Position]["mean"], attributeMeans["Carrying"][player.Position]["stddev"], player.Year)
	CatchingGrade := util.GetLetterGrade(player.Catching, attributeMeans["Catching"][player.Position]["mean"], attributeMeans["Catching"][player.Position]["stddev"], player.Year)
	RouteRunningGrade := util.GetLetterGrade(player.RouteRunning, attributeMeans["RouteRunning"][player.Position]["mean"], attributeMeans["RouteRunning"][player.Position]["stddev"], player.Year)
	ZoneCoverageGrade := util.GetLetterGrade(player.ZoneCoverage, attributeMeans["ZoneCoverage"][player.Position]["mean"], attributeMeans["ZoneCoverage"][player.Position]["stddev"], player.Year)
	ManCoverageGrade := util.GetLetterGrade(player.ManCoverage, attributeMeans["ManCoverage"][player.Position]["mean"], attributeMeans["ManCoverage"][player.Position]["stddev"], player.Year)
	StrengthGrade := util.GetLetterGrade(player.Strength, attributeMeans["Strength"][player.Position]["mean"], attributeMeans["Strength"][player.Position]["stddev"], player.Year)
	TackleGrade := util.GetLetterGrade(player.Tackle, attributeMeans["Tackle"][player.Position]["mean"], attributeMeans["Tackle"][player.Position]["stddev"], player.Year)
	PassBlockGrade := util.GetLetterGrade(player.PassBlock, attributeMeans["PassBlock"][player.Position]["mean"], attributeMeans["PassBlock"][player.Position]["stddev"], player.Year)
	RunBlockGrade := util.GetLetterGrade(player.RunBlock, attributeMeans["RunBlock"][player.Position]["mean"], attributeMeans["RunBlock"][player.Position]["stddev"], player.Year)
	PassRushGrade := util.GetLetterGrade(player.PassRush, attributeMeans["PassRush"][player.Position]["mean"], attributeMeans["PassRush"][player.Position]["stddev"], player.Year)
	RunDefenseGrade := util.GetLetterGrade(player.RunDefense, attributeMeans["RunDefense"][player.Position]["mean"], attributeMeans["RunDefense"][player.Position]["stddev"], player.Year)
	ThrowPowerGrade := util.GetLetterGrade(player.ThrowPower, attributeMeans["ThrowPower"][player.Position]["mean"], attributeMeans["ThrowPower"][player.Position]["stddev"], player.Year)
	ThrowAccuracyGrade := util.GetLetterGrade(player.ThrowAccuracy, attributeMeans["ThrowAccuracy"][player.Position]["mean"], attributeMeans["ThrowAccuracy"][player.Position]["stddev"], player.Year)
	KickPowerGrade := util.GetLetterGrade(player.KickPower, attributeMeans["KickPower"][player.Position]["mean"], attributeMeans["KickPower"][player.Position]["stddev"], player.Year)
	KickAccuracyGrade := util.GetLetterGrade(player.KickAccuracy, attributeMeans["KickAccuracy"][player.Position]["mean"], attributeMeans["KickAccuracy"][player.Position]["stddev"], player.Year)
	PuntPowerGrade := util.GetLetterGrade(player.PuntPower, attributeMeans["PuntPower"][player.Position]["mean"], attributeMeans["PuntPower"][player.Position]["stddev"], player.Year)
	PuntAccuracyGrade := util.GetLetterGrade(player.PuntAccuracy, attributeMeans["PuntAccuracy"][player.Position]["mean"], attributeMeans["PuntAccuracy"][player.Position]["stddev"], player.Year)

	return CollegePlayerCSV{
		FirstName:          player.FirstName,
		LastName:           player.LastName,
		Position:           player.Position,
		Archetype:          player.Archetype,
		Year:               Year,
		Age:                player.Age,
		Stars:              player.Stars,
		HighSchool:         player.HighSchool,
		City:               player.City,
		State:              player.State,
		Height:             player.Height,
		Weight:             player.Weight,
		OverallGrade:       OverallGrade,
		StaminaGrade:       StaminaGrade,
		InjuryGrade:        InjuryGrade,
		FootballIQGrade:    FootballIQGrade,
		SpeedGrade:         SpeedGrade,
		CarryingGrade:      CarryingGrade,
		AgilityGrade:       AgilityGrade,
		CatchingGrade:      CatchingGrade,
		RouteRunningGrade:  RouteRunningGrade,
		ZoneCoverageGrade:  ZoneCoverageGrade,
		ManCoverageGrade:   ManCoverageGrade,
		StrengthGrade:      StrengthGrade,
		TackleGrade:        TackleGrade,
		PassBlockGrade:     PassBlockGrade,
		RunBlockGrade:      RunBlockGrade,
		PassRushGrade:      PassRushGrade,
		RunDefenseGrade:    RunDefenseGrade,
		ThrowPowerGrade:    ThrowPowerGrade,
		ThrowAccuracyGrade: ThrowAccuracyGrade,
		KickAccuracyGrade:  KickAccuracyGrade,
		KickPowerGrade:     KickPowerGrade,
		PuntAccuracyGrade:  PuntAccuracyGrade,
		PuntPowerGrade:     PuntPowerGrade,
		PotentialGrade:     player.PotentialGrade,
		RedshirtStatus:     RedShirtStatus,
		Stats:              player.Stats,
	}
}

func MapPlayerToCSVModel(player structs.CollegePlayer) CollegePlayerCSV {

	attributeMeans := config.AttributeMeans()
	Year, RedShirtStatus := util.GetYearAndRedshirtStatus(player.Year, player.IsRedshirt)
	OverallGrade := util.GetOverallGrade(player.Overall, player.Year)
	StaminaGrade := util.GetLetterGrade(player.Stamina, attributeMeans["Stamina"][player.Position]["mean"], attributeMeans["Stamina"][player.Position]["stddev"], player.Year)
	InjuryGrade := util.GetLetterGrade(player.Injury, attributeMeans["Injury"][player.Position]["mean"], attributeMeans["Injury"][player.Position]["stddev"], player.Year)
	SpeedGrade := util.GetLetterGrade(player.Speed, attributeMeans["Speed"][player.Position]["mean"], attributeMeans["Speed"][player.Position]["stddev"], player.Year)
	FootballIQGrade := util.GetLetterGrade(player.FootballIQ, attributeMeans["FootballIQ"][player.Position]["mean"], attributeMeans["FootballIQ"][player.Position]["stddev"], player.Year)
	AgilityGrade := util.GetLetterGrade(player.Agility, attributeMeans["Agility"][player.Position]["mean"], attributeMeans["Agility"][player.Position]["stddev"], player.Year)
	CarryingGrade := util.GetLetterGrade(player.Carrying, attributeMeans["Carrying"][player.Position]["mean"], attributeMeans["Carrying"][player.Position]["stddev"], player.Year)
	CatchingGrade := util.GetLetterGrade(player.Catching, attributeMeans["Catching"][player.Position]["mean"], attributeMeans["Catching"][player.Position]["stddev"], player.Year)
	RouteRunningGrade := util.GetLetterGrade(player.RouteRunning, attributeMeans["RouteRunning"][player.Position]["mean"], attributeMeans["RouteRunning"][player.Position]["stddev"], player.Year)
	ZoneCoverageGrade := util.GetLetterGrade(player.ZoneCoverage, attributeMeans["ZoneCoverage"][player.Position]["mean"], attributeMeans["ZoneCoverage"][player.Position]["stddev"], player.Year)
	ManCoverageGrade := util.GetLetterGrade(player.ManCoverage, attributeMeans["ManCoverage"][player.Position]["mean"], attributeMeans["ManCoverage"][player.Position]["stddev"], player.Year)
	StrengthGrade := util.GetLetterGrade(player.Strength, attributeMeans["Strength"][player.Position]["mean"], attributeMeans["Strength"][player.Position]["stddev"], player.Year)
	TackleGrade := util.GetLetterGrade(player.Tackle, attributeMeans["Tackle"][player.Position]["mean"], attributeMeans["Tackle"][player.Position]["stddev"], player.Year)
	PassBlockGrade := util.GetLetterGrade(player.PassBlock, attributeMeans["PassBlock"][player.Position]["mean"], attributeMeans["PassBlock"][player.Position]["stddev"], player.Year)
	RunBlockGrade := util.GetLetterGrade(player.RunBlock, attributeMeans["RunBlock"][player.Position]["mean"], attributeMeans["RunBlock"][player.Position]["stddev"], player.Year)
	PassRushGrade := util.GetLetterGrade(player.PassRush, attributeMeans["PassRush"][player.Position]["mean"], attributeMeans["PassRush"][player.Position]["stddev"], player.Year)
	RunDefenseGrade := util.GetLetterGrade(player.RunDefense, attributeMeans["RunDefense"][player.Position]["mean"], attributeMeans["RunDefense"][player.Position]["stddev"], player.Year)
	ThrowPowerGrade := util.GetLetterGrade(player.ThrowPower, attributeMeans["ThrowPower"][player.Position]["mean"], attributeMeans["ThrowPower"][player.Position]["stddev"], player.Year)
	ThrowAccuracyGrade := util.GetLetterGrade(player.ThrowAccuracy, attributeMeans["ThrowAccuracy"][player.Position]["mean"], attributeMeans["ThrowAccuracy"][player.Position]["stddev"], player.Year)
	KickPowerGrade := util.GetLetterGrade(player.KickPower, attributeMeans["KickPower"][player.Position]["mean"], attributeMeans["KickPower"][player.Position]["stddev"], player.Year)
	KickAccuracyGrade := util.GetLetterGrade(player.KickAccuracy, attributeMeans["KickAccuracy"][player.Position]["mean"], attributeMeans["KickAccuracy"][player.Position]["stddev"], player.Year)
	PuntPowerGrade := util.GetLetterGrade(player.PuntPower, attributeMeans["PuntPower"][player.Position]["mean"], attributeMeans["PuntPower"][player.Position]["stddev"], player.Year)
	PuntAccuracyGrade := util.GetLetterGrade(player.PuntAccuracy, attributeMeans["PuntAccuracy"][player.Position]["mean"], attributeMeans["PuntAccuracy"][player.Position]["stddev"], player.Year)

	return CollegePlayerCSV{
		FirstName:          player.FirstName,
		LastName:           player.LastName,
		Position:           player.Position,
		Archetype:          player.Archetype,
		Year:               Year,
		Age:                player.Age,
		Stars:              player.Stars,
		HighSchool:         player.HighSchool,
		City:               player.City,
		State:              player.State,
		Height:             player.Height,
		Weight:             player.Weight,
		OverallGrade:       OverallGrade,
		StaminaGrade:       StaminaGrade,
		InjuryGrade:        InjuryGrade,
		FootballIQGrade:    FootballIQGrade,
		SpeedGrade:         SpeedGrade,
		CarryingGrade:      CarryingGrade,
		AgilityGrade:       AgilityGrade,
		CatchingGrade:      CatchingGrade,
		RouteRunningGrade:  RouteRunningGrade,
		ZoneCoverageGrade:  ZoneCoverageGrade,
		ManCoverageGrade:   ManCoverageGrade,
		StrengthGrade:      StrengthGrade,
		TackleGrade:        TackleGrade,
		PassBlockGrade:     PassBlockGrade,
		RunBlockGrade:      RunBlockGrade,
		PassRushGrade:      PassRushGrade,
		RunDefenseGrade:    RunDefenseGrade,
		ThrowPowerGrade:    ThrowPowerGrade,
		ThrowAccuracyGrade: ThrowAccuracyGrade,
		KickAccuracyGrade:  KickAccuracyGrade,
		KickPowerGrade:     KickPowerGrade,
		PuntAccuracyGrade:  PuntAccuracyGrade,
		PuntPowerGrade:     PuntPowerGrade,
		PotentialGrade:     player.PotentialGrade,
		RedshirtStatus:     RedShirtStatus,
		Shotgun:            player.Shotgun,
		Team:               player.TeamAbbr,
	}
}

func MapNFLPlayerToCSVModel(player structs.NFLPlayer) CollegePlayerCSV {

	attributeMeans := config.AttributeMeans()
	Year := util.GetNFLYear(player.Experience)
	OverallGrade := strconv.Itoa(player.Overall)
	StaminaGrade := strconv.Itoa(player.Stamina)
	InjuryGrade := strconv.Itoa(player.Injury)
	SpeedGrade := strconv.Itoa(player.Speed)
	FootballIQGrade := strconv.Itoa(player.FootballIQ)
	AgilityGrade := strconv.Itoa(player.Agility)
	CarryingGrade := strconv.Itoa(player.Carrying)
	CatchingGrade := strconv.Itoa(player.Catching)
	RouteRunningGrade := strconv.Itoa(player.RouteRunning)
	ZoneCoverageGrade := strconv.Itoa(player.ZoneCoverage)
	ManCoverageGrade := strconv.Itoa(player.ManCoverage)
	StrengthGrade := strconv.Itoa(player.Strength)
	TackleGrade := strconv.Itoa(player.Tackle)
	PassBlockGrade := strconv.Itoa(player.PassBlock)
	RunBlockGrade := strconv.Itoa(player.RunBlock)
	PassRushGrade := strconv.Itoa(player.PassRush)
	RunDefenseGrade := strconv.Itoa(player.RunDefense)
	ThrowPowerGrade := strconv.Itoa(player.ThrowPower)
	ThrowAccuracyGrade := strconv.Itoa(player.ThrowAccuracy)
	KickPowerGrade := strconv.Itoa(player.KickPower)
	KickAccuracyGrade := strconv.Itoa(player.KickAccuracy)
	PuntPowerGrade := strconv.Itoa(player.PuntPower)
	PuntAccuracyGrade := strconv.Itoa(player.PuntAccuracy)

	if player.Experience < 2 {
		OverallGrade = util.GetOverallGrade(player.Overall, int(player.Experience))
		StaminaGrade = util.GetLetterGrade(player.Stamina, attributeMeans["Stamina"][player.Position]["mean"], attributeMeans["Stamina"][player.Position]["stddev"], int(player.Experience))
		InjuryGrade = util.GetLetterGrade(player.Injury, attributeMeans["Injury"][player.Position]["mean"], attributeMeans["Injury"][player.Position]["stddev"], int(player.Experience))
		SpeedGrade = util.GetLetterGrade(player.Speed, attributeMeans["Speed"][player.Position]["mean"], attributeMeans["Speed"][player.Position]["stddev"], int(player.Experience))
		FootballIQGrade = util.GetLetterGrade(player.FootballIQ, attributeMeans["FootballIQ"][player.Position]["mean"], attributeMeans["FootballIQ"][player.Position]["stddev"], int(player.Experience))
		AgilityGrade = util.GetLetterGrade(player.Agility, attributeMeans["Agility"][player.Position]["mean"], attributeMeans["Agility"][player.Position]["stddev"], int(player.Experience))
		CarryingGrade = util.GetLetterGrade(player.Carrying, attributeMeans["Carrying"][player.Position]["mean"], attributeMeans["Carrying"][player.Position]["stddev"], int(player.Experience))
		CatchingGrade = util.GetLetterGrade(player.Catching, attributeMeans["Catching"][player.Position]["mean"], attributeMeans["Catching"][player.Position]["stddev"], int(player.Experience))
		RouteRunningGrade = util.GetLetterGrade(player.RouteRunning, attributeMeans["RouteRunning"][player.Position]["mean"], attributeMeans["RouteRunning"][player.Position]["stddev"], int(player.Experience))
		ZoneCoverageGrade = util.GetLetterGrade(player.ZoneCoverage, attributeMeans["ZoneCoverage"][player.Position]["mean"], attributeMeans["ZoneCoverage"][player.Position]["stddev"], int(player.Experience))
		ManCoverageGrade = util.GetLetterGrade(player.ManCoverage, attributeMeans["ManCoverage"][player.Position]["mean"], attributeMeans["ManCoverage"][player.Position]["stddev"], int(player.Experience))
		StrengthGrade = util.GetLetterGrade(player.Strength, attributeMeans["Strength"][player.Position]["mean"], attributeMeans["Strength"][player.Position]["stddev"], int(player.Experience))
		TackleGrade = util.GetLetterGrade(player.Tackle, attributeMeans["Tackle"][player.Position]["mean"], attributeMeans["Tackle"][player.Position]["stddev"], int(player.Experience))
		PassBlockGrade = util.GetLetterGrade(player.PassBlock, attributeMeans["PassBlock"][player.Position]["mean"], attributeMeans["PassBlock"][player.Position]["stddev"], int(player.Experience))
		RunBlockGrade = util.GetLetterGrade(player.RunBlock, attributeMeans["RunBlock"][player.Position]["mean"], attributeMeans["RunBlock"][player.Position]["stddev"], int(player.Experience))
		PassRushGrade = util.GetLetterGrade(player.PassRush, attributeMeans["PassRush"][player.Position]["mean"], attributeMeans["PassRush"][player.Position]["stddev"], int(player.Experience))
		RunDefenseGrade = util.GetLetterGrade(player.RunDefense, attributeMeans["RunDefense"][player.Position]["mean"], attributeMeans["RunDefense"][player.Position]["stddev"], int(player.Experience))
		ThrowPowerGrade = util.GetLetterGrade(player.ThrowPower, attributeMeans["ThrowPower"][player.Position]["mean"], attributeMeans["ThrowPower"][player.Position]["stddev"], int(player.Experience))
		ThrowAccuracyGrade = util.GetLetterGrade(player.ThrowAccuracy, attributeMeans["ThrowAccuracy"][player.Position]["mean"], attributeMeans["ThrowAccuracy"][player.Position]["stddev"], int(player.Experience))
		KickPowerGrade = util.GetLetterGrade(player.KickPower, attributeMeans["KickPower"][player.Position]["mean"], attributeMeans["KickPower"][player.Position]["stddev"], int(player.Experience))
		KickAccuracyGrade = util.GetLetterGrade(player.KickAccuracy, attributeMeans["KickAccuracy"][player.Position]["mean"], attributeMeans["KickAccuracy"][player.Position]["stddev"], int(player.Experience))
		PuntPowerGrade = util.GetLetterGrade(player.PuntPower, attributeMeans["PuntPower"][player.Position]["mean"], attributeMeans["PuntPower"][player.Position]["stddev"], int(player.Experience))
		PuntAccuracyGrade = util.GetLetterGrade(player.PuntAccuracy, attributeMeans["PuntAccuracy"][player.Position]["mean"], attributeMeans["PuntAccuracy"][player.Position]["stddev"], int(player.Experience))
	}

	return CollegePlayerCSV{
		FirstName:          player.FirstName,
		LastName:           player.LastName,
		Position:           player.Position,
		Archetype:          player.Archetype,
		Year:               Year,
		Age:                player.Age,
		Stars:              player.Stars,
		HighSchool:         player.HighSchool,
		City:               player.Hometown,
		State:              player.State,
		Height:             player.Height,
		Weight:             player.Weight,
		OverallGrade:       OverallGrade,
		StaminaGrade:       StaminaGrade,
		InjuryGrade:        InjuryGrade,
		FootballIQGrade:    FootballIQGrade,
		SpeedGrade:         SpeedGrade,
		CarryingGrade:      CarryingGrade,
		AgilityGrade:       AgilityGrade,
		CatchingGrade:      CatchingGrade,
		RouteRunningGrade:  RouteRunningGrade,
		ZoneCoverageGrade:  ZoneCoverageGrade,
		ManCoverageGrade:   ManCoverageGrade,
		StrengthGrade:      StrengthGrade,
		TackleGrade:        TackleGrade,
		PassBlockGrade:     PassBlockGrade,
		RunBlockGrade:      RunBlockGrade,
		PassRushGrade:      PassRushGrade,
		RunDefenseGrade:    RunDefenseGrade,
		ThrowPowerGrade:    ThrowPowerGrade,
		ThrowAccuracyGrade: ThrowAccuracyGrade,
		KickAccuracyGrade:  KickAccuracyGrade,
		KickPowerGrade:     KickPowerGrade,
		PuntAccuracyGrade:  PuntAccuracyGrade,
		PuntPowerGrade:     PuntPowerGrade,
		PotentialGrade:     player.PotentialGrade,
	}
}
