package structs

type DashboardTeamProfileResponse struct {
	TeamProfile  RecruitingTeamProfile
	TeamNeedsMap map[string]int
}

func (d *DashboardTeamProfileResponse) SetTeamProfile(profile RecruitingTeamProfile) {
	d.TeamProfile = profile
}

func (d *DashboardTeamProfileResponse) SetTeamNeedsMap(obj map[string]int) {
	d.TeamNeedsMap = obj
}
