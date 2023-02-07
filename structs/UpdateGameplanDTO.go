package structs

type UpdateGameplanDTO struct {
	GameplanID         string
	UpdatedGameplan    CollegeGameplan
	UpdatedNFLGameplan NFLGameplan
	Username           string
	TeamName           string
}
