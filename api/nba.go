package main

import "net/http"

func (app *application) nbaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Will be the root of NBA stats page\n"))
}