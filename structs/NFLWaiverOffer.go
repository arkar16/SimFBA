package structs

type NFLWaiverOffDTO struct {
	ID          uint
	NFLPlayerID uint
	TeamID      uint
	Team        string
	WaiverOrder uint
	IsActive    bool
}

type NFLWaiverOffer struct {
	ID          uint
	TeamID      uint
	Team        string
	WaiverOrder uint
	NFLPlayerID uint
	IsActive    bool
}

func (wo *NFLWaiverOffer) AssignID(id uint) {
	wo.ID = id
}

func (wo *NFLWaiverOffer) AssignNewWaiverOrder(val uint) {
	wo.WaiverOrder = val
}

func (wo *NFLWaiverOffer) Map(offer NFLWaiverOffDTO) {
	wo.TeamID = offer.TeamID
	wo.Team = offer.Team
	wo.NFLPlayerID = offer.NFLPlayerID
	wo.WaiverOrder = offer.WaiverOrder
	wo.IsActive = true
}

func (wo *NFLWaiverOffer) DeactivateWaiverOffer() {
	wo.IsActive = false
}
