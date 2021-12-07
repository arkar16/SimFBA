package structs

// Request - A player request to sign for a team
type CreateRequestDTO struct {
	ID         int
	TeamID     int
	TeamName   string
	TeamAbbr   string
	Username   string
	Conference string
	IsApproved bool
}
