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
	event.FourtyYardDash = Run40YardDash(draftee.Speed, event.IsCombine)
	event.BenchPress = RunBenchPress(draftee.Strength, event.IsCombine, draftee.Position)
	event.Shuttle = RunShuttle(draftee.Agility, event.IsCombine)
	event.ThreeCone = Run3Cone(draftee.Agility, event.IsCombine)
	event.VerticalJump = RunVertJump(draftee.Agility, draftee.Strength, draftee.Weight, event.IsCombine)
	event.BroadJump = RunBroadJump(draftee.Agility, draftee.Strength, draftee.Weight, event.IsCombine)

	if event.IsCombine {
		event.Wonderlic = RunWonderlic(draftee.FootballIQ)
	}

	return event
}

func RunPositionEvents(draftee models.NFLDraftee, shouldHidePerformance bool, event models.EventResults) models.EventResults {
	// CREATE POSITIONAL EVENTS
	event.ThrowingAccuracy
}

func Run40YardDash(speed int, isCombine bool) float32 {
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

func RunBenchPress(strength int, isCombine bool, position string) uint8 {
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

func RunShuttle(agility int, isCombine bool) float32 {
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

func Run3Cone(agility int, isCombine bool) float32 {
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

func RunVertJump(agility int, strength int, weight int, isCombine bool) uint8 {
	delta := GetDelta(isCombine)
	newStrength := float64(strength) + delta
	if newStrength > 99.0 {
		newStrength = 99.0
	}
	delta = GetDelta(isCombine)
	newAgility := float64(agility) + delta
	if newAgility > 99.0 {
		newAgility = 99.0
	}
	temp := ((newStrength + newAgility) / float64(weight))
	temp = 1500 * temp
	temp = math.Sqrt(temp)
	temp = temp + 11.0
	return uint8(temp)
}

func RunBroadJump(agility int, strength int, weight int, isCombine bool) uint8 {
	delta := GetDelta(isCombine)
	newStrength := float64(strength) + delta
	if newStrength > 99.0 {
		newStrength = 99.0
	}
	delta = GetDelta(isCombine)
	newAgility := float64(agility) + delta
	if newAgility > 99.0 {
		newAgility = 99.0
	}
	temp := ((newStrength + newAgility) / float64(weight))
	temp = 20000 * temp
	temp = math.Sqrt(temp)
	temp = temp / 2.0
	temp = temp + 79.0
	return uint8(temp)
}

func RunWonderlic(fbIQ int) uint8 {
	delta := GetDelta(true)
	temp := float64(fbIQ) + delta
	if temp > 99.0 {
		temp = 99.0
	}
	temp = temp - 130.0
	temp = math.Pow(temp, 2)
	temp = temp / 25000.0
	temp = temp + 51.0
	return uint8(temp)
}

func RunQBAccuracy(throwAcc int, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	temp := float64(throwAcc) + delta
	// Make sure they can't score more than 10.00
	if temp > 80.0 {
		temp = 80.0
	}
	temp = temp / 80.0
	temp = temp * 10.0
	return float32(temp)
}

func RunQBDistance(throwPow int, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	temp := float64(throwPow) + delta
	// Make sure they can't score more than 10.00
	if temp > 85.0 {
		temp = 85.0
	}
	temp = temp / 85.0
	temp = temp * 10.0
	return float32(temp)
}

func RunInsideRun(speed int, strength int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	newSpeed := float64(speed) + delta
	delta = GetDelta(isCombine)
	newStrength := float64(strength) + delta
	temp := newSpeed + newStrength
	// Make sure they can't score more than 10.00
	if temp > 170.0 {
		temp = 170.0
	}
	temp = temp / 170.0
	temp = temp * 10.0
	return float32(temp)
}

func RunOutsideRun(speed int, agility int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	newSpeed := float64(speed) + delta
	delta = GetDelta(isCombine)
	newAgility := float64(agility) + delta
	temp := newSpeed + newAgility
	// Make sure they can't score more than 10.00
	if temp > 180.0 {
		temp = 180.0
	}
	temp = temp / 180.0
	temp = temp * 10.0
	return float32(temp)
}

func RunCatching(catching int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	temp := float64(catching) + delta
	// Make sure they can't score more than 10.00
	if temp > 80.0 {
		temp = 80.0
	}
	temp = temp / 80.0
	temp = temp * 10.0
	return float32(temp)
}

func RunRouteRunning(routeRun int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	temp := float64(routeRun) + delta
	// Make sure they can't score more than 10.00
	if temp > 65.0 {
		temp = 65.0
	}
	temp = temp / 65.0
	temp = temp * 10.0
	return float32(temp)
}

func RunRunBlocking(runBlock int, isCombine bool, position string) float32 {
	delta := GetDelta(isCombine)
	temp := 0.0
	// Make sure they can't score more than 10.00
	if strings.Contains(strings.ToLower(position), strings.ToLower("FB")) {
		temp = float64(runBlock) + delta - 15.0
	} else if strings.Contains(strings.ToLower(position), strings.ToLower("TE")) {
		temp = float64(runBlock) + delta - 8.0
	} else {
		temp = float64(runBlock) + delta
	}
	if temp > 85.0 {
		temp = 85.0
	}
	temp = temp / 85.0
	temp = temp * 10.0
	return float32(temp)
}

func RunPassBlocking(passBlock int, isCombine bool, position string) float32 {
	delta := GetDelta(isCombine)
	temp := 0.0
	// Make sure they can't score more than 10.00
	if strings.Contains(strings.ToLower(position), strings.ToLower("FB")) {
		temp = float64(passBlock) + delta - 15.0
	} else if strings.Contains(strings.ToLower(position), strings.ToLower("TE")) {
		temp = float64(passBlock) + delta - 8.0
	} else {
		temp = float64(passBlock) + delta
	}
	if temp > 85.0 {
		temp = 85.0
	}
	temp = temp / 85.0
	temp = temp * 10.0
	return float32(temp)
}

func RunRunStop(runDef int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	temp := float64(runDef) + delta
	// Make sure they can't score more than 10.00
	if temp > 85.0 {
		temp = 85.0
	}
	temp = temp / 85.0
	temp = temp * 10.0
	return float32(temp)
}

func RunPassRush(PassRush int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	temp := float64(PassRush) + delta
	// Make sure they can't score more than 10.00
	if temp > 80.0 {
		temp = 80.0
	}
	temp = temp / 80.0
	temp = temp * 10.0
	return float32(temp)
}

func RunLBCoverage(manCov int, zonCov int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	newManCov := float64(manCov) + delta
	delta = GetDelta(isCombine)
	newZonCov := float64(zonCov) + delta
	temp := newManCov + newZonCov
	if temp > 150.0 {
		temp = 150.0
	}
	temp = temp / 150.0
	temp = temp * 10
	return float32(temp)
}

func RunManCoverage(manCov int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	temp := float64(manCov) + delta
	if temp > 90.0 {
		temp = 90.0
	}
	temp = temp / 90.0
	temp = temp * 10
	return float32(temp)
}

func RunZoneCoverage(zonCov int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	temp := float64(zonCov) + delta
	if temp > 80.0 {
		temp = 80.0
	}
	temp = temp / 80.0
	temp = temp * 10
	return float32(temp)
}

func RunKickoffDrill(kickPow int, puntPow int, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	// Get larger of the two kicking values
	temp := math.Max(float64(kickPow), float64(puntPow)) + delta
	if temp > 75.0 {
		temp = 75.0
	}
	temp = temp / 75.0
	temp = temp * 10
	return float32(temp)
}

func RunFieldGoalDrill(kickPow int, kickAcc int, isCombine bool) float32 {
	delta := GetDelta(isCombine)
	newKickPow := float64(kickPow) + delta
	delta = GetDelta(isCombine)
	newKickAcc := float64(kickAcc) + delta
	temp := newKickPow + newKickAcc
	if temp > 155.0 {
		temp = 155.0
	}
	temp = temp / 155.0
	temp = temp * 10
	return float32(temp)
}

func RunPuntDistance(puntPow int, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	// Get larger of the two kicking values
	temp := float64(puntPow) + delta
	if temp > 60.0 {
		temp = 60.0
	}
	temp = temp / 60.0
	temp = temp * 10
	return float32(temp)
}

func RunCoffinPunt(puntAcc int, isCombine bool) float32 {
	delta := GetDelta(isCombine)

	// Get larger of the two kicking values
	temp := float64(puntAcc) + delta
	if temp > 66.0 {
		temp = 66.0
	}
	temp = temp / 66.0
	temp = temp * 10
	return float32(temp)
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
