package structs

import "github.com/jinzhu/gorm"

// RecruitingTeamProfile - The profile for a team for recruiting
type RecruitingTeamProfile struct {
	gorm.Model
	TeamID                    int
	Team                      string
	TeamAbbreviation          string
	State                     string
	ScholarshipsAvailable     int
	WeeklyPoints              int
	SpentPoints               int
	TotalCommitments          int
	BaseEfficiencyScore       float64
	RecruitingEfficiencyScore float64
	PreviousOverallWinPer     float64
	PreviousConferenceWinPer  float64
	CurrentOverallWinPer      float64
	CurrentConferenceWinPer   float64
	ESPNScore                 float64
	RivalsScore               float64
	Rank247Score              float64
	CompositeScore            float64
	RecruitingClassRank       int
	Recruits                  []RecruitPlayerProfile `gorm:"foreignKey:ProfileID"`
	Affinities                []ProfileAffinity      `gorm:"foreignKey:ProfileID"`
}

type SaveProfile interface{}

func (r *RecruitingTeamProfile) AssignRES(res float64) {
	r.RecruitingEfficiencyScore = res
}

func (r *RecruitingTeamProfile) SubtractScholarshipsAvailable() {
	if r.ScholarshipsAvailable > 0 {
		r.ScholarshipsAvailable--
	}
}

func (r *RecruitingTeamProfile) ReallocateScholarship() {
	if r.ScholarshipsAvailable < 40 {
		r.ScholarshipsAvailable++
	}
}

func (r *RecruitingTeamProfile) ResetScholarshipCount() {
	r.ScholarshipsAvailable = 40
}

func (r *RecruitingTeamProfile) AllocateSpentPoints(points int) {
	r.SpentPoints = points
}

func (r *RecruitingTeamProfile) ResetWeeklyPoints(points int) {
	r.WeeklyPoints = points
}

func (r *RecruitingTeamProfile) AddRecruitsToProfile(croots []RecruitPlayerProfile) {
	r.Recruits = croots
}

func (r *RecruitingTeamProfile) AssignRivalsRank(score float64) {
	r.RivalsScore = score
}

func (r *RecruitingTeamProfile) Assign247Rank(score float64) {
	r.Rank247Score = score
}

func (r *RecruitingTeamProfile) AssignESPNRank(score float64) {
	r.ESPNScore = score
}

func (r *RecruitingTeamProfile) AssignCompositeRank(score float64) {
	r.CompositeScore = score
}
