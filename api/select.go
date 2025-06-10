package main

import (
	"log"
	"net/http"

	"github.com/jdetok/web/internal/db"
)

func (app *application) selectHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Testing selecting players from database via HTTP request\n"))
		
    database, err := db.Connect()
    if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
        log.Printf("An error occured: %s", err)
    }

    js, err := db.Select(database, db.CarrerStats, false)
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
	

	w.Write([]byte(string(js)))
}
