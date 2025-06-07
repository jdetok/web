package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jdetok/web/internal/env"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		 log.Println("dotenv didn't work")
	}

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
	fmt.Println("Connected to MariaDB!")
	
	return db, nil
}

func Select(db *sql.DB) {
	// q := `
	// select a.player, b.team
	// from player a
	// inner join team b on b.team_id = a.team_id
	// where a.active = 1
	// and a.lg = "NBA"
	// and player = ?
	// `

	q := `
select a.player, b.team, sum(c.pts) as career_pts
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
	}

	for rows.Next() {
		var player, team string
		var career_pts int
		// err := rows.Scan(&player)
		err := rows.Scan(&player, &team, &career_pts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Player: %s | Team: %s | Career Points: %d\n", player, team, career_pts)
	}

// PRINT THE COLUMNS
	// cols, err := rows.Columns()
	// if err != nil {
	// 	print("Error getting columns")
	// }

	// for i := range(len(cols)) {
	// 	fmt.Printf("Column: %v\n", cols[i])
	// }

	// fmt.Printf("Columns: %v\n", cols)
}