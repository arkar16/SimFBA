package controller

import (
	"fmt"

	"github.com/CalebRose/SimFBA/managers"
)

func CronTest() {
	fmt.Println("PING!")
}

func FillAIBoardsViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron && !ts.IsOffSeason && !ts.CollegeSeasonOver {
		managers.FillAIRecruitingBoards()
	}
}

func SyncAIBoardsViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron && !ts.IsOffSeason && !ts.CollegeSeasonOver {
		managers.ResetAIBoardsForCompletedTeams()
		managers.AllocatePointsToAIBoards()
	}
}

func SyncRecruitingViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron && !ts.CollegeSeasonOver && !ts.CFBSpringGames && ts.CollegeWeek > 0 && ts.CollegeWeek < 21 {
		managers.SyncRecruiting(ts)
	}
}

func SyncFreeAgencyViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron && !ts.IsDraftTime {
		managers.SyncFreeAgencyOffers()
		managers.MoveUpInOffseasonFreeAgency()
	}
}

func SyncToNextWeekViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron {
		if !ts.IsOffSeason || !ts.IsNFLOffSeason {
			managers.MoveUpWeek()
		}
	}
}

func RunAISchemeAndDCViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron && !ts.IsOffSeason && !ts.CollegeSeasonOver {
		managers.DetermineAIGameplan()
	}
}

func RunAIGameplanViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron && !ts.IsOffSeason && !ts.CollegeSeasonOver {
		managers.SetAIGameplan()
	}
}

func RunTheGamesViaCron() {
	ts := managers.GetTimestamp()
	if ts.RunCron {
		if !ts.IsOffSeason && !ts.RunGames {
			managers.RunTheGames()
		}
	}
}

func ShowCFBThursdayViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.ThursdayGames {
		timeslot = "Thursday Night"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowNFLThursdayViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.NFLThursday {
		timeslot = "Thursday Night Football"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowCFBFridayViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.FridayGames {
		timeslot = "Friday Night"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowCFBSatMornViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.SaturdayMorning {
		timeslot = "Saturday Morning"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowCFBSatAftViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.SaturdayNoon {
		timeslot = "Saturday Afternoon"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowCFBSatEveViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.SaturdayEvening {
		timeslot = "Saturday Evening"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowCFBSatNitViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.ThursdayGames {
		timeslot = "Thursday Night"
	} else if !ts.NFLThursday {
		timeslot = "Thursday Night Football"
	} else if !ts.SaturdayMorning {
		timeslot = "Saturday Morning"
	} else if !ts.SaturdayNoon {
		timeslot = "Saturday Afternoon"
	} else if !ts.SaturdayEvening {
		timeslot = "Saturday Evening"
	} else if !ts.SaturdayNight {
		timeslot = "Saturday Night"
	} else if !ts.NFLSundayNoon {
		timeslot = "Sunday Noon"
	} else if !ts.NFLSundayAfternoon {
		timeslot = "Sunday Afternoon"
	} else if !ts.NFLSundayEvening {
		timeslot = "Sunday Night Football"
	} else if !ts.NFLMondayEvening {
		timeslot = "Monday Night Football"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowNFLSunNoonViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.NFLSundayNoon {
		timeslot = "Sunday Noon"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowNFLSunAftViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.NFLSundayAfternoon {
		timeslot = "Sunday Afternoon"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowNFLSunNitViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.NFLSundayEvening {
		timeslot = "Sunday Night Football"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}

func ShowNFLMonNitViaCron() {
	ts := managers.GetTimestamp()
	timeslot := ""
	if !ts.NFLMondayEvening {
		timeslot = "Monday Night Football"
	}
	if ts.RunCron {
		managers.SyncTimeslot(timeslot)
	}
}
