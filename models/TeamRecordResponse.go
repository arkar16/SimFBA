package models

type TeamRecordResponse struct {
	OverallWins             int
	OverallLosses           int
	CurrentSeasonWins       int
	CurrentSeasonLosses     int
	BowlWins                int
	BowlLosses              int
	ConferenceChampionships []string
	DivisionTitles          []string
	NationalChampionships   []string
}
