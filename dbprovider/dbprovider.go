package dbprovider

import (
	"fmt"
	"log"
	"sync"

	config "github.com/CalebRose/SimFBA/secrets"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Provider struct {
}

var db *gorm.DB
var once sync.Once
var instance *Provider

func GetInstance() *Provider {
	once.Do(func() {
		instance = &Provider{}
	})
	return instance
}

func (p *Provider) InitDatabase() bool {
	fmt.Println("Database initializing...")
	var err error
	c := config.Config() // c["cs"]
	db, err = gorm.Open(mysql.Open(c["cs"]), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return false
	}

	// AutoMigrations -- uncomment when needing to update a table
	// db.AutoMigrate(&structs.CollegePlayer{})
	// db.AutoMigrate(&structs.CollegePlayerSeasonStats{})
	// db.AutoMigrate(&structs.HistoricCollegePlayer{})
	// db.AutoMigrate(&structs.Player{})
	// db.AutoMigrate(&structs.CollegePlayerStats{})
	// db.AutoMigrate(&structs.UnsignedPlayer{})

	// db.AutoMigrate(&structs.CollegeTeam{})
	// db.AutoMigrate(&structs.CollegeRival{})
	// db.AutoMigrate(&structs.League{})
	// db.AutoMigrate(&structs.CollegeConference{})
	// db.AutoMigrate(&structs.CollegeDivision{})
	// db.AutoMigrate(&structs.CollegeTeamStats{})
	// db.AutoMigrate(&structs.CollegeTeamSeasonStats{})
	// db.AutoMigrate(&structs.CollegeTeamDepthChart{})
	// db.AutoMigrate(&structs.CollegeDepthChartPosition{})
	// db.AutoMigrate(&structs.CollegeGameplan{})
	// db.AutoMigrate(&structs.CollegeGame{})
	// db.AutoMigrate(&structs.CollegeCoach{})
	// db.AutoMigrate(&structs.CollegeStandings{})
	// db.AutoMigrate(&structs.Recruit{})
	// db.AutoMigrate(&structs.RecruitPlayerProfile{})
	// db.AutoMigrate(&structs.RecruitingTeamProfile{})
	// db.AutoMigrate(&structs.RecruitPointAllocation{})
	// db.AutoMigrate(&structs.RecruitRegion{})
	// db.AutoMigrate(&structs.RecruitState{})
	// db.AutoMigrate(&structs.ProfileAffinity{})
	// db.AutoMigrate(&structs.Stadium{})
	// db.AutoMigrate(&structs.CollegeWeek{})
	// db.AutoMigrate(&structs.CollegeSeason{})
	// db.AutoMigrate(&structs.NFLPlayer{})
	// db.AutoMigrate(&structs.NFLContract{})
	// db.AutoMigrate(&structs.NFLUser{})
	// db.AutoMigrate(&structs.NFLTeam{})
	// db.AutoMigrate(&structs.NFLGame{})
	// db.AutoMigrate(&structs.NFLGameplan{})
	// db.AutoMigrate(&structs.NFLStandings{})
	// db.AutoMigrate(&structs.NFLRequest{})
	// db.AutoMigrate(&structs.AdminRecruitModifier{})
	// db.AutoMigrate(&structs.Affinity{})
	// db.AutoMigrate(&structs.TeamRequest{})
	// db.AutoMigrate(&structs.Timestamp{})
	// db.AutoMigrate(&structs.NFLDraftee{})
	// db.AutoMigrate(&structs.NewsLog{})
	return true
}

func (p *Provider) GetDB() *gorm.DB {
	return db
}
