package managers

import (
	"errors"
	"fmt"
	"sort"

	"github.com/CalebRose/SimFBA/structs"
)

// =============================================================================
// SCHEDULING HELPER FUNCTIONS
// =============================================================================
// This file contains utility functions for the college football scheduling system.
// Each function is designed to handle a specific aspect of constraint-based scheduling.

// BuildContendingMap creates a comprehensive scheduling map from existing games
// This is the foundation of the scheduling system - it tracks which teams are
// available during which weeks based on their existing conference games.
//
// Go Learning Note: We return both a result and an error (Go's standard error handling pattern)
// The function signature uses named return values for clarity
func BuildContendingMap(existingGames []structs.CollegeGame, teams []structs.CollegeTeam, numWeeks int) (structs.SchedulingMaps, error) {
	// Initialize the SchedulingMaps struct with all required maps
	// Go Learning: make() creates maps - the key difference from structs is that maps are reference types
	sm := structs.SchedulingMaps{
		ContendingMap:         make(map[int]map[uint]bool),        // week -> teamID -> is_busy
		CFBTeamMap:            make(map[uint]structs.CollegeTeam), // teamID -> team_data
		FBSPlayedFCSMap:       make(map[uint]bool),                // FBS teams that played FCS
		FCSPlayedFBSMap:       make(map[uint]bool),                // FCS teams that played FBS
		TeamRivalries:         make(map[uint][]uint),              // teamID -> rival_team_ids
		TeamRivalryPriorities: make(map[uint]map[uint]uint),       // teamID -> rivalTeamID -> priority
		HomeGamesMap:          make(map[uint][]int),               // teamID -> home_game_weeks
		AwayGamesMap:          make(map[uint][]int),               // teamID -> away_game_weeks
		ByeGamesMap:           make(map[uint][]int),               // teamID -> bye_weeks
		OpenWeeksMap:          make(map[uint][]int),               // teamID -> available_weeks
		OpponentHistoryMap:    make(map[uint]map[uint]bool),       // teamA -> teamB -> played_already
	}

	// Step 1: Initialize all weeks with empty availability maps
	// Go Learning: range creates a loop from 0 to numWeeks-1
	for week := 0; week < numWeeks; week++ {
		sm.ContendingMap[week] = make(map[uint]bool)
	}

	// Step 2: Populate the team map for quick lookups
	// Go Learning: range over slices gives index, value. We ignore index with _
	for _, team := range teams {
		sm.CFBTeamMap[team.ID] = team

		// Initialize opponent history map for this team
		sm.OpponentHistoryMap[team.ID] = make(map[uint]bool)

		// Initialize all weeks as open initially
		openWeeks := make([]int, 0, numWeeks)
		for week := 0; week < numWeeks; week++ {
			openWeeks = append(openWeeks, week)
		}
		sm.OpenWeeksMap[team.ID] = openWeeks
	}

	// Step 3: Process existing games to mark teams as busy
	for _, game := range existingGames {
		weekID := int(game.WeekID) // Convert to int for map key

		if weekID >= numWeeks || weekID < 0 {
			// TODO add debug here
			continue
		}
		// Mark both teams as busy this week
		sm.ContendingMap[weekID][uint(game.HomeTeamID)] = true
		sm.ContendingMap[weekID][uint(game.AwayTeamID)] = true

		// Track home/away games
		sm.HomeGamesMap[uint(game.HomeTeamID)] = append(sm.HomeGamesMap[uint(game.HomeTeamID)], weekID)
		sm.AwayGamesMap[uint(game.AwayTeamID)] = append(sm.AwayGamesMap[uint(game.AwayTeamID)], weekID)

		// Track opponent history (prevent duplicate matchups)
		sm.OpponentHistoryMap[uint(game.HomeTeamID)][uint(game.AwayTeamID)] = true
		sm.OpponentHistoryMap[uint(game.AwayTeamID)][uint(game.HomeTeamID)] = true

		// Remove this week from open weeks for both teams
		sm.OpenWeeksMap[uint(game.HomeTeamID)] = removeWeekFromSlice(sm.OpenWeeksMap[uint(game.HomeTeamID)], weekID)
		sm.OpenWeeksMap[uint(game.AwayTeamID)] = removeWeekFromSlice(sm.OpenWeeksMap[uint(game.AwayTeamID)], weekID)

		// Track FBS/FCS crossover games
		homeTeam := sm.CFBTeamMap[uint(game.HomeTeamID)]
		awayTeam := sm.CFBTeamMap[uint(game.AwayTeamID)]

		if homeTeam.IsFBS && !awayTeam.IsFBS {
			sm.FBSPlayedFCSMap[uint(game.HomeTeamID)] = true
			sm.FCSPlayedFBSMap[uint(game.AwayTeamID)] = true
		} else if !homeTeam.IsFBS && awayTeam.IsFBS {
			sm.FCSPlayedFBSMap[uint(game.HomeTeamID)] = true
			sm.FBSPlayedFCSMap[uint(game.AwayTeamID)] = true
		}
	}

	// Step 4: Calculate bye weeks for each team
	// TODO: You'll implement this logic based on your bye week rules
	for teamID := range sm.CFBTeamMap {
		byeWeeks := calculateByeWeeks(teamID, sm.OpenWeeksMap[teamID], numWeeks)
		sm.ByeGamesMap[teamID] = byeWeeks
	}

	return sm, nil
}

// removeWeekFromSlice removes a specific week from a slice of weeks
// Go Learning: This demonstrates slice manipulation - a common Go pattern
func removeWeekFromSlice(weeks []int, weekToRemove int) []int {
	result := make([]int, 0, len(weeks))
	for _, week := range weeks {
		if week != weekToRemove {
			result = append(result, week)
		}
	}
	return result
}

// calculateByeWeeks determines which weeks should be bye weeks for a team
// TODO: Implement your bye week logic here based on your game requirements
func calculateByeWeeks(teamID uint, openWeeks []int, totalWeeks int) []int {
	// Placeholder implementation - you'll need to add your specific logic
	// This could involve checking team preferences, conference requirements, etc.
	var byeWeeks []int

	// Example logic: if team has fewer than expected games, some open weeks become byes
	expectedGames := 12 // Adjust based on your rules
	gamesScheduled := totalWeeks - len(openWeeks)

	if gamesScheduled < expectedGames {
		// Need more games, fewer byes
		maxByes := len(openWeeks) - (expectedGames - gamesScheduled)
		if maxByes > 0 && len(openWeeks) > maxByes {
			// Take some open weeks as byes (you can add more sophisticated logic here)
			byeWeeks = openWeeks[:maxByes]
		}
	}

	return byeWeeks
}

// ValidateScheduleConstraints checks if the current schedule violates any constraints
// This function is crucial for ensuring the generated schedule meets all requirements
//
// Go Learning: Multiple return values are common in Go - (result, error) pattern
func ValidateScheduleConstraints(games []structs.CollegeGame, constraints structs.SchedulingConstraints, teams []structs.CollegeTeam) ([]string, error) {
	violations := make([]string, 0) // Use slice for dynamic violations list

	// Create team lookup map for efficiency
	// Go Learning: Building lookup maps is a common optimization pattern
	teamMap := make(map[uint]structs.CollegeTeam)
	for _, team := range teams {
		teamMap[team.ID] = team
	}

	// Track games per team
	teamGameCounts := make(map[uint]int)
	teamHomeGames := make(map[uint]int)
	teamAwayGames := make(map[uint]int)
	teamGamesByWeek := make(map[uint]map[int]bool) // teamID -> week -> hasGame

	// Initialize tracking maps
	for _, team := range teams {
		teamGameCounts[team.ID] = 0
		teamHomeGames[team.ID] = 0
		teamAwayGames[team.ID] = 0
		teamGamesByWeek[team.ID] = make(map[int]bool)
	}

	// Process all games to count violations
	for _, game := range games {
		week := int(game.WeekID)

		// Count games per team
		teamGameCounts[uint(game.HomeTeamID)]++
		teamGameCounts[uint(game.AwayTeamID)]++

		// Count home/away games
		teamHomeGames[uint(game.HomeTeamID)]++
		teamAwayGames[uint(game.AwayTeamID)]++

		// Track games by week for each team
		teamGamesByWeek[uint(game.HomeTeamID)][week] = true
		teamGamesByWeek[uint(game.AwayTeamID)][week] = true
	}

	// Check constraint violations for each team
	for _, team := range teams {
		teamID := team.ID

		// Check max games per team
		if teamGameCounts[teamID] > constraints.MaxGamesPerTeam {
			violations = append(violations,
				fmt.Sprintf("Team %s exceeds max games: %d > %d",
					team.TeamName, teamGameCounts[teamID], constraints.MaxGamesPerTeam))
		}

		// Check max home games
		if teamHomeGames[teamID] > constraints.MaxHomeGames {
			violations = append(violations,
				fmt.Sprintf("Team %s exceeds max home games: %d > %d",
					team.TeamName, teamHomeGames[teamID], constraints.MaxHomeGames))
		}

		// Check max away games
		if teamAwayGames[teamID] > constraints.MaxAwayGames {
			violations = append(violations,
				fmt.Sprintf("Team %s exceeds max away games: %d > %d",
					team.TeamName, teamAwayGames[teamID], constraints.MaxAwayGames))
		}

		// Check for back-to-back bye weeks (if constraint is enabled)
		if constraints.PreventBackToBackByes {
			consecutiveByes := checkConsecutiveByes(teamID, teamGamesByWeek[teamID])
			if consecutiveByes {
				violations = append(violations,
					fmt.Sprintf("Team %s has back-to-back bye weeks", team.TeamName))
			}
		}

		// TODO: Add more constraint checks as needed
		// - Consecutive away games
		// - Conference scheduling conflicts
		// - Rivalry game requirements
	}

	return violations, nil
}

// checkConsecutiveByes identifies if a team has consecutive bye weeks
// Go Learning: This shows how to work with maps and implement business logic
func checkConsecutiveByes(teamID uint, gamesByWeek map[int]bool) bool {
	// Sort weeks to check them in order
	weeks := make([]int, 0, len(gamesByWeek))
	for week := range gamesByWeek {
		weeks = append(weeks, week)
	}
	sort.Ints(weeks) // Go Learning: sort package provides common sorting functions

	// Check for consecutive weeks without games
	for i := 0; i < len(weeks)-1; i++ {
		if weeks[i+1] == weeks[i]+1 {
			// Consecutive weeks with games, reset bye counter
		} else if weeks[i+1] == weeks[i]+2 {
			// One week gap (one bye week)
		} else if weeks[i+1] > weeks[i]+2 {
			// Multiple week gap
			gapSize := weeks[i+1] - weeks[i] - 1
			if gapSize >= 2 {
				return true // Found consecutive byes
			}
		}
	}

	return false
}

// CalculateTeamBalance analyzes home/away game distribution for a team
// This helps ensure fair scheduling and can influence future game assignments
func CalculateTeamBalance(teamID uint, games []structs.CollegeGame) structs.TeamBalance {
	var homeGames, awayGames int
	var homeWeeks, awayWeeks []int
	var longestHomeStreak, longestAwayStreak int
	var currentHomeStreak, currentAwayStreak int

	// Create a map of weeks to game types for this team
	weekToGameType := make(map[int]string) // "home", "away", or "bye"

	for _, game := range games {
		week := int(game.WeekID)
		if game.HomeTeamID == int(teamID) {
			homeGames++
			homeWeeks = append(homeWeeks, week)
			weekToGameType[week] = "home"
		} else if game.AwayTeamID == int(teamID) {
			awayGames++
			awayWeeks = append(awayWeeks, week)
			weekToGameType[week] = "away"
		}
	}

	// Sort weeks to analyze streaks
	allWeeks := append(homeWeeks, awayWeeks...)
	sort.Ints(allWeeks)

	// Calculate streaks
	for _, week := range allWeeks {
		gameType := weekToGameType[week]

		if gameType == "home" {
			currentHomeStreak++
			if currentHomeStreak > longestHomeStreak {
				longestHomeStreak = currentHomeStreak
			}
			currentAwayStreak = 0
		} else if gameType == "away" {
			currentAwayStreak++
			if currentAwayStreak > longestAwayStreak {
				longestAwayStreak = currentAwayStreak
			}
			currentHomeStreak = 0
		}
	}

	// Calculate balance ratio
	totalGames := homeGames + awayGames
	var homeRatio float64
	if totalGames > 0 {
		homeRatio = float64(homeGames) / float64(totalGames)
	}

	// TODO: Define TeamBalance struct in structs package
	return structs.TeamBalance{
		TeamID:            teamID,
		HomeGames:         homeGames,
		AwayGames:         awayGames,
		HomeRatio:         homeRatio,
		LongestHomeStreak: longestHomeStreak,
		LongestAwayStreak: longestAwayStreak,
		IsBalanced:        homeRatio >= 0.4 && homeRatio <= 0.6, // 40-60% is considered balanced
	}
}

// FindAvailableWeek finds the next available week for a team to play
// This is used during the scheduling process to assign games to open slots
func FindAvailableWeek(teamID uint, contendingMap map[int]map[uint]bool, startWeek, endWeek int) (int, error) {
	for week := startWeek; week <= endWeek; week++ {
		// Check if team is available this week
		if weekMap, exists := contendingMap[week]; exists {
			if !weekMap[teamID] { // false means team is available
				return week, nil
			}
		} else {
			// Week doesn't exist in map, team is available
			return week, nil
		}
	}

	// Go Learning: Custom error creation using errors.New()
	return -1, errors.New(fmt.Sprintf("No available week found for team %d between weeks %d-%d", teamID, startWeek, endWeek))
}

// SortTeamsBySchedulingPriority sorts teams based on their scheduling constraints
// Teams with more constraints (harder to schedule) should be scheduled first
func SortTeamsBySchedulingPriority(teams []structs.CollegeTeam, schedulingMaps structs.SchedulingMaps) []structs.CollegeTeam {
	// Create a copy to avoid modifying the original slice
	// Go Learning: make() with capacity optimization
	sortedTeams := make([]structs.CollegeTeam, len(teams))
	copy(sortedTeams, teams)

	// Sort using a custom comparison function
	// Go Learning: sort.Slice() with anonymous function (closure)
	sort.Slice(sortedTeams, func(i, j int) bool {
		teamA := sortedTeams[i]
		teamB := sortedTeams[j]

		// Priority factors (higher values = higher priority):
		// 1. Number of rivalry games to schedule
		// 2. Number of existing constraints
		// 3. Conference requirements

		priorityA := calculateSchedulingPriority(teamA, schedulingMaps)
		priorityB := calculateSchedulingPriority(teamB, schedulingMaps)

		return priorityA > priorityB // Higher priority teams first
	})

	return sortedTeams
}

// calculateSchedulingPriority determines how difficult a team will be to schedule
// Higher values indicate teams that should be scheduled first
func calculateSchedulingPriority(team structs.CollegeTeam, maps structs.SchedulingMaps) int {
	priority := 0

	// Add priority for rivalry games
	if rivalries, exists := maps.TeamRivalries[team.ID]; exists {
		priority += len(rivalries) * 10 // Each rivalry adds significant priority
	}

	// Add priority based on available weeks (fewer available = higher priority)
	if openWeeks, exists := maps.OpenWeeksMap[team.ID]; exists {
		priority += (15 - len(openWeeks)) * 5 // Fewer open weeks = higher priority
	}

	// Add priority for special constraints
	if team.IsFBS {
		priority += 2 // FBS teams have more complex scheduling requirements
	}

	// TODO: Add more priority factors based on your specific requirements
	// - Conference championship implications
	// - Television scheduling preferences
	// - Geographic considerations

	return priority
}

// FilterCompatibleOpponents returns teams that can be scheduled against the given team
// This enforces division rules, FBS/FCS limits, and other compatibility constraints
func FilterCompatibleOpponents(teamID uint, allTeams []structs.CollegeTeam, schedulingMaps structs.SchedulingMaps, constraints structs.SchedulingConstraints) []structs.CollegeTeam {
	var compatibleTeams []structs.CollegeTeam

	// Get the source team
	sourceTeam, exists := schedulingMaps.CFBTeamMap[teamID]
	if !exists {
		return compatibleTeams // Return empty if team not found
	}

	for _, potentialOpponent := range allTeams {
		// Skip self
		if potentialOpponent.ID == teamID {
			continue
		}

		// Skip if already played this season
		if alreadyPlayed, exists := schedulingMaps.OpponentHistoryMap[teamID][potentialOpponent.ID]; exists && alreadyPlayed {
			continue
		}

		// Skip same conference teams (assuming non-conference scheduling)
		if sourceTeam.ConferenceID == potentialOpponent.ConferenceID {
			continue
		}

		// Check FBS/FCS crossover limits
		if sourceTeam.IsFBS && !potentialOpponent.IsFBS {
			// FBS team trying to play FCS team
			if hasPlayedFCS, exists := schedulingMaps.FBSPlayedFCSMap[teamID]; exists && hasPlayedFCS {
				continue // Already played an FCS team this season
			}
		} else if !sourceTeam.IsFBS && potentialOpponent.IsFBS {
			// FCS team trying to play FBS team
			if hasPlayedFBS, exists := schedulingMaps.FCSPlayedFBSMap[teamID]; exists && hasPlayedFBS {
				continue // Already played an FBS team this season
			}
		}

		// TODO: Add more compatibility checks:
		// - Geographic constraints
		// - Television scheduling conflicts
		// - Academic calendar conflicts
		// - Stadium availability

		compatibleTeams = append(compatibleTeams, potentialOpponent)
	}

	return compatibleTeams
}
