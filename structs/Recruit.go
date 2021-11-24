package structs

import "github.com/jinzhu/gorm"

type Recruit struct {
	gorm.Model
	PlayerID int
	TeamID   int
	BasePlayer
	HighSchool            string
	City                  string
	State                 string
	AffinityOne           string
	AffinityTwo           string
	IsSigned              bool
	CommitmentChoiceVal   float32
	CalculatedRanking     float32
	RecruitPlayerProfiles []RecruitPlayerProfile   `gorm:"foreignKey:RecruitI"`
	RecruitPoints         []RecruitPointAllocation `gorm:"foreignKey:RecruitID"`
}

func (r *Recruit) UpdatePlayerID(id int) {
	r.PlayerID = id
}

func (r *Recruit) UpdateTeamID(id int) {
	r.TeamID = id
}

func (r *Recruit) UpdateSigningStatus() {
	r.IsSigned = !r.IsSigned
}

func (r *Recruit) SetCommitmentChoiceVal(val float32) {
	r.CommitmentChoiceVal = val
}

func (r *Recruit) Map(createRecruitDTO CreateRecruitDTO) {
	r.FirstName = createRecruitDTO.FirstName
	r.LastName = createRecruitDTO.LastName
	r.Position = createRecruitDTO.Position
	r.Archetype = createRecruitDTO.Archetype
	r.Age = createRecruitDTO.Age
	r.Height = createRecruitDTO.Height
	r.Weight = createRecruitDTO.Weight
	r.Stars = createRecruitDTO.Stars
	r.Overall = createRecruitDTO.Overall
	r.Stamina = createRecruitDTO.Stamina
	r.Injury = createRecruitDTO.Injury
	r.FootballIQ = createRecruitDTO.FootballIQ
	r.WorkEthic = createRecruitDTO.WorkEthic
	r.Speed = createRecruitDTO.Speed
	r.Carrying = createRecruitDTO.Carrying
	r.Agility = createRecruitDTO.Agility
	r.Catching = createRecruitDTO.Catching
	r.RouteRunning = createRecruitDTO.RouteRunning
	r.ZoneCoverage = createRecruitDTO.ZoneCoverage
	r.ManCoverage = createRecruitDTO.ManCoverage
	r.Strength = createRecruitDTO.Strength
	r.Tackle = createRecruitDTO.Tackle
	r.PassBlock = createRecruitDTO.PassBlock
	r.RunBlock = createRecruitDTO.RunBlock
	r.PassRush = createRecruitDTO.PassRush
	r.RunDefense = createRecruitDTO.RunDefense
	r.ThrowPower = createRecruitDTO.ThrowPower
	r.ThrowAccuracy = createRecruitDTO.ThrowAccuracy
	r.KickAccuracy = createRecruitDTO.KickAccuracy
	r.KickPower = createRecruitDTO.KickPower
	r.PuntAccuracy = createRecruitDTO.PuntAccuracy
	r.PuntPower = createRecruitDTO.PuntPower
	r.Progression = createRecruitDTO.Progression
	r.PotentialGrade = createRecruitDTO.PotentialGrade
	r.HighSchool = createRecruitDTO.HighSchool
	r.City = createRecruitDTO.City
	r.State = createRecruitDTO.State
	r.AffinityOne = createRecruitDTO.AffinityOne
	r.AffinityTwo = createRecruitDTO.AffinityTwo
	r.IsSigned = false
	r.CommitmentChoiceVal = 0
	r.CalculatedRanking = 0
}

func (r *Recruit) AssignPlayerID(ID int) {
	r.PlayerID = ID
}
