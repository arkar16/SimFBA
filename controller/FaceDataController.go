package controller

import (
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func MigrateFaceData(w http.ResponseWriter, r *http.Request) {
	managers.MigrateFaceDataToRecruits()
	managers.MigrateFaceDataToCollegePlayers()
	managers.MigrateFaceDataToHistoricCollegePlayers()
	managers.MigrateFaceDataToNFLPlayers()
	managers.MigrateFaceDataToRetiredPlayers()

	fmt.Println("All Faces have been generated")
	w.WriteHeader(http.StatusOK)
}
