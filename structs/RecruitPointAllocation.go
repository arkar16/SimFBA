package structs

import "github.com/jinzhu/gorm"

type RecruitPointAllocation struct {
	gorm.Model
	RecruitID          int
	TeamProfileID      int
	RecruitProfileID   int
	WeekID             int
	Points             int
	AffinityOneApplied bool
	AffinityTwoApplied bool
	CaughtCheating     bool
}

func (rpa *RecruitPointAllocation) UpdatePointsSpent(points int) {
	rpa.Points = points
}

func (rpa *RecruitPointAllocation) ApplyAffinityOne() {
	rpa.AffinityOneApplied = true
}

func (rpa *RecruitPointAllocation) ApplyAffinityTwo() {
	rpa.AffinityTwoApplied = true
}

func (rpa *RecruitPointAllocation) ApplyCaughtCheating() {
	rpa.CaughtCheating = true
}
