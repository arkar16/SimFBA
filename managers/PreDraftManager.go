package managers

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

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

			playerEvents := GenerateEvent(player, event)

			// Run events on them
			playerEvents = RunEvents(player, hidePerformance, playerEvents)
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

func GenerateEvent(draftee models.NFLDraftee, event models.PreDraftEvent) models.EventResults {
	var newEvent models.EventResults
	newEvent.PlayerID = uint(draftee.PlayerID)
	newEvent.IsCombine = event.IsCombine
	newEvent.Name = event.Name
	return newEvent
}

func RunEvents(draftee models.NFLDraftee, shouldHidePerformance bool, event models.EventResults) models.EventResults {
	event = RunUniversalEvents(draftee, shouldHidePerformance, event)
	event = RunPositionEvents(draftee, shouldHidePerformance, event)
	return event
}

func RunUniversalEvents(draftee models.NFLDraftee, shouldHidePerformance bool, event models.EventResults) models.EventResults {
	event.FourtyYardDash = Run40YardDash(uint(draftee.Speed), event.IsCombine)
	event.BenchPress = RunBenchPress()
	event.Shuttle = RunShuttle()
	event.ThreeCone = Run3Cone()
	event.VerticalJump = RunVertJump()
	event.BroadJump = RunBroadJump()

	if event.IsCombine {
		event.Wonderlic = RunWonderlic()
	}
}

func RunPositionEvents(draftee models.NFLDraftee, shouldHidePerformance bool, event models.EventResults) models.EventResults {
	// CREATE POSITIONAL EVENTS
}

func Run40YardDash(speed uint, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	temp := float64(speed) + delta

	if temp > 99.0 {
		temp = 99.0
	}

	temp = 100 - temp
	temp = math.Pow(temp, 2)
	temp = temp / 4000
	temp = temp + 4.3

	return float32(temp)
}

func RunBenchPress(strength uint, isCombine bool, position string) uint8 {
	delta := GetDelta(isCombine)

	temp := 0.0

	if strings.Contains(strings.ToLower(position), strings.ToLower("FB")) {
		temp = float64(strength) + delta - 10.0
	} else {
		temp = float64(strength) + delta
	}
	if temp > 99.0 {
		temp = 99.0
	}
	temp = 185.0 - temp
	temp = math.Pow(temp, 2)
	temp = temp / 600.0
	temp = temp * -1.0
	temp = temp + 66.0
	return uint8(temp)
}

func RunShuttle(agility uint, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	temp := float64(agility) + delta
	if temp > 99.0 {
		temp = 99.0
	}
	temp = 100.0 - temp
	temp = math.Pow(temp, 2)
	temp = temp / 6000.0
	temp = temp + 3.7
	return float32(temp)
}

func Run3Cone(agility uint, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	temp := float64(agility) + delta
	if temp > 99.0 {
		temp = 99.0
	}
	temp = 100.0 - temp
	temp = math.Pow(temp, 2)
	temp = temp / 4200.0
	temp = temp + 6.28
	return float32(temp)
}

func RunVertJump(agility uint, strength uint, weight uint, isCombine bool) uint8 {
	delta := GetDelta(isCombine)
	newStrength := float64(strength) + delta
	delta = GetDelta(isCombine)
	newAgility := float64(agility) + delta
	temp := ((newStrength + newAgility) / float64(weight))
	temp = 1500 * temp
	temp = math.Sqrt(temp)
	temp = temp + 11.0
	return uint8(temp)
}

func RunBroadJump(agility uint, strength uint, weight uint, isCombine bool) uint8 {
	delta := GetDelta(isCombine)
	newStrength := float64(strength) + delta
	delta = GetDelta(isCombine)
	newAgility := float64(agility) + delta
	temp := ((newStrength + newAgility) / float64(weight))
	temp = 20000 * temp
	temp = math.Sqrt(temp)
	temp = temp / 2.0
	temp = temp + 79.0
	return uint8(temp)
}

func RunWonderlic() uint8 {

}

func GetDelta(isCombine bool) float64 {
	min := 0
	max := 0

	if isCombine {
		min = -10
		max = 10
	} else {
		min = -5
		max = 15
	}

	return float64(rand.Intn((max - min) + min))
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
