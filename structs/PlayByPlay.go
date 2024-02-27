package structs

type PlayByPlay struct {
	ID                   uint
	GameID               uint
	WeekID               uint
	SeasonID             uint
	HomeTeamID           uint
	HomeTeamScore        uint8
	AwayTeamID           uint
	AwayTeamScore        uint8
	Quarter              uint8
	TimeRemaining        string // Will always be at a maximum of 5 characters: "00:00" format
	Down                 uint
	Distance             uint
	LineOfScrimmage      string // Could make into a uint? Keep track of which side of field? Maybe as a bool? Enum?
	PlayTypeID           uint8  // Enum using specific IDs as reference
	PlayNameID           uint8
	OffFormationID       uint8
	OffensivePoA         uint
	DefensiveFormationID uint8
	DefensiveTendency    uint
	BlitzNumber          uint
	LBCoverage           bool  // 0 == Man, 1 == Zone
	CBCoverage           bool  // 0 == Man, 1 == Zone
	SCoverage            bool  // 0 == Man, 1 == Zone
	QBPlayerID           uint  // Action player is the player invoking an action. In these instances, Kickers, Punters, QBs, RBs, and FBs. WRs in some instances
	BallCarrierID        uint  // The player that attains from the result of the acting player. RBs, WRs, DBs, LBs
	CatcherID            uint  // The player that attains from the result of the acting player. RBs, WRs, DBs, LBs
	ResultPlayerID       uint  // The player that attains from the result of the acting player. RBs, WRs, DBs, LBs
	Tackler1ID           uint  // ID of the player doing the tackle
	Tackler2ID           uint  // ID of the player doing the tackle
	ResultYards          int16 // Must be int due to negative values
	ResultID             uint8 // Reference ID for the different type of results from the play. Running out of bounds, dropped pass, catch, touchdown, FG, XP, tackled, int, fumble, etc.
	TurnoverPlayerID     uint  // ID of player turning over the ball
	PenaltyID            uint  //If greater than 0, there was a penalty on the play. Could also be used as reference for the type of penalty
	PenaltyPlayerID      uint  // The player from the resulting penalty
	PenaltyYards         int   // Resulting yards from the penalty
	InjuredPlayerID      uint  // If a player was injured, show the player
	InjuryType           uint  // Show injury type. The UINT stands for the enum
	InjurySeverity       uint  // In the event that there are different categories of an injury type, this column is here
	InjuryDuration       uint8 // Duration should by a number
}
