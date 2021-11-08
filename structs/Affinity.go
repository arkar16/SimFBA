package structs

import "github.com/jinzhu/gorm"

type Affinity struct {
	gorm.Model
	AffinityName string
}
