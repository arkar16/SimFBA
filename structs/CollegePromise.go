package structs

import "gorm.io/gorm"

type CollegePromise struct {
	gorm.Model
	TeamID          uint
	CollegePlayerID uint
	PromiseType     string // Snaps, Wins, Bowl Game, Conf Championship, Playoffs, National Championship, Scheme Change
	PromiseWeight   string
	IsFullfilled    bool
}
