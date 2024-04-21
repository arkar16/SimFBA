package structs

type GameStatDTO struct {
	GameID            uint
	HomeTeam          TeamStatDTO
	AwayTeam          TeamStatDTO
	HomePlayers       []PlayerStatDTO
	AwayPlayers       []PlayerStatDTO
	HomeScore         int
	AwayScore         int
	Plays             []PlayByPlayDTO
	PlayerSnapTracker PlayerSnapTracker
}

type PlayerSnapTracker struct {
	PlayerSnapCounts map[int]map[string]int
}
