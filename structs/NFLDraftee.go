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

func (n *NFLDraftee) MapUnsignedPlayer(up UnsignedPlayer) {
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
}
