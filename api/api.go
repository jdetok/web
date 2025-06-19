package main

import (
	"fmt"
	"net/http"
	"time"
)

type application struct {
	config config
	StartTime time.Time
	lastUpdate time.Time
}

type config struct {
	addr string
	cachePath string
}

func (app *application) setStartTime() {
	app.StartTime = time.Now()
}

func (app *application) JSONWriter(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) run(mux *http.ServeMux) error {
	
// server configuration
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	// set the time for caching
	app.setStartTime()

	fmt.Printf("http server configured and starting at %v...\n", 
		app.StartTime.Format("2006-01-02 15:04:05"))
	return srv.ListenAndServe()
}

// returns type ServeMux for a router
func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

// TESTING UPDATED ENDPOINTS 06/19
	// for when a user doesn't specify a player -- ?lg=&stype=
	mux.HandleFunc("GET /bball/players", app.getStats)

// testing kicking off the select via request
	mux.HandleFunc("GET /select", app.selectPlayersH)
	mux.HandleFunc("GET /select/games", app.selectGameHandler)
	mux.HandleFunc("GET /select/players", app.selectPlayersH)
	mux.HandleFunc("GET /select/player", app.selectPlayerHTest)
	//mux.HandleFunc("GET /select/player", app.selectPlayerH)

// SERVES STATIC SITE IN WEB DIRECTORY, DON'T CACHE JS & CSS
	mux.Handle("/js/", http.HandlerFunc(app.jsNoCache))
	mux.Handle("/css/", http.HandlerFunc(app.cssNoCache))
	mux.HandleFunc("/", app.rootHandler)
	
// return mux instance - call app.mount() to get mux then app.run(mux) to run server
	return mux
}

