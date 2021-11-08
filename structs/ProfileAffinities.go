package structs

import "github.com/jinzhu/gorm"

type ProfileAffinity struct {
	gorm.Model
	AffinityID   int
	ProfileID    int
	AffinityName string
	IsApplicable bool
}

func (pa *ProfileAffinity) ToggleApplicability() {
	pa.IsApplicable = !pa.IsApplicable
}
