package structs

type BasePlayerGameSnaps struct {
	ID       uint
	SeasonID uint
	PlayerID uint
	GameID   uint
	WeekID   uint
	QBSnaps  uint8
	RBSnaps  uint8
	FBSnaps  uint8
	WRSnaps  uint8
	TESnaps  uint8
	OTSnaps  uint8
	OGSnaps  uint8
	CSnaps   uint8
	DTSnaps  uint8
	DESnaps  uint8
	OLBSnaps uint8
	ILBSnaps uint8
	CBSnaps  uint8
	FSSnaps  uint8
	SSSnaps  uint8
	KSnaps   uint8
	PSnaps   uint8
}

func (g *BasePlayerGameSnaps) MapSnapsToPosition(pos string, snaps int) {
	if pos == "QB" {
		g.QBSnaps += uint8(snaps)
	} else if pos == "RB" {
		g.RBSnaps += uint8(snaps)
	} else if pos == "FB" {
		g.FBSnaps += uint8(snaps)
	} else if pos == "WR" {
		g.WRSnaps += uint8(snaps)
	} else if pos == "TE" {
		g.TESnaps += uint8(snaps)
	} else if pos == "OT" {
		g.OTSnaps += uint8(snaps)
	} else if pos == "OG" {
		g.OGSnaps += uint8(snaps)
	} else if pos == "C" {
		g.CSnaps += uint8(snaps)
	} else if pos == "DE" {
		g.DESnaps += uint8(snaps)
	} else if pos == "DT" {
		g.DTSnaps += uint8(snaps)
	} else if pos == "OLB" {
		g.OLBSnaps += uint8(snaps)
	} else if pos == "ILB" {
		g.ILBSnaps += uint8(snaps)
	} else if pos == "CB" {
		g.CBSnaps += uint8(snaps)
	} else if pos == "FS" {
		g.FSSnaps += uint8(snaps)
	} else if pos == "SS" {
		g.SSSnaps += uint8(snaps)
	} else if pos == "P" {
		g.PSnaps += uint8(snaps)
	} else if pos == "K" {
		g.KSnaps += uint8(snaps)
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
}

func (s *BasePlayerSeasonSnaps) AddToSeason(g BasePlayerGameSnaps) {
	s.QBSnaps += uint16(g.QBSnaps)
	s.RBSnaps += uint16(g.RBSnaps)
	s.FBSnaps += uint16(g.FBSnaps)
	s.WRSnaps += uint16(g.WRSnaps)
	s.TESnaps += uint16(g.TESnaps)
	s.OTSnaps += uint16(g.OTSnaps)
	s.OGSnaps += uint16(g.OGSnaps)
	s.CSnaps += uint16(g.CSnaps)
	s.DTSnaps += uint16(g.DTSnaps)
	s.DESnaps += uint16(g.DESnaps)
	s.OLBSnaps += uint16(g.OLBSnaps)
	s.ILBSnaps += uint16(g.ILBSnaps)
	s.CBSnaps += uint16(g.CBSnaps)
	s.FSSnaps += uint16(g.FSSnaps)
	s.SSSnaps += uint16(g.SSSnaps)
	s.PSnaps += uint16(g.PSnaps)
	s.KSnaps += uint16(g.KSnaps)
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
