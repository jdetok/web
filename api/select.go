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

    db.Select(database)
}

