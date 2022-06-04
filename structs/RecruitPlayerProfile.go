package structs

import (
	"github.com/jinzhu/gorm"
)

// RecruitPlayerProfile - Individual points profile for a Team's Recruiting Portfolio
type RecruitPlayerProfile struct {
	gorm.Model
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
	Recruit                   Recruit                  `gorm:"foreignKey:RecruitID"`
	RecruitPoints             []RecruitPointAllocation `gorm:"foreignKey:RecruitID"`
}

func (rp *RecruitPlayerProfile) AllocateCurrentWeekPoints(points float64) {
	rp.CurrentWeeksPoints = points
}

func (rp *RecruitPlayerProfile) AddCurrentWeekPointsToTotal(CurrentPoints float64) {
	// If user spends points on a recruit
	if CurrentPoints > 0 {
		rp.TotalPoints += CurrentPoints
		// rp.SpendingCount += 1
		// if rp.SpendingCount > 4 {
		// 	rp.SpendingCount = 0
		// 	if rp.AffinityOneEligible {
		// 		rp.TotalPoints += 25
		// 	}
		// 	if rp.AffinityTwoEligible {
		// 		rp.TotalPoints += 25
		// 	}
		// }
	} else {
		rp.TotalPoints = 0
		rp.CaughtCheating = true
		rp.SpendingCount = 0
	}
	rp.CurrentWeeksPoints = 0
}

func (rp *RecruitPlayerProfile) ToggleAffinityOne() {
	rp.AffinityOneEligible = !rp.AffinityOneEligible
}

func (rp *RecruitPlayerProfile) ToggleAffinityTwo() {
	rp.AffinityTwoEligible = !rp.AffinityTwoEligible
}

func (rp *RecruitPlayerProfile) ToggleRemoveFromBoard() {
	rp.RemovedFromBoard = !rp.RemovedFromBoard
}

func (rp *RecruitPlayerProfile) ToggleScholarship(rewardScholarship bool, revokeScholarship bool) {
	rp.Scholarship = rewardScholarship
	rp.ScholarshipRevoked = revokeScholarship
}

func (rp *RecruitPlayerProfile) SetWinningTeamAbbreviation(team string) {
	rp.TeamAbbreviation = team
}

func (rp *RecruitPlayerProfile) SignPlayer() {
	if rp.Scholarship {
		rp.IsSigned = true
		rp.IsLocked = true
	}
}

func (rp *RecruitPlayerProfile) LockPlayer() {
	rp.IsLocked = true
}

func (rp *RecruitPlayerProfile) AssignRES(res float64) {
	rp.RecruitingEfficiencyScore = res
}

// Sorting Funcs
type ByPoints []RecruitPlayerProfile

func (rp ByPoints) Len() int      { return len(rp) }
func (rp ByPoints) Swap(i, j int) { rp[i], rp[j] = rp[j], rp[i] }
func (rp ByPoints) Less(i, j int) bool {
	return rp[i].TotalPoints > rp[j].TotalPoints
}
