package repository

import (
	"fmt"
	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
)

type SchedulingRepository struct {
	db *gorm.DB
}

func NewSchedulingRepository(db *gorm.DB) *SchedulingRepository {
	return &SchedulingRepository{db: db}
}

func (r *SchedulingRepository) GetExistingConferenceGames(seasonID uint) ([]structs.CollegeGame, error) {
	var games []structs.CollegeGame
	err := r.db.Where("season_id = ? AND is_conference = ?", int(seasonID), true).Find(&games).Error
	return games, err
}

// TODO might not need this function

func (r *SchedulingRepository) GetExistingNonConGames(seasonID uint) ([]structs.CollegeGame, error) {
	var games []structs.CollegeGame
	err := r.db.Where("season_id = ? and is_conference = ?", seasonID, false).Find(&games).Error
	return games, err
}

func (r *SchedulingRepository) GetAllTeams(seasonID uint) ([]structs.CollegeTeam, error) {
	var teams []structs.CollegeTeam
	err := r.db.Where("is_active = ?", true).Find(&teams).Error
	return teams, err
}

func (r *SchedulingRepository) GetRivalries() (map[uint][]uint, error) {
	// TODO: Update this to use the correct rivalry struct
	// Currently returning placeholder data since CollegeRival struct needs to be defined

	// PSEUDOCODE: Query rivalry relationships from database
	// var rivalries []structs.CollegeRival
	// teamRivals := CREATE empty map[int][]int
	//
	// err := QUERY database for all rivalry relationships
	// IF error THEN return nil, error
	//
	// FOR each rivalry in rivalries:
	//     ADD rivalry.TeamTwoID to teamRivals[rivalry.TeamOneID]
	//     ADD rivalry.TeamOneID to teamRivals[rivalry.TeamTwoID] (bidirectional)
	//
	// RETURN teamRivals, nil

	// Load all rivalry relationships from database
	var rivalries []structs.CollegeRival
	err := r.db.Find(&rivalries).Error

	if err != nil {
		return nil, fmt.Errorf("failed to load rivalries: %w", err)
	}

	teamRivals := make(map[uint][]uint)

	// Build bidirectional rivalry map with priority tracking
	for _, rivalry := range rivalries {
		// Add TeamTwoID to TeamOneID's rivalry list
		teamRivals[rivalry.TeamOneID] = append(teamRivals[rivalry.TeamOneID], rivalry.TeamTwoID)
		
		// Add TeamOneID to TeamTwoID's rivalry list (bidirectional)
		teamRivals[rivalry.TeamTwoID] = append(teamRivals[rivalry.TeamTwoID], rivalry.TeamOneID)
	}

	return teamRivals, nil
}

// GetRivalryPriorities returns a map of team rivalries with their priority levels
// This is useful for scheduling decisions where some rivalries have higher importance
func (r *SchedulingRepository) GetRivalryPriorities() (map[uint]map[uint]uint, error) {
	var rivalries []structs.CollegeRival
	err := r.db.Find(&rivalries).Error

	if err != nil {
		return nil, fmt.Errorf("failed to load rivalry priorities: %w", err)
	}

	// teamID -> rivalTeamID -> priority
	teamRivalPriorities := make(map[uint]map[uint]uint)

	for _, rivalry := range rivalries {
		// Initialize maps if they don't exist
		if teamRivalPriorities[rivalry.TeamOneID] == nil {
			teamRivalPriorities[rivalry.TeamOneID] = make(map[uint]uint)
		}
		if teamRivalPriorities[rivalry.TeamTwoID] == nil {
			teamRivalPriorities[rivalry.TeamTwoID] = make(map[uint]uint)
		}

		// Set priorities for both teams
		teamRivalPriorities[rivalry.TeamOneID][rivalry.TeamTwoID] = rivalry.TeamOnePriority
		teamRivalPriorities[rivalry.TeamTwoID][rivalry.TeamOneID] = rivalry.TeamTwoPriority
	}

	return teamRivalPriorities, nil
}

// GetTeamNameByID retrieves the team name for a given team ID
// This is used for CSV export functionality to show team names instead of IDs
func (r *SchedulingRepository) GetTeamNameByID(teamID uint) (string, error) {
	var team structs.CollegeTeam
	err := r.db.Select("team_name").Where("id = ?", teamID).First(&team).Error
	if err != nil {
		return "", err
	}
	return team.TeamName, nil
}

// GetAvailableSeasons returns a list of seasons that have scheduling data
// This is useful for the scheduling status endpoint
func (r *SchedulingRepository) GetAvailableSeasons() ([]int, error) {
	var seasons []int

	// PSEUDOCODE: Query distinct seasons from games or teams table
	// err := QUERY "SELECT DISTINCT season_id FROM college_games ORDER BY season_id DESC"
	// IF error THEN return nil, error
	// RETURN seasons, nil

	err := r.db.Model(&structs.CollegeGame{}).
		Distinct("season_id").
		Order("season_id DESC").
		Pluck("season_id", &seasons).Error

	return seasons, err
}

// SaveScheduleToCSV saves the generated schedule to a CSV file
// This method handles the file I/O operations for schedule export
func (r *SchedulingRepository) SaveScheduleToCSV(games []structs.CollegeGame, filepath string) error {
	// PSEUDOCODE: CSV File Creation
	// csvFile := CREATE new file at filepath
	// IF error creating file THEN return error
	//
	// csvWriter := CREATE CSV writer for csvFile
	//
	// WRITE header row: ["Week", "Home Team", "Away Team", "Conference Game", "Rivalry", "Season"]
	//
	// FOR each game in games:
	//     homeTeamName := CALL GetTeamNameByID(game.HomeTeamID)
	//     awayTeamName := CALL GetTeamNameByID(game.AwayTeamID)
	//
	//     row := [
	//         CONVERT game.WeekID to string,
	//         homeTeamName,
	//         awayTeamName,
	//         CONVERT game.IsConference to string,
	//         CONVERT game.IsRivalry to string,
	//         CONVERT game.SeasonID to string
	//     ]
	//
	//     WRITE row to CSV
	//     IF error THEN return error
	//
	// FLUSH and CLOSE csvWriter
	// CLOSE csvFile
	// RETURN nil

	// TODO: Implement actual CSV file creation
	// You'll need to:
	// 1. Import "encoding/csv" and "os"
	// 2. Create the file at the specified filepath
	// 3. Write CSV headers and data rows
	// 4. Handle file cleanup and error cases

	// Placeholder implementation
	return nil
}
