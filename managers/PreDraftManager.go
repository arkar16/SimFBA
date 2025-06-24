package managers

import (
	"fmt"

	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/util"
)

func RunPreDraftEvents() {
	draftees := GetAllNFLDraftees()

	// Create a list of events (Pro Days and Combines)
	eventList := GenerateTypicalListOfEvents()

	// Add the participants to each list
	eventList = AddParticipants(util.GetParticipantIDS(), eventList, draftees)

	// For each event, go through each ID in the participant IDs and grab the corresponding draftee and add it to the event

	// For each event, create a result for each player
	// For some % of draftees, create results based on their advertised grades, not their real grades.
}

func AddParticipants(json map[string][]uint, events []models.PreDraftEvent, players []models.NFLDraftee) []models.PreDraftEvent {
	// For each event, find that events list of players, and get that player from the draftee list and add to participants
	for _, event := range events {
		for _, player := range json[event.Name] {
			participant, found := FindParticipant(player, players)

			if found {
				event.Participants = append(event.Participants, participant)
			} else {
				// Participant not found
				fmt.Println("EVENT PARTICIPANT NOT FOUND!!!")
			}
		}
	}

	return events
}

func FindParticipant(x uint, list []models.NFLDraftee) (models.NFLDraftee, bool) {
	for _, n := range list {
		if x == uint(n.PlayerID) {
			return n, true
		}
	}
	return models.NFLDraftee{}, false
}

func GenerateTypicalListOfEvents() []models.PreDraftEvent {
	var tempList []models.PreDraftEvent
	var tempEvent models.PreDraftEvent
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
