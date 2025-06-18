package main

import (
	"log"
	"time"

	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/store"
	"github.com/joho/godotenv"
) 

func main() {

    err := godotenv.Load()
	if err != nil {
		 log.Println("dotenv didn't work")
	}

    // configs go here - 8080 for testing, will derive real vals from environment
    cfg := config{
        addr: env.GetString("SRV_IP"),
        cachePath: env.GetString("CACHE_PATH"),
    }

    // initialize the app with the configs
    app := &application{
        config: cfg,
    }

    app.lastUpdate = time.Now()

    // checks if cache needs refreshed every 30 seconds, refreshes if 300 sec since last
    go store.CheckCache(&app.lastUpdate, 10*time.Second, 10*time.Second)

    // mount initializes mux (serves/routes HTTP) & handlers
    mux := app.mount()
    log.Fatal(app.run(mux))
}
