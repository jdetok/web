package store

// import (
// 	"fmt"
// 	"time"

// 	"github.com/jdetok/web/internal/db"
// 	"github.com/jdetok/web/internal/errs"
// 	"github.com/jdetok/web/internal/jsonops"
// )

// // runs every interval seconds, updates if time since last update is > threshold
// func OldCheckCache(lastUpdate *time.Time, inteval time.Duration, threshold time.Duration) {
// 	e := errs.ErrInfo{Prefix: "cache check",}
// 	ticker := time.NewTicker(inteval)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		if time.Since(*lastUpdate) > threshold {
// 			fmt.Printf("refreshing cache at %v: %v since last update\n",
// 				time.Now().Format("2006-01-02 15:04:05"), threshold)
// 			updateTime, err := OldUpdateCache()
// 			if err != nil {
// 				e.Msg = "cache update failed"
// 				fmt.Println(e.Error(err))
// 			}
// 			*lastUpdate = *updateTime
// 		}
// 	}
// }

// // query all players from database & save to JSON
// func OldUpdateCache() (*time.Time, error) {
// 	e := errs.ErrInfo{Prefix: "cache update",}

// 	fmt.Println("updating /nba_rs_totals.json...")
// 	rsTots, err := db.NewSelect("select * from v_nba_rs_totals", false)
// 	if err != nil {
// 		e.Msg = "database query failed"
// 		return nil, e.Error(err)
// 	}
// 	err = jsonops.SaveJSON(cachePath + "/nba_rs_totals.json", rsTots)
// 	if err != nil {
// 		e.Msg = "saving db response to json failed"
// 		return nil, e.Error(err)
// 	}

// 	fmt.Println("updating /nba_rs_avgs.json...")
// 	rsAvgs, err := db.NewSelect("select * from v_nba_rs_avgs", false)
// 	if err != nil {
// 		e.Msg = "database query failed"
// 		return nil, e.Error(err)
// 	}
// 	err = jsonops.SaveJSON(cachePath + "/nba_rs_avgs.json", rsAvgs)
// 	if err != nil {
// 		e.Msg = "saving db response to json failed"
// 		return nil, e.Error(err)
// 	}

// 	fmt.Println("updating /nba_po_totals.json...")
// 	poTots, err := db.NewSelect("select * from v_nba_po_totals", false)
// 	if err != nil {
// 		e.Msg = "database query failed"
// 		return nil, e.Error(err)
// 	}
// 	err = jsonops.SaveJSON(cachePath + "/nba_po_totals.json", poTots)
// 	if err != nil {
// 		e.Msg = "saving db response to json failed"
// 		return nil, e.Error(err)
// 	}

// 	fmt.Println("updating /nba_po_avgs.json...")
// 	poAvgs, err := db.NewSelect("select * from v_nba_po_avgs", false)
// 	if err != nil {
// 		e.Msg = "database query failed"
// 		return nil, e.Error(err)
// 	}
// 	err = jsonops.SaveJSON(cachePath + "/nba_po_avgs.json", poAvgs)
// 	if err != nil {
// 		e.Msg = "saving db response to json failed"
// 		return nil, e.Error(err)
// 	}

// 	updateTime := time.Now()
// 	return &updateTime, nil
// }



