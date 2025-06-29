// // RUN THIS TO TEST DATABASE QUERIES
// // go run ./internal/db/test
package main

import (
	"database/sql"
	"fmt"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/errs"
)

type Player struct {
	PlayerId uint32
	Name string
}

// QUERY FOR PLAYER ID, PLAYER AND SAVE TO A LIST OF PLAYER STRUCTS
func GetPlayers(db *sql.DB) []Player {
	e := errs.ErrInfo{Prefix: "saving players to structs"}
	rows, err := db.Query(`
		select player_id, player 
		from player 
		where lg in ("NBA", "WNBA") 
		group by player_id, player
	`)
	if err != nil {
		e.Msg = "query failed"
		e.Error(err)
	}
	var players []Player
	for rows.Next() {
		var p Player
		rows.Scan(&p.PlayerId, &p.Name)
		players = append(players, p)
	}
	return players
} 

func SearchPlayers(players []Player, pSearch string) bool {
	for _, p := range players {
		if p.Name == pSearch {
			fmt.Println(p)
			return true
		}
	}
	return false
}

func main() {
	database := db.InitDB()
	players := GetPlayers(database)

	if SearchPlayers(players, "Caitlin Clark") {
		fmt.Println("Player exists")
	}
	
	// fmt.Println(players)
	// fmt.Println(len(players))


	// js, _ := db.SelectDB(database, `select player_id, player from player where lg in ("NBA", "WNBA") group by player_id, player`)
	// fmt.Println(string(js))
	// fmt.Println(len(js))
	// db.SelectDB(database, `select team_name from team where lg = "NBA"`)
}

