package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/store"
	"github.com/joho/godotenv"
) 

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    slog.SetDefault(logger)
    
    // load environment variabels
    err := godotenv.Load()
	if err != nil {
        slog.Error("failed to get environment variables")
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

    app.players, err = store.GetPlayers(app.database)
    fmt.Println(app.players)

    if err != nil {
        slog.Error("error getting players")
    }
    // fmt.Println(app.players)

    // checks if cache needs refreshed every 30 seconds, refreshes if 60 sec since last
    go store.CheckCache(app.database, &app.lastUpdate, &app.players, 30*time.Second, 300*time.Second)

    mux := app.mount()
    if err := app.run(mux); err != nil {
        slog.Error("error running server")
    }
}    


    // log.Fatal(app.run(mux))
