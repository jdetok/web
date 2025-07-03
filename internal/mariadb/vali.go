package mariadb

import (
	"database/sql"
	"net/http"

	"github.com/jdetok/go-api-jdeko.me/internal/errs"
)

func ValiPlayer(db *sql.DB, w *http.ResponseWriter, player string) []byte {
	e := errs.ErrInfo{Prefix: "players select",}
	player_id, err := SelectPlayers(db, player)
	if err != nil {
		e.Msg = "player validation failed"
		errs.HTTPErr(*w, e.Error(err))
	}
	
	return player_id
}