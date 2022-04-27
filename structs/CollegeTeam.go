package structs

import "github.com/jinzhu/gorm"

type CollegeTeam struct {
	gorm.Model
	BaseTeam
	ConferenceID      int
	Conference        string
	DivisionID        int
	Division          string
	ProgramPrestige   int
	AcademicPrestige  int
	Facilities        int
	RecruitingProfile RecruitingTeamProfile `gorm:"foreignKey:TeamID"`
	TeamStats         []CollegeTeamStats    `gorm:"foreignKey:TeamID"`
	CollegeRivals     []CollegeRival        `gorm:"foreignKey:TeamID"`
	TeamRecord        CollegeTeamRecords    `gorm:"foreignKey:TeamID"`
	TeamGameplan      CollegeGameplan       `gorm:"foreignKey:TeamID"`
	TeamDepthChart    CollegeTeamDepthChart `gorm:"foreignKey:TeamID"`
	TeamStandings     []CollegeStandings    `gorm:"foreignKey:TeamID"`
}
