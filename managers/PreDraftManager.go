package managers

import (
	"github.com/CalebRose/SimFBA/structs"
)

func RunPreDraftEvents() {
	draftees := GetAllNFLDraftees()

	// Create a list of events (Pro Days and Combines)
	eventList := GenerateTypicalListOfEvents()

	event

	// Create a list of players invited to each event (Read in json)
	// For each event, create a result for each player
	// For some % of draftees, create results based on their advertised grades, not their real grades.
}

func GetStartEventID() uint {
	return 0
}

func GenerateTypicalListOfEvents() []structs.PreDraftEvent {
	// GET EVENT ID
	var eventID = GetStartEventID()

	var tempList []structs.PreDraftEvent
	var tempEvent structs.PreDraftEvent
	tempEvent.Name = "AAC Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "ACC Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "Big Ten Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "Big XII Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "C-USA Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "MAC Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "MWC Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "SEC Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "Sun Belt Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "FCS Pro Day"
	tempEvent.IsCombine = false
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	tempEvent.Name = "NFL Combine"
	tempEvent.IsCombine = true
	tempEvent.EventID = eventID
	tempList = append(tempList, tempEvent)
	eventID++

	return tempList
}
