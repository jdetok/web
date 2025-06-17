package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
)

func (app *application) selectPlayersH(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Printf("Referer: %s\n", r.Referer())
	
	lg := r.URL.Query().Get("lg")
	var js []byte	

	// query the database for WNBA players
	if lg == "WNBA" {
		database, err := db.Connect()
		if err != nil {
			errs.HTTPErr(w, r, err)
			http.Error(w, "Error retrieving data", http.StatusInternalServerError)
			log.Printf("An error occured: %s", err)
		}

		js, err = db.SelectArg(database, db.CarrerStatsByLg, false, lg)
		if err != nil {
			http.Error(w, "Error retrieving data", http.StatusInternalServerError)
			w.Write([]byte("Error occured getting data from database"))
			return
		}
	} else {
		js = jsonops.ReadJSON(app.config.cachePath + "/players.json")
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
	
	lg := r.URL.Query().Get("lg")

	// var c = store.CacheJSON{}

	// js, err := c.LoadPlayers(db.CarrerStatsByLg, lg)
	// if err != nil {
	// 	fmt.Printf("Error getting players: %s", err)
	// }

	// c.AllPlayers = js
		
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
