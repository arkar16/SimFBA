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
}

func (rpa *RecruitPointAllocation) UpdatePointsSpent(points int) {
	rpa.Points = points
}
