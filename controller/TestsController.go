package controller

import (
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func TestCFBProgressionAlgorithm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")

	managers.CFBProgressionExport(w)
}

func TestNFLProgressionAlgorithm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")

	managers.NFLProgressionExport(w)
}
