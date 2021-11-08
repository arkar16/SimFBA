package structs

type UpdateRecruitPointsDto struct {
	RecruitPointsID   int
	RecruitID         int
	ProfileID         int
	WeekID            int
	AllocationID      int
	SpentPoints       int
	RewardScholarship bool
	RevokeScholarship bool
}
