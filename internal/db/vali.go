package db

import (
	"net/http"

	"github.com/jdetok/web/internal/errs"
)

func ValiPlayer(w *http.ResponseWriter, player, lg string) []byte {
	e := errs.ErrInfo{Prefix: "players select",}
	player_id, err := SelectPlayers(player, lg)
	if err != nil {
		e.Msg = "player validation failed"
		errs.HTTPErr(*w, e.Error(err))
	}
	
	return player_id
}