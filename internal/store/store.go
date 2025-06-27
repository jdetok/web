package store

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
)

var cachePath string = env.GetString("CACHE_PATH")

type fPath struct {
	Query string
	File string
}

func (p fPath) Construct() string {
	return (cachePath + p.File)
}

var paths = []fPath{
	{
		Query: "select * from v_nba_rs_avgs", 
		File: "/nba_rs_avgs.json",
	},
	{
		Query: "select * from v_nba_rs_totals", 
		File: "/nba_rs_totals.json",
	},
	{
		Query: "select * from v_nba_po_avgs", 
		File: "/nba_po_avgs.json",
	},
	{
		Query: "select * from v_nba_po_totals", 
		File: "/nba_po_totals.json",
	},
	{
		Query: "select * from v_wnba_rs_avgs", 
		File: "/wnba_rs_avgs.json",
	},
	{
		Query: "select * from v_wnba_rs_totals", 
		File: "/wnba_rs_totals.json",
	},
	{
		Query: "select * from v_wnba_po_avgs", 
		File: "/wnba_po_avgs.json",
	},
	{
		Query: "select * from v_wnba_po_totals", 
		File: "/wnba_po_totals.json",
	},
}

// runs every interval seconds, updates if time since last update is > threshold
func CheckCache(db *sql.DB, lastUpdate *time.Time, players *[]Player, inteval time.Duration, threshold time.Duration) {
	e := errs.ErrInfo{Prefix: "cache check",}
	ticker := time.NewTicker(inteval)
	defer ticker.Stop()

	for range ticker.C {
		if time.Since(*lastUpdate) > threshold {
			fmt.Printf("refreshing cache at %v: %v since last update\n", 
				time.Now().Format("2006-01-02 15:04:05"), threshold)
			newPlayers, err := GetPlayers(db)
			if err != nil {
				e.Msg = "failed to get players"
			}
			*players = newPlayers
			// fmt.Println(players)
			updateTime, err := UpdateManyCache(db, paths)
			if err != nil {
				e.Msg = "cache update failed"
				fmt.Println(e.Error(err))
			}
			*lastUpdate = *updateTime
			fmt.Printf("finished refreshing cache at %v\n", updateTime)
		}
	}
} 
// TODO - connect once, pass connection
func UpdateManyCache(db *sql.DB, paths []fPath) (*time.Time, error) {
	
	var wg sync.WaitGroup
	for _, p := range paths {
		wg.Add(1)
		// run UpdateCacheNew concurrently for each query
		go func(p fPath){
			defer wg.Done()
			fmt.Printf("updating %s at %v\n", p.File, time.Now().Format("2006-01-02 15:04:05"))
			if err := UpdateCache(db, p.Query, p.Construct()); err != nil {
				fmt.Println(err)
			}
		}(p)
	}
	wg.Wait()
	// update the time after all have finished
	updateTime := time.Now()
	return &updateTime, nil
}

func UpdateCache(database *sql.DB, q string, path string) error{
	e := errs.ErrInfo{Prefix: ("cache update for " + path),}
	js, err := db.SelectDB(database, q)
	if err != nil {
		e.Msg = "database query failed"
		return e.Error(err)
	}
	err = jsonops.SaveJSON(path, js)
	if err != nil {
		e.Msg = "saving db response to json failed"
		return e.Error(err)
	}
	return nil
}