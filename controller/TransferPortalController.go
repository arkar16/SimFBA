package controller

import (
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func ProcessTransferIntention(w http.ResponseWriter, r *http.Request) {
	managers.ProcessTransferIntention(w)
}
