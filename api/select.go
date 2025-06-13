package main

import (
	"log"
	"net/http"

	"github.com/jdetok/web/internal/db"
)

func (app *application) selectPlayersH(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Testing selecting players from database via HTTP request\n"))
	
	lg := r.URL.Query().Get("lg")

	database, err := db.Connect()
    if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
        log.Printf("An error occured: %s", err)
    }

    js, err := db.SelectArg(database, db.CarrerStatsByLg, false, lg)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		w.Write([]byte("Error occured getting data from database"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	

	w.Write(js)
}

// need to write function to accept player name as string, search if in DB, then query this with id
func (app *application) selectPlayerH(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Testing selecting players from database via HTTP request\n"))
	
	lg := r.URL.Query().Get("lg")
	player := r.URL.Query().Get("player")

	database, err := db.Connect()
    if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
        log.Printf("An error occured: %s", err)
    }

    js, err := db.SelectArgs(database, db.LgPlayerStat.Q, false, lg, player)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		w.Write([]byte("Error occured getting data from database"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	

	w.Write(js)
}

func (app *application) selectHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Testing selecting players from database via HTTP request\n"))

	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	log.Printf("User-Agent: %s", r.UserAgent())
	log.Printf("Referer: %s", r.Referer())
	log.Printf("Host: %s", r.Host)
	
	lg := r.URL.Query().Get("lg")
		
    database, err := db.Connect()
    if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
        log.Printf("An error occured: %s", err)
    }

    js, err := db.SelectArg(database, db.CarrerStatsByLg, false, lg)
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		w.Write([]byte("Error occured getting data from database"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	

	w.Write(js)
}

func (app *application) selectGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing selecting games from database via HTTP request\n"))
		
    database, err := db.Connect()
    if err != nil {
        log.Printf("An error occured: %s", err)
    }

    js, err := db.Select(database, db.Games, true)
	if err != nil {
		w.Write([]byte("Error occured getting data from database"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	

	w.Write(js)
}
