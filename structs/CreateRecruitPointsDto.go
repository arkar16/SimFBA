package structs

type CreateRecruitPointsDto struct {
	SeasonID            int
	RecruitID           int
	ProfileID           int
	Team                string
	AffinityOneEligible bool
	AffinityTwoEligible bool
}
