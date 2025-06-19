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
func respFromFile(f string) ([]byte, error) {
	e := errs.ErrInfo{Prefix: "json file read"}
	js, err := jsonops.ReadJSON(env.GetString("CACHE_PATH") + f)
	if err != nil {
		e.Msg = ("error reading json file: " + f)
		return nil, e.Error(err)
	}
	return js, nil
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
				js, err := respFromFile("/nba_rs_totals.json")
				if err != nil {
					errs.HTTPErr(w, r, err)
				}
				app.JSONWriter(w, js)

			default: // NBA RS TOTALS SPECIFIC PLAYER
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

		// NBA AVERAGES
		case "avg":
			switch player {
			case "all": // NBA RS AVG ALL PLAYERS
				js, err := respFromFile("/nba_rs_avgs.json")
				if err != nil {
					errs.HTTPErr(w, r, err)
				}
				app.JSONWriter(w, js)

			default:  // NBA RS AVG SPECIFIC PLAYER
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
		// WNBA SWITCH TOTALS/AVERAGES
		switch sType {
		case "tot":
			// WNBA TOTALS
			switch player {
			case "all": // ALL WNBA TOTALS
				js, err := respFromFile("/wnba_rs_totals.json")
				if err != nil {
					errs.HTTPErr(w, r, err)
				}
				app.JSONWriter(w, js)

			default: // SPECIFIC WNBA PLAYER TOTALS
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
		// WNBA AVERAGES
		case "avg":
			switch player {
			case "all": // ALL WNBA AVERAGES
				js, err := respFromFile("/wnba_rs_avgs.json")
				if err != nil {
					errs.HTTPErr(w, r, err)
				}
				app.JSONWriter(w, js)

			default: // SPECIFIC WNBA PLAYER AVERAGES
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
	}
}