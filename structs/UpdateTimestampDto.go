package structs

type UpdateTimestampDto struct {
	MoveUpCollegeWeek   bool
	MoveUpCollegeSeason bool
	MoveUpNFLWeek       bool
	MoveUpNFLSeason     bool
	ThursdayGames       bool
	FridayGames         bool
	SaturdayMorning     bool
	SaturdayNoon        bool
	SaturdayEvening     bool
	SaturdayNight       bool
	RESSynced           bool
	RecruitingSynced    bool
	GMActionsCompleted  bool
}
