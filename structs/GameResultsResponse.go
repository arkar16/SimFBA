package structs

type GameResultsResponse struct {
	HomePlayers []GameResultsPlayer
	AwayPlayers []GameResultsPlayer
	PlayByPlays []PlayByPlayResponse
	Score       ScoreBoard
}

type GameResultsPlayer struct {
	ID                   uint
	FirstName            string
	LastName             string
	Position             string
	Archetype            string
	TeamAbbr             string
	League               string
	Year                 uint
	PassingYards         int
	PassAttempts         int
	PassCompletions      int
	PassingTDs           int
	Interceptions        int
	LongestPass          int
	Sacks                int
	RushAttempts         int
	RushingYards         int
	RushingTDs           int
	Fumbles              int
	LongestRush          int
	Targets              int
	Catches              int
	ReceivingYards       int
	ReceivingTDs         int
	LongestReception     int
	SoloTackles          float64
	AssistedTackles      float64
	TacklesForLoss       float64
	SacksMade            float64
	ForcedFumbles        int
	RecoveredFumbles     int
	PassDeflections      int
	InterceptionsCaught  int
	Safeties             int
	DefensiveTDs         int
	FGMade               int
	FGAttempts           int
	LongestFG            int
	ExtraPointsMade      int
	ExtraPointsAttempted int
	KickoffTouchbacks    int
	Punts                int
	PuntTouchbacks       int
	PuntsInside20        int
	KickReturns          int
	KickReturnTDs        int
	KickReturnYards      int
	PuntReturns          int
	PuntReturnTDs        int
	PuntReturnYards      int
	STSoloTackles        float64
	STAssistedTackles    float64
	PuntsBlocked         int
	FGBlocked            int
	Snaps                int
	Pancakes             int
	SacksAllowed         int
	PlayedGame           int
	StartedGame          int
	WasInjured           bool
	WeeksOfRecovery      uint
	InjuryType           string
}

type StreamResponse struct {
	GameID              uint
	HomeTeamID          uint
	HomeTeam            string
	HomeTeamCoach       string
	HomeTeamRank        uint
	HomeOffensiveScheme string
	HomeDefensiveScheme string
	AwayTeamID          uint
	AwayTeam            string
	AwayTeamCoach       string
	AwayTeamRank        uint
	AwayOffensiveScheme string
	AwayDefensiveScheme string
	GameTemp            float64
	GameCloud           string
	GameWind            string
	GamePrecip          string
	GameWindSpeed       float64
	Streams             []PlayByPlayResponse
}

type ScoreBoard struct {
	Q1Home  int
	Q2Home  int
	Q3Home  int
	Q4Home  int
	OT1Home int
	OT2Home int
	OT3Home int
	OT4Home int
	Q1Away  int
	Q2Away  int
	Q3Away  int
	Q4Away  int
	OT1Away int
	OT2Away int
	OT3Away int
	OT4Away int
}
