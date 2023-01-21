package structs

import "github.com/jinzhu/gorm"

type NFLGameplan struct {
	gorm.Model
	TeamID                uint
	OffensiveScheme       string
	OffRunToPassRatio     int
	OffFormation1Name     string
	OffFormation2Name     string
	OffFormation3Name     string
	OffFormation1         int
	OffFormation2         int
	OffFormation3         int
	RunnerDistributionQB  int
	RunnerDistributionBK1 int
	RunnerDistributionBK2 int
	RunnerDistributionBK3 int
	RunOutsideLeft        int
	RunOutsideRight       int
	RunInsideLeft         int
	RunInsideRight        int
	RunPowerLeft          int
	RunPowerRight         int
	RunDrawLeft           int
	RunDrawRight          int
	PassQuick             int
	PassShort             int
	PassLong              int
	PassScreen            int
	PassPAShort           int
	PassPALong            int
	TargetingWR1          int
	TargetingWR2          int
	TargetingWR3          int
	TargetingWR4          int
	TargetingWR5          int
	TargetingTE1          int
	TargetingTE2          int
	TargetingTE3          int
	TargetingRB1          int
	TargetingRB2          int
	TargetingRB3          int
	DefensiveScheme       string
	DefPersonnelOne       int
	DefPersonnelTwo       int
	DefPersonnelThree     int
	DefRunToPassRatio     int
	BlitzRatio            int
	BlitzAggressiveness   string
	BlitzSafeties         bool
	BlitzCorners          bool
	LinebackerCoverage    string
	CornersCoverage       string
	SafetiesCoverage      string
	PrimaryHB             int
	MaximumFGDistance     int
	GoFor4AndShort        int
	GoFor4AndLong         int
}

func (ng *NFLGameplan) UpdateGameplan(dto CollegeGameplan) {
	// Validation is done in UI, so we're just passing data along in API
	ng.OffensiveScheme = dto.OffensiveScheme
	ng.OffRunToPassRatio = dto.OffRunToPassRatio
	ng.OffFormation1 = dto.OffFormation1
	ng.OffFormation2 = dto.OffFormation2
	ng.OffFormation3 = dto.OffFormation3
	ng.RunnerDistributionQB = dto.RunnerDistributionQB
	ng.RunnerDistributionBK1 = dto.RunnerDistributionBK1
	ng.RunnerDistributionBK2 = dto.RunnerDistributionBK2
	ng.RunnerDistributionBK3 = dto.RunnerDistributionBK3
	ng.RunOutsideLeft = dto.RunOutsideLeft
	ng.RunOutsideRight = dto.RunOutsideRight
	ng.RunInsideLeft = dto.RunInsideLeft
	ng.RunInsideRight = dto.RunInsideRight
	ng.RunPowerLeft = dto.RunPowerLeft
	ng.RunPowerRight = dto.RunPowerRight
	ng.RunDrawLeft = dto.RunDrawLeft
	ng.RunDrawRight = dto.RunDrawRight
	ng.PassQuick = dto.PassQuick
	ng.PassShort = dto.PassShort
	ng.PassLong = dto.PassLong
	ng.PassScreen = dto.PassScreen
	ng.PassPAShort = dto.PassPAShort
	ng.PassPALong = dto.PassPALong
	ng.TargetingWR1 = dto.TargetingWR1
	ng.TargetingWR2 = dto.TargetingWR2
	ng.TargetingWR3 = dto.TargetingWR3
	ng.TargetingWR4 = dto.TargetingWR4
	ng.TargetingWR5 = dto.TargetingWR5
	ng.TargetingTE1 = dto.TargetingTE1
	ng.TargetingTE2 = dto.TargetingTE2
	ng.TargetingTE3 = dto.TargetingTE3
	ng.TargetingRB1 = dto.TargetingRB1
	ng.TargetingRB2 = dto.TargetingRB2
	ng.TargetingRB3 = dto.TargetingRB3
	ng.DefensiveScheme = dto.DefensiveScheme
	ng.DefRunToPassRatio = dto.DefRunToPassRatio
	ng.DefPersonnelOne = dto.DefPersonnelOne
	ng.DefPersonnelTwo = dto.DefPersonnelTwo
	ng.DefPersonnelThree = dto.DefPersonnelThree
	ng.BlitzRatio = dto.BlitzRatio
	ng.BlitzAggressiveness = dto.BlitzAggressiveness
	ng.BlitzSafeties = dto.BlitzSafeties
	ng.BlitzCorners = dto.BlitzCorners
	ng.LinebackerCoverage = dto.LinebackerCoverage
	ng.CornersCoverage = dto.CornersCoverage
	ng.SafetiesCoverage = dto.SafetiesCoverage
	ng.PrimaryHB = dto.PrimaryHB
	ng.MaximumFGDistance = dto.MaximumFGDistance
	ng.GoFor4AndLong = dto.GoFor4AndLong
	ng.GoFor4AndShort = dto.GoFor4AndShort
}
