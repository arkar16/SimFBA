package managers

import (
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
	"gorm.io/gorm"
)

func GetTradeBlockDataByTeamID(TeamID string) structs.NFLTradeBlockResponse {
	var waitgroup sync.WaitGroup
	waitgroup.Add(4)
	nflTeamChan := make(chan structs.NFLTeam)
	playersChan := make(chan []structs.NFLPlayer)
	picksChan := make(chan []structs.NFLDraftPick)
	proposalsChan := make(chan structs.NFLTeamProposals)

	go func() {
		waitgroup.Wait()
		close(nflTeamChan)
		close(playersChan)
		close(picksChan)
		close(proposalsChan)
	}()

	go func() {
		defer waitgroup.Done()
		team := GetNFLTeamWithCapsheetByTeamID(TeamID)
		nflTeamChan <- team
	}()

	go func() {
		defer waitgroup.Done()
		players := GetTradableNFLPlayersByTeamID(TeamID)
		playersChan <- players
	}()

	go func() {
		defer waitgroup.Done()
		picks := GetDraftPicksByTeamID(TeamID)
		picksChan <- picks
	}()

	go func() {
		defer waitgroup.Done()
		proposals := GetTradeProposalsByNFLID(TeamID)
		proposalsChan <- proposals
	}()

	nflTeam := <-nflTeamChan
	tradablePlayers := <-playersChan
	draftPicks := <-picksChan
	teamProposals := <-proposalsChan

	// close(nflTeamChan)
	// close(playersChan)
	// close(picksChan)
	// close(proposalsChan)

	return structs.NFLTradeBlockResponse{
		Team:                   nflTeam,
		TradablePlayers:        tradablePlayers,
		DraftPicks:             draftPicks,
		SentTradeProposals:     teamProposals.SentTradeProposals,
		ReceivedTradeProposals: teamProposals.ReceivedTradeProposals,
	}
}

func GetOnlyTradeProposalByProposalID(proposalID string) structs.NFLTradeProposal {
	db := dbprovider.GetInstance().GetDB()

	proposal := structs.NFLTradeProposal{}

	db.Preload("NFLTeamTradeOptions").Preload("RecepientTeamTradeOptions").Where("id = ?", proposalID).Find(&proposal)

	return proposal
}

func GetRejectedTradeProposals() []structs.NFLTradeProposal {
	db := dbprovider.GetInstance().GetDB()

	proposals := []structs.NFLTradeProposal{}

	db.Preload("NFLTeamTradeOptions").Preload("RecepientTeamTradeOptions").Where("is_trade_rejected = ?", true).Find(&proposals)

	return proposals
}

// GetTradeProposalsByNFLID -- Returns all trade proposals that were either sent or received
func GetTradeProposalsByNFLID(TeamID string) structs.NFLTeamProposals {
	db := dbprovider.GetInstance().GetDB()

	proposals := []structs.NFLTradeProposal{}

	db.Preload("NFLTeamTradeOptions").Preload("RecepientTeamTradeOptions").Where("nfl_team_id = ? OR recepient_team_id = ?", TeamID, TeamID).Find(&proposals)

	SentProposals := []structs.NFLTradeProposalDTO{}
	ReceivedProposals := []structs.NFLTradeProposalDTO{}

	id := uint(util.ConvertStringToInt(TeamID))

	sentOptions := []structs.NFLTradeOptionObj{}
	receivedOptions := []structs.NFLTradeOptionObj{}

	for _, proposal := range proposals {
		for _, sentOption := range proposal.NFLTeamTradeOptions {
			opt := structs.NFLTradeOptionObj{
				ID:               sentOption.Model.ID,
				TradeProposalID:  sentOption.TradeProposalID,
				NFLTeamID:        sentOption.NFLTeamID,
				SalaryPercentage: sentOption.SalaryPercentage,
			}
			if sentOption.NFLPlayerID > 0 {
				player := GetNFLPlayerRecord(strconv.Itoa(int(sentOption.NFLPlayerID)))
				opt.AssignPlayer(player)
			} else if sentOption.NFLDraftPickID > 0 {
				draftPick := GetDraftPickByDraftPickID(strconv.Itoa((int(sentOption.NFLDraftPickID))))
				opt.AssignPick(draftPick)
			}
			sentOptions = append(sentOptions, opt)
		}

		for _, receivedOption := range proposal.RecepientTeamTradeOptions {
			opt := structs.NFLTradeOptionObj{
				ID:               receivedOption.Model.ID,
				TradeProposalID:  receivedOption.TradeProposalID,
				NFLTeamID:        receivedOption.NFLTeamID,
				SalaryPercentage: receivedOption.SalaryPercentage,
			}
			if receivedOption.NFLPlayerID > 0 {
				player := GetNFLPlayerRecord(strconv.Itoa(int(receivedOption.NFLPlayerID)))
				opt.AssignPlayer(player)
			} else if receivedOption.NFLDraftPickID > 0 {
				draftPick := GetDraftPickByDraftPickID(strconv.Itoa((int(receivedOption.NFLDraftPickID))))
				opt.AssignPick(draftPick)
			}
			receivedOptions = append(receivedOptions, opt)
		}

		proposalResponse := structs.NFLTradeProposalDTO{
			ID:                        proposal.Model.ID,
			NFLTeamID:                 proposal.NFLTeamID,
			NFLTeam:                   proposal.NFLTeam,
			RecepientTeamID:           proposal.RecepientTeamID,
			RecepientTeam:             proposal.RecepientTeam,
			IsTradeAccepted:           proposal.IsTradeAccepted,
			IsTradeRejected:           proposal.IsTradeRejected,
			NFLTeamTradeOptions:       sentOptions,
			RecepientTeamTradeOptions: receivedOptions,
		}

		if proposal.NFLTeamID == id {
			SentProposals = append(SentProposals, proposalResponse)
		} else if proposal.RecepientTeamID == id {
			ReceivedProposals = append(ReceivedProposals, proposalResponse)
		}
	}
	return structs.NFLTeamProposals{
		SentTradeProposals:     SentProposals,
		ReceivedTradeProposals: ReceivedProposals,
	}
}

func PlaceNFLPlayerOnTradeBlock(playerID string) {
	db := dbprovider.GetInstance().GetDB()

	player := GetNFLPlayerRecord(playerID)

	player.ToggleTradeBlock()

	db.Save(&player)
}

func CreateTradeProposal(TradeProposal structs.NFLTradeProposalDTO) {
	db := dbprovider.GetInstance().GetDB()
	latestID := GetLatestProposalInDB(db)

	// Create Trade Options
	SentTradeOptions := TradeProposal.NFLTeamTradeOptions
	ReceivedTradeOptions := TradeProposal.RecepientTeamTradeOptions

	for _, sentOption := range SentTradeOptions {
		tradeOption := structs.NFLTradeOption{
			TradeProposalID:  latestID,
			NFLTeamID:        TradeProposal.NFLTeamID,
			NFLPlayerID:      sentOption.NFLPlayerID,
			NFLDraftPickID:   sentOption.NFLDraftPickID,
			SalaryPercentage: sentOption.SalaryPercentage,
		}
		db.Create(&tradeOption)
	}

	for _, recepientOption := range ReceivedTradeOptions {
		tradeOption := structs.NFLTradeOption{
			TradeProposalID:  latestID,
			NFLTeamID:        TradeProposal.NFLTeamID,
			NFLPlayerID:      recepientOption.NFLPlayerID,
			NFLDraftPickID:   recepientOption.NFLDraftPickID,
			SalaryPercentage: recepientOption.SalaryPercentage,
		}
		db.Create(&tradeOption)
	}

	// Create Trade Proposal Object
	proposal := structs.NFLTradeProposal{
		NFLTeamID:       TradeProposal.NFLTeamID,
		NFLTeam:         TradeProposal.NFLTeam,
		RecepientTeamID: TradeProposal.RecepientTeamID,
		RecepientTeam:   TradeProposal.NFLTeam,
		IsTradeAccepted: false,
		IsTradeRejected: false,
	}
	proposal.AssignID(latestID)

	db.Create(&proposal)
}

func AcceptTradeProposal(proposalID string) {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()

	proposal := GetOnlyTradeProposalByProposalID(proposalID)

	proposal.AcceptTrade()

	// Create News Log
	newsLogMessage := proposal.RecepientTeam + " has accepted a trade offer from " + proposal.NFLTeam + " for trade the following players:\n\n"

	for _, options := range proposal.NFLTeamTradeOptions {
		if options.NFLPlayerID > 0 {
			playerRecord := GetNFLPlayerRecord(strconv.Itoa(int(options.NFLPlayerID)))
			newsLogMessage += playerRecord.Position + " " + playerRecord.FirstName + " " + playerRecord.LastName + " to " + proposal.RecepientTeam + "\n"
		} else if options.NFLDraftPickID > 0 {
			draftPick := GetDraftPickByDraftPickID(strconv.Itoa(int(options.NFLDraftPickID)))
			pickRound := strconv.Itoa(int(draftPick.Round))
			roundAbbreviation := util.GetRoundAbbreviation(pickRound)
			season := strconv.Itoa(int(draftPick.Season))
			newsLogMessage += season + " " + roundAbbreviation + " pick to " + proposal.RecepientTeam + "\n"
		}
	}
	newsLogMessage += "\n"

	for _, options := range proposal.RecepientTeamTradeOptions {
		if options.NFLPlayerID > 0 {
			playerRecord := GetNFLPlayerRecord(strconv.Itoa(int(options.NFLPlayerID)))
			newsLogMessage += playerRecord.Position + " " + playerRecord.FirstName + " " + playerRecord.LastName + " to " + proposal.NFLTeam + "\n"
		} else if options.NFLDraftPickID > 0 {
			draftPick := GetDraftPickByDraftPickID(strconv.Itoa(int(options.NFLDraftPickID)))
			pickRound := strconv.Itoa(int(draftPick.Round))
			roundAbbreviation := util.GetRoundAbbreviation(pickRound)
			season := strconv.Itoa(int(draftPick.Season))
			newsLogMessage += season + " " + roundAbbreviation + " pick to " + proposal.NFLTeam + "\n"
		}
	}

	newsLog := structs.NewsLog{
		WeekID:      ts.NFLWeekID,
		Week:        ts.NFLWeek,
		SeasonID:    ts.NFLSeasonID,
		League:      "NFL",
		MessageType: "Trade",
		Message:     newsLogMessage,
	}

	db.Create(&newsLog)
	db.Save(&proposal)
}

func RejectTradeProposal(proposalID string) {
	db := dbprovider.GetInstance().GetDB()
	ts := GetTimestamp()

	proposal := GetOnlyTradeProposalByProposalID(proposalID)

	proposal.RejectTrade()
	newsLog := structs.NewsLog{
		WeekID:      ts.NFLWeekID,
		Week:        ts.NFLWeek,
		SeasonID:    ts.NFLSeasonID,
		League:      "NFL",
		MessageType: "Trade",
		Message:     proposal.RecepientTeam + " has rejected a trade from " + proposal.NFLTeam,
	}

	db.Create(&newsLog)
	db.Save(&proposal)
}

func GetLatestProposalInDB(db *gorm.DB) uint {
	var latestProposal structs.NFLTradeProposal

	err := db.Last(&latestProposal).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1
		}
		log.Fatalln("ERROR! Could not find latest record" + err.Error())
	}

	return latestProposal.ID + 1
}

func RemoveRejectedTrades() {
	db := dbprovider.GetInstance().GetDB()

	rejectedProposals := GetRejectedTradeProposals()

	for _, proposal := range rejectedProposals {
		sentOptions := proposal.NFLTeamTradeOptions
		recepientOptions := proposal.RecepientTeamTradeOptions
		deleteOptions(db, sentOptions)
		deleteOptions(db, recepientOptions)

		// Delete Proposal
		db.Delete(&proposal)
	}
}

func SyncAcceptedTrade(proposalID string) {
	db := dbprovider.GetInstance().GetDB()

	proposal := GetOnlyTradeProposalByProposalID(proposalID)
	SentOptions := proposal.NFLTeamTradeOptions
	RecepientOptions := proposal.RecepientTeamTradeOptions

	syncAcceptedOptions(db, SentOptions, proposal.NFLTeamID, proposal.NFLTeam, proposal.RecepientTeamID, proposal.RecepientTeam)
	syncAcceptedOptions(db, RecepientOptions, proposal.RecepientTeamID, proposal.RecepientTeam, proposal.NFLTeamID, proposal.NFLTeam)

	proposal.ToggleSyncStatus()

	db.Save(&proposal)
}

func syncAcceptedOptions(db *gorm.DB, options []structs.NFLTradeOption, senderID uint, senderTeam string, recepientID uint, recepientTeam string) {
	for _, option := range options {
		if option.NFLPlayerID > 0 {
			playerRecord := GetNFLPlayerRecord(strconv.Itoa(int(option.NFLPlayerID)))
			playerRecord.TradePlayer(recepientID, recepientTeam)
			// Contract
			contract := playerRecord.Contract
			percentage := option.SalaryPercentage

			// Subtract Contract from Senders's Capsheet
			sendersPercentage := 100 - percentage
			SendersCapsheet := GetCapsheetByTeamID(strconv.Itoa(int(senderID)))
			SendersCapsheet.SubtractFromCapsheet(contract)
			SendersCapsheet.NegotiateSalaryDifference(contract.Y1BaseSalary, float64(contract.Y1BaseSalary*sendersPercentage))

			db.Save(&SendersCapsheet)

			// Add to Recepient Capsheet
			recepientCapsheet := GetCapsheetByTeamID(strconv.Itoa(int(recepientID)))
			recepientCapsheet.AddContractViaTrade(contract, float64(percentage*contract.Y1BaseSalary))
			db.Save(&recepientCapsheet)

		} else if option.NFLDraftPickID > 0 {
			draftPick := GetDraftPickByDraftPickID(strconv.Itoa(int(option.NFLDraftPickID)))
			draftPick.TradePick(recepientID, recepientTeam)
			db.Save(&draftPick)
		}

		db.Delete(&option)
	}
}

func VetoTrade(proposalID string) {
	db := dbprovider.GetInstance().GetDB()

	proposal := GetOnlyTradeProposalByProposalID(proposalID)
	SentOptions := proposal.NFLTeamTradeOptions
	RecepientOptions := proposal.RecepientTeamTradeOptions

	deleteOptions(db, SentOptions)
	deleteOptions(db, RecepientOptions)

	db.Delete(&proposal)
}

func deleteOptions(db *gorm.DB, options []structs.NFLTradeOption) {
	// Delete Recepient Trade Options
	for _, option := range options {
		// Deletes the option
		db.Delete(&option)
	}
}
