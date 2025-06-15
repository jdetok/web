package main

import (
	"fmt"
	"time"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/jsonops"
)

func (app *application) CheckTime() time.Duration {
	return time.Since(app.StartTime)
}

// runs every interval seconds, updates if time since last update is > threshold
func (app *application) checkCache(inteval time.Duration, threshold time.Duration) {
	ticker := time.NewTicker(inteval)
	defer ticker.Stop()

	for range ticker.C {
		if time.Since(app.lastUpdate) > threshold {
			fmt.Printf("Refreshing cache at %v...\n", time.Now().Format("2006-01-02 15:04:05"))
			app.updateCache()
		}
	}
} 

// query all players from database & save to JSON
func (app *application) updateCache() {
	database, err := db.Connect()
	if err != nil {
		fmt.Println("Error connection to databse: ", err)
		return
	}
	js, err := db.Select(database, db.CarrerStats, false)
	if err != nil {
		fmt.Println("Error getting data from databse: ", err)
	}

	jsonops.SaveJSON(app.config.cachePath + "/players.json", js)
	app.lastUpdate = time.Now()
}

