package structs

import "github.com/jinzhu/gorm"

type NFLDraftee struct {
	gorm.Model
	BasePlayer
	PlayerID   int
	HighSchool string
	College    string
	City       string
	State      string
}

func (n *NFLDraftee) Map(cp CollegePlayer) {
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
