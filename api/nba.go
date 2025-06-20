package main

import (
	"net/http"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
	"github.com/jdetok/web/internal/logs"
)

// RETURN CONTENTS OF JSON FILE AS []byte
func respFromFile(w *http.ResponseWriter, f string) []byte {
	e := errs.ErrInfo{Prefix: "json file read"}
	js, err := jsonops.ReadJSON(env.GetString("CACHE_PATH") + f)
	if err != nil {
		e.Msg = ("error reading json file: " + f)
		errs.HTTPErr(*w, e.Error(err))
	}
	return js
}

// HANDLE /bball/players REQUESTS
// EX. QUERY STRING - ALL PLAYERS: ?lg=nba&stype=tot&player=all
// EX. QUERY STRING - SPECIFIC PLAYER: ?lg=nba&stype=avg&player=tyrese%20haliburton
func (app *application) getStats(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)

	lg := r.URL.Query().Get("lg")
	sType := r.URL.Query().Get("stype")
	player := r.URL.Query().Get("player")
	// OUTER SWTICH: NBA/WNBA	
	switch lg {
	case "nba":
		// NBA SWITCH TOTALS/AVERAGES
		switch sType {
		case "tot":
			// NBA TOTALS
			switch player {
			case "all": // NBA RS TOTALS ALL PLAYERS
				js := respFromFile(&w, "/nba_rs_totals.json")
				app.JSONWriter(w, js)
			default: // NBA RS TOTALS SPECIFIC PLAYER
				playerId := db.ValiPlayer(&w, player, lg)
				js := db.SelectLgPlayer(&w, db.LgPlayerStat.Q, lg, string(playerId))

				// w.Header().Set("Content-Type", "image/png")
				// http.ServeFile(w, r, env.GetString("HS_PATH") + string(playerId) + ".png")
				// fmt.Println(env.GetString("HS_PATH") + string(playerId) + ".png")
				app.JSONWriter(w, js)
			}

		// NBA AVERAGES
		case "avg":
			switch player {
			case "all": // NBA RS AVG ALL PLAYERS
				js := respFromFile(&w, "/nba_rs_avgs.json")
				app.JSONWriter(w, js)

			default:  // NBA RS AVG SPECIFIC PLAYER
				playerId := db.ValiPlayer(&w, player, lg)
				js := db.SelectLgPlayer(&w, db.LgPlayerAvg.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		}
	case "wnba":
		// WNBA SWITCH TOTALS/AVERAGES
		switch sType {
		case "tot":
			// WNBA TOTALS
			switch player {
			case "all": // ALL WNBA TOTALS
				js := respFromFile(&w, "/wnba_rs_totals.json")
				app.JSONWriter(w, js)

			default: // SPECIFIC WNBA PLAYER TOTALS
				playerId := db.ValiPlayer(&w, player, lg)
				js := db.SelectLgPlayer(&w, db.LgPlayerAvg.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		// WNBA AVERAGES
		case "avg":
			switch player {
			case "all": // ALL WNBA AVERAGES
				js := respFromFile(&w, "/wnba_rs_avgs.json")
				app.JSONWriter(w, js)

			default: // SPECIFIC WNBA PLAYER AVERAGES
				playerId := db.ValiPlayer(&w, player, lg)
				js := db.SelectLgPlayer(&w, db.LgPlayerAvg.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		}
	}
}