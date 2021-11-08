package structs

import "github.com/jinzhu/gorm"

type Timestamp struct {
	gorm.Model
	CollegeWeekID      int
	CollegeWeek        int
	CollegeSeasonID    int
	NFLSeasonID        int
	NFLWeekID          int
	NFLWeek            int
	Season             int
	ThursdayGames      bool
	FridayGames        bool
	SaturdayMorning    bool
	SaturdayNoon       bool
	SaturdayEvening    bool
	SaturdayNight      bool
	RecruitingSynced   bool
	GMActionsCompleted bool
	IsOffSeason        bool
}

func (t *Timestamp) MoveUpWeekCollege() {
	t.CollegeWeekID++
	t.CollegeWeek++
}

func (t *Timestamp) MoveUpSeason() {
	t.CollegeSeasonID++
	t.Season++
	t.CollegeWeek = 0
	t.CollegeWeekID++
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

func (t *Timestamp) ToggleRecruiting() {
	t.RecruitingSynced = !t.RecruitingSynced
}

func (t *Timestamp) ToggleGMActions() {
	t.GMActionsCompleted = !t.GMActionsCompleted
}

func (t *Timestamp) SyncToNextWeek() {
	t.MoveUpWeekCollege()
	t.MoveUpSeason()
	// Reset Toggles
	t.ToggleThursdayGames()
	t.ToggleFridayGames()
	t.ToggleSaturdayMorningGames()
	t.ToggleSaturdayNoonGames()
	t.ToggleSaturdayEveningGames()
	t.ToggleSaturdayNightGames()
	t.ToggleRecruiting()
	t.ToggleGMActions()
}
