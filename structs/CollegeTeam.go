package structs

import "github.com/jinzhu/gorm"

type CollegeTeam struct {
	gorm.Model
	BaseTeam
	ConferenceID      int
	Conference        string
	DivisionID        int
	Division          string
	Color1            string
	Color2            string
	Color3            string
	RecruitingProfile RecruitingTeamProfile
	Gameplan          CollegeGameplan
	TeamStats         []CollegeTeamStats    `gorm:"foreignKey:TeamID"`
	CollegeRivals     []CollegeRival        `gorm:"foreignKey:TeamID"`
	TeamRecord        CollegeTeamRecords    `gorm:"foreignKey:TeamID"`
	TeamGameplan      CollegeGameplan       `gorm:"foreignKey:TeamID"`
	TeamDepthChart    CollegeTeamDepthChart `gorm:"foreignKey:TeamID"`
	TeamStandings     []CollegeStandings    `gorm:"foreignKey:TeamID"`
}
