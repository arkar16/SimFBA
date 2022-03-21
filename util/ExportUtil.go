package util

import (
	"github.com/CalebRose/SimFBA/models"
	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
)

func MapPlayerToCSVModel(player structs.CollegePlayer) models.CollegePlayerCSV {

	attributeMeans := config.AttributeMeans()
	Year, RedShirtStatus := GetYearAndRedshirtStatus(player.Year, player.IsRedshirt)
	OverallGrade := GetOverallGrade(player.Overall)
	StaminaGrade := GetLetterGrade(player.Stamina, attributeMeans["Stamina"][player.Position]["mean"], attributeMeans["Stamina"][player.Position]["stddev"])
	InjuryGrade := GetLetterGrade(player.Injury, attributeMeans["Injury"][player.Position]["mean"], attributeMeans["Injury"][player.Position]["stddev"])
	SpeedGrade := GetLetterGrade(player.Speed, attributeMeans["Speed"][player.Position]["mean"], attributeMeans["Speed"][player.Position]["stddev"])
	FootballIQGrade := GetLetterGrade(player.FootballIQ, attributeMeans["FootballIQ"][player.Position]["mean"], attributeMeans["FootballIQ"][player.Position]["stddev"])
	AgilityGrade := GetLetterGrade(player.Agility, attributeMeans["Agility"][player.Position]["mean"], attributeMeans["Agility"][player.Position]["stddev"])
	CarryingGrade := GetLetterGrade(player.Carrying, attributeMeans["Carrying"][player.Position]["mean"], attributeMeans["Carrying"][player.Position]["stddev"])
	CatchingGrade := GetLetterGrade(player.Catching, attributeMeans["Catching"][player.Position]["mean"], attributeMeans["Catching"][player.Position]["stddev"])
	RouteRunningGrade := GetLetterGrade(player.RouteRunning, attributeMeans["RouteRunning"][player.Position]["mean"], attributeMeans["RouteRunning"][player.Position]["stddev"])
	ZoneCoverageGrade := GetLetterGrade(player.ZoneCoverage, attributeMeans["ZoneCoverage"][player.Position]["mean"], attributeMeans["ZoneCoverage"][player.Position]["stddev"])
	ManCoverageGrade := GetLetterGrade(player.ManCoverage, attributeMeans["ManCoverage"][player.Position]["mean"], attributeMeans["ManCoverage"][player.Position]["stddev"])
	StrengthGrade := GetLetterGrade(player.Strength, attributeMeans["Strength"][player.Position]["mean"], attributeMeans["Strength"][player.Position]["stddev"])
	TackleGrade := GetLetterGrade(player.Tackle, attributeMeans["Tackle"][player.Position]["mean"], attributeMeans["Tackle"][player.Position]["stddev"])
	PassBlockGrade := GetLetterGrade(player.PassBlock, attributeMeans["PassBlock"][player.Position]["mean"], attributeMeans["PassBlock"][player.Position]["stddev"])
	RunBlockGrade := GetLetterGrade(player.RunBlock, attributeMeans["RunBlock"][player.Position]["mean"], attributeMeans["RunBlock"][player.Position]["stddev"])
	PassRushGrade := GetLetterGrade(player.PassRush, attributeMeans["PassRush"][player.Position]["mean"], attributeMeans["PassRush"][player.Position]["stddev"])
	RunDefenseGrade := GetLetterGrade(player.RunDefense, attributeMeans["RunDefense"][player.Position]["mean"], attributeMeans["RunDefense"][player.Position]["stddev"])
	ThrowPowerGrade := GetLetterGrade(player.ThrowPower, attributeMeans["ThrowPower"][player.Position]["mean"], attributeMeans["ThrowPower"][player.Position]["stddev"])
	ThrowAccuracyGrade := GetLetterGrade(player.ThrowAccuracy, attributeMeans["ThrowAccuracy"][player.Position]["mean"], attributeMeans["ThrowAccuracy"][player.Position]["stddev"])
	KickPowerGrade := GetLetterGrade(player.KickPower, attributeMeans["KickPower"][player.Position]["mean"], attributeMeans["KickPower"][player.Position]["stddev"])
	KickAccuracyGrade := GetLetterGrade(player.KickAccuracy, attributeMeans["KickAccuracy"][player.Position]["mean"], attributeMeans["KickAccuracy"][player.Position]["stddev"])
	PuntPowerGrade := GetLetterGrade(player.PuntPower, attributeMeans["PuntPower"][player.Position]["mean"], attributeMeans["PuntPower"][player.Position]["stddev"])
	PuntAccuracyGrade := GetLetterGrade(player.PuntAccuracy, attributeMeans["PuntAccuracy"][player.Position]["mean"], attributeMeans["PuntAccuracy"][player.Position]["stddev"])

	return models.CollegePlayerCSV{
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
	}
}

func MapNFLDrafteeToModel(player structs.NFLDraftee) models.NFLDrafteeCSV {

	attributeMeans := config.AttributeMeans()
	OverallGrade := GetNFLOverallGrade(player.Overall)
	StaminaGrade := GetLetterGrade(player.Stamina, attributeMeans["Stamina"][player.Position]["mean"], attributeMeans["Stamina"][player.Position]["stddev"])
	InjuryGrade := GetLetterGrade(player.Injury, attributeMeans["Injury"][player.Position]["mean"], attributeMeans["Injury"][player.Position]["stddev"])
	SpeedGrade := GetLetterGrade(player.Speed, attributeMeans["Speed"][player.Position]["mean"], attributeMeans["Speed"][player.Position]["stddev"])
	FootballIQGrade := GetLetterGrade(player.FootballIQ, attributeMeans["FootballIQ"][player.Position]["mean"], attributeMeans["FootballIQ"][player.Position]["stddev"])
	AgilityGrade := GetLetterGrade(player.Agility, attributeMeans["Agility"][player.Position]["mean"], attributeMeans["Agility"][player.Position]["stddev"])
	CarryingGrade := GetLetterGrade(player.Carrying, attributeMeans["Carrying"][player.Position]["mean"], attributeMeans["Carrying"][player.Position]["stddev"])
	CatchingGrade := GetLetterGrade(player.Catching, attributeMeans["Catching"][player.Position]["mean"], attributeMeans["Catching"][player.Position]["stddev"])
	RouteRunningGrade := GetLetterGrade(player.RouteRunning, attributeMeans["RouteRunning"][player.Position]["mean"], attributeMeans["RouteRunning"][player.Position]["stddev"])
	ZoneCoverageGrade := GetLetterGrade(player.ZoneCoverage, attributeMeans["ZoneCoverage"][player.Position]["mean"], attributeMeans["ZoneCoverage"][player.Position]["stddev"])
	ManCoverageGrade := GetLetterGrade(player.ManCoverage, attributeMeans["ManCoverage"][player.Position]["mean"], attributeMeans["ManCoverage"][player.Position]["stddev"])
	StrengthGrade := GetLetterGrade(player.Strength, attributeMeans["Strength"][player.Position]["mean"], attributeMeans["Strength"][player.Position]["stddev"])
	TackleGrade := GetLetterGrade(player.Tackle, attributeMeans["Tackle"][player.Position]["mean"], attributeMeans["Tackle"][player.Position]["stddev"])
	PassBlockGrade := GetLetterGrade(player.PassBlock, attributeMeans["PassBlock"][player.Position]["mean"], attributeMeans["PassBlock"][player.Position]["stddev"])
	RunBlockGrade := GetLetterGrade(player.RunBlock, attributeMeans["RunBlock"][player.Position]["mean"], attributeMeans["RunBlock"][player.Position]["stddev"])
	PassRushGrade := GetLetterGrade(player.PassRush, attributeMeans["PassRush"][player.Position]["mean"], attributeMeans["PassRush"][player.Position]["stddev"])
	RunDefenseGrade := GetLetterGrade(player.RunDefense, attributeMeans["RunDefense"][player.Position]["mean"], attributeMeans["RunDefense"][player.Position]["stddev"])
	ThrowPowerGrade := GetLetterGrade(player.ThrowPower, attributeMeans["ThrowPower"][player.Position]["mean"], attributeMeans["ThrowPower"][player.Position]["stddev"])
	ThrowAccuracyGrade := GetLetterGrade(player.ThrowAccuracy, attributeMeans["ThrowAccuracy"][player.Position]["mean"], attributeMeans["ThrowAccuracy"][player.Position]["stddev"])
	KickPowerGrade := GetLetterGrade(player.KickPower, attributeMeans["KickPower"][player.Position]["mean"], attributeMeans["KickPower"][player.Position]["stddev"])
	KickAccuracyGrade := GetLetterGrade(player.KickAccuracy, attributeMeans["KickAccuracy"][player.Position]["mean"], attributeMeans["KickAccuracy"][player.Position]["stddev"])
	PuntPowerGrade := GetLetterGrade(player.PuntPower, attributeMeans["PuntPower"][player.Position]["mean"], attributeMeans["PuntPower"][player.Position]["stddev"])
	PuntAccuracyGrade := GetLetterGrade(player.PuntAccuracy, attributeMeans["PuntAccuracy"][player.Position]["mean"], attributeMeans["PuntAccuracy"][player.Position]["stddev"])

	return models.NFLDrafteeCSV{
		PlayerID:           player.PlayerID,
		FirstName:          player.FirstName,
		LastName:           player.LastName,
		Position:           player.Position,
		Archetype:          player.Archetype,
		Age:                player.Age,
		Stars:              player.Stars,
		College:            player.College,
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
	}
}

func GetNFLOverallGrade(value int) string {
	if value > 63 {
		return "A+"
	} else if value > 61 {
		return "A"
	} else if value > 59 {
		return "A-"
	} else if value > 57 {
		return "B+"
	} else if value > 55 {
		return "B"
	} else if value > 53 {
		return "B-"
	} else if value > 51 {
		return "C+"
	} else if value > 49 {
		return "C"
	} else if value > 47 {
		return "C-"
	} else if value > 45 {
		return "D+"
	} else if value > 43 {
		return "D"
	} else if value > 41 {
		return "D-"
	}
	return "F"
}

func GetOverallGrade(value int) string {
	if value > 44 {
		return "A"
	} else if value > 34 {
		return "B"
	} else if value > 24 {
		return "C"
	} else if value > 14 {
		return "D"
	}
	return "F"
}

func GetLetterGrade(Attribute int, mean float32, stddev float32) string {
	if mean == 0 || stddev == 0 {
		return GetOverallGrade(Attribute)
	}

	val := float32(Attribute)
	dev := stddev * 2
	if val > mean+dev {
		return "A"
	}
	dev = stddev
	if val > mean+dev {
		return "B"
	}
	if val > mean {
		return "C"
	}
	dev = stddev * -1
	if val > mean+dev {
		return "D"
	}
	return "F"
}

func GetYearAndRedshirtStatus(year int, redshirt bool) (string, string) {
	status := ""
	if redshirt {
		status = "Redshirt"
	} else {
		status = ""
	}

	if year == 1 && !redshirt {
		return "Fr", status
	} else if year == 2 && redshirt {
		return "(Fr)", status
	} else if year == 2 && !redshirt {
		return "So", status
	} else if year == 3 && redshirt {
		return "(So)", status
	} else if year == 3 && !redshirt {
		return "Jr", status
	} else if year == 4 && redshirt {
		return "(Jr)", status
	} else if year == 4 && !redshirt {
		return "Sr", status
	} else if year == 5 && redshirt {
		return "(Sr)", status
	}
	return "Super Sr", status
}
