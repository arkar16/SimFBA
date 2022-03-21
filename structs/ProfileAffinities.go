package structs

import "github.com/jinzhu/gorm"

type ProfileAffinity struct {
	gorm.Model
	AffinityID        int
	ProfileID         int
	AffinityName      string
	IsApplicable      bool
	IsDynamicAffinity bool
	IsCheckedWeekly   bool
	IsCheckedSeasonal bool
	AffinityValue     float32
}

func (pa *ProfileAffinity) ToggleApplicability() {
	pa.IsApplicable = !pa.IsApplicable
}
