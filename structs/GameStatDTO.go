package structs

type GameStatDTO struct {
	HomeTeam    TeamStatDTO
	AwayTeam    TeamStatDTO
	HomePlayers []PlayerStatDTO
	AwayPlayers []PlayerStatDTO
	HomeScore   int
	AwayScore   int
}
