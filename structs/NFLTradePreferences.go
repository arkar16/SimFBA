package structs

import "github.com/jinzhu/gorm"

type NFLTradePreferences struct {
	gorm.Model
	NFLTeamID             uint
	Quarterbacks          bool
	QuarterbackType       string
	Runningbacks          bool
	RunningbackType       string
	Fullbacks             bool
	FullbackType          string
	WideReceivers         bool
	WideReceiverType      string
	TightEnds             bool
	TightEndType          string
	OffensiveTackles      bool
	OffensiveGuards       bool
	Centers               bool
	OffensiveTackleType   string
	OffensiveGuardType    string
	CenterType            string
	DefensiveTackles      bool
	DefensiveTackleType   string
	DefensiveEnds         bool
	DefensiveEndType      string
	OutsideLinebackers    bool
	OutsideLinebackerType string
	InsideLinebackers     bool
	InsideLinebackerType  string
	Cornerbacks           bool
	CornerbackType        string
	FreeSafeties          bool
	FreeSafetyType        string
	StrongSafeties        bool
	StrongSafetyType      string
	Kickers               bool
	KickerType            string
	Punters               bool
	PunterType            string
}

type NFLTradePreferencesDTO struct {
	NFLTeamID             uint
	Quarterbacks          bool
	QuarterbackType       string
	Runningbacks          bool
	RunningbackType       string
	Fullbacks             bool
	FullbackType          string
	WideReceivers         bool
	WideReceiverType      string
	TightEnds             bool
	TightEndType          string
	OffensiveTackles      bool
	OffensiveGuards       bool
	Centers               bool
	OffensiveTackleType   string
	OffensiveGuardType    string
	CenterType            string
	DefensiveTackles      bool
	DefensiveTackleType   string
	DefensiveEnds         bool
	DefensiveEndType      string
	OutsideLinebackers    bool
	OutsideLinebackerType string
	InsideLinebackers     bool
	InsideLinebackerType  string
	Cornerbacks           bool
	CornerbackType        string
	FreeSafeties          bool
	FreeSafetyType        string
	StrongSafeties        bool
	StrongSafetyType      string
	Kickers               bool
	KickerType            string
	Punters               bool
	PunterType            string
	DraftPicks            bool
	DraftPickType         string
}

func (tp *NFLTradePreferences) UpdatePreferences(pref NFLTradePreferencesDTO) {
	tp.Quarterbacks = pref.Quarterbacks
	if tp.Quarterbacks {
		tp.QuarterbackType = pref.QuarterbackType
	}
	tp.Runningbacks = pref.Runningbacks
	if tp.Runningbacks {
		tp.RunningbackType = pref.RunningbackType
	}
	tp.Fullbacks = pref.Fullbacks
	if tp.Fullbacks {
		tp.FullbackType = pref.FullbackType
	}
	tp.TightEnds = pref.TightEnds
	if tp.TightEnds {
		tp.TightEndType = pref.TightEndType
	}
	tp.WideReceivers = pref.WideReceivers
	if tp.WideReceivers {
		tp.WideReceiverType = pref.WideReceiverType
	}
	tp.OffensiveTackles = pref.OffensiveTackles
	if tp.OffensiveTackles {
		tp.OffensiveTackleType = pref.OffensiveTackleType
	}
	tp.OffensiveGuards = pref.OffensiveGuards
	if tp.OffensiveGuards {
		tp.OffensiveGuardType = pref.OffensiveGuardType
	}
	tp.Centers = pref.Centers
	if tp.Centers {
		tp.CenterType = pref.CenterType
	}
	tp.DefensiveEnds = pref.DefensiveEnds
	if tp.DefensiveEnds {
		tp.DefensiveEndType = pref.DefensiveEndType
	}
	tp.DefensiveTackles = pref.DefensiveTackles
	if tp.DefensiveTackles {
		tp.DefensiveTackleType = pref.DefensiveTackleType
	}
	tp.OutsideLinebackers = pref.OutsideLinebackers
	if tp.OutsideLinebackers {
		tp.OutsideLinebackerType = pref.OutsideLinebackerType
	}
	tp.InsideLinebackers = pref.InsideLinebackers
	if tp.InsideLinebackers {
		tp.InsideLinebackerType = pref.InsideLinebackerType
	}
	tp.Cornerbacks = pref.Cornerbacks
	if tp.Cornerbacks {
		tp.CornerbackType = pref.CornerbackType
	}
	tp.FreeSafeties = pref.FreeSafeties
	if tp.FreeSafeties {
		tp.FreeSafetyType = pref.FreeSafetyType
	}
	tp.StrongSafeties = pref.StrongSafeties
	if tp.StrongSafeties {
		tp.StrongSafetyType = pref.StrongSafetyType
	}
	tp.Punters = pref.Punters
	if tp.Punters {
		tp.PunterType = pref.PunterType
	}
	tp.Kickers = pref.Kickers
	if tp.Kickers {
		tp.KickerType = pref.KickerType
	}
}
