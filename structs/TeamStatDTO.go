package structs

type TeamStatDTO struct {
	Abbreviation     string
	Points           int
	TwoPointTries    int
	TwoPointSucceed  int
	PassAttempts     int
	PassCompletions  int
	PassYards        int
	PassTDS          int
	PassINTs         int
	LongestPass      int
	TimesSacked      int
	RushAttempts     int
	RushYards        int
	RushTDs          int
	LongestRush      int
	Fumbles          int
	RushYardsAllowed int
	PassYardsAllowed int
	PointsAllowed    int
	SoloTackles      int
	AssistedTackles  int
	TacklesForLoss   float64
	Sacks            float64
	ForcedFumbles    int
	RecoveredFumbles int
	INTs             int
	Safeties         int
	DefensiveTDs     int
	TurnoverYards    int
	FGAttempts       int
	FGMade           int
	LongestFG        int
	XPAttempts       int
	XPMade           int
	KickoffTBs       int
	PuntTBs          int
	PuntYards        int
	PuntInside20     int
	KickReturnYards  int
	KickReturnTDs    int
	PuntReturnYards  int
	PuntReturnTDs    int
	OffPenalties     int
	DefPenalties     int
	OffPenaltyYards  int
	DefPenaltyYards  int
}

func (t *TeamStatDTO) GetAbbreviation() string {
	return t.Abbreviation
}

func (t *TeamStatDTO) MapToBaseTeamStatsObject() BaseTeamStats {
	return BaseTeamStats{
		PointsScored:           t.Points,
		PointsAgainst:          t.PointsAllowed,
		TwoPointTries:          t.TwoPointTries,
		TwoPointSucceed:        t.TwoPointSucceed,
		PassingYards:           t.PassYards,
		PassingAttempts:        t.PassAttempts,
		PassingCompletions:     t.PassCompletions,
		PassingTouchdowns:      t.PassTDS,
		PassingInterceptions:   t.PassINTs,
		LongestPass:            t.LongestPass,
		QBSacks:                t.TimesSacked,
		RushAttempts:           t.RushAttempts,
		RushingYards:           t.RushYards,
		RushingTouchdowns:      t.RushTDs,
		RushingFumbles:         t.Fumbles,
		LongestRush:            t.LongestRush,
		PassingYardsAllowed:    t.PassYardsAllowed,
		RushingYardsAllowed:    t.RushYardsAllowed,
		TacklesForLoss:         float64(t.TacklesForLoss),
		SoloTackles:            t.SoloTackles,
		AssistedTackles:        t.AssistedTackles,
		DefensiveSacks:         t.Sacks,
		ForcedFumbles:          t.ForcedFumbles,
		FumblesRecovered:       t.RecoveredFumbles,
		DefensiveInterceptions: t.INTs,
		TurnoverYards:          t.TurnoverYards,
		DefensiveTDs:           t.DefensiveTDs,
		Safeties:               t.Safeties,
		ExtraPointsMade:        t.XPMade,
		ExtraPointsAttempted:   t.XPAttempts,
		FieldGoalsMade:         t.FGMade,
		FieldGoalsAttempted:    t.FGAttempts,
		LongestFieldGoal:       t.LongestFG,
		PuntYards:              t.PuntYards,
		PuntsInside20:          t.PuntInside20,
		KickoffTBs:             t.KickoffTBs,
		PuntTBs:                t.PuntTBs,
		KickReturnYards:        t.KickReturnYards,
		KickReturnTDs:          t.KickReturnTDs,
		PuntReturnYards:        t.PuntReturnYards,
		PuntReturnTDs:          t.PuntReturnTDs,
		OffensivePenalties:     t.OffPenalties,
		OffPenaltyYards:        t.OffPenaltyYards,
		DefensivePenalties:     t.DefPenalties,
		DefPenaltyYards:        t.DefPenaltyYards,
	}
}
