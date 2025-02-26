package repository

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func FindAllCFBTeamRequests() []structs.TeamRequest {
	db := dbprovider.GetInstance().GetDB()
	var teamRequests []structs.TeamRequest

	//NFL Team Requests
	db.Where("is_approved = false").Find(&teamRequests)

	return teamRequests
}

func FindAllNFLTeamRequests() []structs.NFLRequest {
	db := dbprovider.GetInstance().GetDB()
	var nflTeamRequests []structs.NFLRequest

	//NFL Team Requests
	db.Where("is_approved = false").Find(&nflTeamRequests)

	return nflTeamRequests
}
