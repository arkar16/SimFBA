package models

import "github.com/CalebRose/SimFBA/structs"

type CrootProfile struct {
	ID                        uint
	SeasonID                  int
	RecruitID                 int
	ProfileID                 int
	TotalPoints               float64
	CurrentWeeksPoints        float64
	SpendingCount             int
	RecruitingEfficiencyScore float64
	Scholarship               bool
	ScholarshipRevoked        bool
	AffinityOneEligible       bool
	AffinityTwoEligible       bool
	TeamAbbreviation          string
	RemovedFromBoard          bool
	IsSigned                  bool
	IsLocked                  bool
	CaughtCheating            bool
	Recruit                   Croot
}

func (cp *CrootProfile) Map(rp structs.RecruitPlayerProfile, c Croot) {
	cp.ID = rp.ID
	cp.SeasonID = rp.SeasonID
	cp.RecruitID = rp.RecruitID
	cp.ProfileID = rp.ProfileID
	cp.TotalPoints = rp.TotalPoints
	cp.CurrentWeeksPoints = rp.CurrentWeeksPoints
	cp.SpendingCount = rp.SpendingCount
	cp.RecruitingEfficiencyScore = rp.RecruitingEfficiencyScore
	cp.Scholarship = rp.Scholarship
	cp.ScholarshipRevoked = rp.ScholarshipRevoked
	cp.AffinityOneEligible = rp.AffinityOneEligible
	cp.AffinityTwoEligible = rp.AffinityTwoEligible
	cp.TeamAbbreviation = rp.TeamAbbreviation
	cp.RemovedFromBoard = rp.RemovedFromBoard
	cp.IsSigned = rp.IsSigned
	cp.IsLocked = rp.IsLocked
	cp.CaughtCheating = rp.CaughtCheating
	cp.Recruit = c
}
