package dbprovider

import (
	"fmt"
	"log"
	"sync"

	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	c := config.Config()
	db, err = gorm.Open(c["db"], c["cs"])
	if err != nil {
		log.Fatal(err)
		return false
	}
	// AutoMigrations -- uncomment when needing to update a table
	// db.AutoMigrate(&structs.CollegePlayer{})
	// db.AutoMigrate(&structs.Recruit{})
	// db.AutoMigrate(&structs.Player{})
	// db.AutoMigrate(&structs.CollegePlayerStats{})
	// db.AutoMigrate(&structs.CollegeTeam{})
	// db.AutoMigrate(&structs.CollegeRival{})
	// db.AutoMigrate(&structs.CollegeConference{})
	// db.AutoMigrate(&structs.CollegeDivision{})
	// db.AutoMigrate(&structs.CollegeTeamStats{})
	// db.AutoMigrate(&structs.CollegeTeamDepthChart{})
	db.AutoMigrate(&structs.CollegeDepthChartPosition{})
	// db.AutoMigrate(&structs.CollegeGameplan{})
	// db.AutoMigrate(&structs.CollegeGame{})

	// db.AutoMigrate(&structs.CollegeStandings{})
	// db.AutoMigrate(&structs.RecruitingTeamProfile{})
	// db.AutoMigrate(&structs.RecruitPlayerProfile{})
	// db.AutoMigrate(&structs.RecruitPointAllocation{})
	// db.AutoMigrate(&structs.RecruitRegion{})
	// db.AutoMigrate(&structs.RecruitState{})
	// db.AutoMigrate(&structs.ProfileAffinity{})

	// db.AutoMigrate(&structs.CollegeWeek{})
	// db.AutoMigrate(&structs.CollegeSeason{})
	// db.AutoMigrate(&structs.NFLPlayer{})
	// db.AutoMigrate(&structs.AdminRecruitModifier{})
	// db.AutoMigrate(&structs.Affinity{})
	// db.AutoMigrate(&structs.TeamRequest{})
	// db.AutoMigrate(&structs.Timestamp{})
	// db.AutoMigrate(&structs.NFLDraftee{})
	return true
}

func (p *Provider) GetDB() *gorm.DB {
	return db
}
