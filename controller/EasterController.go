package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	config "github.com/CalebRose/SimFBA/secrets"
	"github.com/CalebRose/SimFBA/structs"
)

// Collusion Button
func CollusionButton(w http.ResponseWriter, r *http.Request) {
	apiUrl := "https://cors-anywhere.herokuapp.com/https://www.simfba.com/index.php?api/threads/"
	var collusionButton structs.CollusionDto
	err := json.NewDecoder(r.Body).Decode(&collusionButton)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sfaConfig := config.SFAConfig()

	data := url.Values{}
	data.Set("node_id", "75")
	title := "BREAKING NEWS: COACH " + collusionButton.User + " HAS BEEN CAUGHT COLLUDING ON THE INTERFACE!"
	data.Set("title", title)
	message := "THIS JUST IN: COACH " + collusionButton.User + " OF THE " + collusionButton.Team + " " + collusionButton.Mascot + " WAS CAUGHT ATTEMPTING TO COLLUDE FOR 200 POINTS ON THE SIM. THE FANBASE IS NOW IN DISGRACE WITH THIS ATROCITY. WHY? WHY CAN'T YOU USE THE INTERFACE LIKE EVERY OTHER RESPONSIBLE USER?"
	data.Set("message", message)

	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("XF-Api-Key", sfaConfig["sfaKey"])
	req.Header.Add("XF-Api-User", sfaConfig["sfaUser"])

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
}
