package main

import (
	"net/http"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
	"github.com/jdetok/web/internal/logs"
)

func (app *application) getStats(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)

	lg := r.URL.Query().Get("lg")
	sType := r.URL.Query().Get("stype")
	player := r.URL.Query().Get("player")
	
	
	

	// TODO - switch statements to their own functions
	// if player != "all" {
	// 	player, err := db.ValiPlayer(player, lg)
	// 	if err != nil {
	// 	http.Error(w, "Player not found in database", 500)
	// 	}

	// }
	
	switch lg {
	case "nba":
		switch sType {
		case "tot":
			switch player {
			case "all":
				js, err := jsonops.ReadJSON(app.config.cachePath + "/nba_rs_totals.json")
				if err != nil {
					errs.HTTPErr(w, r, err)
				}
				app.JSONWriter(w, js)

			default:
				playerId, err := db.ValiPlayer(player, lg)
			 	if err != nil {
					http.Error(w, "Player not found in database", 500)
				}
				js, err := db.SelectLgPlayer(db.LgPlayerStat.Q, lg, string(playerId))
				if err != nil {
					http.Error(w, "Error getting player stats", 500)
				}
				app.JSONWriter(w, js)
			}
			
		case "avg":
			switch player {
			case "all":
				js, err := jsonops.ReadJSON(app.config.cachePath + "/nba_rs_avgs.json")
				if err != nil {
					errs.HTTPErr(w, r, err)
				}
				app.JSONWriter(w, js)

			default:
				playerId, err := db.ValiPlayer(player, lg)
			 	if err != nil {
					http.Error(w, "Player not found in database", 500)
				}
				js, err := db.SelectLgPlayer(db.LgPlayerAvg.Q, lg, string(playerId))
				if err != nil {
					http.Error(w, "Error getting player stats", 500)
				}
				app.JSONWriter(w, js)
			}
		}
	case "wnba":
		switch sType {
		case "tot":
			js, err := getResp("select * from v_wnba_rs_totals")
			if err != nil {
				errs.HTTPErr(w, r, err)
			}
			app.JSONWriter(w, js)
		case "avg":
			js, err := getResp("select * from v_wnba_rs_avgs")
			if err != nil {
				errs.HTTPErr(w, r, err)
			}
			app.JSONWriter(w, js)
		}
	}
}

func getResp(q string) ([]byte, error) {
	e := errs.ErrInfo{Prefix: "HTTP database request"}
	js, err := db.NewSelect(q, false)
	if err != nil {
		e.Msg = "failed getting response from database"
		return nil, e.Error(err)
	} 
	return js, nil
}

// 	var js []byte	
// 	var err error
// 	// switch for stype / lg

// 	// query the database for WNBA players
// 	if lg == "wnba" {
// 		js, err = db.NewSelect("select * from v_wnba_rs_totals", false)
// 		if err != nil {
// 			errs.HTTPErr(w, r, err)
// 			return
// 		} 
// 	// return the cached json for nba players
// 	} else {
// 		js = jsonops.ReadJSON(app.config.cachePath + "/nba_rs_totals.json")
// 	}
// 	app.JSONWriter(w, js)
// }