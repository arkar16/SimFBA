package structs

import "github.com/jinzhu/gorm"

type NFLRequest struct {
	gorm.Model
	Username            string
	NFLTeamID           uint
	NFLTeam             string
	NFLTeamAbbreviation string
	IsOwner             bool
	IsManager           bool
	IsCoach             bool
	IsAssistant         bool
	IsApproved          bool
}

func (r *NFLRequest) ApproveTeamRequest() {
	r.IsApproved = true
}

func (r *NFLRequest) RejectTeamRequest() {
	r.IsApproved = false
}
