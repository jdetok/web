package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
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

	log.Printf("Server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}

// returns type ServeMux for a router
func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/js/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.StripPrefix("/js/", http.FileServer(
			http.Dir("/home/jdeto/go/github.com/jdetok/web/www/src/js"))).ServeHTTP(w, r)
	}))

	mux.Handle("/css/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.StripPrefix("/css/", http.FileServer(
			http.Dir("/home/jdeto/go/github.com/jdetok/web/www/src/css"))).ServeHTTP(w, r)
	}))
	
// health check endpoint
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

// base nba endpoint
	mux.HandleFunc("GET /nba", app.nbaHandler)

// testing kicking off the select via request
	mux.HandleFunc("GET /select", app.selectHandler)

	mux.HandleFunc("GET /select/games", app.selectGameHandler)

	mux.HandleFunc("GET /select/players", app.selectPlayersH)

	mux.HandleFunc("GET /select/player", app.selectPlayerH)

// SERVES STATIC SITE IN WEB DIRECTORY
	mux.Handle("/", http.FileServer(http.Dir("/home/jdeto/go/github.com/jdetok/web/www/src")))
	
// return mux instance - call app.mount() to get mux then app.run(mux) to run server
	return mux
}

