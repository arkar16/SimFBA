package controller

import (
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func AssignAllRecruitRanks(w http.ResponseWriter, r *http.Request) {

	managers.AssignAllRecruitRanks()

	fmt.Println("All Croots now have ranks")
	w.WriteHeader(http.StatusOK)
}
