package main

import (
	"log"
	"net/http"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
	"github.com/jdetok/web/internal/logs"
)

func (app *application) selectPlayersH(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	
	lg := r.URL.Query().Get("lg")
	var js []byte	

	// query the database for WNBA players
	if lg == "WNBA" {
		database, err := db.Connect()
		if err != nil {
			// TODO!!!
			errs.HTTPErr(w, r, err)
			return
		}

		js, err = db.SelectArg(database, db.CarrerStatsByLg, false, lg)
		if err != nil {
			errs.HTTPErr(w, r, err)
			return
		}
	} else {
		// read cached json for nba 
		js = jsonops.ReadJSON(app.config.cachePath + "/players.json")
	}
	app.JSONWriter(w, js)
}

// need to write function to accept player name as string, search if in DB, then query this with id
func (app *application) selectPlayerH(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Testing selecting players from database via HTTP request\n"))
	logs.LogHTTP(r)
	lg := r.URL.Query().Get("lg")
	player := r.URL.Query().Get("player")

	database, err := db.Connect()
    if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
        log.Printf("An error occured: %s", err)
    }

    js, err := db.SelectArgs(database, db.LgPlayerStat.Q, false, lg, player)
	if err != nil {
		errs.HTTPErr(w, r, err)
	}
	app.JSONWriter(w, js)
}

func (app *application) selectHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Testing selecting players from database via HTTP request\n"))
	logs.LogHTTP(r)
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

	app.JSONWriter(w, js)
}

func (app *application) selectGameHandler(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	// w.Write([]byte("Testing selecting games from database via HTTP request\n"))
		
    database, err := db.Connect()
    if err != nil {
        log.Printf("An error occured: %s", err)
    }

    js, err := db.Select(database, db.Games, true)
	if err != nil {
		w.Write([]byte("Error occured getting data from database"))
		return
	}
	app.JSONWriter(w, js)
}
