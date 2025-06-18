package db

import (
	"github.com/jdetok/web/internal/errs"
)

func ValiPlayer(player, lg string) ([]byte, error) {
	e := errs.ErrInfo{Prefix: "players select",}
	player_id, err := SelectPlayers(player, lg)
	if err != nil {
		e.Msg = "player validation failed"
		return nil, e.Error(err)
	}
	
	return player_id, nil
}