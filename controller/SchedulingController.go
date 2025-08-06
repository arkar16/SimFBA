package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/gorilla/mux"
)

// =============================================================================
// SCHEDULING CONTROLLER - HTTP API ENDPOINTS
// =============================================================================
// This file contains the HTTP handlers for the scheduling system.
// It follows your existing controller patterns and integrates with the
// SchedulingManager for business logic processing.
//
// Go Learning: This demonstrates the Controller pattern - HTTP request handling
// that delegates to manager classes for business logic processing.

// GenerateNonConferenceSchedule handles the HTTP request to generate a new non-conference schedule
// This endpoint processes the schedule generation request and returns either JSON data or CSV download
//
// API ENDPOINT: GET /api/scheduling/generate/nonconference/{seasonID}
// QUERY PARAMETERS:
//   - download=true: Returns CSV file download instead of JSON response
//   - format=csv: Alternative way to specify CSV download
//
// PSEUDOCODE FLOW:
//  1. EXTRACT seasonID from URL path parameter
//  2. VALIDATE seasonID is a valid integer
//  3. CREATE new SchedulingManager instance
//  4. CALL SchedulingManager.GenerateOptimalSchedule(seasonID)
//  5. IF error occurred THEN return HTTP error response
//  6. IF download requested THEN
//     a. GENERATE CSV file from schedule data
//     b. SET appropriate HTTP headers for file download
//     c. STREAM CSV data to client
//  7. ELSE return JSON response with schedule data
func GenerateNonConferenceSchedule(w http.ResponseWriter, r *http.Request) {
	// Go Learning: mux.Vars extracts path parameters from gorilla/mux router
	vars := mux.Vars(r)
	seasonIDStr := vars["seasonID"]

	// Convert seasonID string to integer with error handling
	// Go Learning: strconv.Atoi returns (int, error) - classic Go error handling pattern
	seasonID, err := strconv.Atoi(seasonIDStr)
	log.Printf("Season ID: %d", seasonID)
	if err != nil {
		// Return HTTP 400 Bad Request for invalid seasonID
		http.Error(w, fmt.Sprintf("Invalid season ID: %s", seasonIDStr), http.StatusBadRequest)
		return
	}

	// PSEUDOCODE: Initialize scheduling manager with database connection
	// db := GET database connection from provider
	// schedulingManager := CREATE new SchedulingManager(db)

	db := dbprovider.GetInstance().GetDB()
	schedulingManager := managers.NewSchedulingManager(db)

	// PSEUDOCODE: Generate the optimal schedule
	// scheduleResponse := CALL schedulingManager.GenerateOptimalSchedule(seasonID)
	// IF error THEN
	//     LOG error details
	//     RETURN HTTP 500 Internal Server Error

	scheduleResponse, err := schedulingManager.GenerateOptimalSchedule(seasonID)
	if err != nil {
		// Log the detailed error for debugging
		fmt.Printf("Schedule generation failed for season %d: %v\n", seasonID, err)

		// Return user-friendly error message
		http.Error(w, fmt.Sprintf("Failed to generate schedule for season %d", seasonID),
			http.StatusInternalServerError)
		return
	}

	// Check if client requested CSV download
	// PSEUDOCODE: Check for download request parameters
	// downloadRequested := CHECK query parameter "download" == "true" OR "format" == "csv"
	// IF downloadRequested THEN
	//     CALL handleCSVDownload(scheduleResponse, seasonID)
	//     RETURN

	downloadRequested := r.URL.Query().Get("download") == "true" ||
		r.URL.Query().Get("format") == "csv"

	if downloadRequested {
		handleScheduleCSVDownload(w, scheduleResponse, seasonID)
		return
	}

	// PSEUDOCODE: Return JSON response for non-download requests
	// SET Content-Type header to "application/json"
	// ENCODE scheduleResponse as JSON
	// WRITE JSON response to client

	// Go Learning: Setting HTTP headers before writing response body
	w.Header().Set("Content-Type", "application/json")

	// TODO: You'll need to import "encoding/json" and use json.NewEncoder
	// For now, return a simple success message
	_, err = fmt.Fprintf(w, `{
			"success": true,
			"season_id": %d,
			"message": "Schedule generated successfully",
			"games_count": %d,
			"weeks_scheduled": %d
		}`, seasonID, getTotalGamesCount(scheduleResponse), len(scheduleResponse.ScheduledGamesByWeekMap))

	if err != nil {
		return
	}
}

// handleScheduleCSVDownload processes CSV file download requests
// This function converts the schedule data to CSV format and streams it to the client
//
// PSEUDOCODE ALGORITHM:
// 1. EXTRACT all games from the schedule response
// 2. GENERATE appropriate filename with timestamp
// 3. SET HTTP headers for file download (Content-Disposition, Content-Type)
// 4. CREATE CSV writer and write header row
// 5. FOR each game, write game data as CSV row
// 6. FLUSH and close the CSV writer
func handleScheduleCSVDownload(w http.ResponseWriter, scheduleResponse *structs.ScheduleResponse, seasonID int) {
	// PSEUDOCODE: Collect all games from the schedule
	// allGames := CREATE empty slice
	// FOR weekID, games := range scheduleResponse.ScheduledGamesByWeekMap:
	//     allGames = APPEND allGames, games...

	var allGames []structs.CollegeGame
	for _, games := range scheduleResponse.ScheduledGamesByWeekMap {
		allGames = append(allGames, games...)
	}

	// PSEUDOCODE: Generate filename with current timestamp
	// currentTime := GET current time
	// filename := FORMAT "nonconference_schedule_{seasonID}_{timestamp}.csv"

	filename := fmt.Sprintf("nonconference_schedule_season_%d.csv", seasonID)

	// PSEUDOCODE: Set HTTP headers for file download
	// SET "Content-Type" header to "text/csv"
	// SET "Content-Disposition" header to "attachment; filename={filename}"
	// SET "Cache-Control" header to "no-cache"

	// Go Learning: HTTP headers must be set before writing response body
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Cache-Control", "no-cache")

	// PSEUDOCODE: Write CSV data
	// csvWriter := CREATE new CSV writer for response writer
	// WRITE header row: ["Week", "Home Team", "Away Team", "Is Conference", "Is Rivalry", "Season"]
	//
	// FOR each game in allGames:
	//     homeTeamName := LOOKUP team name by game.HomeTeamID (TODO: implement lookup)
	//     awayTeamName := LOOKUP team name by game.AwayTeamID (TODO: implement lookup)
	//     row := [game.WeekID, homeTeamName, awayTeamName, game.IsConference, game.IsRivalry, game.SeasonID]
	//     WRITE row to CSV
	//
	// FLUSH csvWriter

	// TODO: Implement actual CSV writing logic
	// You'll need to:
	// 1. Import "encoding/csv"
	// 2. Create csv.NewWriter(w)
	// 3. Implement team name lookup (probably via repository)
	// 4. Write actual CSV rows

	// For now, write a simple CSV placeholder
	csvContent := "Week,Home Team,Away Team,Is Conference,Is Rivalry,Season\n"
	for _, game := range allGames {
		// TODO: Replace with actual team name lookups
		csvContent += fmt.Sprintf("%d,Team_%d,Team_%d,%t,%t,%d\n",
			game.WeekID, game.HomeTeamID, game.AwayTeamID,
			game.IsConference, game.IsRivalryGame, game.SeasonID)
	}

	// Write the CSV content
	_, err := w.Write([]byte(csvContent))
	if err != nil {
		return
	}
}

// GetSchedulingStatus provides information about the current scheduling system status
// This can be useful for debugging and monitoring the scheduling system
//
// API ENDPOINT: GET /api/scheduling/status
//
// PSEUDOCODE RESPONSE:
//
//	{
//	  "system_status": "operational",
//	  "available_seasons": [2024, 2025, 2026],
//	  "scheduling_constraints": { ... },
//	  "last_generated": "2024-01-15T10:30:00Z"
//	}
func GetSchedulingStatus(w http.ResponseWriter, r *http.Request) {
	// PSEUDOCODE: Gather system status information
	// constraints := GET default scheduling constraints
	// availableSeasons := QUERY database for available seasons
	// lastGenerated := GET timestamp of most recent schedule generation
	// systemHealth := CHECK if all components are working

	w.Header().Set("Content-Type", "application/json")

	// TODO: Implement actual status checking logic
	// This should query your database and check system health

	fmt.Fprintf(w, `{
		"system_status": "operational",
		"message": "Scheduling system is ready",
		"available_features": [
			"non_conference_scheduling",
			"rivalry_preservation", 
			"constraint_validation",
			"csv_export"
		],
		"default_constraints": {
			"max_games_per_team": 12,
			"max_home_games": 7,
			"max_away_games": 7,
			"prevent_back_to_back_byes": true,
			"enforce_rivalries": true
		}
	}`)
}

// Helper functions

// getTotalGamesCount calculates the total number of games in a schedule response
func getTotalGamesCount(response *structs.ScheduleResponse) int {
	total := 0
	for _, games := range response.ScheduledGamesByWeekMap {
		total += len(games)
	}
	return total
}

// TODO: You'll need to add these endpoints to your main.go router:
//
// PSEUDOCODE: Add to your main.go handleRequests() function:
// apiRouter.HandleFunc("/scheduling/generate/nonconference/{seasonID}", controller.GenerateNonConferenceSchedule).Methods("GET")
// apiRouter.HandleFunc("/scheduling/status", controller.GetSchedulingStatus).Methods("GET")
//
// Go Learning: The {seasonID} syntax creates a path parameter that can be extracted with mux.Vars()
// The .Methods("GET") restricts this endpoint to only handle GET requests
