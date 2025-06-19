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
	// TODO - build in selector in HTML for avg/total
	lg := r.URL.Query().Get("lg")
	var js []byte	
	var err error
	// query the database for WNBA players
	if lg == "WNBA" {
		js, err = db.NewSelect("select * from v_wnba_rs_totals", false)
		if err != nil {
			errs.HTTPErr(w, r, err)
			return
		} 
	// return the cached json for nba players
	} else {
		js = jsonops.ReadJSON(app.config.cachePath + "/nba_rs_totals.json")
	}
	app.JSONWriter(w, js)
}