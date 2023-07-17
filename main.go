package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func rockPaperScissors(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "rock-paper-scissors.html")
}

func playRockPaperScissors(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c")[:1]) //take first char of query string
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc(("/rock-paper-scissors"), rockPaperScissors)
	http.HandleFunc("/play_rps", playRockPaperScissors)
	log.Println("Starting web server on 8080")
	http.ListenAndServe(":8080", nil)
}
