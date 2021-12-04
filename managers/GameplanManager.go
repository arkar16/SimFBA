package managers

import (
	"fmt"
	"log"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetGameplanByTeamID(teamID string) structs.CollegeGameplan {
	db := dbprovider.GetInstance().GetDB()

	var gamePlan structs.CollegeGameplan

	err := db.Where("team_id = ?", teamID).Find(&gamePlan)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Gameplan does not exist for team.")
	}
	return gamePlan
}

func GetDepthchartByTeamID(teamID string) structs.CollegeTeamDepthChart {
	db := dbprovider.GetInstance().GetDB()

	var depthChart structs.CollegeTeamDepthChart

	// Preload Depth Chart Positions
	err := db.Preload("DepthChartPlayers").Where("team_id = ?", teamID).Find(&depthChart)
	if err != nil {
		fmt.Println(err)
		panic("Depthchart does not exist for team.")
	}
	return depthChart
}

func GetDepthChartPositionPlayersByDepthchartID(depthChartID string) []structs.CollegeDepthChartPosition {
	db := dbprovider.GetInstance().GetDB()

	var positionPlayers []structs.CollegeDepthChartPosition

	err := db.Where("id = ?", depthChartID).Find(&positionPlayers)
	if err != nil {
		fmt.Println(err)
		panic("Depth Chart does not exist for this ID")
	}

	return positionPlayers
}
