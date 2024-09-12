package models

import (
	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
	"github.com/jinzhu/gorm"
)

type NFLDraftee struct {
	gorm.Model
	structs.BasePlayer
	PlayerID           int
	HighSchool         string
	CollegeID          uint
	College            string
	DraftedTeamID      uint
	DraftedTeam        string
	DraftedRound       uint
	DraftPickID        uint
	DraftedPick        uint
	City               string
	State              string
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
	BoomOrBust         bool
	BoomOrBustStatus   string
}

func (n *NFLDraftee) Map(cp structs.CollegePlayer) {
	n.ID = cp.ID
	n.PlayerID = cp.PlayerID
	n.HighSchool = cp.HighSchool
	n.College = cp.TeamAbbr
	n.City = cp.City
	n.State = cp.State
	n.FirstName = cp.FirstName
	n.LastName = cp.LastName
	n.Position = cp.Position
	n.Archetype = cp.Archetype
	n.Height = cp.Height
	n.Weight = cp.Weight
	n.Age = cp.Age
	n.Stars = cp.Stars
	n.Overall = cp.Overall
	n.Stamina = cp.Stamina
	n.Injury = cp.Injury
	n.FootballIQ = cp.FootballIQ
	n.Speed = cp.Speed
	n.Carrying = cp.Carrying
	n.Agility = cp.Agility
	n.Catching = cp.Catching
	n.RouteRunning = cp.RouteRunning
	n.ZoneCoverage = cp.ZoneCoverage
	n.ManCoverage = cp.ManCoverage
	n.Strength = cp.Strength
	n.Tackle = cp.Tackle
	n.PassBlock = cp.PassBlock
	n.RunBlock = cp.RunBlock
	n.PassRush = cp.PassRush
	n.RunDefense = cp.RunDefense
	n.ThrowPower = cp.ThrowPower
	n.ThrowAccuracy = cp.ThrowAccuracy
	n.KickAccuracy = cp.KickAccuracy
	n.KickPower = cp.KickPower
	n.PuntAccuracy = cp.PuntAccuracy
	n.PuntPower = cp.PuntPower
	n.Progression = cp.Progression
	n.Discipline = cp.Discipline
	n.PotentialGrade = cp.PotentialGrade
	n.FreeAgency = cp.FreeAgency
	n.Personality = cp.Personality
	n.RecruitingBias = cp.RecruitingBias
	n.WorkEthic = cp.WorkEthic
	n.AcademicBias = cp.AcademicBias
}

func (n *NFLDraftee) GetLetterGrades() {
	attributeMeans := config.NFLAttributeMeans()
	rangeNum := 7
	OverallGrade := util.GetNFLOverallGrade(util.GenerateIntFromRange(n.Overall-rangeNum, n.Overall+rangeNum))
	StaminaGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Stamina-rangeNum, n.Stamina+rangeNum), attributeMeans["Stamina"][n.Position]["mean"], attributeMeans["Stamina"][n.Position]["stddev"], 5)
	InjuryGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Injury-rangeNum, n.Injury+rangeNum), attributeMeans["Injury"][n.Position]["mean"], attributeMeans["Injury"][n.Position]["stddev"], 5)
	SpeedGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Speed-rangeNum, n.Speed+rangeNum), attributeMeans["Speed"][n.Position]["mean"], attributeMeans["Speed"][n.Position]["stddev"], 5)
	FootballIQGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.FootballIQ-rangeNum, n.FootballIQ+rangeNum), attributeMeans["FootballIQ"][n.Position]["mean"], attributeMeans["FootballIQ"][n.Position]["stddev"], 5)
	AgilityGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Agility-rangeNum, n.Agility+rangeNum), attributeMeans["Agility"][n.Position]["mean"], attributeMeans["Agility"][n.Position]["stddev"], 5)
	CarryingGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Carrying-rangeNum, n.Carrying+rangeNum), attributeMeans["Carrying"][n.Position]["mean"], attributeMeans["Carrying"][n.Position]["stddev"], 5)
	CatchingGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Catching-rangeNum, n.Catching+rangeNum), attributeMeans["Catching"][n.Position]["mean"], attributeMeans["Catching"][n.Position]["stddev"], 5)
	RouteRunningGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.RouteRunning-rangeNum, n.RouteRunning+rangeNum), attributeMeans["RouteRunning"][n.Position]["mean"], attributeMeans["RouteRunning"][n.Position]["stddev"], 5)
	ZoneCoverageGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.ZoneCoverage-rangeNum, n.ZoneCoverage+rangeNum), attributeMeans["ZoneCoverage"][n.Position]["mean"], attributeMeans["ZoneCoverage"][n.Position]["stddev"], 5)
	ManCoverageGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.ManCoverage-rangeNum, n.ManCoverage+rangeNum), attributeMeans["ManCoverage"][n.Position]["mean"], attributeMeans["ManCoverage"][n.Position]["stddev"], 5)
	StrengthGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Strength-rangeNum, n.Strength+rangeNum), attributeMeans["Strength"][n.Position]["mean"], attributeMeans["Strength"][n.Position]["stddev"], 5)
	TackleGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.Tackle-rangeNum, n.Tackle+rangeNum), attributeMeans["Tackle"][n.Position]["mean"], attributeMeans["Tackle"][n.Position]["stddev"], 5)
	PassBlockGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.PassBlock-rangeNum, n.PassBlock+rangeNum), attributeMeans["PassBlock"][n.Position]["mean"], attributeMeans["PassBlock"][n.Position]["stddev"], 5)
	RunBlockGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.RunBlock-rangeNum, n.RunBlock+rangeNum), attributeMeans["RunBlock"][n.Position]["mean"], attributeMeans["RunBlock"][n.Position]["stddev"], 5)
	PassRushGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.PassRush-rangeNum, n.PassRush+rangeNum), attributeMeans["PassRush"][n.Position]["mean"], attributeMeans["PassRush"][n.Position]["stddev"], 5)
	RunDefenseGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.RunDefense-rangeNum, n.RunDefense+rangeNum), attributeMeans["RunDefense"][n.Position]["mean"], attributeMeans["RunDefense"][n.Position]["stddev"], 5)
	ThrowPowerGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.ThrowPower-rangeNum, n.ThrowPower+rangeNum), attributeMeans["ThrowPower"][n.Position]["mean"], attributeMeans["ThrowPower"][n.Position]["stddev"], 5)
	ThrowAccuracyGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.ThrowAccuracy-rangeNum, n.ThrowAccuracy+rangeNum), attributeMeans["ThrowAccuracy"][n.Position]["mean"], attributeMeans["ThrowAccuracy"][n.Position]["stddev"], 5)
	KickPowerGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.KickPower-rangeNum, n.KickPower+rangeNum), attributeMeans["KickPower"][n.Position]["mean"], attributeMeans["KickPower"][n.Position]["stddev"], 5)
	KickAccuracyGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.KickAccuracy-rangeNum, n.KickAccuracy+rangeNum), attributeMeans["KickAccuracy"][n.Position]["mean"], attributeMeans["KickAccuracy"][n.Position]["stddev"], 5)
	PuntPowerGrade := util.GetLetterGrade(util.GenerateIntFromRange(n.PuntPower-rangeNum, n.PuntPower+rangeNum), attributeMeans["PuntPower"][n.Position]["mean"], attributeMeans["PuntPower"][n.Position]["stddev"], 5)
	PuntAccuracyGrade := util.GetLetterGrade(n.PuntAccuracy, attributeMeans["PuntAccuracy"][n.Position]["mean"], attributeMeans["PuntAccuracy"][n.Position]["stddev"], 5)
	n.OverallGrade = OverallGrade
	n.StaminaGrade = StaminaGrade
	n.InjuryGrade = InjuryGrade
	n.FootballIQGrade = FootballIQGrade
	n.SpeedGrade = SpeedGrade
	n.AgilityGrade = AgilityGrade
	n.CarryingGrade = CarryingGrade
	n.CatchingGrade = CatchingGrade
	n.RouteRunningGrade = RouteRunningGrade
	n.ZoneCoverageGrade = ZoneCoverageGrade
	n.ManCoverageGrade = ManCoverageGrade
	n.StrengthGrade = StrengthGrade
	n.TackleGrade = TackleGrade
	n.PassBlockGrade = PassBlockGrade
	n.RunBlockGrade = RunBlockGrade
	n.PassRushGrade = PassRushGrade
	n.RunDefenseGrade = RunDefenseGrade
	n.ThrowPowerGrade = ThrowPowerGrade
	n.ThrowAccuracyGrade = ThrowAccuracyGrade
	n.KickPowerGrade = KickPowerGrade
	n.KickAccuracyGrade = KickAccuracyGrade
	n.PuntPowerGrade = PuntPowerGrade
	n.PuntAccuracyGrade = PuntAccuracyGrade
}

func (n *NFLDraftee) MapUnsignedPlayer(up structs.UnsignedPlayer) {
	attributeMeans := config.NFLAttributeMeans()
	n.ID = up.ID
	n.PlayerID = int(up.PlayerID)
	n.HighSchool = up.HighSchool
	n.College = up.TeamAbbr
	n.City = up.City
	n.State = up.State
	n.FirstName = up.FirstName
	n.LastName = up.LastName
	n.Position = up.Position
	n.Archetype = up.Archetype
	n.Height = up.Height
	n.Weight = up.Weight
	n.Age = up.Age
	n.Stars = up.Stars
	n.Overall = up.Overall
	n.Stamina = up.Stamina
	n.Injury = up.Injury
	n.FootballIQ = up.FootballIQ
	n.Speed = up.Speed
	n.Carrying = up.Carrying
	n.Agility = up.Agility
	n.Catching = up.Catching
	n.RouteRunning = up.RouteRunning
	n.ZoneCoverage = up.ZoneCoverage
	n.ManCoverage = up.ManCoverage
	n.Strength = up.Strength
	n.Tackle = up.Tackle
	n.PassBlock = up.PassBlock
	n.RunBlock = up.RunBlock
	n.PassRush = up.PassRush
	n.RunDefense = up.RunDefense
	n.ThrowPower = up.ThrowPower
	n.ThrowAccuracy = up.ThrowAccuracy
	n.KickAccuracy = up.KickAccuracy
	n.KickPower = up.KickPower
	n.PuntAccuracy = up.PuntAccuracy
	n.PuntPower = up.PuntPower
	n.Progression = up.Progression
	n.Discipline = up.Discipline
	n.PotentialGrade = up.PotentialGrade
	n.FreeAgency = up.FreeAgency
	n.Personality = up.Personality
	n.RecruitingBias = up.RecruitingBias
	n.WorkEthic = up.WorkEthic
	n.AcademicBias = up.AcademicBias
	rangeNum := 7
	OverallGrade := util.GetNFLOverallGrade(util.GenerateIntFromRange(up.Overall-rangeNum, up.Overall+rangeNum))
	StaminaGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Stamina-rangeNum, up.Stamina+rangeNum), attributeMeans["Stamina"][up.Position]["mean"], attributeMeans["Stamina"][up.Position]["stddev"], up.Year)
	InjuryGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Injury-rangeNum, up.Injury+rangeNum), attributeMeans["Injury"][up.Position]["mean"], attributeMeans["Injury"][up.Position]["stddev"], up.Year)
	SpeedGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Speed-rangeNum, up.Speed+rangeNum), attributeMeans["Speed"][up.Position]["mean"], attributeMeans["Speed"][up.Position]["stddev"], up.Year)
	FootballIQGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.FootballIQ-rangeNum, up.FootballIQ+rangeNum), attributeMeans["FootballIQ"][up.Position]["mean"], attributeMeans["FootballIQ"][up.Position]["stddev"], up.Year)
	AgilityGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Agility-rangeNum, up.Agility+rangeNum), attributeMeans["Agility"][up.Position]["mean"], attributeMeans["Agility"][up.Position]["stddev"], up.Year)
	CarryingGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Carrying-rangeNum, up.Carrying+rangeNum), attributeMeans["Carrying"][up.Position]["mean"], attributeMeans["Carrying"][up.Position]["stddev"], up.Year)
	CatchingGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Catching-rangeNum, up.Catching+rangeNum), attributeMeans["Catching"][up.Position]["mean"], attributeMeans["Catching"][up.Position]["stddev"], up.Year)
	RouteRunningGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.RouteRunning-rangeNum, up.RouteRunning+rangeNum), attributeMeans["RouteRunning"][up.Position]["mean"], attributeMeans["RouteRunning"][up.Position]["stddev"], up.Year)
	ZoneCoverageGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.ZoneCoverage-rangeNum, up.ZoneCoverage+rangeNum), attributeMeans["ZoneCoverage"][up.Position]["mean"], attributeMeans["ZoneCoverage"][up.Position]["stddev"], up.Year)
	ManCoverageGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.ManCoverage-rangeNum, up.ManCoverage+rangeNum), attributeMeans["ManCoverage"][up.Position]["mean"], attributeMeans["ManCoverage"][up.Position]["stddev"], up.Year)
	StrengthGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Strength-rangeNum, up.Strength+rangeNum), attributeMeans["Strength"][up.Position]["mean"], attributeMeans["Strength"][up.Position]["stddev"], up.Year)
	TackleGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.Tackle-rangeNum, up.Tackle+rangeNum), attributeMeans["Tackle"][up.Position]["mean"], attributeMeans["Tackle"][up.Position]["stddev"], up.Year)
	PassBlockGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.PassBlock-rangeNum, up.PassBlock+rangeNum), attributeMeans["PassBlock"][up.Position]["mean"], attributeMeans["PassBlock"][up.Position]["stddev"], up.Year)
	RunBlockGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.RunBlock-rangeNum, up.RunBlock+rangeNum), attributeMeans["RunBlock"][up.Position]["mean"], attributeMeans["RunBlock"][up.Position]["stddev"], up.Year)
	PassRushGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.PassRush-rangeNum, up.PassRush+rangeNum), attributeMeans["PassRush"][up.Position]["mean"], attributeMeans["PassRush"][up.Position]["stddev"], up.Year)
	RunDefenseGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.RunDefense-rangeNum, up.RunDefense+rangeNum), attributeMeans["RunDefense"][up.Position]["mean"], attributeMeans["RunDefense"][up.Position]["stddev"], up.Year)
	ThrowPowerGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.ThrowPower-rangeNum, up.ThrowPower+rangeNum), attributeMeans["ThrowPower"][up.Position]["mean"], attributeMeans["ThrowPower"][up.Position]["stddev"], up.Year)
	ThrowAccuracyGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.ThrowAccuracy-rangeNum, up.ThrowAccuracy+rangeNum), attributeMeans["ThrowAccuracy"][up.Position]["mean"], attributeMeans["ThrowAccuracy"][up.Position]["stddev"], up.Year)
	KickPowerGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.KickPower-rangeNum, up.KickPower+rangeNum), attributeMeans["KickPower"][up.Position]["mean"], attributeMeans["KickPower"][up.Position]["stddev"], up.Year)
	KickAccuracyGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.KickAccuracy-rangeNum, up.KickAccuracy+rangeNum), attributeMeans["KickAccuracy"][up.Position]["mean"], attributeMeans["KickAccuracy"][up.Position]["stddev"], up.Year)
	PuntPowerGrade := util.GetLetterGrade(util.GenerateIntFromRange(up.PuntAccuracy-rangeNum, up.PuntPower+rangeNum), attributeMeans["PuntPower"][up.Position]["mean"], attributeMeans["PuntPower"][up.Position]["stddev"], up.Year)
	PuntAccuracyGrade := util.GetLetterGrade(up.PuntAccuracy, attributeMeans["PuntAccuracy"][up.Position]["mean"], attributeMeans["PuntAccuracy"][up.Position]["stddev"], up.Year)
	n.OverallGrade = OverallGrade
	n.StaminaGrade = StaminaGrade
	n.InjuryGrade = InjuryGrade
	n.FootballIQGrade = FootballIQGrade
	n.SpeedGrade = SpeedGrade
	n.AgilityGrade = AgilityGrade
	n.CarryingGrade = CarryingGrade
	n.CatchingGrade = CatchingGrade
	n.RouteRunningGrade = RouteRunningGrade
	n.ZoneCoverageGrade = ZoneCoverageGrade
	n.ManCoverageGrade = ManCoverageGrade
	n.StrengthGrade = StrengthGrade
	n.TackleGrade = TackleGrade
	n.PassBlockGrade = PassBlockGrade
	n.RunBlockGrade = RunBlockGrade
	n.PassRushGrade = PassRushGrade
	n.RunDefenseGrade = RunDefenseGrade
	n.ThrowPowerGrade = ThrowPowerGrade
	n.ThrowAccuracyGrade = ThrowAccuracyGrade
	n.KickPowerGrade = KickPowerGrade
	n.KickAccuracyGrade = KickAccuracyGrade
	n.PuntPowerGrade = PuntPowerGrade
	n.PuntAccuracyGrade = PuntAccuracyGrade
}

func (n *NFLDraftee) MapNewOverallGrade(grade string) {
	n.OverallGrade = grade
}

func (n *NFLDraftee) AssignDraftedTeam(num uint, pickID uint, teamID uint, team string) {
	n.DraftedPick = num
	n.DraftPickID = pickID
	n.DraftedTeamID = teamID
	n.DraftedTeam = team
}

func (n *NFLDraftee) AssignBoomBustStatus(status string) {
	n.BoomOrBust = true
	n.BoomOrBustStatus = status
}

func (np *NFLDraftee) Progress(attr structs.CollegePlayerProgressions) {
	np.Agility = attr.Agility
	np.Speed = attr.Speed
	np.FootballIQ = attr.FootballIQ
	np.Carrying = attr.Carrying
	np.Catching = attr.Catching
	np.RouteRunning = attr.RouteRunning
	np.PassBlock = attr.PassBlock
	np.RunBlock = attr.RunBlock
	np.PassRush = attr.PassRush
	np.RunDefense = attr.RunDefense
	np.Tackle = attr.Tackle
	np.Strength = attr.Strength
	np.ManCoverage = attr.ManCoverage
	np.ZoneCoverage = attr.ZoneCoverage
	np.KickAccuracy = attr.KickAccuracy
	np.KickPower = attr.KickPower
	np.PuntAccuracy = attr.PuntAccuracy
	np.PuntPower = attr.PuntPower
	np.ThrowAccuracy = attr.ThrowAccuracy
	np.ThrowPower = attr.ThrowPower
}
