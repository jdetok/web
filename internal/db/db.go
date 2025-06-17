package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/errs"
)

func Connect() (*sql.DB, error) {
// DECLARE ERRINFO TYPE
	e := errs.ErrInfo{Prefix: "database conenction",}

// get conn. vars from .env & build connection string
	dbUser := env.GetString("DB_USER")
	dbHost := env.GetString("DB_HOST")
	database := env.GetString("DB")
	connStr := dbUser + "@tcp(" + dbHost + ")/" + database

// attempt to open database connection
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		e.Msg = "sql.Open() failed"
		return nil, e.Error(err)
	}
	
// ping to confirm connection - exits with error if ping is not successful
	if err := db.Ping(); err != nil {
		e.Msg = "db.Ping() failed"
		return nil, e.Error(err)
	}
// return the open database as *sql.DB type
	return db, nil
}

func Select(db *sql.DB, q string, indent_resp bool) ([]byte, error) {
// query db - returns sql.Rows type
	e := errs.ErrInfo{Prefix: "database query",}
	rows, err := db.Query(q)
	if err != nil {
		e.Msg = "db.Query failed"
		return nil, e.Error(err)
	}
	
// convert the response to json
	js, err := RowsToJSON(rows, indent_resp)
	if err != nil {
		e.Msg = "func RowsToJSON() failed"
		return nil, e.Error(err)
	}
// return the response as json
	return js, nil
}

func SelectArg(db *sql.DB, q string, indent_resp bool, r string) ([]byte, error) {
// query db - returns sql.Rows type
	e := errs.ErrInfo{Prefix: "database query (arg)",}
	rows, err := db.Query(q, r)
	if err != nil {
		e.Msg = "db.Query failed"
		return nil, e.Error(err)
	}
	
// return the response as json
	js, err := RowsToJSON(rows, indent_resp)
	if err != nil {
		e.Msg = "func RowsToJSON() failed"
		return nil, e.Error(err)
	}
	return js, nil
}

func SelectArgs(db *sql.DB, q string, indent_resp bool, r1, r2 string) ([]byte, error) {
// query db - returns sql.Rows type
	e := errs.ErrInfo{Prefix: "database query (2 args)",}
	rows, err := db.Query(q, r1, r2)
	if err != nil {
		e.Msg = "db.Query failed"
		return nil, e.Error(err)
	}
	
// return the response as json
	js, err := RowsToJSON(rows, indent_resp)
	if err != nil {
		e.Msg = "func RowsToJSON() failed"
		return nil, e.Error(err)
	}
	return js, nil
}