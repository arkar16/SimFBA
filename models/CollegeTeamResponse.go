package models

import "github.com/CalebRose/SimFBA/structs"

type CollegeTeamResponse struct {
	ID int
	structs.BaseTeam
	ConferenceID int
	Conference   string
	DivisionID   int
	Division     string
	TeamStats    []structs.CollegeTeamStats
	SeasonStats  TeamSeasonStats
}

func (ctr *CollegeTeamResponse) MapSeasonalStats() {
	var ss TeamSeasonStats
	for _, stat := range ctr.TeamStats {
		ss.PassingYards = ss.PassingYards + stat.PassingYards
		ss.PassingAttempts = ss.PassingAttempts + stat.PassingAttempts
		ss.PassingCompletions = ss.PassingCompletions + stat.PassingCompletions
		ss.PassingTouchdowns = ss.PassingTouchdowns + stat.PassingTouchdowns
		ss.PassingInterceptions = ss.PassingInterceptions + stat.PassingInterceptions
		ss.QBSacks = ss.QBSacks + stat.QBSacks
		ss.LongestPass = ss.LongestPass + stat.LongestPass
		ss.RushAttempts = ss.RushAttempts + stat.RushAttempts
		ss.RushingTouchdowns = ss.RushingTouchdowns + stat.RushingTouchdowns
		ss.RushingYards = ss.RushingYards + stat.RushingYards
		ss.Fumbles = ss.Fumbles + stat.ReceivingFumbles + ss.RushingFumbles
		ss.ReceivingTargets = ss.ReceivingTargets + stat.ReceivingTargets
		ss.ReceivingCatches = ss.ReceivingCatches + stat.ReceivingCatches
		ss.ReceivingTouchdowns = ss.ReceivingTouchdowns + stat.ReceivingTouchdowns
		ss.ReceivingYards = ss.ReceivingYards + stat.ReceivingYards
		ss.AssistedTackles = ss.AssistedTackles + stat.AssistedTackles
		ss.SoloTackles = ss.SoloTackles + stat.SoloTackles
		ss.FumblesRecovered = ss.FumblesRecovered + stat.FumblesRecovered
		ss.DefensiveSacks = ss.DefensiveSacks + stat.DefensiveSacks
		ss.ForcedFumbles = ss.ForcedFumbles + stat.ForcedFumbles
		ss.TacklesForLoss = ss.TacklesForLoss + stat.TacklesForLoss
		ss.DefensiveInterceptions = ss.DefensiveInterceptions + stat.DefensiveInterceptions
		ss.Safeties = ss.Safeties + stat.Safeties
		ss.DefensiveTDs = ss.DefensiveTDs + stat.DefensiveTDs
		ss.FieldGoalsMade = ss.FieldGoalsMade + stat.FieldGoalsMade
		ss.FieldGoalsAttempted = ss.FieldGoalsAttempted + stat.FieldGoalsAttempted
		ss.LongestFieldGoal = ss.LongestFieldGoal + stat.LongestFieldGoal
		ss.ExtraPointsAttempted = ss.ExtraPointsAttempted + stat.ExtraPointsAttempted
		ss.ExtraPointsMade = ss.ExtraPointsMade + stat.ExtraPointsMade
		ss.KickoffTBs = ss.KickoffTBs + stat.KickoffTBs
		ss.Punts = ss.Punts + stat.Punts
		ss.PuntTBs = ss.PuntTBs + stat.PuntTBs
		ss.PuntsInside20 = ss.PuntsInside20 + stat.PuntsInside20
		ss.KickReturnYards = ss.KickReturnYards + stat.KickReturnYards
		ss.KickReturnTDs = ss.KickReturnTDs + stat.KickReturnTDs
		ss.PuntReturnYards = ss.PuntReturnYards + stat.PuntReturnYards
		ss.PuntReturnTDs = ss.PuntReturnTDs + stat.PuntReturnTDs
		ss.PointsScored = ss.PointsScored + stat.PointsScored
		ss.PointsAgainst = ss.PointsAgainst + stat.PointsAgainst
		ss.TwoPointTries = ss.TwoPointTries + stat.TwoPointTries
		ss.TwoPointSucceed = ss.TwoPointSucceed + stat.TwoPointSucceed
		ss.PassingCompletionsAllowed = ss.PassingCompletionsAllowed + stat.PassingCompletionsAllowed
		ss.PassingTDsAllowed = ss.PassingTDsAllowed + stat.PassingTDsAllowed
		ss.PassingYardsAllowed = ss.PassingYardsAllowed + stat.PassingYardsAllowed
		ss.RushingTDsAllowed = ss.RushingTDsAllowed + stat.RushingTDsAllowed
		ss.RushingYardsAllowed = ss.RushingYardsAllowed + stat.RushingYardsAllowed
	}
	ss.TotalOffensiveYards = ss.PassingYards + ss.RushingYards
	ss.TotalYardsAllowed = ss.PassingYardsAllowed + ss.RushingYardsAllowed
	ctr.SeasonStats = ss
}
