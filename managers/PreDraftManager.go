package managers

import (
	"github.com/CalebRose/SimFBA/structs"
)

func RunPreDraftEvents() {
	draftees := GetAllNFLDraftees()

	// Create a list of events (Pro Days and Combines)
	eventList := GenerateTypicalListOfEvents()

	// Create a list of players invited to each event (Read in json)
	// For each event, go through each ID in the participant IDs and grab the corresponding draftee and add it to the event
	// For each event, create a result for each player
	// For some % of draftees, create results based on their advertised grades, not their real grades.
}

func GenerateTypicalListOfEvents() []structs.PreDraftEvent {
	var tempList []structs.PreDraftEvent
	var tempEvent structs.PreDraftEvent
	tempEvent.Name = "AAC Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "ACC Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "Big Ten Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "Big XII Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "C-USA Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "MAC Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "MWC Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "SEC Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "Sun Belt Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "FCS Pro Day"
	tempEvent.IsCombine = false
	tempList = append(tempList, tempEvent)

	tempEvent.Name = "NFL Combine"
	tempEvent.IsCombine = true
	tempList = append(tempList, tempEvent)

	return tempList
}
