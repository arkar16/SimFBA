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
	CFBTeamMap         map[uint]CollegeTeam       // teamID to CollegeTeam
	ContendingMap      map[int]map[uint]bool      // week to teamID to bool (true if playing)
	FBSPlayedFCSMap    map[uint]bool              // teamID to bool (true if FBS team has played an FCS team)
	FCSPlayedFBSMap    map[uint]bool              // teamID to bool (true if FCS team has played an FCS team)
	TeamRivalries      map[uint][]uint            // teamID to slice of rival teamIDs
	TeamRivalryPriorities map[uint]map[uint]uint  // teamID to rivalTeamID to priority level
	HomeGamesMap       map[uint][]int             // teamID to slice of weeks (home games)
	AwayGamesMap       map[uint][]int             // teamID to slice of weeks (away games)
	ByeGamesMap        map[uint][]int             // bye game map
	OpenWeeksMap       map[uint][]int             // teamID to all open weeks
	OpponentHistoryMap map[uint]map[uint]bool     // teamID (a) to teamID (b) to bool to ensure no double scheduling
}

type ScheduleRequest struct {
	SeasonID    int
	Constraints SchedulingConstraints
	Teams       []CollegeTeam
	Rivalries   map[uint][]uint
}

type ScheduleResponse struct {
	ScheduledGamesByWeekMap map[int][]CollegeGame
}

type TeamBalance struct {
	TeamID            uint    `json:"team_id"`
	HomeGames         int     `json:"home_games"`
	AwayGames         int     `json:"away_games"`
	HomeRatio         float64 `json:"home_ratio"`          // Percentage of games played at home (0.0 to 1.0)
	LongestHomeStreak int     `json:"longest_home_streak"` // Most consecutive home games
	LongestAwayStreak int     `json:"longest_away_streak"` // Most consecutive away games
	IsBalanced        bool    `json:"is_balanced"`         // True if home ratio is between 0.4-0.6
}

