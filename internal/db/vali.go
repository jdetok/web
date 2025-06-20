package db

import (
	"database/sql"
	"net/http"

	"github.com/jdetok/web/internal/errs"
)

func ValiPlayer(db *sql.DB, w *http.ResponseWriter, player, lg string) []byte {
	e := errs.ErrInfo{Prefix: "players select",}
	player_id, err := SelectPlayers(db, player, lg)
	if err != nil {
		e.Msg = "player validation failed"
		errs.HTTPErr(*w, e.Error(err))
	}
	
	return player_id
}