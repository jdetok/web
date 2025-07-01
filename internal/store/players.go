package store

// import (
// 	"database/sql"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/jdetok/web/internal/errs"
// )

// type Player struct {
// 	PlayerId uint64
// 	Name string
// 	League string
// }

// // QUERY FOR PLAYER ID, PLAYER AND SAVE TO A LIST OF PLAYER STRUCTS
// func GetPlayers(db *sql.DB) ([]Player, error) {
// 	fmt.Println("querying players & saving to struct")
// 	e := errs.ErrInfo{Prefix: "saving players to structs"}
// 	rows, err := db.Query(`
// 		select player_id, player, lg
// 		from player
// 		where lg in ("NBA", "WNBA")
// 		group by player_id, player, lg
// 	`)
// 	if err != nil {
// 		e.Msg = "query failed"
// 		return nil, e.Error(err)
// 	}
// 	var players []Player
// 	for rows.Next() {
// 		var p Player
// 		rows.Scan(&p.PlayerId, &p.Name, &p.League)
// 		// convert to lowercase to match requests
// 		p.Name = strings.ToLower(p.Name)
// 		p.League = strings.ToLower(p.League)
// 		players = append(players, p)
// 	}
// 	return players, nil
// }

// func SearchPlayers(players []Player, pSearch string) string {
// 	for _, p := range players {
// 		if p.Name == pSearch { // return match playerid (uint32) as string
// 			return strconv.FormatUint(p.PlayerId, 10)
// 		}
// 	}
// 	return ""
// }