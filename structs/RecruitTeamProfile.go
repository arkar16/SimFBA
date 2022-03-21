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
	RecruitingEfficiencyScore float32
	Recruits                  []RecruitPlayerProfile `gorm:"foreignKey:ProfileID"`
	Affinities                []ProfileAffinity      `gorm:"foreignKey:ProfileID"`
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
