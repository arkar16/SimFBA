package structs

import "gorm.io/gorm"

type CollegePromise struct {
	gorm.Model
	TeamID          uint
	CollegePlayerID uint
	PromiseType     string // Snaps, Wins, Bowl Game, Conf Championship, Playoffs, National Championship, Scheme Change
	PromiseWeight   string // The impact the promise will have on their decision
	Benchmark       int    // The value that must be met
	PromiseMade     bool   // The player has agreed to the premise of the promise
	IsFullfilled    bool
}
