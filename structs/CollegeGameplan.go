package structs

import "github.com/jinzhu/gorm"

type CollegeGameplan struct {
	gorm.Model
	TeamID                int
	OffensiveScheme       string
	OffRunToPassRatio     int
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
	BlitzRatio            int
	BlitzAggressiveness   string
	BlitzSafeties         bool
	BlitzCorners          bool
	LinebackerCoverage    string
	CornersCoverage       string
	SafetiesCoverage      string
}

func (cg *CollegeGameplan) UpdateGameplan(dto UpdateGameplanDTO) {
	// Validation is done in UI, so we're just passing data along in API
	cg.OffensiveScheme = dto.OffensiveScheme
	cg.OffRunToPassRatio = dto.OffRunToPassRatio
	cg.OffFormation1 = dto.OffFormation1
	cg.OffFormation2 = dto.OffFormation2
	cg.OffFormation3 = dto.OffFormation3
	cg.RunnerDistributionQB = dto.RunnerDistributionQB
	cg.RunnerDistributionBK1 = dto.RunnerDistributionBK1
	cg.RunnerDistributionBK2 = dto.RunnerDistributionBK2
	cg.RunnerDistributionBK3 = dto.RunnerDistributionBK3
	cg.RunOutsideLeft = dto.RunOutsideLeft
	cg.RunOutsideRight = dto.RunOutsideRight
	cg.RunInsideLeft = dto.RunInsideLeft
	cg.RunInsideRight = dto.RunInsideRight
	cg.RunPowerLeft = dto.RunPowerLeft
	cg.RunPowerRight = dto.RunPowerRight
	cg.RunDrawLeft = dto.RunDrawLeft
	cg.RunDrawRight = dto.RunDrawRight
	cg.PassQuick = dto.PassQuick
	cg.PassShort = dto.PassShort
	cg.PassLong = dto.PassLong
	cg.PassScreen = dto.PassScreen
	cg.PassPAShort = dto.PassPAShort
	cg.PassPALong = dto.PassPALong
	cg.TargetingWR1 = dto.TargetingWR1
	cg.TargetingWR2 = dto.TargetingWR2
	cg.TargetingWR3 = dto.TargetingWR3
	cg.TargetingWR4 = dto.TargetingWR4
	cg.TargetingWR5 = dto.TargetingWR5
	cg.TargetingTE1 = dto.TargetingTE1
	cg.TargetingTE2 = dto.TargetingTE2
	cg.TargetingTE3 = dto.TargetingTE3
	cg.TargetingRB1 = dto.TargetingRB1
	cg.TargetingRB2 = dto.TargetingRB2
	cg.TargetingRB3 = dto.TargetingRB3
	cg.DefensiveScheme = dto.DefensiveScheme
	cg.DefPersonnelOne = dto.DefPersonnelOne
	cg.DefPersonnelTwo = dto.DefPersonnelTwo
	cg.DefPersonnelThree = dto.DefPersonnelThree
	cg.BlitzRatio = dto.BlitzRatio
	cg.BlitzAggressiveness = dto.BlitzAggressiveness
	cg.BlitzSafeties = dto.BlitzSafeties
	cg.BlitzCorners = dto.BlitzCorners
	cg.LinebackerCoverage = dto.LinebackerCoverage
	cg.CornersCoverage = dto.CornersCoverage
	cg.SafetiesCoverage = dto.SafetiesCoverage
}
