package managers

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/CalebRose/SimFBA/repository"
	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
)

// =============================================================================
// SCHEDULING MANAGER - MAIN BUSINESS LOGIC ORCHESTRATOR
// =============================================================================
// This file contains the core scheduling algorithm that coordinates all the
// helper functions to generate an optimal non-conference football schedule.
//
// Go Learning: This demonstrates the Manager pattern - a central orchestrator
// that coordinates multiple helper functions and repository operations.

// SchedulingManager handles the complex task of generating football schedules
// while respecting numerous constraints and business rules
type SchedulingManager struct {
	repository *repository.SchedulingRepository
	db         *gorm.DB
}

// NewSchedulingManager creates a new instance of the scheduling manager
// Go Learning: Constructor pattern in Go - factory function that returns configured struct
func NewSchedulingManager(db *gorm.DB) *SchedulingManager {
	return &SchedulingManager{
		repository: repository.NewSchedulingRepository(db),
		db:         db,
	}
}

// GenerateOptimalSchedule is the main entry point for schedule generation
// This orchestrates the entire scheduling process following the algorithm outlined in guidelines.md
//
// PSEUDOCODE ALGORITHM:
// 1. LOAD existing conference games from database
// 2. BUILD comprehensive scheduling maps (availability, rivalries, constraints)
// 3. VALIDATE input constraints and team data
// 4. SCHEDULE rivalry games first (highest priority)
// 5. FILL remaining non-conference slots with constraint satisfaction
// 6. VALIDATE final schedule against all constraints
// 7. EXPORT results to CSV format
// 8. RETURN schedule data for API response
func (sm *SchedulingManager) GenerateOptimalSchedule(seasonID int) (*structs.ScheduleResponse, error) {
	log.Printf("Starting schedule generation for season %d", seasonID)

	// ==== STEP 1: DATA LOADING PHASE ====
	// TODO: Implement these data loading operations in your repository

	// PSEUDOCODE: Load existing conference games
	// existingConferenceGames := CALL repository.GetExistingConferenceGames(seasonID)
	// IF error THEN return error
	// LOG "Loaded X existing conference games"

	existingConferenceGames, err := sm.repository.GetExistingConferenceGames(uint(seasonID))
	if err != nil {
		return nil, fmt.Errorf("failed to load existing conference games: %w", err)
	}
	log.Printf("Loaded %d existing conference games", len(existingConferenceGames))

	// PSEUDOCODE: Load all teams for this season
	// allTeams := CALL repository.GetAllTeams(seasonID)
	// IF error THEN return error
	// LOG "Loaded X teams"

	allTeams, err := sm.repository.GetAllTeams(uint(seasonID))
	if err != nil {
		return nil, fmt.Errorf("failed to load teams: %w", err)
	}
	log.Printf("Loaded %d teams", len(allTeams))

	// PSEUDOCODE: Load rivalry relationships
	// rivalries := CALL repository.GetRivalries()
	// IF error THEN return error
	// LOG "Loaded X rivalry relationships"

	rivalries, err := sm.repository.GetRivalries()
	if err != nil {
		return nil, fmt.Errorf("failed to load rivalries: %w", err)
	}
	log.Printf("Loaded rivalry data for %d teams", len(rivalries))

	// ==== STEP 2: SCHEDULING MAPS CONSTRUCTION ====
	// This builds the foundation data structures needed for constraint-based scheduling

	// PSEUDOCODE: Build comprehensive scheduling maps
	// schedulingMaps := CALL BuildContendingMap(existingConferenceGames, allTeams, 15)
	// schedulingMaps.TeamRivalries = rivalries
	// IF error THEN return error
	// LOG "Built scheduling maps with availability data"

	const NUM_WEEKS = 15 // Standard college football season length 0-14
	schedulingMaps, err := BuildContendingMap(existingConferenceGames, allTeams, NUM_WEEKS)
	if err != nil {
		return nil, fmt.Errorf("failed to build scheduling maps: %w", err)
	}

	// Load rivalry priorities for scheduling decisions
	rivalryPriorities, err := sm.repository.GetRivalryPriorities()
	if err != nil {
		return nil, fmt.Errorf("failed to load rivalry priorities: %w", err)
	}

	// Integrate rivalry data into scheduling maps
	schedulingMaps.TeamRivalries = rivalries
	schedulingMaps.TeamRivalryPriorities = rivalryPriorities
	log.Printf("Built comprehensive scheduling maps with %d rivalry relationships", len(rivalryPriorities))

	// ==== STEP 3: CONSTRAINT SETUP ====
	// Define the business rules that the schedule must satisfy

	// PSEUDOCODE: Define scheduling constraints
	// constraints := CREATE SchedulingConstraints{
	//     MaxGamesPerTeam: 12,
	//     MaxHomeGames: 7,
	//     MaxAwayGames: 7,
	//     MaxConsecutiveAwayGames: 3,
	//     PreventBackToBackByes: true,
	//     EnforceRivalries: true
	// }

	constraints := structs.SchedulingConstraints{
		MaxGamesPerTeam:         12,   // Standard college football season
		MaxGamesPerWeek:         1,    // One game per team per week
		MaxHomeGames:            7,    // Roughly half should be home games
		MaxAwayGames:            7,    // Roughly half should be away games
		MaxConsecutiveAwayGames: 3,    // Don't want too many road trips in a row
		PreventBackToBackByes:   true, // Avoid consecutive bye weeks
		EnforceRivalries:        true, // Must schedule traditional rivalry games
	}
	log.Printf("Defined scheduling constraints")

	// ==== STEP 4: PRIORITY-BASED SCHEDULING ====
	// Schedule the most constrained teams first (rivalry games, limited availability)

	// PSEUDOCODE: Sort teams by scheduling difficulty
	// prioritizedTeams := CALL SortTeamsBySchedulingPriority(allTeams, schedulingMaps)
	// LOG "Sorted teams by scheduling priority"

	prioritizedTeams := SortTeamsBySchedulingPriority(allTeams, schedulingMaps)
	log.Printf("Prioritized %d teams for scheduling", len(prioritizedTeams))

	// Initialize the new games slice to collect our generated schedule
	var newNonConferenceGames []structs.CollegeGame

	// ==== STEP 5: RIVALRY GAMES SCHEDULING ====
	// Handle rivalry games first since they're the hardest constraints to satisfy

	// PSEUDOCODE: Schedule all rivalry games first
	// rivalryGames := CALL scheduleRivalryGames(rivalries, schedulingMaps, constraints)
	// newNonConferenceGames = APPEND newNonConferenceGames, rivalryGames
	// UPDATE schedulingMaps with newly scheduled games
	// LOG "Scheduled X rivalry games"

	rivalryGames, updatedMaps, err := sm.scheduleRivalryGames(rivalries, schedulingMaps, constraints, seasonID)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule rivalry games: %w", err)
	}
	newNonConferenceGames = append(newNonConferenceGames, rivalryGames...)
	schedulingMaps = updatedMaps
	log.Printf("Scheduled %d rivalry games", len(rivalryGames))

	// ==== STEP 6: REMAINING SLOTS SCHEDULING ====
	// Fill the remaining open slots with compatible opponents

	// PSEUDOCODE: Fill remaining non-conference slots
	// FOR each team in prioritizedTeams:
	//     WHILE team needs more games:
	//         compatibleOpponents := CALL FilterCompatibleOpponents(team, allTeams, schedulingMaps, constraints)
	//         IF no compatible opponents THEN continue to next team
	//         bestOpponent := CALL selectBestOpponent(team, compatibleOpponents, schedulingMaps)
	//         availableWeek := CALL FindAvailableWeek(team, bestOpponent, schedulingMaps)
	//         IF no available week THEN continue to next opponent
	//         CREATE new game between team and bestOpponent in availableWeek
	//         UPDATE schedulingMaps with new game
	//         ADD game to newNonConferenceGames

	remainingGames, finalMaps, err := sm.fillRemainingNonConferenceSlots(prioritizedTeams, schedulingMaps, constraints, seasonID)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule remaining games: %w", err)
	}
	newNonConferenceGames = append(newNonConferenceGames, remainingGames...)
	schedulingMaps = finalMaps
	log.Printf("Scheduled %d additional non-conference games", len(remainingGames))

	// ==== STEP 7: VALIDATION PHASE ====
	// Ensure the final schedule meets all business constraints

	// PSEUDOCODE: Validate the complete schedule
	// allGames := COMBINE existingConferenceGames + newNonConferenceGames
	// violations := CALL ValidateScheduleConstraints(allGames, constraints, allTeams)
	// IF violations exist THEN
	//     LOG warnings about violations
	//     OPTIONALLY attempt to fix violations

	allScheduledGames := append(existingConferenceGames, newNonConferenceGames...)
	violations, err := ValidateScheduleConstraints(allScheduledGames, constraints, allTeams)
	if err != nil {
		return nil, fmt.Errorf("schedule validation failed: %w", err)
	}

	if len(violations) > 0 {
		log.Printf("Schedule has %d constraint violations:", len(violations))
		for _, violation := range violations {
			log.Printf("  - %s", violation)
		}
		// TODO: Implement constraint violation resolution
		// This could involve backtracking, game rescheduling, or manual review flags
	} else {
		log.Printf("Schedule passes all constraint validations")
	}

	// ==== STEP 8: RESPONSE PREPARATION ====
	// Organize the results for API consumption and CSV export

	// PSEUDOCODE: Organize games by week for easy consumption
	// gamesByWeek := CREATE empty map[int][]CollegeGame
	// FOR each game in newNonConferenceGames:
	//     week := game.WeekID
	//     gamesByWeek[week] = APPEND gamesByWeek[week], game

	gamesByWeek := make(map[int][]structs.CollegeGame)
	for _, game := range newNonConferenceGames {
		week := int(game.WeekID)
		gamesByWeek[week] = append(gamesByWeek[week], game)
	}

	// Create the final response
	response := &structs.ScheduleResponse{
		ScheduledGamesByWeekMap: gamesByWeek,
	}

	log.Printf("Successfully generated schedule with %d new games across %d weeks",
		len(newNonConferenceGames), len(gamesByWeek))

	return response, nil
}

// scheduleRivalryGames handles the high-priority task of scheduling traditional rivalry games
// These games have the most constraints and should be scheduled first
//
// PSEUDOCODE ALGORITHM:
// 1. ITERATE through all rivalry relationships
// 2. FOR each rivalry pair, find mutual available weeks
// 3. DETERMINE home/away rotation based on historical data
// 4. ASSIGN game to optimal week considering TV schedules and traditions
// 5. UPDATE scheduling maps to reflect new assignments
func (sm *SchedulingManager) scheduleRivalryGames(
	rivalries map[uint][]uint,
	schedulingMaps structs.SchedulingMaps,
	constraints structs.SchedulingConstraints,
	seasonID int,
) ([]structs.CollegeGame, structs.SchedulingMaps, error) {

	var rivalryGames []structs.CollegeGame
	updatedMaps := schedulingMaps // Copy the maps for updates

	processedPairs := make(map[string]bool)
	var smallerID, largerID uint

	for teamA, rivalTeams := range rivalries {
		for _, teamB := range rivalTeams {
			if teamA < teamB {
				smallerID = teamA
				largerID = teamB
			} else {
				smallerID = teamB
				largerID = teamA
			}

			pairKey := fmt.Sprintf("%d-%d", smallerID, largerID)

			// todo make sure the teams haven't already been scheduled
			if processedPairs[pairKey] {
				continue
			}

			processedPairs[pairKey] = true
			// need to pull in TeamRivalryPriorities
			teamA := GetTeamByTeamID(strconv.Itoa(int(teamA)))
			teamB := GetTeamByTeamID(strconv.Itoa(int(teamB)))

			if teamA.Conference == teamB.Conference {
				processedPairs[pairKey] = true // since we dont want to schedule in conference matchups
				continue
			}

			mutualWeeks := sm.findMutualAvailableWeeks(teamA, teamB, updatedMaps)

			if len(mutualWeeks) == 0 {
				log.Printf("No available weeks for rivalry between %s and %s", teamA.TeamName, teamB.TeamName)
				continue
			}

			// FIXME placeholder for now
			optimalWeek := slices.Max(mutualWeeks)
			// determine home vs away and update maps
			// create game and append to rivalry games
			// update maps
			updatedMaps.ContendingMap[optimalWeek][teamA.ID] = true
			updatedMaps.ContendingMap[optimalWeek][teamB.ID] = true
			updatedMaps.OpenWeeksMap[teamA.ID] = removeWeekFromSlice(updatedMaps.OpenWeeksMap[teamA.ID], optimalWeek)
			updatedMaps.OpenWeeksMap[teamB.ID] = removeWeekFromSlice(updatedMaps.OpenWeeksMap[teamB.ID], optimalWeek)
			log.Printf("Scheduled rivalry for %s vs %s - Week %d", teamA.TeamName, teamB.TeamName, optimalWeek)
		}
	}
	// PSEUDOCODE: Process each rivalry relationship
	// processedPairs := CREATE empty set to avoid double-scheduling
	// FOR teamA, rivalTeams := range rivalries:
	//     FOR each rivalTeamB in rivalTeams:
	//         IF pair already processed THEN continue
	//         MARK pair as processed
	//
	//         availableWeeksA := GET open weeks for teamA
	//         availableWeeksB := GET open weeks for rivalTeamB
	//         mutualWeeks := FIND intersection of availableWeeksA and availableWeeksB
	//
	//         IF no mutual weeks THEN
	//             LOG warning about unschedulable rivalry
	//             continue
	//
	//         optimalWeek := CALL selectOptimalRivalryWeek(mutualWeeks, teamA, rivalTeamB)
	//         homeTeam, awayTeam := CALL determineRivalryHomeAway(teamA, rivalTeamB, seasonID)
	//
	//         rivalryGame := CREATE CollegeGame{
	//             HomeTeamID: homeTeam,
	//             AwayTeamID: awayTeam,
	//             WeekID: optimalWeek,
	//             SeasonID: seasonID,
	//             IsRivalry: true,
	//             IsConference: false
	//         }
	//
	//         ADD rivalryGame to rivalryGames
	//         UPDATE updatedMaps to mark teams as busy in optimalWeek
	//         REMOVE optimalWeek from both teams' available weeks

	// TODO: Implement the actual rivalry scheduling logic here
	// This is where you'll add the specific business rules for:
	// - Determining home/away rotation
	// - Selecting optimal weeks (traditional rivalry week preferences)
	// - Handling conflicts when multiple rivalries compete for the same week

	log.Printf("Rivalry scheduling logic placeholder - implement based on your specific rules")

	return rivalryGames, updatedMaps, nil
}

// fillRemainingNonConferenceSlots handles the bulk scheduling of non-rivalry games
// This uses a constraint satisfaction approach to fill remaining open slots
//
// PSEUDOCODE ALGORITHM:
// 1. FOR each team in priority order:
// 2.   CALCULATE how many more games the team needs
// 3.   WHILE team needs more games:
// 4.     FIND compatible opponents (different conference, FBS/FCS rules, not already played)
// 5.     EVALUATE each opponent for optimal matchup (geographic, competitive balance)
// 6.     SELECT best available opponent
// 7.     FIND mutual available week for both teams
// 8.     DETERMINE home/away based on team balance needs
// 9.     CREATE and schedule the game
// 10.    UPDATE availability maps
func (sm *SchedulingManager) fillRemainingNonConferenceSlots(
	prioritizedTeams []structs.CollegeTeam,
	schedulingMaps structs.SchedulingMaps,
	constraints structs.SchedulingConstraints,
	seasonID int,
) ([]structs.CollegeGame, structs.SchedulingMaps, error) {

	var newGames []structs.CollegeGame
	updatedMaps := schedulingMaps

	// PSEUDOCODE: Main scheduling loop
	// FOR each team in prioritizedTeams:
	//     currentGameCount := COUNT games already scheduled for team
	//     gamesNeeded := constraints.MaxGamesPerTeam - currentGameCount
	//
	//     IF gamesNeeded <= 0 THEN continue to next team
	//
	//     FOR attempt := 1 to gamesNeeded:
	//         compatibleOpponents := CALL FilterCompatibleOpponents(team, prioritizedTeams, updatedMaps, constraints)
	//         IF no compatible opponents THEN break out of attempts loop
	//
	//         bestOpponent := CALL selectBestOpponent(team, compatibleOpponents, updatedMaps)
	//         mutualWeeks := CALL findMutualAvailableWeeks(team, bestOpponent, updatedMaps)
	//         IF no mutual weeks THEN continue to next opponent
	//
	//         optimalWeek := CALL selectOptimalWeek(mutualWeeks, team, bestOpponent)
	//         homeTeam, awayTeam := CALL determineHomeAway(team, bestOpponent, updatedMaps)
	//
	//         newGame := CREATE CollegeGame with the determined parameters
	//         ADD newGame to newGames
	//         UPDATE updatedMaps with new game assignment
	//         REMOVE week from availability for both teams
	//         MARK opponents as having played each other

	// TODO: Implement the actual slot-filling algorithm
	// Key considerations for your implementation:
	// - Geographic preferences (minimize travel costs)
	// - Competitive balance (avoid mismatched games)
	// - Television scheduling windows
	// - Stadium capacity and revenue considerations
	// - Conference strength of schedule requirements

	log.Printf("Non-conference slot filling logic placeholder - implement based on your specific rules")

	return newGames, updatedMaps, nil
}

// SaveScheduleToCSV exports the generated schedule to a CSV file for download
// This integrates with your existing CSV export infrastructure
//
// Go Learning: This demonstrates integration with existing codebase patterns
// Look for similar ExportToCSV functions in your codebase to maintain consistency
func (sm *SchedulingManager) SaveScheduleToCSV(games []structs.CollegeGame, filepath string) error {
	// PSEUDOCODE: CSV Export Process
	// csvFile := CREATE new CSV file at filepath
	// writer := CREATE CSV writer for csvFile
	//
	// WRITE header row: ["Week", "Home Team", "Away Team", "Conference Game", "Rivalry", "Season"]
	//
	// FOR each game in games:
	//     homeTeam := LOOKUP team name by game.HomeTeamID
	//     awayTeam := LOOKUP team name by game.AwayTeamID
	//     row := [game.WeekID, homeTeam, awayTeam, game.IsConference, game.IsRivalry, game.SeasonID]
	//     WRITE row to CSV
	//
	// CLOSE csvFile
	// RETURN success/error

	// TODO: Implement CSV export using your existing patterns
	// Look at other CSV export functions in your codebase (like ExportRosterToCSV)
	// to maintain consistency with your existing file export architecture

	log.Printf("CSV export logic placeholder - filepath: %s, games: %d", filepath, len(games))

	return nil
}

// Additional helper methods you'll need to implement:

// selectBestOpponent evaluates potential opponents and returns the optimal matchup
// Consider factors like geographic distance, competitive balance, historical matchups
func (sm *SchedulingManager) selectBestOpponent(
	team structs.CollegeTeam,
	candidates []structs.CollegeTeam,
	maps structs.SchedulingMaps,
) structs.CollegeTeam {
	// TODO: Implement opponent selection algorithm
	// Factors to consider:
	// - Geographic proximity (reduce travel costs)
	// - Competitive balance (similar strength teams)
	// - Historical matchup interest
	// - Revenue potential (stadium capacity, fanbase size)
	// - Television scheduling preferences

	if len(candidates) > 0 {
		return candidates[0] // Placeholder - return first available
	}
	return structs.CollegeTeam{}
}

// findMutualAvailableWeeks returns weeks when both teams are available to play
func (sm *SchedulingManager) findMutualAvailableWeeks(
	teamA, teamB structs.CollegeTeam,
	maps structs.SchedulingMaps,
) []int {
	// Return intersection of both teams' available weeks

	openWeeksTeamA := maps.OpenWeeksMap[teamA.ID]
	openWeeksTeamB := maps.OpenWeeksMap[teamB.ID]

	var mutualAvailableWeeks []int

	for _, wA := range openWeeksTeamA {
		for _, wB := range openWeeksTeamB {
			if (wA == wB) && (wA != 0) {
				mutualAvailableWeeks = append(mutualAvailableWeeks, wA)
			}
		}
	}

	return mutualAvailableWeeks
}

// determineHomeAway decides which team should host based on balance needs
func (sm *SchedulingManager) determineHomeAway(
	teamA, teamB structs.CollegeTeam,
	maps structs.SchedulingMaps,
) (homeTeamID, awayTeamID uint) {
	// TODO: Implement home/away determination logic
	// Consider:
	// - Current home/away balance for each team
	// - Revenue sharing agreements
	// - Stadium capacity differences
	// - Traditional hosting patterns

	return teamA.ID, teamB.ID // Placeholder
}
