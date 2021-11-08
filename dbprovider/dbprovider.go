package dbprovider

import (
	"fmt"
	"log"
	"sync"

	"github.com/CalebRose/SimFBA/config"
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
	// AutoMigrations
	db.AutoMigrate(&structs.CollegePlayer{})
	db.AutoMigrate(&structs.CollegeTeam{})
	db.AutoMigrate(&structs.CollegeConference{})
	db.AutoMigrate(&structs.CollegeDivision{})
	db.AutoMigrate(&structs.RecruitingTeamProfile{})
	db.AutoMigrate(&structs.CollegeTeamStats{})

	db.AutoMigrate(&structs.Player{})

	return true
}

func (p *Provider) GetDB() *gorm.DB {
	return db
}
