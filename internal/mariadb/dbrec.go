package mariadb

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jdetok/go-api-jdeko.me/internal/env"
// 	"github.com/jdetok/go-api-jdeko.me/internal/errs"
// )

// func InitDB() *sql.DB {
// 	// e := errs.ErrInfo{Prefix: "database conenction",}

// // get conn. vars from .env & build connection string
// 	dbUser := env.GetString("DB_USER")
// 	dbHost := env.GetString("DB_HOST")
// 	database := env.GetString("DB")
// 	connStr := dbUser + "@tcp(" + dbHost + ")/" + database

// // attempt to open database connection
// 	db, err := sql.Open("mysql", connStr)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// // ping to confirm connection - exits with error if ping is not successful
// 	if err := db.Ping(); err != nil {
// 		fmt.Println(err)
// 	}
// 	db.SetMaxIdleConns(20)
// 	db.SetMaxOpenConns(200)
// 	return db
// }
// func SelectConnPool(database *sql.DB, q string) ([]byte, error) {
// 	e := errs.ErrInfo{Prefix: "database query",}
// 	rows, err := database.Query(q)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	js, err := RowsToJSON(rows, false)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		return nil, e.Error(err)
// 	}
// // return the response as json
// 	return js, nil
// }
// // func OpenDB() (*sql.DB, error) {
// // // DECLARE ERRINFO TYPE
// // 	e := errs.ErrInfo{Prefix: "database conenction",}

// // // get conn. vars from .env & build connection string
// // 	dbUser := env.GetString("DB_USER")
// // 	dbHost := env.GetString("DB_HOST")
// // 	database := env.GetString("DB")
// // 	connStr := dbUser + "@tcp(" + dbHost + ")/" + database

// // // attempt to open database connection
// // 	db, err := sql.Open("mysql", connStr)
// // 	if err != nil {
// // 		e.Msg = "sql.Open() failed"
// // 		return nil, e.Error(err)
// // 	}

// // // ping to confirm connection - exits with error if ping is not successful
// // 	if err := db.Ping(); err != nil {
// // 		e.Msg = "db.Ping() failed"
// // 		return nil, e.Error(err)
// // 	}
// // // return the open database as *sql.DB type
// // 	return db, nil
// // }

// func Connect() (*sql.DB, error) {
// // DECLARE ERRINFO TYPE
// 	e := errs.ErrInfo{Prefix: "database conenction",}

// // get conn. vars from .env & build connection string
// 	dbUser := env.GetString("DB_USER")
// 	dbHost := env.GetString("DB_HOST")
// 	database := env.GetString("DB")
// 	connStr := dbUser + "@tcp(" + dbHost + ")/" + database

// // attempt to open database connection
// 	db, err := sql.Open("mysql", connStr)
// 	if err != nil {
// 		e.Msg = "sql.Open() failed"
// 		return nil, e.Error(err)
// 	}

// // ping to confirm connection - exits with error if ping is not successful
// 	if err := db.Ping(); err != nil {
// 		e.Msg = "db.Ping() failed"
// 		return nil, e.Error(err)
// 	}
// // return the open database as *sql.DB type
// 	return db, nil
// }

// func NewSelect(db sql.DB, q string, indent_resp bool) ([]byte, error) {
// // query db - returns sql.Rows type
// 	// e := errs.ErrInfo{Prefix: "database query",}
// 	// db, err := Connect()
// 	// if err != nil {
// 	// 	e.Msg = "database connection failed"
// 	// 	return nil, e.Error(err)
// 	// }
// 	// defer db.Close()
// 	e := errs.ErrInfo{Prefix: "database query",}
// 	rows, err := db.Query(q)
// 	if err != nil {
// 		e.Msg = "query failed"
// 		return nil, e.Error(err)
// 	}
// // convert the response to json
// 	js, err := RowsToJSON(rows, indent_resp)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		return nil, e.Error(err)
// 	}
// // return the response as json
// 	return js, nil
// }

// func NewOldSelect(q string, indent_resp bool) ([]byte, error) {
// // query db - returns sql.Rows type
// 	e := errs.ErrInfo{Prefix: "database query",}
// 	db, err := Connect()
// 	if err != nil {
// 		e.Msg = "database connection failed"
// 		return nil, e.Error(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query(q)
// 	if err != nil {
// 		e.Msg = "query failed"
// 		return nil, e.Error(err)
// 	}
// // convert the response to json
// 	js, err := RowsToJSON(rows, indent_resp)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		return nil, e.Error(err)
// 	}
// // return the response as json
// 	return js, nil
// }

// func SelectLgPlayer(w *http.ResponseWriter, q string, lg string, pl string) []byte {
// // query db - returns sql.Rows type
// 	e := errs.ErrInfo{Prefix: "database query (arg)",}
// 	db, err := Connect()
// 	if err != nil {
// 		e.Msg = "database connection failed"
// 		errs.HTTPErr(*w, e.Error(err))
// 	}

// 	rows, err := db.Query(q, lg, pl)
// 	if err != nil {
// 		e.Msg = "db.Query failed"
// 		errs.HTTPErr(*w, e.Error(err))
// 	}

// // return the response as json
// 	js, err := RowsToJSON(rows, false)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		errs.HTTPErr(*w, e.Error(err))
// 	}
// 	return js
// }

// // func SelectLgPlayer(q string, lg string, pl string) ([]byte, error) {
// // // query db - returns sql.Rows type
// // 	e := errs.ErrInfo{Prefix: "database query (arg)",}
// // 	db, err := Connect()
// // 	if err != nil {
// // 		e.Msg = "database connection failed"
// // 		return nil, e.Error(err)
// // 	}

// // 	rows, err := db.Query(q, lg, pl)
// // 	if err != nil {
// // 		e.Msg = "db.Query failed"
// // 		return nil, e.Error(err)
// // 	}

// // // return the response as json
// // 	js, err := RowsToJSON(rows, false)
// // 	if err != nil {
// // 		e.Msg = "func RowsToJSON() failed"
// // 		return nil, e.Error(err)
// // 	}
// // 	return js, nil
// // }

// // accept the player from query string, query db & return player id if player exists
// func SelectPlayers(player, lg string) ([]byte, error) {
// 	e := errs.ErrInfo{Prefix: "players select",}
// 	db, err := Connect()
// 	if err != nil {
// 		e.Msg = "database connection failed"
// 		return nil, e.Error(err)
// 	}
// 	var playerId []byte
// 	err = db.QueryRow(Players.Q, player, lg).Scan(&playerId)
// 	if err != nil {
// 		e.Msg = "query failed"
// 		return nil, e.Error(err)
// 	}
// 	return playerId, nil
// }

// func Select(db *sql.DB, q string, indent_resp bool) ([]byte, error) {
// // query db - returns sql.Rows type
// 	e := errs.ErrInfo{Prefix: "database query",}
// 	rows, err := db.Query(q)
// 	if err != nil {
// 		e.Msg = "query failed"
// 		return nil, e.Error(err)
// 	}

// // convert the response to json
// 	js, err := RowsToJSON(rows, indent_resp)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		return nil, e.Error(err)
// 	}
// // return the response as json
// 	return js, nil
// }

// func SelectList(q string) ([]string, error) {
// // query db - returns sql.Rows type
// 	e := errs.ErrInfo{Prefix: "database query - list of players",}

// 	db, err := Connect()
// 	if err != nil {
// 		e.Msg = "database connection failed"
// 		return nil, e.Error(err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(q)
// 	if err != nil {
// 		e.Msg = "query failed"
// 		return nil, e.Error(err)
// 	}

// 	var playerIds []string
// 	for rows.Next() {
// 		var pId string
// 		if err := rows.Scan(&pId); err != nil {
// 			fmt.Println("error scanning rows")
// 		}
// 		playerIds = append(playerIds, pId)
// 	}
// 	return playerIds, nil
// }

// // convert the response to json
// 	// js, err := RowsToJSON(rows, indent_resp)
// 	// if err != nil {
// 	// 	e.Msg = "func RowsToJSON() failed"
// 	// 	return nil, e.Error(err)
// 	// }
// // return the response as json

// func SelectArg(db *sql.DB, q string, indent_resp bool, r string) ([]byte, error) {
// // query db - returns sql.Rows type
// 	e := errs.ErrInfo{Prefix: "database query (arg)",}
// 	rows, err := db.Query(q, r)
// 	if err != nil {
// 		e.Msg = "db.Query failed"
// 		return nil, e.Error(err)
// 	}

// // return the response as json
// 	js, err := RowsToJSON(rows, indent_resp)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		return nil, e.Error(err)
// 	}
// 	return js, nil
// }

// func SelectArgs(db *sql.DB, q string, indent_resp bool, r1, r2 string) ([]byte, error) {
// // query db - returns sql.Rows type
// 	e := errs.ErrInfo{Prefix: "database query (2 args)",}
// 	rows, err := db.Query(q, r1, r2)
// 	if err != nil {
// 		e.Msg = "db.Query failed"
// 		return nil, e.Error(err)
// 	}

// // return the response as json
// 	js, err := RowsToJSON(rows, indent_resp)
// 	if err != nil {
// 		e.Msg = "func RowsToJSON() failed"
// 		return nil, e.Error(err)
// 	}
// 	return js, nil
// }