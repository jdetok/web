package db

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/errs"
)

func InitDB() *sql.DB {
	// e := errs.ErrInfo{Prefix: "database conenction",}

// get conn. vars from .env & build connection string
	dbUser := env.GetString("DB_USER")
	dbHost := env.GetString("DB_HOST")
	database := env.GetString("DB")
	connStr := dbUser + "@tcp(" + dbHost + ")/" + database

// attempt to open database connection
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
	}
	
// ping to confirm connection - exits with error if ping is not successful
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(200)
	return db
}
func SelectDB(database *sql.DB, q string) ([]byte, error) {
	e := errs.ErrInfo{Prefix: "database query",}
	rows, err := database.Query(q)
	if err != nil {
		fmt.Println(err)
	}	
	js, err := RowsToJSON(rows, false)
	if err != nil {
		e.Msg = "func RowsToJSON() failed"
		return nil, e.Error(err)
	}
// return the response as json
	return js, nil
}

func SelectLgPlayer(db *sql.DB, w *http.ResponseWriter, q string, lg string, pl string) []byte {
	e := errs.ErrInfo{Prefix: "database query (arg)",}
	
	rows, err := db.Query(q, lg, pl)
	if err != nil {
		e.Msg = "db.Query failed"
		errs.HTTPErr(*w, e.Error(err))
	}
	
// return the response as json
	js, err := RowsToJSON(rows, false)
	if err != nil {
		e.Msg = "func RowsToJSON() failed"
		errs.HTTPErr(*w, e.Error(err))
	}
	return js
}

// accept the player from query string, query db & return player id if player exists
func SelectPlayers(db *sql.DB, player, lg string) ([]byte, error) {
	e := errs.ErrInfo{Prefix: "players select",}
	var playerId []byte
	err := db.QueryRow(Players.Q, player, lg).Scan(&playerId)
	if err != nil {
		e.Msg = "query failed"
		return nil, e.Error(err)
	}
	return playerId, nil
}

func SelectList(db *sql.DB, q string) ([]string, error) {
	e := errs.ErrInfo{Prefix: "database query - list of players",}
	rows, err := db.Query(q)
	if err != nil {
		e.Msg = "query failed"
		return nil, e.Error(err)
	}
	
	var playerIds []string
	for rows.Next() {
		var pId string
		if err := rows.Scan(&pId); err != nil {
			fmt.Println("error scanning rows")
		}
		playerIds = append(playerIds, pId)
	}
	return playerIds, nil
}