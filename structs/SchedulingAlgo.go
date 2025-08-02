package structs

type SchedulingConstraints struct {
	MaxGamesPerTeam         int
	MaxGamesPerWeek         int
	MaxHomeGames            int
	MaxAwayGames            int
	MaxConsecutiveAwayGames int
	PreventBackToBackByes   bool
	EnforceRivalries        bool
}

type SchedulingMaps struct {
	CFBTeamMap         map[uint]CollegeTeam   // teamID to CollegeTeam
	ContendingMap      map[int]map[uint]bool  // week to teamID to bool (true if playing)
	FBSPlayedFCSMap    map[uint]bool          // teamID to bool (true if FBS team has played an FCS team)
	FCSPlayedFBSMap    map[uint]bool          // teamID to bool (true if FCS team has played an FBS team)
	TeamRivalries      map[uint][]uint        // teamID to slice of rival teamIDs
	HomeGamesMap       map[uint][]int         // teamID to slice of weeks (home games)
	AwayGamesMap       map[uint][]int         // teamID to slice of weeks (away games)
	ByeGamesMap        map[uint][]int         // bye game map
	OpenWeeksMap       map[uint][]int         // teamID to all open weeks
	OpponentHistoryMap map[uint]map[uint]bool // teamID (a) to teamID (b) to bool to ensure no double scheduling
}

type ScheduleRequest struct {
	SeasonID    uint
	Constraints SchedulingConstraints
	Teams       []CollegeTeam
	Rivalries   map[uint][]uint
}

type ScheduledGame struct {
	Week       int
	HomeTeamID uint
	AwayTeamID uint
	IsRivalry  bool
}

type ScheduleResponse struct {
	ScheduledGamesByWeekMap map[int][]ScheduledGame
}
