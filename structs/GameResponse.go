package structs

type GameResponse struct {
	GameID                   uint
	WeekID                   int
	Week                     int
	SeasonID                 int
	HomeTeamID               int
	HomeTeam                 string
	HomeTeamCoach            string
	AwayTeamID               int
	AwayTeam                 string
	AwayTeamCoach            string
	TimeSlot                 string
	StadiumID                uint
	Stadium                  string
	City                     string
	State                    string
	Region                   string
	LowTemp                  float64
	HighTemp                 float64
	GameTemp                 float64
	Cloud                    string
	Precip                   string
	WindSpeed                float64
	WindCategory             string
	IsNeutral                bool
	IsDomed                  bool
	IsNightGame              bool
	IsConference             bool
	IsDivisional             bool
	IsConferenceChampionship bool
	IsBowlGame               bool
	IsPlayoffGame            bool
	IsNationalChampionship   bool
	IsSuperBowl              bool
	IsRivalryGame            bool
	GameTitle                string // For rivalry match-ups, bowl games, championships, and more
	League                   string
}
