package structs

type BasePlayerGameSnaps struct {
	ID       uint
	SeasonID uint
	PlayerID uint
	GameID   uint
	WeekID   uint
	QBSnaps  uint16
	RBSnaps  uint16
	FBSnaps  uint16
	WRSnaps  uint16
	TESnaps  uint16
	OTSnaps  uint16
	OGSnaps  uint16
	CSnaps   uint16
	DTSnaps  uint16
	DESnaps  uint16
	OLBSnaps uint16
	ILBSnaps uint16
	CBSnaps  uint16
	FSSnaps  uint16
	SSSnaps  uint16
	KSnaps   uint16
	PSnaps   uint16
	STSnaps  uint16
	KRSnaps  uint16
	PRSnaps  uint16
	KOSSnaps uint16
}

func (g *BasePlayerGameSnaps) MapSnapsToPosition(pos string, snaps int) {
	if pos == "QB" {
		g.QBSnaps += uint16(snaps)
	} else if pos == "RB" {
		g.RBSnaps += uint16(snaps)
	} else if pos == "FB" {
		g.FBSnaps += uint16(snaps)
	} else if pos == "WR" {
		g.WRSnaps += uint16(snaps)
	} else if pos == "TE" {
		g.TESnaps += uint16(snaps)
	} else if pos == "OT" {
		g.OTSnaps += uint16(snaps)
	} else if pos == "OG" {
		g.OGSnaps += uint16(snaps)
	} else if pos == "C" {
		g.CSnaps += uint16(snaps)
	} else if pos == "DE" {
		g.DESnaps += uint16(snaps)
	} else if pos == "DT" {
		g.DTSnaps += uint16(snaps)
	} else if pos == "OLB" {
		g.OLBSnaps += uint16(snaps)
	} else if pos == "ILB" {
		g.ILBSnaps += uint16(snaps)
	} else if pos == "CB" {
		g.CBSnaps += uint16(snaps)
	} else if pos == "FS" {
		g.FSSnaps += uint16(snaps)
	} else if pos == "SS" {
		g.SSSnaps += uint16(snaps)
	} else if pos == "P" {
		g.PSnaps += uint16(snaps)
	} else if pos == "K" {
		g.KSnaps += uint16(snaps)
	} else if pos == "ST" {
		g.STSnaps += uint16(snaps)
	} else if pos == "KR" {
		g.KRSnaps += uint16(snaps)
	} else if pos == "PR" {
		g.PRSnaps += uint16(snaps)
	} else if pos == "KOS" {
		g.KOSSnaps += uint16(snaps)
	}
}

type BasePlayerSeasonSnaps struct {
	ID       uint
	SeasonID uint
	PlayerID uint
	QBSnaps  uint16
	RBSnaps  uint16
	FBSnaps  uint16
	WRSnaps  uint16
	TESnaps  uint16
	OTSnaps  uint16
	OGSnaps  uint16
	CSnaps   uint16
	DTSnaps  uint16
	DESnaps  uint16
	OLBSnaps uint16
	ILBSnaps uint16
	CBSnaps  uint16
	FSSnaps  uint16
	SSSnaps  uint16
	KSnaps   uint16
	PSnaps   uint16
	STSnaps  uint16
	KRSnaps  uint16
	PRSnaps  uint16
	KOSSnaps uint16
}

func (s *BasePlayerSeasonSnaps) AddToSeason(g BasePlayerGameSnaps) {
	s.QBSnaps += g.QBSnaps
	s.RBSnaps += g.RBSnaps
	s.FBSnaps += g.FBSnaps
	s.WRSnaps += g.WRSnaps
	s.TESnaps += g.TESnaps
	s.OTSnaps += g.OTSnaps
	s.OGSnaps += g.OGSnaps
	s.CSnaps += g.CSnaps
	s.DTSnaps += g.DTSnaps
	s.DESnaps += g.DESnaps
	s.OLBSnaps += g.OLBSnaps
	s.ILBSnaps += g.ILBSnaps
	s.CBSnaps += g.CBSnaps
	s.FSSnaps += g.FSSnaps
	s.SSSnaps += g.SSSnaps
	s.PSnaps += g.PSnaps
	s.KSnaps += g.KSnaps
	s.STSnaps += g.STSnaps
	s.PRSnaps += g.PRSnaps
	s.KRSnaps += g.KRSnaps
	s.KOSSnaps += g.KOSSnaps
}

func (s *BasePlayerSeasonSnaps) GetTotalSnaps() int {
	return int(s.QBSnaps + s.RBSnaps + s.FBSnaps + s.WRSnaps +
		s.TESnaps + s.OTSnaps + s.OGSnaps + s.CSnaps + s.DTSnaps +
		s.DESnaps + s.OLBSnaps + s.ILBSnaps + s.CBSnaps + s.FSSnaps +
		s.SSSnaps + s.PSnaps + s.KSnaps + s.STSnaps + s.PRSnaps +
		s.KRSnaps + s.KOSSnaps)
}

type CollegePlayerSeasonSnaps struct {
	BasePlayerSeasonSnaps
}

type NFLPlayerSeasonSnaps struct {
	BasePlayerSeasonSnaps
}

type CollegePlayerGameSnaps struct {
	BasePlayerGameSnaps
}

type NFLPlayerGameSnaps struct {
	BasePlayerGameSnaps
}
