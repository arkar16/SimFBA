package structs

type PlayByPlay struct {
	ID               uint
	GameID           uint
	WeekID           uint
	SeasonID         uint
	HomeTeamID       uint
	HomeTeamScore    uint
	AwayTeamID       uint
	AwayTeamScore    uint
	Quarter          uint
	TimeRemaining    string // Will always be at a maximum of 5 characters: "00:00" format
	Down             uint
	Distance         uint
	LineOfScrimmage  string // Could make into a uint? Keep track of which side of field? Maybe as a bool? Enum?
	PlayTypeID       uint   // Enum using specific IDs as reference
	ActionPlayerID   uint   // Action player is the player invoking an action. In these instances, Kickers, Punters, QBs, RBs, and FBs. WRs in some instances
	ResultPlayerID   uint   // The player that attains from the result of the acting player. RBs, WRs, DBs, LBs
	ResultID         uint   // Reference ID for the different type of results from the play. Running out of bounds, dropped pass, catch, touchdown, FG, XP, tackled, int, fumble, etc.
	TacklerID        uint   // ID of the player doing the tackle
	ResultYards      int    // Must be int due to negative values
	TurnoverPlayerID uint   // ID of player turning over the ball
	PenaltyID        uint   //If greater than 0, there was a penalty on the play. Could also be used as reference for the type of penalty
	PenaltyPlayerID  uint   // The player from the resulting penalty
	PenaltyYards     int    // Resulting yards from the penalty
}

type PlayByPlayB struct {
	ID              uint
	GameID          uint
	WeekID          uint
	SeasonID        uint
	HomeTeamID      uint
	HomeTeamScore   uint
	AwayTeamID      uint
	AwayTeamScore   uint
	Quarter         uint
	TimeRemaining   string // Will always be at a maximum of 5 characters: "00:00" format
	Down            uint
	Distance        uint
	LineOfScrimmage string // Could make into a uint? Keep track of which side of field? Maybe as a bool? Enum?
	PlayTypeID      uint   // Enum using specific IDs as reference
	Result          string // Large varchar that should hold more than 255 characters
}
