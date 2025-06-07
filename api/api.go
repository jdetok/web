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
	
// handles root requests -- requests to http://url should be routed to /v1
// SERVES STATIC SITE IN WEB DIRECTORY
	mux.Handle("/", http.FileServer(http.Dir("/home/jdeto/go/src/go-api/web/src")))
	
	// mux.HandleFunc("GET /", app.rootHandler)
	
// health check endpoint
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

// base nba endpoint
	mux.HandleFunc("GET /nba", app.nbaHandler)

// testing kicking off the select via request
	mux.HandleFunc("GET /select", app.selectHandler)
	
// return mux instance - call app.mount() to get mux then app.run(mux) to run server
	return mux
}

