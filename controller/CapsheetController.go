package controller

import (
	"fmt"
	"net/http"

	"github.com/CalebRose/SimFBA/managers"
)

func GenerateCapsheets(w http.ResponseWriter, r *http.Request) {
	managers.AllocateCapsheets()
	fmt.Println(w, "Congrats, you generated the Capsheets!")
}
