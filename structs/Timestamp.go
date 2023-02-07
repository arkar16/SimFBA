package structs

import "github.com/jinzhu/gorm"

type Timestamp struct {
	gorm.Model
	CollegeWeekID              int
	CollegeWeek                int
	CollegeSeasonID            int
	Season                     int
	NFLSeasonID                int
	NFLWeekID                  int
	NFLWeek                    int
	ThursdayGames              bool
	FridayGames                bool
	SaturdayMorning            bool
	SaturdayNoon               bool
	SaturdayEvening            bool
	SaturdayNight              bool
	NFLThursday                bool
	NFLSundayNoon              bool
	NFLSundayAfternoon         bool
	NFLSundayEvening           bool
	NFLMondayEvening           bool
	NFLTradingAllowed          bool
	NFLPreseason               bool
	RecruitingEfficiencySynced bool
	RecruitingSynced           bool
	GMActionsCompleted         bool
	IsOffSeason                bool
	IsNFLOffSeason             bool
	IsRecruitingLocked         bool
	IsFreeAgencyLocked         bool
	IsDraftTime                bool
	Y1Capspace                 float64
	Y2Capspace                 float64
	Y3Capspace                 float64
	Y4Capspace                 float64
	Y5Capspace                 float64
}

func (t *Timestamp) MoveUpWeekCollege() {
	t.CollegeWeekID++
	t.CollegeWeek++
}

func (t *Timestamp) MoveUpWeekNFL() {
	t.NFLWeekID++
	t.NFLWeek++
}

func (t *Timestamp) MoveUpSeason() {
	t.CollegeSeasonID++
	t.Season++
	t.CollegeWeek = 0
	t.NFLWeek = 0
	t.NFLSeasonID++
}

func (t *Timestamp) ToggleThursdayGames() {
	t.ThursdayGames = !t.ThursdayGames
}

func (t *Timestamp) ToggleFridayGames() {
	t.FridayGames = !t.FridayGames
}

func (t *Timestamp) ToggleSaturdayMorningGames() {
	t.SaturdayMorning = !t.SaturdayMorning
}

func (t *Timestamp) ToggleSaturdayNoonGames() {
	t.SaturdayNoon = !t.SaturdayNoon
}

func (t *Timestamp) ToggleSaturdayEveningGames() {
	t.SaturdayEvening = !t.SaturdayEvening
}

func (t *Timestamp) ToggleSaturdayNightGames() {
	t.SaturdayNight = !t.SaturdayNight
}

func (t *Timestamp) ToggleRES() {
	t.RecruitingEfficiencySynced = !t.RecruitingEfficiencySynced
}

func (t *Timestamp) ToggleRecruiting() {
	t.RecruitingSynced = !t.RecruitingSynced
	t.IsRecruitingLocked = false
}

func (t *Timestamp) ToggleGMActions() {
	t.GMActionsCompleted = !t.GMActionsCompleted
}

func (t *Timestamp) ToggleLockRecruiting() {
	t.IsRecruitingLocked = !t.IsRecruitingLocked
}

func (t *Timestamp) SyncToNextWeek() {
	t.MoveUpWeekCollege()
	if t.CollegeWeek > 20 {
		t.MoveUpSeason()
	}
	// Reset Toggles
	// t.ToggleThursdayGames()
	// t.ToggleFridayGames()
	// t.ToggleSaturdayMorningGames()
	// t.ToggleSaturdayNoonGames()
	// t.ToggleSaturdayEveningGames()
	// t.ToggleSaturdayNightGames()
	t.ToggleRES()
	t.ToggleRecruiting()
	// t.ToggleGMActions()

	// Migrate game results ?
}
