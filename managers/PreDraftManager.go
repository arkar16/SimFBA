package managers

import (
	"fmt"
	"math/rand"

	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/util"
)

func RunPreDraftEvents() {
	draftees := GetAllNFLDraftees()

	// Create a list of events (Pro Days and Combines)
	eventList := GenerateTypicalListOfEvents()

	// Add the participants to each list
	eventList = AddParticipants(util.GetParticipantIDS(), eventList, draftees)

	// For each event, create a result for each player
	for _, event := range eventList {
		for _, player := range event.Participants {
			// For some % of draftees, create results based on their advertised grades, not their real grades.
			hidePerformance := ShouldHidePerformance()

			// Run events on them
			player = RunEvents(player, hidePerformance)
		}
	}

	// Export the results
	blah
}

// Set % of draftees that only perform based on their advertised grades, not real grades
func ShouldHidePerformance() bool {
	// 10% chance
	chance := 10

	roll := rand.Intn(100)

	if roll < chance {
		return true
	} else {
		return false
	}
}

func RunEvents(draftee models.NFLDraftee, shouldHidePerformance bool) models.NFLDraftee {
	draftee = RunUniversalEvents(draftee, shouldHidePerformance)
	draftee = RunPositionEvents(draftee, shouldHidePerformance)
	return draftee
}

func RunUniversalEvents(draftee models.NFLDraftee, shouldHidePerformance bool) models.NFLDraftee {

}

func RunPositionEvents(draftee models.NFLDraftee, shouldHidePerformance bool) models.NFLDraftee {

}

func Run40YardDash(speed uint) float32 {

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
