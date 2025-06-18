package store

import (
	"fmt"
	"time"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/jsonops"
)

// runs every interval seconds, updates if time since last update is > threshold
func CheckCache(lastUpdate *time.Time, inteval time.Duration, threshold time.Duration) {
	ticker := time.NewTicker(inteval)
	defer ticker.Stop()

	for range ticker.C {
		if time.Since(*lastUpdate) > threshold {
			fmt.Printf("Refreshing cache at %v...\n", time.Now().Format("2006-01-02 15:04:05"))
			if newTime := UpdateCache(); newTime != nil {
			*lastUpdate = *newTime
			}
		}
	}
} 

// query all players from database & save to JSON
func UpdateCache() *time.Time {
	cachePath:= env.GetString("CACHE_PATH")
	database, err := db.Connect()
	if err != nil {
		fmt.Println("Error connection to databse: ", err)
		return nil
	}
	js, err := db.Select(database, db.CarrerStats, false)
	if err != nil {
		fmt.Println("Error getting data from databse: ", err)
		return nil
	}

	jsonops.SaveJSON(cachePath + "/players.json", js)
	
	updateTime := time.Now()

	return &updateTime
}

