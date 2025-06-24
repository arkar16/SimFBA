package models

import (
	"gorm.io/gorm"
)

type PreDraftEvent struct {
	Name      string
	IsCombine bool
	// List of draftees participating
	Participants []NFLDraftee
	// Event Results struct
	Results []EventResults
}

type EventResults struct {
	gorm.Model
	Name             string
	IsCombine        bool
	PlayerID         uint
	FourtyYardDash   float32
	Shuttle          float32
	ThreeCone        float32
	VerticalJump     uint8
	BroadJump        uint8
	BenchPress       uint8
	ThrowingAccuracy float32
	ThrowingDistance float32
	Catching         float32
	RouteRunning     float32
	InsideRun        float32
	OutsideRun       float32
	RunBlocking      float32
	PassBlocking     float32
	PassRush         float32
	RunStop          float32
	LBCoverage       float32
	ManCoverage      float32
	ZoneCoverage     float32
	Kickoff          float32
	Fieldgoal        float32
	PuntDistance     float32
	CoffinPunt       float32
}
