package main

import (
	"fmt"
	"net/http"

	"github.com/jdetok/go-api-jdeko.me/internal/logs"
	"github.com/jdetok/go-api-jdeko.me/internal/mariadb"
)


func (app *application) getLeaders(w http.ResponseWriter, r *http.Request) {
	league := r.URL.Query().Get("league")
	season := r.URL.Query().Get("season")
	team := r.URL.Query().Get("team")
	logs.LogHTTP(r)
	// fmt.Printf("Season: %v | Team: %v", season, team)
	
	// q := `
	// 	select * from v_szn_avgs
	// 	where lg = ?
	// 	and season_id = ?
	// 	and team = ?
	// `

	qAll := `
		select * from v_szn_avgs
		where lg = ?
		and season_id = ?
	`

	if team == "all" {
		resp, err := mariadb.DBJSONResposne(app.database, qAll, league, season)
		if err != nil {
			fmt.Printf("Error occured querying db: %v\n", err)
		}
		app.JSONWriter(w, resp)
	} else {
		resp, err := mariadb.DBJSONResposne(app.database, mariadb.Test.Q, league, season, team)
		if err != nil {
			fmt.Printf("Error occured querying db: %v\n", err)
		}
		app.JSONWriter(w, resp)
	}
}