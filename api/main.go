package main

import (
	"log"
	"time"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/store"
	"github.com/joho/godotenv"
) 

func main() {
    // load environment variabels
    err := godotenv.Load()
	if err != nil {
		 log.Fatal("dotenv failed to get environment variables")
	}

    // configs go here - 8080 for testing, will derive real vals from environment
    cfg := config{
        addr: env.GetString("SRV_IP"),
        cachePath: env.GetString("CACHE_PATH"),
        // TODO - ADD IN DB CONNECTION POOl
    }

    // initialize the app with the configs
    app := &application{
        config: cfg,
        database: db.InitDB(),
    }

    // checks if cache needs refreshed every 30 seconds, refreshes if 60 sec since last
    go store.CheckCache(app.database, &app.lastUpdate, 5*time.Second, 10*time.Second)

    mux := app.mount()
    log.Fatal(app.run(mux))
}