package models

import "github.com/CalebRose/SimFBA/structs"

type DashboardTeamProfileResponse struct {
	TeamProfile  structs.RecruitingTeamProfile
	TeamNeedsMap map[string]int
}

func (d *DashboardTeamProfileResponse) SetTeamProfile(profile structs.RecruitingTeamProfile) {
	d.TeamProfile = profile
}

func (d *DashboardTeamProfileResponse) SetTeamNeedsMap(obj map[string]int) {
	d.TeamNeedsMap = obj
}

type TeamBoardTeamProfileResponse struct {
	TeamProfile  SimTeamBoardResponse
	TeamNeedsMap map[string]int
}

func (t *TeamBoardTeamProfileResponse) SetTeamProfile(profile SimTeamBoardResponse) {
	t.TeamProfile = profile
}

func (t *TeamBoardTeamProfileResponse) SetTeamNeedsMap(obj map[string]int) {
	t.TeamNeedsMap = obj
}
