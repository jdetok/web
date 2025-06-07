package main

import (
	"log"
	"net/http"

	"github.com/jdetok/web/internal/db"
)

func (app *application) selectHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing selecting from database via HTTP request\n"))
		
    database, err := db.Connect()
    if err != nil {
        log.Printf("An error occured: %s", err)
    }

    js, err := db.Select(database)
	if err != nil {
		w.Write([]byte("Error occured getting data from database"))
		return
	}

	w.Write([]byte(string(js)))
}

