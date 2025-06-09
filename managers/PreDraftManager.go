package managers

import (
	"github.com/CalebRose/SimFBA/structs"
)

func RunPreDraftEvents() {
	draftees := GetAllNFLDraftees()

	// Create a list of events (Pro Days and Combines)
	eventList := GenerateTypicalListOfEvents()

	event

	// HOW TO GET EVENTID FOR THE FUTURE???

	// Create a list of players invited to each event (HOW?)
	// For each event, create a result for each player
	// For some % of draftees, create results based on their advertised grades, not their real grades.
}

func GenerateTypicalListOfEvents() []structs.PreDraftEvent {
	var tempList []structs.PreDraftEvent
	var tempEvent structs.PreDraftEvent
	tempEvent.Name = "AAC Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	// DO ALL OTHER TYPICAL EVENTS!!!

	return tempList
}
