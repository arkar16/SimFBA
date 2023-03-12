package structs

import "github.com/jinzhu/gorm"

type NFLContract struct {
	gorm.Model
	PlayerID        int
	NFLPlayerID     int
	TeamID          uint
	Team            string
	OriginalTeamID  uint
	OriginalTeam    string
	ContractLength  int
	Y1BaseSalary    float64
	Y1Bonus         float64
	Y2BaseSalary    float64
	Y2Bonus         float64
	Y3BaseSalary    float64
	Y3Bonus         float64
	Y4BaseSalary    float64
	Y4Bonus         float64
	Y5BaseSalary    float64
	Y5Bonus         float64
	BonusPercentage float64
	ContractType    string // Pro Bowl, Starter, Veteran, New ?
	ContractValue   float64
	IsActive        bool
	IsComplete      bool
	IsExtended      bool
}

func (c *NFLContract) DeactivateContract() {
	c.IsActive = false
}

func (c *NFLContract) TradePlayer(TeamID uint, Team string) {
	c.TeamID = TeamID
	c.Team = Team
}

func (c *NFLContract) ProgressContract() {
	c.Y1BaseSalary = c.Y2BaseSalary
	c.Y1Bonus = c.Y2Bonus
	c.Y2BaseSalary = c.Y3BaseSalary
	c.Y2Bonus = c.Y3Bonus
	c.Y3BaseSalary = c.Y4BaseSalary
	c.Y3Bonus = c.Y4Bonus
	c.Y4BaseSalary = c.Y5BaseSalary
	c.Y4Bonus = c.Y5Bonus
	c.Y5BaseSalary = 0
	c.Y5Bonus = 0

	if c.Y1BaseSalary == 0 && c.Y1Bonus == 0 {
		c.IsComplete = true
		c.DeactivateContract()
	}
}

func (c *NFLContract) CalculateContract() {
	// Calculate Value
	y1SalaryVal := c.Y1BaseSalary * 0.8
	y1BonusVal := c.Y1Bonus * 1
	y2SalaryVal := c.Y2BaseSalary * 0.4
	y2BonusVal := c.Y2Bonus * 0.9
	y3SalaryVal := c.Y3BaseSalary * 0.2
	y3BonusVal := c.Y3Bonus * 0.8
	y4SalaryVal := c.Y4BaseSalary * 0.1
	y4BonusVal := c.Y4Bonus * 0.7
	y5SalaryVal := c.Y5BaseSalary * 0.05
	y5BonusVal := c.Y5Bonus * 0.6
	c.ContractValue = y1SalaryVal + y1BonusVal + y2SalaryVal + y2BonusVal + y3SalaryVal + y3BonusVal + y4SalaryVal + y4BonusVal + y5SalaryVal + y5BonusVal
}
