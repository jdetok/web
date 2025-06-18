package main

import (
	"fmt"
	"net/http"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/jsonops"
	"github.com/jdetok/web/internal/logs"
)

func (app *application) selectPlayersH(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	
	lg := r.URL.Query().Get("lg")
	var js []byte	

	// query the database for WNBA players
	if lg == "WNBA" {
		database, err := db.Connect()
		if err != nil {
			// TODO!!!
			errs.HTTPErr(w, r, err)
			return
		}

		js, err = db.SelectArg(database, db.CarrerStatsByLg, false, lg)
		if err != nil {
			errs.HTTPErr(w, r, err)
			return
		}
	} else {
	// read cached json for nba 
		js = jsonops.ReadJSON(app.config.cachePath + "/players.json")
	}
	app.JSONWriter(w, js)
}

// need to write function to accept player name as string, search if in DB, then query this with id
func (app *application) selectPlayerH(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	lg := r.URL.Query().Get("lg")
	// TODO - middleware func that accepts the player name entered and returns whether they are valid
	player := r.URL.Query().Get("player")

	database, err := db.Connect()
    if err != nil {
		errs.HTTPErr(w, r, err)
    }

    js, err := db.SelectArgs(database, db.LgPlayerStat.Q, false, lg, player)
	if err != nil {
		errs.HTTPErr(w, r, err)
	}
	app.JSONWriter(w, js)
}

// need to write function to accept player name as string, search if in DB, then query this with id
func (app *application) selectPlayerHTest(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
	lg := r.URL.Query().Get("lg")
	// TODO - middleware func that accepts the player name entered and returns whether they are valid
	player := r.URL.Query().Get("player")
	fmt.Println(player)
	fmt.Println(lg)
	playerId, err := db.ValiPlayer(player, lg)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("no player with that name"))
	}

	database, err := db.Connect()
    if err != nil {
		errs.HTTPErr(w, r, err)
    }

	js, err := db.SelectArgs(database, db.LgPlayerStat.Q, false, lg, string(playerId))
	if err != nil {
		errs.HTTPErr(w, r, err)
	}

	fmt.Println(string(js))
	app.JSONWriter(w, js)
	
	
	// database, err := db.Connect()
    // if err != nil {
	// 	errs.HTTPErr(w, r, err)
    // }

    // js, err := db.SelectArgs(database, db.LgPlayerStat.Q, false, lg, player)
	// if err != nil {
	// 	errs.HTTPErr(w, r, err)
	// }
	
}

func (app *application) selectGameHandler(w http.ResponseWriter, r *http.Request) {
	logs.LogHTTP(r)
		
    database, err := db.Connect()
    if err != nil {
        errs.HTTPErr(w, r, err)
    }

    js, err := db.Select(database, db.Games, true)
	if err != nil {
		errs.HTTPErr(w, r, err)
	}
	app.JSONWriter(w, js)
}
