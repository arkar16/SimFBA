package controller

import (
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func ExportDrafteesToCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")

	managers.ExportDrafteesToCSV(w)
}
