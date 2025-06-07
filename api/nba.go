package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) nbaHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/nba HTTP request: %v", http.StatusOK)
	fmt.Println(r.Context())
	w.Write([]byte("Will be the root of NBA stats page\n"))
}