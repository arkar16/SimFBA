package structs

import "github.com/jinzhu/gorm"

type NewsLog struct {
	gorm.Model
	WeekID      int
	Week        int
	SeasonID    int
	MessageType string
	Message     string
	League      string
}
