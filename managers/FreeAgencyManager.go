package managers

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
)

func GetAllFreeAgents() []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	fas := []structs.NFLPlayer{}

	db.Where("is_free_agent = ?", true).Find(&fas)

	return fas
}

func GetAllAvailableNFLPlayers(TeamID string) models.FreeAgencyResponse {
	FAs := GetAllFreeAgentsWithOffers()
	WaiverPlayers := GetAllWaiverWirePlayers()
	Offers := GetFreeAgentOffersByTeamID(TeamID)

	return models.FreeAgencyResponse{
		FreeAgents:    FAs,
		WaiverPlayers: WaiverPlayers,
		TeamOffers:    Offers,
	}
}

// GetAllFreeAgentsWithOffers -- For Free Agency UI Page.
func GetAllFreeAgentsWithOffers() []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	fas := []structs.NFLPlayer{}

	db.Preload("Offers", func(db *gorm.DB) *gorm.DB {
		return db.Order("contract_value DESC").Where("is_active = true")
	}).Order("overall desc").Where("is_free_agent = ?", true).Find(&fas)

	return fas
}

func GetAllWaiverWirePlayers() []structs.NFLPlayer {
	db := dbprovider.GetInstance().GetDB()

	WaivedPlayers := []structs.NFLPlayer{}

	db.Where("is_waived = ?", true).Find(&WaivedPlayers)

	return WaivedPlayers
}

func GetFreeAgentOffersByTeamID(TeamID string) []structs.FreeAgencyOffer {
	db := dbprovider.GetInstance().GetDB()

	offers := []structs.FreeAgencyOffer{}

	err := db.Where("team_id = ? AND is_active = ?", TeamID, true).Find(&offers).Error
	if err != nil {
		return offers
	}

	return offers
}

func GetFreeAgentOffersByPlayerID(PlayerID string) []structs.FreeAgencyOffer {
	db := dbprovider.GetInstance().GetDB()

	offers := []structs.FreeAgencyOffer{}

	err := db.Where("nfl_player_id = ? AND is_active = ?", PlayerID, true).Find(&offers).Error
	if err != nil {
		return offers
	}

	return offers
}

func GetFreeAgentOfferByOfferID(OfferID string) structs.FreeAgencyOffer {
	db := dbprovider.GetInstance().GetDB()

	offer := structs.FreeAgencyOffer{}

	err := db.Where("id = ?", OfferID).Find(&offer).Error
	if err != nil {
		return offer
	}

	return offer
}

func CreateFAOffer(offer structs.FreeAgencyOfferDTO) structs.FreeAgencyOffer {
	db := dbprovider.GetInstance().GetDB()

	freeAgentOffer := GetFreeAgentOfferByOfferID(strconv.Itoa(int(offer.ID)))

	if freeAgentOffer.ID == 0 {
		id := GetLatestOfferInDB(db)
		freeAgentOffer.AssignID(id)
	}

	freeAgentOffer.CalculateOffer(offer)

	db.Save(&freeAgentOffer)

	fmt.Println("Creating offer!")

	return freeAgentOffer
}

func CancelOffer(offer structs.FreeAgencyOfferDTO) {
	db := dbprovider.GetInstance().GetDB()

	OfferID := strconv.Itoa(int(offer.ID))

	freeAgentOffer := GetFreeAgentOfferByOfferID(OfferID)

	freeAgentOffer.CancelOffer()

	db.Save(&freeAgentOffer)
}

func SignFreeAgent(offer structs.FreeAgencyOffer, FreeAgent structs.NFLPlayer, ts structs.Timestamp) {
	db := dbprovider.GetInstance().GetDB()

	NFLTeam := GetNFLTeamByTeamID(strconv.Itoa(int(offer.TeamID)))
	FreeAgent.SignPlayer(int(NFLTeam.ID), NFLTeam.TeamAbbr)

	Contract := structs.NFLContract{
		PlayerID:       FreeAgent.PlayerID,
		NFLPlayerID:    FreeAgent.PlayerID,
		TeamID:         NFLTeam.ID,
		Team:           NFLTeam.TeamAbbr,
		OriginalTeamID: NFLTeam.ID,
		OriginalTeam:   NFLTeam.TeamAbbr,
		ContractLength: offer.ContractLength,
		Y1BaseSalary:   offer.Y1BaseSalary,
		Y1Bonus:        offer.Y1Bonus,
		Y2BaseSalary:   offer.Y2BaseSalary,
		Y2Bonus:        offer.Y2Bonus,
		Y3BaseSalary:   offer.Y3BaseSalary,
		Y3Bonus:        offer.Y3Bonus,
		Y4BaseSalary:   offer.Y4BaseSalary,
		Y4Bonus:        offer.Y4Bonus,
		Y5BaseSalary:   offer.Y5BaseSalary,
		Y5Bonus:        offer.Y5Bonus,
		ContractValue:  offer.ContractValue,
		IsActive:       true,
		IsComplete:     false,
		IsExtended:     false,
	}

	db.Create(&Contract)
	db.Save(&FreeAgent)

	// News Log
	message := "FA " + FreeAgent.Position + " " + FreeAgent.FirstName + " " + FreeAgent.LastName + " has signed with the " + NFLTeam.TeamName + " with a contract worth approximately $" + strconv.Itoa(int(Contract.ContractValue)) + " Million Dollars."
	newsLog := structs.NewsLog{
		WeekID:      ts.NFLWeekID,
		SeasonID:    ts.NFLSeasonID,
		MessageType: "Free Agency",
		Message:     message,
		League:      "NFL",
	}

	db.Create(&newsLog)
}

func SyncFreeAgencyOffers() {
	db := dbprovider.GetInstance().GetDB()

	ts := GetTimestamp()

	FreeAgents := GetAllFreeAgents()

	for _, FA := range FreeAgents {
		// If the Free Agent is not available in off-season free agency anymore
		if ts.IsNFLOffSeason && !FA.IsNegotiating && !FA.IsAcceptingOffers {
			continue
		}

		// Check if still accepting offers
		if ts.IsNFLOffSeason && FA.IsAcceptingOffers && ts.FreeAgencyRound < FA.NegotiationRound {
			continue
		}

		if ts.IsNFLOffSeason && FA.IsAcceptingOffers && ts.FreeAgencyRound >= FA.NegotiationRound {
			FA.ToggleIsNegotiating()
			db.Save(&FA)
			continue
		}

		// Check if still negotiation
		if ts.IsNFLOffSeason && FA.IsNegotiating && ts.FreeAgencyRound < FA.SigningRound {
			continue
		}

		// Is Ready to Sign
		Offers := GetFreeAgentOffersByPlayerID(strconv.Itoa(int(FA.ID)))

		// Sort by highest contract value
		sort.Sort(structs.ByContractValue(Offers))

		WinningOffer := structs.FreeAgencyOffer{}

		for _, Offer := range Offers {
			// Get the Contract with the best value for the FA
			if Offer.IsActive && WinningOffer.ID == 0 {
				WinningOffer = Offer
			}
			if Offer.IsActive {
				Offer.CancelOffer()
			}

			db.Save(&Offer)
		}

		if WinningOffer.ID > 0 {
			SignFreeAgent(WinningOffer, FA, ts)
		} else {
			FA.WaitUntilAfterDraft()
			db.Save(&FA)
		}
	}
}

func GetLatestOfferInDB(db *gorm.DB) uint {
	var latestOffer structs.FreeAgencyOffer

	err := db.Last(&latestOffer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1
		}
		log.Fatalln("ERROR! Could not find latest record" + err.Error())
	}

	return latestOffer.ID + 1
}
