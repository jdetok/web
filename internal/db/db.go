package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jdetok/web/internal/env"
)

func Connect() (*sql.DB, error) {
// get components of connection string from .env
	dbUser := env.GetString("DB_USER")
	dbHost := env.GetString("DB_HOST")
	dbPort := env.GetString("DB_PORT")
	database := env.GetString("DB")

// build connection string & attempt to connect
	connStr := dbUser + "@tcp(" + dbHost + ":" + dbPort + ")/" + database
	
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		return nil, err
	}
// ping to confirm connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// fmt.Println("Connected to MariaDB!")
	
	return db, nil
}

// TODO - split out player part & make this a more general select
func Select(db *sql.DB, q string, indent_resp bool) ([]byte, error) {
// query db - returns sql.Rows type
	rows, err := db.Query(q)
	if err != nil {
		fmt.Printf("Error querying: %s", err)
		log.Fatal(err)
		return nil, err
	}
	
// return the response as json
	js, err := RowsToJSON(rows, indent_resp)
	if err != nil {
		fmt.Println("Error occured converting to JSON")
		return nil, err
	}
	return js, nil
}

func SelectArg(db *sql.DB, q string, indent_resp bool, r string) ([]byte, error) {
// query db - returns sql.Rows type
	

	rows, err := db.Query(q, r)
	if err != nil {
		fmt.Printf("Error querying: %s", err)
		log.Fatal(err)
		return nil, err
	}
	
// return the response as json
	js, err := RowsToJSON(rows, indent_resp)
	if err != nil {
		fmt.Println("Error occured converting to JSON")
		return nil, err
	}
	return js, nil
}

func SelectArgs(db *sql.DB, q string, indent_resp bool, r1, r2 string) ([]byte, error) {
// query db - returns sql.Rows type
	

	rows, err := db.Query(q, r1, r2)
	if err != nil {
		fmt.Printf("Error querying: %s", err)
		log.Fatal(err)
		return nil, err
	}
	
// return the response as json
	js, err := RowsToJSON(rows, indent_resp)
	if err != nil {
		fmt.Println("Error occured converting to JSON")
		return nil, err
	}
	return js, nil
}

// TODO - split out player part & make this a more general select
func TestSelect(db *sql.DB) ([]byte, error) {
	q := `
	select a.player, b.team, sum(c.pts) as pts, avg(c.pts) as pts_pg 
	from player a
	inner join team b on b.team_id = a.team_id
	inner join p_box c on c.player_id = a.player_id
	inner join season d on d.season_id = c.season_id
	where a.active = 1
	and a.lg = "NBA"
	and d.season like "%RS"
	and player = ?
	group by a.player, b.team	
	`

	var plr string = "LeBron James"

	rows, err := db.Query(q, plr)
	if err != nil {
		fmt.Printf("Error querying: %s", err)
		log.Fatal(err)
		return nil, err
	}
	
	js, err := RowsToJSON(rows, false)
	if err != nil {
		fmt.Println("Error occured")
		return nil, err
	}

	return js, nil
}