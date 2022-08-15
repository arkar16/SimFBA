package managers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func GetGameplanByTeamID(teamID string) structs.CollegeGameplan {
	db := dbprovider.GetInstance().GetDB()

	var gamePlan structs.CollegeGameplan

	err := db.Where("id = ?", teamID).Find(&gamePlan).Error
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Gameplan does not exist for team.")
	}
	return gamePlan
}

func GetGameplanByGameplanID(gameplanID string) structs.CollegeGameplan {
	db := dbprovider.GetInstance().GetDB()

	var gamePlan structs.CollegeGameplan

	err := db.Where("id = ?", gameplanID).Find(&gamePlan).Error
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
	err := db.Preload("DepthChartPlayers.CollegePlayer").Where("team_id = ?", teamID).Find(&depthChart).Error
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Depthchart does not exist for team.")
	}
	return depthChart
}

func GetDepthChartPositionPlayersByDepthchartID(depthChartID string) []structs.CollegeDepthChartPosition {
	db := dbprovider.GetInstance().GetDB()

	var positionPlayers []structs.CollegeDepthChartPosition

	err := db.Where("depth_chart_id = ?", depthChartID).Find(&positionPlayers).Error
	if err != nil {
		fmt.Println(err)
		panic("Depth Chart does not exist for this ID")
	}

	return positionPlayers
}

func UpdateGameplan(updateGameplanDto structs.UpdateGameplanDTO) {
	db := dbprovider.GetInstance().GetDB()

	gameplanID := updateGameplanDto.GameplanID

	currentGameplan := GetGameplanByGameplanID(gameplanID)

	if currentGameplan.OffensiveScheme != updateGameplanDto.UpdatedGameplan.OffensiveScheme {
		ts := GetTimestamp()

		newsLog := structs.NewsLog{
			WeekID:      ts.CollegeWeekID,
			SeasonID:    ts.CollegeSeasonID,
			MessageType: "Gameplan",
			Message:     "Coach " + updateGameplanDto.Username + " has updated " + updateGameplanDto.TeamName + "'s offensive scheme from " + currentGameplan.OffensiveScheme + " to " + updateGameplanDto.UpdatedGameplan.OffensiveScheme,
		}

		db.Save(&newsLog)
	}

	currentGameplan.UpdateGameplan(updateGameplanDto.UpdatedGameplan)

	db.Save(&currentGameplan)
}

func UpdateDepthChart(updateDepthchartDTO structs.UpdateDepthChartDTO) {

	depthChartID := strconv.Itoa(updateDepthchartDTO.DepthChartID)
	depthChartPlayers := GetDepthChartPositionPlayersByDepthchartID(depthChartID)

	updatedPlayers := updateDepthchartDTO.UpdatedPlayerPositions
	updateCounter := 0

	fmt.Println(len(depthChartPlayers))
	fmt.Println(len(updatedPlayers))
	db := dbprovider.GetInstance().GetDB()

	for i := 0; i < len(depthChartPlayers); i++ {
		player := depthChartPlayers[i]

		updatedPlayer := GetPlayerFromDClist(player.ID, updatedPlayers)

		if player.ID == updatedPlayer.ID &&
			player.PlayerID == updatedPlayer.PlayerID &&
			player.OriginalPosition == updatedPlayer.OriginalPosition {
			continue
		}

		player.UpdateDepthChartPosition(updatedPlayer)

		updateCounter++

		if updateCounter == len(updatedPlayers) {
			break
		}
		db.Save(&player)
	}
}

func GetPlayerFromDClist(id uint, updatedPlayers []structs.CollegeDepthChartPosition) structs.CollegeDepthChartPosition {
	var player structs.CollegeDepthChartPosition

	for i := 0; i < len(updatedPlayers); i++ {
		if updatedPlayers[i].ID == id {
			player = updatedPlayers[i]
			break
		}
	}

	return player
}
