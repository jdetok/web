package main

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
	"github.com/jdetok/web/internal/logs"
	"github.com/jdetok/web/internal/mariadb"
	"github.com/jdetok/web/internal/store"
)

func (app *application) getSeasons(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	season := r.URL.Query().Get("szn")
	w.Header().Set("Content-Type", "application/json") 
	if season == "" {
		json.NewEncoder(w).Encode(app.seasons)	
	} else {
		for _, szn := range app.seasons {
			if season == szn.SeasonId {
				json.NewEncoder(w).Encode(map[string]string{
				"szn": season,
				})
			}
		}
	}	
}

// /teams for all or /teams?team=LAL for specifc team
func (app *application) getTeams(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	team := r.URL.Query().Get("team")
	w.Header().Set("Content-Type", "application/json") 
	if team == "" {
		json.NewEncoder(w).Encode(app.teams)	
	} else {
		for _, tm := range app.teams {
			if team == tm.TeamAbbr {
				tm.LogoUrl = tm.MakeLogoUrl()
				json.NewEncoder(w).Encode(tm)
			}
		}
	}

	// w.Header().Set("Content-Type", "application/json") 
	// json.NewEncoder(w).Encode(app.teams)
}
// func (app *application) getTeam(w http.ResponseWriter, r *http.Request) {
// 	logs.LogHTTP(r)
// 	team := r.URL.Query().Get("team")
// 	// logs.LogDebug("Team Requested: " + team)

// 	w.Header().Set("Content-Type", "application/json") 
// 	json.NewEncoder(w).Encode(`team: ${team}`)
// }
func (app *application) getRandomPlayer(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	numPlayers := len(app.players)
	randNum := rand.IntN(numPlayers)

	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(map[string]string{
		"playerId": strconv.FormatUint(app.players[randNum].PlayerId, 10),
		"player": app.players[randNum].Name,
		"league": app.players[randNum].League,
	})
	// random number in range of len(players) to return random player
}

func (app *application) getPlayerId(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	player := r.URL.Query().Get("player")
	logs.LogDebug("Player Requested: " + player)
	// playerId := db.ValiPlayer(app.database, &w, player)
	
	playerId := store.SearchPlayers(app.players, player)
	logs.LogDebug("PlayerId Return: " + playerId)
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(map[string]string{
		"playerId": playerId,
	})
}

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
	
	switch lg { // OUTER SWTICH: NBA/WNBA	
	case "nba": // NBA SWITCH TOTALS/AVERAGES
		switch sType {
		case "tot": // NBA TOTALS
			switch player {
			case "all": // NBA RS TOTALS ALL PLAYERS
				js := respFromFile(&w, "/nba_rs_totals.json")
				app.JSONWriter(w, js)
			default: // NBA RS TOTALS SPECIFIC PLAYER
				playerId := store.SearchPlayers(app.players, player)
				js := mariadb.SelectLgPlayer(app.database, &w, 
					mariadb.LgPlayerStat.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		case "avg": // NBA AVERAGES
			switch player {
			case "all": // NBA RS AVG ALL PLAYERS
				js := respFromFile(&w, "/nba_rs_avgs.json")
				app.JSONWriter(w, js)
			default:  // NBA RS AVG SPECIFIC PLAYER
				playerId := store.SearchPlayers(app.players, player)
				js := mariadb.SelectLgPlayer(app.database, &w, 
					mariadb.LgPlayerAvg.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		}
	case "wnba":
		switch sType { // WNBA SWITCH TOTALS/AVERAGES
		case "tot": // WNBA TOTALS
			switch player {  // ALL WNBA TOTALS
			case "all":
				js := respFromFile(&w, "/wnba_rs_totals.json")
				app.JSONWriter(w, js)
			default: // SPECIFIC WNBA PLAYER TOTALS
				playerId := store.SearchPlayers(app.players, player)
				js := mariadb.SelectLgPlayer(app.database, &w, 
					mariadb.LgPlayerStat.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		case "avg": // WNBA AVERAGES
			switch player {
			case "all": // ALL WNBA AVERAGES
				js := respFromFile(&w, "/wnba_rs_avgs.json")
				app.JSONWriter(w, js)
			default: // SPECIFIC WNBA PLAYER AVERAGES
				playerId := store.SearchPlayers(app.players, player)
				js := mariadb.SelectLgPlayer(app.database, &w, mariadb.LgPlayerAvg.Q, lg, string(playerId))
				app.JSONWriter(w, js)
			}
		}
	}
}