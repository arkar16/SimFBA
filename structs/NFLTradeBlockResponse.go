package structs

type NFLTradeBlockResponse struct {
	Team                   NFLTeam
	TradablePlayers        []NFLPlayer
	DraftPicks             []NFLDraftPick
	SentTradeProposals     []NFLTradeProposalDTO
	ReceivedTradeProposals []NFLTradeProposalDTO
}
