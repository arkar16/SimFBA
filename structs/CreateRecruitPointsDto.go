package structs

type CreateRecruitProfileDto struct {
	PlayerID            int
	SeasonID            int
	RecruitID           int
	ProfileID           int
	Team                string
	AffinityOneEligible bool
	AffinityTwoEligible bool
	PlayerRecruit       Recruit
}
