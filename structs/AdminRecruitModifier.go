package structs

import "github.com/jinzhu/gorm"

type AdminRecruitModifier struct {
	gorm.Model
	ModifierOne       int
	ModifierTwo       int
	WeeksOfRecruiting int
}

func (ARM *AdminRecruitModifier) SetModifierOne(val int) {
	ARM.ModifierOne = val
}

func (ARM *AdminRecruitModifier) SetModifierTwo(val int) {
	ARM.ModifierTwo = val
}

func (ARM *AdminRecruitModifier) SetWeek(val int) {
	ARM.WeeksOfRecruiting = val
}
