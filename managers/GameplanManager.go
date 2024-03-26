package managers

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"sync"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
	"gorm.io/gorm"
)

func GetAllCollegeGameplans() []structs.CollegeGameplan {
	db := dbprovider.GetInstance().GetDB()

	gameplans := []structs.CollegeGameplan{}

	db.Find(&gameplans)

	return gameplans
}

func GetAllNFLGameplans() []structs.NFLGameplan {
	db := dbprovider.GetInstance().GetDB()

	gameplans := []structs.NFLGameplan{}

	db.Find(&gameplans)

	return gameplans
}

func UpdateGameplanPenalties() {
	db := dbprovider.GetInstance().GetDB()

	collegeGPs := GetAllCollegeGameplans()

	for _, gp := range collegeGPs {
		if gp.HasSchemePenalty {
			gp.LowerPenalty()
			db.Save(&gp)
		}
	}

	nflGPs := GetAllNFLGameplans()

	for _, gp := range nflGPs {
		if gp.HasSchemePenalty {
			gp.LowerPenalty()
			db.Save(&gp)
		}
	}
}

func GetGameplanDataByTeamID(teamID string) structs.GamePlanResponse {
	gamePlan := GetGameplanByTeamID(teamID)

	depthChart := GetDepthchartByTeamID(teamID)

	ts := GetTimestamp()
	seasonID := strconv.Itoa(ts.CollegeSeasonID)
	opponentID := ""
	games := GetCollegeGamesByTeamIdAndSeasonId(teamID, seasonID)
	for _, g := range games {
		if g.GameComplete {
			continue
		}
		homeTeamID := strconv.Itoa(g.HomeTeamID)
		awayTeamID := strconv.Itoa(g.AwayTeamID)
		if homeTeamID == teamID {
			opponentID = awayTeamID
		} else {
			opponentID = homeTeamID
		}
		break
	}

	opponentGameplan := GetGameplanByTeamID(opponentID)

	opponentRoster := GetDepthchartByTeamID(opponentID)

	oppDepthChartPlayers := []structs.CollegePlayer{}

	for _, p := range opponentRoster.DepthChartPlayers {
		if p.Position != "WR" && p.Position != "TE" && p.Position != "RB" && p.Position != "FB" {
			continue
		}
		oppDepthChartPlayers = append(oppDepthChartPlayers, p.CollegePlayer)
	}

	return structs.GamePlanResponse{
		CollegeGP:         gamePlan,
		CollegeDC:         depthChart,
		CollegeOpponent:   opponentGameplan,
		CollegeOppPlayers: oppDepthChartPlayers,
	}
}

func GetGameplanByTeamID(teamID string) structs.CollegeGameplan {
	db := dbprovider.GetInstance().GetDB()

	var gamePlan structs.CollegeGameplan

	err := db.Where("id = ?", teamID).Find(&gamePlan).Error
	if err != nil {
		fmt.Println(err)
		return structs.CollegeGameplan{}
	}

	return gamePlan
}

func GetGameplanTESTByTeamID(teamID string) structs.CollegeGameplanTEST {
	db := dbprovider.GetInstance().GetDB()

	var gamePlan structs.CollegeGameplanTEST

	err := db.Where("id = ?", teamID).Find(&gamePlan).Error
	if err != nil {
		fmt.Println(err)
		return structs.CollegeGameplanTEST{}
	}

	return gamePlan
}

func GetDCTESTByTeamID(teamID string) structs.CollegeTeamDepthChartTEST {
	db := dbprovider.GetInstance().GetDB()

	var dc structs.CollegeTeamDepthChartTEST

	err := db.Preload("DepthChartPlayers").Where("id = ?", teamID).Find(&dc).Error
	if err != nil {
		fmt.Println(err)
		return structs.CollegeTeamDepthChartTEST{}
	}

	return dc
}

func GetNFLGameplanDataByTeamID(teamID string) structs.GamePlanResponse {
	ts := GetTimestamp()
	seasonID := strconv.Itoa(ts.NFLSeasonID)
	gamePlan := GetNFLGameplanByTeamID(teamID)
	depthChart := GetNFLDepthchartByTeamID(teamID)
	nflGames := GetNFLGamesByTeamIdAndSeasonId(teamID, seasonID)
	opponentID := ""
	for _, g := range nflGames {
		if g.GameComplete {
			continue
		}
		homeTeamID := strconv.Itoa(g.HomeTeamID)
		awayTeamID := strconv.Itoa(g.AwayTeamID)
		if homeTeamID == teamID {
			opponentID = awayTeamID
		} else {
			opponentID = homeTeamID
		}
		break
	}

	opponentGameplan := GetNFLGameplanByTeamID(opponentID)

	opponentRoster := GetNFLDepthchartByTeamID(opponentID)

	oppDepthChartPlayers := []structs.NFLPlayer{}

	for _, p := range opponentRoster.DepthChartPlayers {
		if p.Position != "WR" && p.Position != "TE" && p.Position != "RB" && p.Position != "FB" {
			continue
		}
		oppDepthChartPlayers = append(oppDepthChartPlayers, p.NFLPlayer)
	}

	return structs.GamePlanResponse{
		NFLGP:         gamePlan,
		NFLDC:         depthChart,
		NFLOpponent:   opponentGameplan,
		NFLOppPlayers: oppDepthChartPlayers,
	}
}

func GetNFLGameplanByTeamID(teamID string) structs.NFLGameplan {
	db := dbprovider.GetInstance().GetDB()

	var gp structs.NFLGameplan

	err := db.Where("id = ?", teamID).Find(&gp).Error
	if err != nil {
		fmt.Println(err)
		return structs.NFLGameplan{}
	}

	return gp
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

func GetNFLDepthchartByTeamID(teamID string) structs.NFLDepthChart {
	db := dbprovider.GetInstance().GetDB()

	var depthChart structs.NFLDepthChart

	// Preload Depth Chart Positions
	err := db.Preload("DepthChartPlayers.NFLPlayer").Where("team_id = ?", teamID).Find(&depthChart).Error
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

func GetNFLDepthChartPositionsByDepthchartID(depthChartID string) []structs.NFLDepthChartPosition {
	db := dbprovider.GetInstance().GetDB()

	var positionPlayers []structs.NFLDepthChartPosition

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

	ts := GetTimestamp()

	schemePenalty := false

	if currentGameplan.OffensiveScheme != updateGameplanDto.UpdatedGameplan.OffensiveScheme && !ts.CFBSpringGames {

		if ts.CollegeWeek != 0 {
			currentGameplan.ApplySchemePenalty(true)
		}
		schemePenalty = true
	}

	if currentGameplan.DefensiveScheme != updateGameplanDto.UpdatedGameplan.DefensiveScheme && !ts.CFBSpringGames {

		if ts.CollegeWeek != 0 {
			currentGameplan.ApplySchemePenalty(false)
		}
		schemePenalty = true
	}

	if schemePenalty {

		newsLog := structs.NewsLog{
			TeamID:      updateGameplanDto.UpdatedGameplan.TeamID,
			WeekID:      ts.CollegeWeekID,
			Week:        ts.CollegeWeek,
			SeasonID:    ts.CollegeSeasonID,
			MessageType: "Gameplan",
			League:      "CFB",
			Message:     "Coach " + updateGameplanDto.Username + " has updated " + updateGameplanDto.TeamName + "'s offensive scheme from " + currentGameplan.OffensiveScheme + " to " + updateGameplanDto.UpdatedGameplan.OffensiveScheme,
		}

		db.Create(&newsLog)
	}

	currentGameplan.UpdateCollegeGameplan(updateGameplanDto.UpdatedGameplan)

	db.Save(&currentGameplan)
}

func UpdateNFLGameplan(updateGameplanDto structs.UpdateGameplanDTO) {
	db := dbprovider.GetInstance().GetDB()

	gameplanID := updateGameplanDto.GameplanID

	currentGameplan := GetNFLGameplanByTeamID(gameplanID)
	UpdatedGameplan := updateGameplanDto.UpdatedNFLGameplan

	schemeChange := false
	ts := GetTimestamp()
	if currentGameplan.OffensiveScheme != UpdatedGameplan.OffensiveScheme && !ts.IsNFLOffSeason && !ts.NFLPreseason {

		if ts.NFLWeek != 0 {
			currentGameplan.ApplySchemePenalty(true)
		}

		schemeChange = true

	}

	if currentGameplan.DefensiveScheme != UpdatedGameplan.DefensiveScheme && !ts.IsNFLOffSeason && !ts.NFLPreseason {

		if ts.NFLWeek != 0 {
			currentGameplan.ApplySchemePenalty(false)
		}
		schemeChange = true
	}

	if schemeChange {

		newsLog := structs.NewsLog{
			TeamID:      updateGameplanDto.UpdatedGameplan.TeamID,
			WeekID:      ts.NFLWeekID,
			Week:        ts.NFLWeek,
			SeasonID:    ts.NFLSeasonID,
			League:      "NFL",
			MessageType: "Gameplan",
			Message:     "Coach " + updateGameplanDto.Username + " has updated " + updateGameplanDto.TeamName + "'s offensive scheme from " + currentGameplan.OffensiveScheme + " to " + updateGameplanDto.UpdatedNFLGameplan.OffensiveScheme,
		}

		db.Create(&newsLog)
	}

	currentGameplan.UpdateNFLGameplan(UpdatedGameplan)

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

func UpdateNFLDepthChart(updateDepthchartDTO structs.UpdateNFLDepthChartDTO) {

	depthChartID := strconv.Itoa(updateDepthchartDTO.DepthChartID)
	depthChartPlayers := GetNFLDepthChartPositionsByDepthchartID(depthChartID)

	updatedPlayers := updateDepthchartDTO.UpdatedPlayerPositions
	updateCounter := 0

	db := dbprovider.GetInstance().GetDB()

	for i := 0; i < len(depthChartPlayers); i++ {
		player := depthChartPlayers[i]

		updatedPlayer := GetPlayerFromNFLDClist(player.ID, updatedPlayers)

		if player.ID == updatedPlayer.ID &&
			uint(player.PlayerID) == updatedPlayer.PlayerID &&
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

func GetPlayerFromNFLDClist(id uint, updatedPlayers []structs.NFLDepthChartPosition) structs.NFLDepthChartPosition {
	var player structs.NFLDepthChartPosition

	for i := 0; i < len(updatedPlayers); i++ {
		if updatedPlayers[i].ID == id {
			player = updatedPlayers[i]
			break
		}
	}

	return player
}

func CheckAllUserDepthChartsForInjuredPlayers() {
	teams := GetAllCollegeTeams()

	for _, team := range teams {
		if len(team.Coach) == 0 || team.Coach == "AI" {
			continue
		}
		teamID := strconv.Itoa(int(team.ID))

		depthchartPositions := GetDepthChartPositionPlayersByDepthchartID(teamID)

		for _, dcp := range depthchartPositions {
			player := dcp.CollegePlayer

			if player.IsInjured {
				fmt.Println(team.TeamName + ": INJURY AT " + dcp.Position + dcp.PositionLevel + ": " + player.FirstName + " " + player.LastName + " injured with " + strconv.Itoa(int(player.WeeksOfRecovery)) + " weeks of recovery.")
			}
		}
	}

}

func MassUpdateGameplanSchemes(off, def string) {
	db := dbprovider.GetInstance().GetDB()
	teams := GetAllCollegeTeams()
	offensiveSchemes := util.GetOffensiveDefaultSchemes()
	defensiveSchemes := util.GetDefensiveDefaultSchemes()
	for _, team := range teams {
		teamID := strconv.Itoa(int(team.ID))
		gp := GetGameplanByTeamID(teamID)
		gp.UpdateSchemes(off, def)
		// offe := GetTestOffensiveSchemesByTeamID(id)
		// defe := GetTestDefensiveSchemesByTeamID(id)
		// Map Default Scheme for offense & defense
		offFormations := offensiveSchemes[off]
		defFormations := defensiveSchemes[def][off]

		dto := structs.CollegeGameplan{
			TeamID: int(team.ID),
			BaseGameplan: structs.BaseGameplan{
				OffensiveScheme:    off,
				DefensiveScheme:    def,
				OffensiveFormation: offFormations,
				DefensiveFormation: defFormations,
				BlitzSafeties:      gp.BlitzSafeties,
				BlitzCorners:       gp.BlitzCorners,
				LinebackerCoverage: gp.LinebackerCoverage,
				MaximumFGDistance:  gp.MaximumFGDistance,
				GoFor4AndShort:     gp.GoFor4AndShort,
				GoFor4AndLong:      gp.GoFor4AndLong,
				DefaultOffense:     gp.DefaultOffense,
				DefaultDefense:     gp.DefaultDefense,
				PrimaryHB:          75,
				PitchFocus:         50,
				DiveFocus:          50,
			},
		}

		gp.UpdateCollegeGameplan(dto)

		// Autosort Depth Chart
		ReAlignCollegeDepthChart(db, teamID, gp)

		db.Save(&gp)
	}
}

// UpdateCollegeAIDepthCharts
func UpdateCollegeAIDepthCharts() {
	db := dbprovider.GetInstance().GetDB()
	teams := GetAllCollegeTeams()
	for _, team := range teams {
		// if len(team.Coach) > 0 && team.Coach != "AI" {
		// 	continue
		// }

		teamID := strconv.Itoa(int(team.ID))
		gp := GetGameplanByTeamID(teamID)
		ReAlignCollegeDepthChart(db, teamID, gp)
	}

	ts := GetTimestamp()
	ts.ToggleAIDepthCharts()
	db.Save(&ts)
}

func ReAlignCollegeDepthChart(db *gorm.DB, teamID string, gp structs.CollegeGameplan) {
	roster := GetAllCollegePlayersByTeamIdWithoutRedshirts(teamID)
	dcPositions := GetDepthChartPositionPlayersByDepthchartID(teamID)
	sort.Sort(structs.ByOverall(roster))
	positionMap := make(map[string][]structs.DepthChartPositionDTO)
	starterMap := make(map[uint]bool)
	backupMap := make(map[uint]bool)
	stuMap := make(map[uint]bool)
	offScheme := gp.OffensiveScheme
	defScheme := gp.DefensiveScheme
	isLT := true
	isLG := true
	isLE := true
	isLOLB := true

	goodFits := GetFitsByScheme(offScheme, false)
	badFits := GetFitsByScheme(defScheme, false)

	// Allocate the Position Map
	for _, cp := range roster {
		if cp.IsInjured || cp.IsRedshirting {
			continue
		}
		pos := cp.Position
		arch := cp.Archetype
		player := arch + " " + pos
		isGoodFit := CheckPlayerFits(player, goodFits)
		isBadFit := CheckPlayerFits(player, badFits)

		// Add to QB List
		if pos == "QB" || pos == "RB" || pos == "FB" || pos == "ATH" {
			score := 0
			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}
			if pos == "QB" {
				score += 100
			} else if pos == "ATH" && (arch == "Triple-Threat" || arch == "Field General") {
				score += 50
			}
			// score += ((cp.ThrowAccuracy + cp.ThrowPower) / 2)
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["QB"] = append(positionMap["QB"], dcpObj)
		}
		// Add to RB List
		if pos == "RB" || pos == "FB" || pos == "WR" || pos == "TE" || pos == "ATH" {
			score := 0
			if pos == "RB" {
				score += 100
			} else if pos == "FB" {
				score += 25
			} else if pos == "ATH" && (arch == "Wingback" || arch == "Soccer Player" || arch == "Triple-Threat") {
				score += 50
			}
			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += ((cp.Speed + cp.Agility + cp.Strength + cp.Carrying) / 4)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["RB"] = append(positionMap["RB"], dcpObj)
		}

		// Add to FB List
		if pos == "FB" || pos == "TE" || pos == "RB" || pos == "ATH" {
			score := 0
			if pos == "FB" {
				score += 100
			} else if pos == "ATH" && (arch == "Wingback") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}
			score += ((cp.Strength + cp.Carrying + cp.PassBlock + cp.RunBlock) / 4)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["FB"] = append(positionMap["FB"], dcpObj)
		}

		// Add to TE List
		if pos == "FB" || pos == "TE" || pos == "ATH" {
			score := 0
			if pos == "TE" {
				score += 100
			} else if pos == "ATH" && (arch == "Slotback") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) + int(float64(cp.RunBlock)*0.125) + int(float64(cp.PassBlock)*0.125) + int(float64(cp.Catching)*0.125) + int(float64(cp.Strength)*0.125)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["TE"] = append(positionMap["TE"], dcpObj)
		}
		// Add to WR List
		if pos == "WR" || pos == "TE" || pos == "RB" || pos == "ATH" {
			score := 0
			if pos == "WR" {
				score += 100
			} else if pos == "ATH" && (arch == "Wingback" || arch == "Slotback") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.4) +
				int(float64(cp.Speed)*0.12) +
				int(float64(cp.Agility)*0.12) +
				int(float64(cp.Catching)*0.12) +
				int(float64(cp.Strength)*0.12) +
				int(float64(cp.RouteRunning)*0.12)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["WR"] = append(positionMap["WR"], dcpObj)
		}
		// Add to LT and RT List
		if pos == "OT" || pos == "OG" || pos == "C" || pos == "ATH" {
			score := 0
			if pos == "OT" {
				score += 100
			} else if pos == "OG" {
				score += 25
			} else if pos == "ATH" && (arch == "Lineman") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.10) +
				int(float64(cp.RunBlock)*0.75) +
				int(float64(cp.PassBlock)*0.75) +
				int(float64(cp.Agility)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLT {
				positionMap["LT"] = append(positionMap["LT"], dcpObj)
			} else {
				positionMap["RT"] = append(positionMap["RT"], dcpObj)
			}
			isLT = !isLT
		}
		// Add to LG and RG List
		if pos == "OT" || pos == "OG" || pos == "C" || pos == "ATH" {
			score := 0
			if pos == "OG" {
				score += 100
			} else if pos == "C" {
				score += 25
			} else if pos == "ATH" && (arch == "Lineman") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 30
			} else if isBadFit && !isGoodFit {
				score -= 30
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.10) +
				int(float64(cp.RunBlock)*0.75) +
				int(float64(cp.PassBlock)*0.75) +
				int(float64(cp.Agility)*0.05)
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLG {
				positionMap["LG"] = append(positionMap["LG"], dcpObj)
			} else {
				positionMap["RG"] = append(positionMap["RG"], dcpObj)
			}
			isLG = !isLG
		}
		// Add to C List
		if pos == "OT" || pos == "OG" || pos == "C" || pos == "ATH" {
			score := 0
			if pos == "C" {
				score += 100
			} else if pos == "OG" {
				score += 25
			} else if pos == "ATH" && (arch == "Lineman") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.10) +
				int(float64(cp.RunBlock)*0.75) +
				int(float64(cp.PassBlock)*0.75) +
				int(float64(cp.Agility)*0.05)
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["C"] = append(positionMap["C"], dcpObj)
		}

		// Add to LE List
		if pos == "DE" || pos == "DT" || pos == "OLB" || pos == "ATH" {
			score := 0
			if pos == "DE" {
				score += 100
			} else if pos == "OLB" {
				score += 25
			} else if pos == "DT" {
				score += 3
			} else if pos == "ATH" && (arch == "Lineman" || arch == "Strongside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.05) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.PassRush)*0.75) +
				int(float64(cp.RunDefense)*0.75) +
				int(float64(cp.Agility)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLE {
				positionMap["LE"] = append(positionMap["LE"], dcpObj)
			} else {
				positionMap["RE"] = append(positionMap["RE"], dcpObj)
			}
			isLE = !isLE
		}

		// Add to DT list
		if pos == "DE" || pos == "DT" || pos == "OLB" || pos == "ATH" {
			score := 0
			if pos == "DT" {
				score += 100
			} else if pos == "DE" {
				score += 25
			} else if pos == "OLB" {
				score += 12
			} else if pos == "ATH" && (arch == "Lineman" || arch == "Strongside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.05) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.PassRush)*0.75) +
				int(float64(cp.RunDefense)*0.75) +
				int(float64(cp.Agility)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["DT"] = append(positionMap["DT"], dcpObj)
		}

		// Add to OLB list
		if pos == "OLB" || pos == "DE" || pos == "ILB" || pos == "SS" || pos == "FS" || pos == "ATH" {
			score := 0
			if pos == "OLB" {
				score += 100
			} else if pos == "DE" {
				score += 10
			} else if pos == "ILB" {
				score += 25
			} else if pos == "SS" {
				score += 3
			} else if pos == "ATH" && (arch == "Weakside" || arch == "Strongside" || arch == "Bandit") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.6) +
				int(float64(cp.Strength)*0.025) +
				int(float64(cp.Tackle)*0.055) +
				int(float64(cp.PassRush)*0.0755) +
				int(float64(cp.RunDefense)*0.0755) +
				int(float64(cp.ManCoverage)*0.075) +
				int(float64(cp.ZoneCoverage)*0.075) +
				int(float64(cp.Agility)*0.025)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLOLB {
				positionMap["LOLB"] = append(positionMap["LOLB"], dcpObj)
			} else {
				positionMap["ROLB"] = append(positionMap["ROLB"], dcpObj)
			}
			isLOLB = !isLOLB
		}

		// Add to ILB list
		if pos == "OLB" || pos == "DE" || pos == "ILB" || pos == "SS" || pos == "FS" || pos == "ATH" {
			score := 0
			if pos == "ILB" {
				score += 100
			} else if pos == "OLB" {
				score += 25
			} else if pos == "SS" {
				score += 8
			} else if pos == "DE" {
				score += 3
			} else if pos == "ATH" && (arch == "Weakside" || arch == "Bandit" || arch == "Field General") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.6) +
				int(float64(cp.Strength)*0.025) +
				int(float64(cp.Tackle)*0.055) +
				int(float64(cp.PassRush)*0.0755) +
				int(float64(cp.RunDefense)*0.0755) +
				int(float64(cp.ManCoverage)*0.075) +
				int(float64(cp.ZoneCoverage)*0.075) +
				int(float64(cp.Agility)*0.025)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["MLB"] = append(positionMap["MLB"], dcpObj)
		}

		// Add to CB List
		if pos == "CB" || pos == "FS" || pos == "SS" || pos == "ATH" {
			score := 0
			if pos == "CB" {
				score += 100
			} else if pos == "FS" {
				score += 10
			} else if pos == "SS" {
				score += 8
			} else if pos == "ATH" && (arch == "Triple-Threat" || arch == "Bandit" || arch == "Weakside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.Agility)*0.1) +
				int(float64(cp.Catching)*0.1) +
				int(float64(cp.ManCoverage)*0.01) +
				int(float64(cp.ZoneCoverage)*0.01) +
				int(float64(cp.Speed)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["CB"] = append(positionMap["CB"], dcpObj)
		}

		// Add to FS list
		if pos == "CB" || pos == "FS" || pos == "SS" || pos == "ATH" {
			score := 0
			if pos == "FS" {
				score += 100
			} else if pos == "CB" {
				score += 25
			} else if pos == "SS" {
				score += 12
			} else if pos == "ATH" && (arch == "Bandit" || arch == "Weakside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.Agility)*0.1) +
				int(float64(cp.Catching)*0.1) +
				int(float64(cp.ManCoverage)*0.01) +
				int(float64(cp.ZoneCoverage)*0.01) +
				int(float64(cp.Speed)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["FS"] = append(positionMap["FS"], dcpObj)
		}

		// Add to SS list
		if pos == "CB" || pos == "FS" || pos == "SS" || pos == "ATH" {
			score := 0
			if pos == "SS" {
				score += 100
			} else if pos == "FS" {
				score += 25
			} else if pos == "CB" {
				score += 12
			} else if pos == "ATH" && (arch == "Bandit" || arch == "Weakside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.Agility)*0.1) +
				int(float64(cp.Catching)*0.1) +
				int(float64(cp.ManCoverage)*0.01) +
				int(float64(cp.ZoneCoverage)*0.01) +
				int(float64(cp.Speed)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["SS"] = append(positionMap["SS"], dcpObj)
		}

		// Add to P list
		if pos == "K" || pos == "P" || pos == "QB" || pos == "ATH" {
			score := 0
			if pos == "P" {
				score += 100
			} else if pos == "K" {
				score += 25
			} else if pos == "ATH" && (arch == "Soccer Player") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += cp.PuntAccuracy + cp.PuntPower

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["P"] = append(positionMap["P"], dcpObj)
		}

		// Add to K list (Field Goal)
		if pos == "K" || pos == "P" || pos == "QB" || pos == "ATH" {
			score := 0
			if pos == "K" {
				score += 100
			} else if pos == "P" {
				score += 25
			} else if pos == "ATH" && (arch == "Soccer Player") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}
			score += cp.KickAccuracy + cp.KickPower
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["K"] = append(positionMap["K"], dcpObj)
		}

		// FG List
		if pos == "K" || pos == "P" || pos == "QB" || pos == "ATH" {
			score := 0
			if pos == "K" {
				score += 100
			} else if pos == "P" {
				score += 25
			} else if pos == "ATH" && (arch == "Soccer Player") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += cp.KickAccuracy + cp.KickPower

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["FG"] = append(positionMap["FG"], dcpObj)
		}

		// PR
		if pos == "WR" || pos == "RB" || pos == "FS" || pos == "SS" || pos == "CB" || pos == "ATH" {
			score := 0
			if pos == "ATH" && arch == "Return Specialist" {
				score += 50
			} else if pos == "WR" || pos == "RB" {
				score += 25
			}
			score += cp.Agility

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["PR"] = append(positionMap["PR"], dcpObj)
		}

		// KR
		if pos == "WR" || pos == "RB" || pos == "FS" || pos == "SS" || pos == "CB" || pos == "ATH" {
			score := 0
			if pos == "ATH" && arch == "Return Specialist" {
				score += 50
			} else if pos == "WR" || pos == "RB" {
				score += 25
			}
			score += cp.Speed

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["KR"] = append(positionMap["KR"], dcpObj)
		}

		// STU
		if pos == "FB" || pos == "TE" || pos == "ILB" || pos == "OLB" || pos == "RB" || pos == "CB" || pos == "FS" || pos == "SS" || pos == "WR" || pos == "ATH" {
			score := 0
			if cp.Year == 2 || cp.Year == 1 {
				score += 50
			} else if cp.Year == 3 && cp.IsRedshirt {
				score += 25
			}

			score += cp.Tackle
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["STU"] = append(positionMap["STU"], dcpObj)
		}
	}

	// Sort Each DC Position
	sort.Sort(structs.ByDCPosition(positionMap["QB"]))
	sort.Sort(structs.ByDCPosition(positionMap["RB"]))
	sort.Sort(structs.ByDCPosition(positionMap["FB"]))
	sort.Sort(structs.ByDCPosition(positionMap["WR"]))
	sort.Sort(structs.ByDCPosition(positionMap["TE"]))
	sort.Sort(structs.ByDCPosition(positionMap["LT"]))
	sort.Sort(structs.ByDCPosition(positionMap["RT"]))
	sort.Sort(structs.ByDCPosition(positionMap["LG"]))
	sort.Sort(structs.ByDCPosition(positionMap["RG"]))
	sort.Sort(structs.ByDCPosition(positionMap["C"]))
	sort.Sort(structs.ByDCPosition(positionMap["DT"]))
	sort.Sort(structs.ByDCPosition(positionMap["LE"]))
	sort.Sort(structs.ByDCPosition(positionMap["RE"]))
	sort.Sort(structs.ByDCPosition(positionMap["LOLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["ROLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["MLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["CB"]))
	sort.Sort(structs.ByDCPosition(positionMap["FS"]))
	sort.Sort(structs.ByDCPosition(positionMap["SS"]))
	sort.Sort(structs.ByDCPosition(positionMap["P"]))
	sort.Sort(structs.ByDCPosition(positionMap["K"]))
	sort.Sort(structs.ByDCPosition(positionMap["PR"]))
	sort.Sort(structs.ByDCPosition(positionMap["KR"]))
	sort.Sort(structs.ByDCPosition(positionMap["FG"]))
	sort.Sort(structs.ByDCPosition(positionMap["STU"]))

	for _, dcp := range dcPositions {
		positionList := positionMap[dcp.Position]
		for _, pos := range positionList {
			if starterMap[pos.CollegePlayer.ID] &&
				dcp.Position != "FG" &&
				dcp.Position != "PR" &&
				dcp.Position != "KR" {
				continue
			}
			if backupMap[pos.CollegePlayer.ID] && dcp.PositionLevel != "1" && dcp.Position != "STU" {
				continue
			}
			if dcp.Position == "STU" && stuMap[pos.CollegePlayer.ID] {
				continue
			}

			if dcp.Position == "WR" {
				runnerDistPostition := gp.RunnerDistributionWRPosition
				positionLabel := dcp.Position + "" + dcp.PositionLevel
				if runnerDistPostition == positionLabel {
					gp.AssignRunnerWRID(dcp.CollegePlayer.ID)
				}
			}

			if dcp.Position == "STU" {
				stuMap[pos.CollegePlayer.ID] = true
			} else if dcp.PositionLevel == "1" && !starterMap[pos.CollegePlayer.ID] {
				starterMap[pos.CollegePlayer.ID] = true
			} else {
				backupMap[pos.CollegePlayer.ID] = true
			}
			dto := structs.CollegeDepthChartPosition{
				DepthChartID:     dcp.DepthChartID,
				PlayerID:         int(pos.CollegePlayer.ID),
				FirstName:        pos.CollegePlayer.FirstName,
				LastName:         pos.CollegePlayer.LastName,
				OriginalPosition: pos.CollegePlayer.Position,
			}
			dto.AssignID(dcp.ID)
			dcp.UpdateDepthChartPosition(dto)
			db.Save(&dcp)
			break
		}
	}
}

// UpdateCollegeAIDepthCharts
func UpdateNFLAIDepthCharts() {
	db := dbprovider.GetInstance().GetDB()
	teams := GetAllNFLTeams()
	for _, team := range teams {
		if len(team.NFLOwnerName) > 0 || len(team.NFLCoachName) > 0 {
			continue
		}
		teamID := strconv.Itoa(int(team.ID))
		gp := GetNFLGameplanByTeamID(teamID)
		depthchartPositions := GetNFLDepthChartPositionsByDepthchartID(teamID)
		ReAlignNFLDepthChart(db, teamID, gp, depthchartPositions)
	}
}

func ReAlignNFLDepthChart(db *gorm.DB, teamID string, gp structs.NFLGameplan, dcPositions []structs.NFLDepthChartPosition) {
	roster := GetNFLPlayersWithContractsByTeamID(teamID)
	positionMap := make(map[string][]structs.DepthChartPositionDTO)
	starterMap := make(map[uint]bool)
	backupMap := make(map[uint]bool)

	offScheme := gp.OffensiveScheme
	defScheme := gp.DefensiveScheme
	isLT := true
	isLG := true
	isLE := true
	isLOLB := true

	// Allocate the Position Map
	for _, cp := range roster {
		if cp.IsInjured || cp.IsPracticeSquad {
			continue
		}
		pos := cp.Position
		arch := cp.Archetype

		if pos == "QB" || pos == "WR" || pos == "TE" || pos == "RB" || pos == "FB" || pos == "K" || pos == "P" {
			// Add to QB List
			score := 0
			if offScheme == "Pro" {
				if pos == "QB" {
					score += 25
					if arch == "Field General" {
						score += 20
					} else if arch == "Pocket" {
						score += 15
					} else if arch == "Balanced" {
						score += 10
					} else if arch == "Scrambler" {
						score += 5
					}
				}
				score += cp.ThrowAccuracy
			} else if offScheme == "Air Raid" {
				if pos == "QB" {
					score += 25
					if arch == "Field General" {
						score += 15
					} else if arch == "Pocket" {
						score += 20
					} else if arch == "Balanced" {
						score += 10
					} else if arch == "Scrambler" {
						score += 5
					}
				}
				score += cp.ThrowPower
			} else if offScheme == "Spread Option" {
				if pos == "QB" {
					score += 25
					if arch == "Field General" {
						score += 10
					} else if arch == "Pocket" {
						score += 5
					} else if arch == "Balanced" {
						score += 20
					} else if arch == "Scrambler" {
						score += 15
					}
				}
				score += cp.ThrowAccuracy
			} else if offScheme == "Double Wing Option" {
				if pos == "QB" {
					score += 20
					if arch == "Field General" {
						score += 10
					} else if arch == "Pocket" {
						score += 5
					} else if arch == "Balanced" {
						score += 15
					} else if arch == "Scrambler" {
						score += 20
					}
				}
				score += cp.Speed
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["QB"] = append(positionMap["QB"], dcpObj)
		}
		// Add to RB List
		if pos == "RB" || pos == "FB" || pos == "WR" || pos == "TE" {
			score := 0
			if offScheme == "Pro" {
				if pos == "RB" {
					score += 45
				} else if pos == "WR" {
					score += 25
				} else if pos == "FB" {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Air Raid" {
				if pos == "RB" {
					score += 25
					if arch == "Receiving" {
						score += 20
					} else {
						score += 10

					}
				} else {
					score += 20
				}
				score += cp.Catching
			} else if offScheme == "Spread Option" {
				if pos == "RB" {
					score += 45
				} else if pos == "WR" {
					score += 25
				} else if pos == "FB" {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Double Wing Option" {
				if pos == "RB" {
					score += 25
					if arch == "Balanced" {
						score += 20
					} else {
						score += 10
					}
				} else {
					score += 20
				}
				score += cp.Overall
			}

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["RB"] = append(positionMap["RB"], dcpObj)
		}

		// Add to FB List
		if pos == "FB" || pos == "TE" || pos == "RB" || pos == "ILB" || pos == "OLB" {
			score := 0
			if offScheme == "Pro" {
				if pos == "FB" {
					score += 25
					if arch == "Blocking" {
						score += 20
					} else if arch == "Receiving" {
						score += 15
					} else {
						score += 10
					}
				} else {
					score += 15
				}
				score += cp.RunBlock
			} else if offScheme == "Air Raid" {
				if pos == "FB" {
					score += 25
					if arch == "Receiving" {
						score += 20
					} else {
						score += 10
					}
				} else {
					score += 15
				}
				score += cp.Catching
			} else if offScheme == "Spread Option" {
				if pos == "FB" {
					score += 25
					if arch == "Receiving" {
						score += 20
					} else if arch == "Rushing" {
						score += 15
					} else {
						score += 10
					}
				} else {
					score += 15
				}
				score += cp.Catching
			} else if offScheme == "Double Wing Option" {
				if pos == "FB" {
					score += 25
					if arch == "Rushing" {
						score += 20
					} else {
						score += 15
					}
				} else {
					score += 10
				}
				score += cp.Strength
			}

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["FB"] = append(positionMap["FB"], dcpObj)
		}

		// Add to TE List
		if pos == "FB" || pos == "TE" || pos == "WR" {
			score := 0
			if offScheme == "Pro" {
				if pos == "TE" {
					score += 25
					if arch == "Receiving" {
						score += 20
					} else {
						score += 15
					}
				} else {
					score += 10
				}
				score += cp.Overall
			} else if offScheme == "Air Raid" {
				if pos == "TE" {
					score += 25
					if arch == "Vertical Threat" {
						score += 20
					} else if arch == "Receiving" {
						score += 15
					} else {
						score += 5
					}
				} else {
					score += 10
				}
				score += cp.Catching
			} else if offScheme == "Spread Option" {
				if pos == "TE" {
					score += 25
					if arch == "Receiving" {
						score += 15
					} else {
						score += 5
					}
				} else {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Double Wing Option" {
				if pos == "TE" {
					score += 25
					if arch == "Blocking" {
						score += 15
					} else {
						score += 5
					}
				} else {
					score += 15
				}
				score += cp.RunBlock
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["TE"] = append(positionMap["TE"], dcpObj)
		}
		// Add to WR List
		if pos == "WR" || pos == "TE" || pos == "RB" || pos == "CB" {
			score := 0
			if offScheme == "Pro" {
				if pos == "WR" {
					score += 10
					if arch == "Possession" {
						score += 5
					}
				}
				score += cp.Catching
			} else if offScheme == "Air Raid" {
				if pos == "WR" {
					score += 10
					if arch == "Speed" {
						score += 5
					}
				} else if pos == "TE" && arch == "Vertical Threat" {
					score += 10
				}
				score += cp.Speed
			} else if offScheme == "Spread Option" {
				if pos == "WR" {
					score += 10
					if arch == "Route Running" {
						score += 5
					}
				}
				score += cp.RouteRunning
			} else if offScheme == "Double Wing Option" {
				if pos == "WR" {
					score += 10
					if arch == "Red Zone Threat" {
						score += 5
					}
				}
				score += cp.RunBlock
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["WR"] = append(positionMap["WR"], dcpObj)
		}
		// Add to LT and RT List
		if pos == "OT" || pos == "OG" || pos == "C" {
			score := 0
			if offScheme == "Pro" {
				if pos == "OT" {
					score += 20
				} else if pos == "OG" {
					score += 5
				}
				score += cp.Overall
			} else if offScheme == "Air Raid" {
				if (pos == "OT" || pos == "OG") && arch == "Pass Blocking" {
					score += 20
				} else if pos == "C" && arch == "Pass Blocking" {
					score += 5
				} else if (pos == "OT" || pos == "OG") && arch != "Pass Blocking" {
					score += 2
				}
				score += cp.PassBlock
			} else if offScheme == "Spread Option" {
				if pos == "OT" {
					score += 20
				} else if pos == "OG" {
					score += 5
				}
				score += cp.Overall
			} else if offScheme == "Double Wing Option" {
				if (pos == "OT" || pos == "OG") && arch == "Run Blocking" {
					score += 20
				} else if pos == "C" && arch == "Run Blocking" {
					score += 5
				} else if (pos == "OT" || pos == "OG") && arch != "Run Blocking" {
					score += 2
				}
				score += cp.RunBlock
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			if isLT {
				positionMap["LT"] = append(positionMap["LT"], dcpObj)
			} else {
				positionMap["RT"] = append(positionMap["RT"], dcpObj)
			}
			isLT = !isLT
		}
		// Add to LG and RG List
		if pos == "OT" || pos == "OG" || pos == "C" {
			score := 0
			if offScheme == "Pro" {
				if pos == "OG" {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Air Raid" {
				if pos == "OG" {
					score += 10
				}
				if arch == "Pass Blocking" {
					score += 10
				}
				score += cp.PassBlock
			} else if offScheme == "Spread Option" {
				if pos == "OG" {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Double Wing Option" {
				if pos == "OG" {
					score += 10
				}
				if arch == "Run Blocking" {
					score += 10
				}
				score += cp.RunBlock
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			if isLG {
				positionMap["LG"] = append(positionMap["LG"], dcpObj)
			} else {
				positionMap["RG"] = append(positionMap["RG"], dcpObj)
			}
			isLG = !isLG
		}
		// Add to C List
		if pos == "OT" || pos == "OG" || pos == "C" {
			score := 0
			if offScheme == "Pro" {
				if pos == "C" {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Air Raid" {
				if pos == "C" {
					score += 15
					if arch == "Pass Blocking" {
						score += 10
					} else {
						score += 5
					}

				}
				score += cp.PassBlock
			} else if offScheme == "Spread Option" {
				if pos == "C" {
					score += 15
				}
				score += cp.Overall
			} else if offScheme == "Double Wing Option" {
				if pos == "C" {
					score += 15
					if arch == "Run Blocking" {
						score += 10
					}

				}
				score += cp.RunBlock
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["C"] = append(positionMap["C"], dcpObj)
		}
		// Add to LE List
		if pos == "DE" || pos == "DT" || pos == "OLB" {
			score := 0
			if defScheme == "4-3" {
				if pos == "DE" {
					score += 20
					if arch == "Speed Rusher" || arch == "Balanced" {
						score += 10
					} else {
						score += 5
					}
				} else if (pos == "DT" && arch == "Pass Rusher") || (pos == "OLB" && (arch == "Pass Rusher" || arch == "Run Stopper")) {
					score += 10
				}
				score += cp.Overall
			} else if defScheme == "3-4" {
				if pos == "DE" {
					score += 20
					if arch == "Run Stopper" || arch == "Balanced" {
						score += 10
					} else {
						score += 5
					}
				} else if pos == "DT" && (arch == "Pass Rusher" || arch == "Balanced") {
					score += 15
				} else {
					score += 5
				}
				score += cp.Overall
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			if isLE {
				positionMap["LE"] = append(positionMap["LE"], dcpObj)
			} else {
				positionMap["RE"] = append(positionMap["RE"], dcpObj)
			}
			isLE = !isLE
		}

		// Add to DT list
		if pos == "DE" || pos == "DT" || pos == "OLB" {
			score := 0
			if defScheme == "4-3" {
				if pos == "DT" {
					score += 20
					if arch == "Pass Rusher" || arch == "Balanced" {
						score += 10
					} else {
						score += 5
					}
				} else if pos == "DE" && (arch == "Balanced" || arch == "Run Stopper") {
					score += 15
				}
				score += cp.Overall
			} else if defScheme == "3-4" {
				if pos == "DT" {
					score += 20
					if arch == "Nose Tackle" {
						score += 10
					}
				} else if pos == "DE" && arch == "Run Stopper" {
					score += 15
				}
				score += cp.RunDefense
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["DT"] = append(positionMap["DT"], dcpObj)
		}

		// Add to OLB list
		if pos == "OLB" || pos == "DE" || pos == "ILB" || pos == "SS" || pos == "FS" {
			score := 0
			if defScheme == "4-3" {
				if pos == "OLB" && (arch == "Coverage" || arch == "Speed") {
					score += 20
				} else if pos == "ILB" && (arch == "Coverage" || arch == "Speed") {
					score += 15
				} else if pos == "OLB" && arch == "Run Stopper" {
					score += 12
				} else if pos == "ILB" && arch == "Field General" {
					score += 10
				} else if (pos == "SS" || pos == "FS") && arch == "Run Stopper" {
					score += 8
				} else if pos == "OLB" && arch == "Pass Rush" {
					score += 5
				}
				score += cp.Speed
			} else if defScheme == "3-4" {
				if pos == "OLB" && (arch == "Pass Rush" || arch == "Run Stopper") {
					score += 10
				} else if pos == "DE" && arch == "Speed Rush" {
					score += 8
				} else if pos == "ILB" && arch == "Run Stopper" {
					score += 5
				}
				score += cp.PassRush
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			if isLOLB {
				positionMap["LOLB"] = append(positionMap["LOLB"], dcpObj)
			} else {
				positionMap["ROLB"] = append(positionMap["ROLB"], dcpObj)
			}
			isLOLB = !isLOLB
		}

		// Add to ILB list
		if pos == "OLB" || pos == "DE" || pos == "ILB" || pos == "SS" || pos == "FS" {
			score := 0
			if defScheme == "4-3" {
				if pos == "ILB" {
					score += 15
				} else if pos == "OLB" && (arch == "Speed" || arch == "Coverage" || arch == "Run Stopper") {
					score += 12
				} else if (pos == "SS" || pos == "FS") && arch == "Run Stopper" {
					score += 8
				}
				score += cp.RunDefense
			} else if defScheme == "3-4" {
				if pos == "ILB" {
					score += 15
				} else if pos == "OLB" && (arch == "Speed" || arch == "Coverage") {
					score += 12
				} else if (pos == "SS" || pos == "FS") && arch == "Run Stopper" {
					score += 8
				}
				score += cp.PassRush
			}
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["ILB"] = append(positionMap["ILB"], dcpObj)
		}

		// Add to CB List
		if pos == "CB" || pos == "FS" || pos == "SS" {
			score := 0
			if pos == "CB" {
				score += 10
			} else if pos == "SS" || pos == "FS" {
				score += 5
			}
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["CB"] = append(positionMap["CB"], dcpObj)
		}

		// Add to FS list
		if pos == "CB" || pos == "FS" || pos == "SS" {
			score := 0
			if pos == "FS" {
				score += 10
			} else if pos == "SS" || pos == "CB" {
				score += 5
			}
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["FS"] = append(positionMap["FS"], dcpObj)
		}

		// Add to SS list
		if pos == "CB" || pos == "FS" || pos == "SS" {
			score := 0
			if pos == "SS" {
				score += 10
			} else if pos == "FS" || pos == "CB" {
				score += 5
			}
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["SS"] = append(positionMap["SS"], dcpObj)
		}
		// Add to P list
		if pos == "K" || pos == "P" || pos == "QB" {
			score := 0
			if pos == "P" {
				score += 15
			} else if pos == "K" {
				score += 5
			}
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["P"] = append(positionMap["P"], dcpObj)
		}
		// Add to K list (Field Goal)
		if pos == "K" || pos == "P" || pos == "QB" {
			score := 0
			if pos == "K" {
				score += 15
			} else if pos == "P" {
				score += 5
			}
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["K"] = append(positionMap["K"], dcpObj)
		}

		if pos == "K" || pos == "P" || pos == "QB" {
			score := 0
			if pos == "K" {
				score += 15
			} else if pos == "P" {
				score += 5
			}
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["FG"] = append(positionMap["FG"], dcpObj)
		}

		// PR
		if pos == "WR" || pos == "RB" || pos == "FS" || pos == "SS" || pos == "CB" {
			score := 0
			if pos == "WR" || pos == "RB" {
				score += 7
			}
			score += cp.Agility

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["PR"] = append(positionMap["PR"], dcpObj)
		}
		// KR
		if pos == "WR" || pos == "RB" || pos == "FS" || pos == "SS" || pos == "CB" {
			score := 0
			if pos == "WR" || pos == "RB" {
				score += 7
			}
			score += cp.Speed

			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["KR"] = append(positionMap["KR"], dcpObj)
		}
		// STU
		if pos == "FB" || pos == "TE" || pos == "ILB" || pos == "OLB" || pos == "RB" || pos == "CB" || pos == "FS" || pos == "SS" {
			score := 0
			if cp.Experience == 2 {
				score += 15
			} else if cp.Experience == 1 {
				score += 10
			} else if cp.Experience == 3 {
				score += 5
			}

			score += cp.Tackle
			dcpObj := structs.DepthChartPositionDTO{
				Position:  pos,
				Archetype: arch,
				Score:     score,
				NFLPlayer: cp,
			}
			positionMap["STU"] = append(positionMap["STU"], dcpObj)
		}
	}

	// Sort Each DC Position
	sort.Sort(structs.ByDCPosition(positionMap["QB"]))
	sort.Sort(structs.ByDCPosition(positionMap["RB"]))
	sort.Sort(structs.ByDCPosition(positionMap["FB"]))
	sort.Sort(structs.ByDCPosition(positionMap["WR"]))
	sort.Sort(structs.ByDCPosition(positionMap["TE"]))
	sort.Sort(structs.ByDCPosition(positionMap["LT"]))
	sort.Sort(structs.ByDCPosition(positionMap["RT"]))
	sort.Sort(structs.ByDCPosition(positionMap["LG"]))
	sort.Sort(structs.ByDCPosition(positionMap["RG"]))
	sort.Sort(structs.ByDCPosition(positionMap["C"]))
	sort.Sort(structs.ByDCPosition(positionMap["DT"]))
	sort.Sort(structs.ByDCPosition(positionMap["LE"]))
	sort.Sort(structs.ByDCPosition(positionMap["RE"]))
	sort.Sort(structs.ByDCPosition(positionMap["LOLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["ROLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["ILB"]))
	sort.Sort(structs.ByDCPosition(positionMap["CB"]))
	sort.Sort(structs.ByDCPosition(positionMap["FS"]))
	sort.Sort(structs.ByDCPosition(positionMap["SS"]))
	sort.Sort(structs.ByDCPosition(positionMap["P"]))
	sort.Sort(structs.ByDCPosition(positionMap["K"]))
	sort.Sort(structs.ByDCPosition(positionMap["PR"]))
	sort.Sort(structs.ByDCPosition(positionMap["KR"]))
	sort.Sort(structs.ByDCPosition(positionMap["FG"]))
	sort.Sort(structs.ByDCPosition(positionMap["STU"]))

	for _, dcp := range dcPositions {
		positionList := positionMap[dcp.Position]
		for _, pos := range positionList {
			if starterMap[pos.NFLPlayer.ID] {
				continue
			}
			if backupMap[pos.NFLPlayer.ID] && dcp.PositionLevel != "1" {
				continue
			}
			if dcp.PositionLevel == "1" && !starterMap[pos.NFLPlayer.ID] {
				starterMap[pos.NFLPlayer.ID] = true
			} else {
				backupMap[pos.NFLPlayer.ID] = true
			}
			dto := structs.NFLDepthChartPosition{
				DepthChartID:     dcp.DepthChartID,
				PlayerID:         pos.NFLPlayer.ID,
				FirstName:        pos.NFLPlayer.FirstName,
				LastName:         pos.NFLPlayer.LastName,
				OriginalPosition: pos.NFLPlayer.Position,
			}
			dto.AssignID(dcp.ID)
			dcp.UpdateDepthChartPosition(dto)
			db.Save(&dcp)
			break
		}
	}
}

func GetDepthChartMap() map[uint]structs.CollegeTeamDepthChart {
	m := &sync.Mutex{}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)
	collegeTeams := GetAllCollegeTeams()
	dcMap := make(map[uint]structs.CollegeTeamDepthChart)

	for _, team := range collegeTeams {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(t structs.CollegeTeam) {
			defer wg.Done()
			id := strconv.Itoa(int(t.ID))
			depthChart := GetDepthchartByTeamID(id)

			m.Lock()
			dcMap[t.ID] = depthChart
			m.Unlock()

			<-semaphore
		}(team)
	}

	wg.Wait()
	close(semaphore)
	return dcMap
}

func GetTestOffensiveSchemesByTeamID(id uint) string {
	if id == 1 || id == 59 || id == 65 || id == 77 || id == 107 {
		return "West Coast"
	}
	if id == 7 || id == 54 || id == 98 {
		return "Power Run"
	}
	if id == 10 || id == 123 || id == 125 {
		return "Double Wing"
	}
	if id == 13 || id == 12 || id == 47 || id == 118 {
		return "Spread Option"
	}
	if id == 15 || id == 34 || id == 78 || id == 80 {
		return "Wing-T"
	}
	if id == 19 || id == 44 || id == 45 {
		return "Flexbone"
	}
	if id == 23 || id == 37 || id == 109 {
		return "Air Raid"
	}
	if id == 55 || id == 88 {
		return "I Option"
	}
	if id == 56 || id == 86 || id == 93 || id == 100 || id == 115 {
		return "Vertical"
	}
	if id == 63 || id == 99 || id == 108 || id == 62 || id == 39 {
		return "Pistol"
	}
	if id == 75 || id == 96 || id == 122 {
		return "Run and Shoot"
	}
	if id == 94 || id == 97 || id == 127 {
		return "Wishbone"
	}
	return ""
}

func GetTestDefensiveSchemesByTeamID(id uint) string {
	if id == 10 || id == 13 || id == 54 || id == 77 || id == 86 || id == 93 || id == 94 || id == 97 || id == 107 {
		return "Old School"
	}
	if id == 15 || id == 19 || id == 44 || id == 55 || id == 56 || id == 63 || id == 98 || id == 118 {
		return "2-Gap"
	}
	if id == 1 || id == 12 || id == 34 || id == 47 || id == 80 || id == 108 || id == 109 || id == 122 || id == 127 {
		return "4-Man Front Spread Stopper"
	}
	if id == 23 || id == 65 || id == 99 || id == 123 {
		return "3-Man Front Spread Stopper"
	}
	if id == 37 || id == 45 || id == 75 || id == 78 || id == 88 || id == 96 || id == 100 || id == 39 || id == 62 {
		return "Speed"
	}
	if id == 7 || id == 59 || id == 115 || id == 125 {
		return "Multiple"
	}
	return ""
}

func MassUpdateGameplanSchemesTEST(off, def string) {
	db := dbprovider.GetInstance().GetDB()
	teams := GetAllCollegeTeams()
	offensiveSchemes := util.GetOffensiveDefaultSchemes()
	defensiveSchemes := util.GetDefensiveDefaultSchemes()
	for _, team := range teams {
		teamID := strconv.Itoa(int(team.ID))
		gp := GetGameplanTESTByTeamID(teamID)
		gp.UpdateSchemes(off, def)
		// offe := GetTestOffensiveSchemesByTeamID(id)
		// defe := GetTestDefensiveSchemesByTeamID(id)
		// Map Default Scheme for offense & defense
		offFormations := offensiveSchemes[off]
		defFormations := defensiveSchemes[def][off]

		dto := structs.CollegeGameplanTEST{
			TeamID: int(team.ID),
			BaseGameplan: structs.BaseGameplan{
				OffensiveScheme:    off,
				DefensiveScheme:    def,
				OffensiveFormation: offFormations,
				DefensiveFormation: defFormations,
				BlitzSafeties:      gp.BlitzSafeties,
				BlitzCorners:       gp.BlitzCorners,
				LinebackerCoverage: gp.LinebackerCoverage,
				MaximumFGDistance:  gp.MaximumFGDistance,
				GoFor4AndShort:     gp.GoFor4AndShort,
				GoFor4AndLong:      gp.GoFor4AndLong,
				DefaultOffense:     gp.DefaultOffense,
				DefaultDefense:     gp.DefaultDefense,
				PrimaryHB:          75,
				PitchFocus:         50,
				DiveFocus:          50,
			},
		}

		gp.UpdateCollegeGameplanTEST(dto)

		// Autosort Depth Chart
		ReAlignCollegeDepthChartTEST(db, teamID, gp)

		db.Save(&gp)
	}
}

func UpdateIndividualGameplanSchemeTEST(teamID, off, def string) {
	db := dbprovider.GetInstance().GetDB()
	offensiveSchemes := util.GetOffensiveDefaultSchemes()
	defensiveSchemes := util.GetDefensiveDefaultSchemes()

	gp := GetGameplanTESTByTeamID(teamID)
	gp.UpdateSchemes(off, def)
	// Map Default Scheme for offense & defense
	offFormations := offensiveSchemes[off]
	defFormations := defensiveSchemes[def][off]

	dto := structs.CollegeGameplanTEST{
		TeamID: gp.TeamID,
		BaseGameplan: structs.BaseGameplan{
			OffensiveScheme:    off,
			DefensiveScheme:    def,
			OffensiveFormation: offFormations,
			DefensiveFormation: defFormations,
			BlitzSafeties:      gp.BlitzSafeties,
			BlitzCorners:       gp.BlitzCorners,
			LinebackerCoverage: gp.LinebackerCoverage,
			MaximumFGDistance:  gp.MaximumFGDistance,
			GoFor4AndShort:     gp.GoFor4AndShort,
			GoFor4AndLong:      gp.GoFor4AndLong,
			DefaultOffense:     gp.DefaultOffense,
			DefaultDefense:     gp.DefaultDefense,
			PrimaryHB:          75,
			PitchFocus:         50,
			DiveFocus:          50,
		},
	}

	gp.UpdateCollegeGameplanTEST(dto)

	// Autosort Depth Chart
	ReAlignCollegeDepthChartTEST(db, teamID, gp)

	db.Save(&gp)

}

func ReAlignCollegeDepthChartTEST(db *gorm.DB, teamID string, gp structs.CollegeGameplanTEST) {
	roster := GetAllCollegePlayersByTeamIdWithoutRedshirts(teamID)
	dcPositions := GetDepthChartPositionPlayersByDepthchartIDTEST(teamID)
	sort.Sort(structs.ByOverall(roster))
	positionMap := make(map[string][]structs.DepthChartPositionDTO)
	starterMap := make(map[uint]bool)
	backupMap := make(map[uint]bool)
	stuMap := make(map[uint]bool)
	offScheme := gp.OffensiveScheme
	defScheme := gp.DefensiveScheme
	isLT := true
	isLG := true
	isLE := true
	isLOLB := true

	goodFits := GetFitsByScheme(offScheme, false)
	badFits := GetFitsByScheme(defScheme, false)

	// Allocate the Position Map
	for _, cp := range roster {
		if cp.IsInjured || cp.IsRedshirting {
			continue
		}
		pos := cp.Position
		arch := cp.Archetype
		player := arch + " " + pos
		isGoodFit := CheckPlayerFits(player, goodFits)
		isBadFit := CheckPlayerFits(player, badFits)

		// Add to QB List
		if pos == "QB" || pos == "RB" || pos == "FB" || pos == "ATH" {
			score := 0
			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}
			if pos == "QB" {
				score += 100
			} else if pos == "ATH" && (arch == "Triple-Threat" || arch == "Field General") {
				score += 50
			}
			// score += ((cp.ThrowAccuracy + cp.ThrowPower) / 2)
			score += cp.Overall

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["QB"] = append(positionMap["QB"], dcpObj)
		}
		// Add to RB List
		if pos == "RB" || pos == "FB" || pos == "WR" || pos == "TE" || pos == "ATH" {
			score := 0
			if pos == "RB" {
				score += 100
			} else if pos == "FB" {
				score += 25
			} else if pos == "ATH" && (arch == "Wingback" || arch == "Soccer Player" || arch == "Triple-Threat") {
				score += 50
			}
			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += ((cp.Speed + cp.Agility + cp.Strength + cp.Carrying) / 4)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["RB"] = append(positionMap["RB"], dcpObj)
		}

		// Add to FB List
		if pos == "FB" || pos == "TE" || pos == "RB" || pos == "ATH" {
			score := 0
			if pos == "FB" {
				score += 100
			} else if pos == "ATH" && (arch == "Wingback") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}
			score += ((cp.Strength + cp.Carrying + cp.PassBlock + cp.RunBlock) / 4)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["FB"] = append(positionMap["FB"], dcpObj)
		}

		// Add to TE List
		if pos == "FB" || pos == "TE" || pos == "ATH" {
			score := 0
			if pos == "TE" {
				score += 100
			} else if pos == "ATH" && (arch == "Slotback") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) + int(float64(cp.RunBlock)*0.125) + int(float64(cp.PassBlock)*0.125) + int(float64(cp.Catching)*0.125) + int(float64(cp.Strength)*0.125)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["TE"] = append(positionMap["TE"], dcpObj)
		}
		// Add to WR List
		if pos == "WR" || pos == "TE" || pos == "RB" || pos == "ATH" {
			score := 0
			if pos == "WR" {
				score += 100
			} else if pos == "ATH" && (arch == "Wingback" || arch == "Slotback") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.4) +
				int(float64(cp.Speed)*0.12) +
				int(float64(cp.Agility)*0.12) +
				int(float64(cp.Catching)*0.12) +
				int(float64(cp.Strength)*0.12) +
				int(float64(cp.RouteRunning)*0.12)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["WR"] = append(positionMap["WR"], dcpObj)
		}
		// Add to LT and RT List
		if pos == "OT" || pos == "OG" || pos == "C" || pos == "ATH" {
			score := 0
			if pos == "OT" {
				score += 100
			} else if pos == "OG" {
				score += 25
			} else if pos == "ATH" && (arch == "Lineman") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.10) +
				int(float64(cp.RunBlock)*0.75) +
				int(float64(cp.PassBlock)*0.75) +
				int(float64(cp.Agility)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLT {
				positionMap["LT"] = append(positionMap["LT"], dcpObj)
			} else {
				positionMap["RT"] = append(positionMap["RT"], dcpObj)
			}
			isLT = !isLT
		}
		// Add to LG and RG List
		if pos == "OT" || pos == "OG" || pos == "C" || pos == "ATH" {
			score := 0
			if pos == "OG" {
				score += 100
			} else if pos == "C" {
				score += 25
			} else if pos == "ATH" && (arch == "Lineman") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 30
			} else if isBadFit && !isGoodFit {
				score -= 30
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.10) +
				int(float64(cp.RunBlock)*0.75) +
				int(float64(cp.PassBlock)*0.75) +
				int(float64(cp.Agility)*0.05)
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLG {
				positionMap["LG"] = append(positionMap["LG"], dcpObj)
			} else {
				positionMap["RG"] = append(positionMap["RG"], dcpObj)
			}
			isLG = !isLG
		}
		// Add to C List
		if pos == "OT" || pos == "OG" || pos == "C" || pos == "ATH" {
			score := 0
			if pos == "C" {
				score += 100
			} else if pos == "OG" {
				score += 25
			} else if pos == "ATH" && (arch == "Lineman") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.10) +
				int(float64(cp.RunBlock)*0.75) +
				int(float64(cp.PassBlock)*0.75) +
				int(float64(cp.Agility)*0.05)
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["C"] = append(positionMap["C"], dcpObj)
		}

		// Add to LE List
		if pos == "DE" || pos == "DT" || pos == "OLB" || pos == "ATH" {
			score := 0
			if pos == "DE" {
				score += 100
			} else if pos == "OLB" {
				score += 25
			} else if pos == "DT" {
				score += 3
			} else if pos == "ATH" && (arch == "Lineman" || arch == "Strongside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.05) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.PassRush)*0.75) +
				int(float64(cp.RunDefense)*0.75) +
				int(float64(cp.Agility)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLE {
				positionMap["LE"] = append(positionMap["LE"], dcpObj)
			} else {
				positionMap["RE"] = append(positionMap["RE"], dcpObj)
			}
			isLE = !isLE
		}

		// Add to DT list
		if pos == "DE" || pos == "DT" || pos == "OLB" || pos == "ATH" {
			score := 0
			if pos == "DT" {
				score += 100
			} else if pos == "DE" {
				score += 25
			} else if pos == "OLB" {
				score += 12
			} else if pos == "ATH" && (arch == "Lineman" || arch == "Strongside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.7) +
				int(float64(cp.Strength)*0.05) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.PassRush)*0.75) +
				int(float64(cp.RunDefense)*0.75) +
				int(float64(cp.Agility)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["DT"] = append(positionMap["DT"], dcpObj)
		}

		// Add to OLB list
		if pos == "OLB" || pos == "DE" || pos == "ILB" || pos == "SS" || pos == "FS" || pos == "ATH" {
			score := 0
			if pos == "OLB" {
				score += 100
			} else if pos == "DE" {
				score += 10
			} else if pos == "ILB" {
				score += 25
			} else if pos == "SS" {
				score += 3
			} else if pos == "ATH" && (arch == "Weakside" || arch == "Strongside" || arch == "Bandit") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.6) +
				int(float64(cp.Strength)*0.025) +
				int(float64(cp.Tackle)*0.055) +
				int(float64(cp.PassRush)*0.0755) +
				int(float64(cp.RunDefense)*0.0755) +
				int(float64(cp.ManCoverage)*0.075) +
				int(float64(cp.ZoneCoverage)*0.075) +
				int(float64(cp.Agility)*0.025)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			if isLOLB {
				positionMap["LOLB"] = append(positionMap["LOLB"], dcpObj)
			} else {
				positionMap["ROLB"] = append(positionMap["ROLB"], dcpObj)
			}
			isLOLB = !isLOLB
		}

		// Add to ILB list
		if pos == "OLB" || pos == "DE" || pos == "ILB" || pos == "SS" || pos == "FS" || pos == "ATH" {
			score := 0
			if pos == "ILB" {
				score += 100
			} else if pos == "OLB" {
				score += 25
			} else if pos == "SS" {
				score += 8
			} else if pos == "DE" {
				score += 3
			} else if pos == "ATH" && (arch == "Weakside" || arch == "Bandit" || arch == "Field General") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.6) +
				int(float64(cp.Strength)*0.025) +
				int(float64(cp.Tackle)*0.055) +
				int(float64(cp.PassRush)*0.0755) +
				int(float64(cp.RunDefense)*0.0755) +
				int(float64(cp.ManCoverage)*0.075) +
				int(float64(cp.ZoneCoverage)*0.075) +
				int(float64(cp.Agility)*0.025)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["MLB"] = append(positionMap["MLB"], dcpObj)
		}

		// Add to CB List
		if pos == "CB" || pos == "FS" || pos == "SS" || pos == "ATH" {
			score := 0
			if pos == "CB" {
				score += 100
			} else if pos == "FS" {
				score += 10
			} else if pos == "SS" {
				score += 8
			} else if pos == "ATH" && (arch == "Triple-Threat" || arch == "Bandit" || arch == "Weakside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.Agility)*0.1) +
				int(float64(cp.Catching)*0.1) +
				int(float64(cp.ManCoverage)*0.01) +
				int(float64(cp.ZoneCoverage)*0.01) +
				int(float64(cp.Speed)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["CB"] = append(positionMap["CB"], dcpObj)
		}

		// Add to FS list
		if pos == "CB" || pos == "FS" || pos == "SS" || pos == "ATH" {
			score := 0
			if pos == "FS" {
				score += 100
			} else if pos == "CB" {
				score += 25
			} else if pos == "SS" {
				score += 12
			} else if pos == "ATH" && (arch == "Bandit" || arch == "Weakside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.Agility)*0.1) +
				int(float64(cp.Catching)*0.1) +
				int(float64(cp.ManCoverage)*0.01) +
				int(float64(cp.ZoneCoverage)*0.01) +
				int(float64(cp.Speed)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["FS"] = append(positionMap["FS"], dcpObj)
		}

		// Add to SS list
		if pos == "CB" || pos == "FS" || pos == "SS" || pos == "ATH" {
			score := 0
			if pos == "SS" {
				score += 100
			} else if pos == "FS" {
				score += 25
			} else if pos == "CB" {
				score += 12
			} else if pos == "ATH" && (arch == "Bandit" || arch == "Weakside") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += int(float64(cp.Overall)*0.5) +
				int(float64(cp.Tackle)*0.05) +
				int(float64(cp.Agility)*0.1) +
				int(float64(cp.Catching)*0.1) +
				int(float64(cp.ManCoverage)*0.01) +
				int(float64(cp.ZoneCoverage)*0.01) +
				int(float64(cp.Speed)*0.05)

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["SS"] = append(positionMap["SS"], dcpObj)
		}

		// Add to P list
		if pos == "K" || pos == "P" || pos == "QB" || pos == "ATH" {
			score := 0
			if pos == "P" {
				score += 100
			} else if pos == "K" {
				score += 25
			} else if pos == "ATH" && (arch == "Soccer Player") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += cp.PuntAccuracy + cp.PuntPower

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["P"] = append(positionMap["P"], dcpObj)
		}

		// Add to K list (Field Goal)
		if pos == "K" || pos == "P" || pos == "QB" || pos == "ATH" {
			score := 0
			if pos == "K" {
				score += 100
			} else if pos == "P" {
				score += 25
			} else if pos == "ATH" && (arch == "Soccer Player") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}
			score += cp.KickAccuracy + cp.KickPower
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["K"] = append(positionMap["K"], dcpObj)
		}

		// FG List
		if pos == "K" || pos == "P" || pos == "QB" || pos == "ATH" {
			score := 0
			if pos == "K" {
				score += 100
			} else if pos == "P" {
				score += 25
			} else if pos == "ATH" && (arch == "Soccer Player") {
				score += 50
			}

			if isGoodFit && !isBadFit {
				score += 50
			} else if isBadFit && !isGoodFit {
				score -= 50
			}

			score += cp.KickAccuracy + cp.KickPower

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["FG"] = append(positionMap["FG"], dcpObj)
		}

		// PR
		if pos == "WR" || pos == "RB" || pos == "FS" || pos == "SS" || pos == "CB" || pos == "ATH" {
			score := 0
			if pos == "ATH" && arch == "Return Specialist" {
				score += 50
			} else if pos == "WR" || pos == "RB" {
				score += 25
			}
			score += cp.Agility

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["PR"] = append(positionMap["PR"], dcpObj)
		}

		// KR
		if pos == "WR" || pos == "RB" || pos == "FS" || pos == "SS" || pos == "CB" || pos == "ATH" {
			score := 0
			if pos == "ATH" && arch == "Return Specialist" {
				score += 50
			} else if pos == "WR" || pos == "RB" {
				score += 25
			}
			score += cp.Speed

			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["KR"] = append(positionMap["KR"], dcpObj)
		}

		// STU
		if pos == "FB" || pos == "TE" || pos == "ILB" || pos == "OLB" || pos == "RB" || pos == "CB" || pos == "FS" || pos == "SS" || pos == "WR" || pos == "ATH" {
			score := 0
			if cp.Year == 2 || cp.Year == 1 {
				score += 50
			} else if cp.Year == 3 && cp.IsRedshirt {
				score += 25
			}

			score += cp.Tackle
			dcpObj := structs.DepthChartPositionDTO{
				Position:      pos,
				Archetype:     arch,
				Score:         score,
				CollegePlayer: cp,
			}
			positionMap["STU"] = append(positionMap["STU"], dcpObj)
		}
	}

	// Sort Each DC Position
	sort.Sort(structs.ByDCPosition(positionMap["QB"]))
	sort.Sort(structs.ByDCPosition(positionMap["RB"]))
	sort.Sort(structs.ByDCPosition(positionMap["FB"]))
	sort.Sort(structs.ByDCPosition(positionMap["WR"]))
	sort.Sort(structs.ByDCPosition(positionMap["TE"]))
	sort.Sort(structs.ByDCPosition(positionMap["LT"]))
	sort.Sort(structs.ByDCPosition(positionMap["RT"]))
	sort.Sort(structs.ByDCPosition(positionMap["LG"]))
	sort.Sort(structs.ByDCPosition(positionMap["RG"]))
	sort.Sort(structs.ByDCPosition(positionMap["C"]))
	sort.Sort(structs.ByDCPosition(positionMap["DT"]))
	sort.Sort(structs.ByDCPosition(positionMap["LE"]))
	sort.Sort(structs.ByDCPosition(positionMap["RE"]))
	sort.Sort(structs.ByDCPosition(positionMap["LOLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["ROLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["MLB"]))
	sort.Sort(structs.ByDCPosition(positionMap["CB"]))
	sort.Sort(structs.ByDCPosition(positionMap["FS"]))
	sort.Sort(structs.ByDCPosition(positionMap["SS"]))
	sort.Sort(structs.ByDCPosition(positionMap["P"]))
	sort.Sort(structs.ByDCPosition(positionMap["K"]))
	sort.Sort(structs.ByDCPosition(positionMap["PR"]))
	sort.Sort(structs.ByDCPosition(positionMap["KR"]))
	sort.Sort(structs.ByDCPosition(positionMap["FG"]))
	sort.Sort(structs.ByDCPosition(positionMap["STU"]))

	for _, dcp := range dcPositions {
		positionList := positionMap[dcp.Position]
		for _, pos := range positionList {
			if starterMap[pos.CollegePlayer.ID] &&
				dcp.Position != "FG" &&
				dcp.Position != "PR" &&
				dcp.Position != "KR" {
				continue
			}
			if backupMap[pos.CollegePlayer.ID] && dcp.PositionLevel != "1" && dcp.Position != "STU" {
				continue
			}
			if dcp.Position == "STU" && stuMap[pos.CollegePlayer.ID] {
				continue
			}

			if dcp.Position == "WR" {
				runnerDistPostition := gp.RunnerDistributionWRPosition
				positionLabel := dcp.Position + "" + dcp.PositionLevel
				if runnerDistPostition == positionLabel {
					gp.AssignRunnerWRID(dcp.CollegePlayer.ID)
				}
			}

			if dcp.Position == "STU" {
				stuMap[pos.CollegePlayer.ID] = true
			} else if dcp.PositionLevel == "1" && !starterMap[pos.CollegePlayer.ID] {
				starterMap[pos.CollegePlayer.ID] = true
			} else {
				backupMap[pos.CollegePlayer.ID] = true
			}
			dto := structs.CollegeDepthChartPositionTEST{
				DepthChartID:     dcp.DepthChartID,
				PlayerID:         int(pos.CollegePlayer.ID),
				FirstName:        pos.CollegePlayer.FirstName,
				LastName:         pos.CollegePlayer.LastName,
				OriginalPosition: pos.CollegePlayer.Position,
			}
			dto.AssignID(dcp.ID)
			dcp.UpdateDepthChartPosition(dto)
			db.Save(&dcp)
			break
		}
	}
}

func GetDepthChartPositionPlayersByDepthchartIDTEST(depthChartID string) []structs.CollegeDepthChartPositionTEST {
	db := dbprovider.GetInstance().GetDB()

	var positionPlayers []structs.CollegeDepthChartPositionTEST

	err := db.Where("depth_chart_id = ?", depthChartID).Find(&positionPlayers).Error
	if err != nil {
		fmt.Println(err)
		panic("Depth Chart does not exist for this ID")
	}

	return positionPlayers
}
