package structs

// UpdateRecruitingBoardDTO - Data Transfer Object from UI to API
type UpdateRecruitingBoardDTO struct {
	Profile  RecruitingTeamProfile
	Recruits []RecruitPlayerProfile
	TeamID   int
}
