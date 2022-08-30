package models

type CollegeGameResponse struct {
	Week                     int
	SeasonID                 int
	HomeTeamID               int
	HomeTeam                 string
	HomeTeamCoach            string
	HomeTeamWin              bool
	AwayTeamID               int
	AwayTeam                 string
	AwayTeamCoach            string
	AwayTeamWin              bool
	HomeTeamScore            int
	AwayTeamScore            int
	Stadium                  string
	City                     string
	State                    string
	IsNeutral                bool
	IsConference             bool
	IsDivisional             bool
	IsConferenceChampionship bool
	IsBowlGame               bool
	IsPlayoffGame            bool
	IsNationalChampionship   bool
	GameComplete             bool
	ShowGame                 bool
}
