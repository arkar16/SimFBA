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
