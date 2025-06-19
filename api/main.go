package main

import (
	"log"
	"time"

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
    }

    // checks if cache needs refreshed every 30 seconds, refreshes if 60 sec since last
    go store.CheckCache(&app.lastUpdate, 5*time.Second, 60*time.Second)

    mux := app.mount()
    log.Fatal(app.run(mux))
}

    // TURN BACK ON ASAP
    // // force a write to cache before server starts to ensure no stale data
    // fmt.Println("refreshing json stores before starting server...")
    // if update, err := store.UpdateCache(); err != nil {
    //     fmt.Println(err)
    // } else {
    //     app.lastUpdate = *update
    // }

    // mount & start server (routers/handlers in api.go)
