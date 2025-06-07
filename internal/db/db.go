package db

import (
	"database/sql"
	"encoding/json"
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
	
	js, err := RowsToJSON(rows)
	if err != nil {
		fmt.Println("Error occured")
	}
	fmt.Println(string(js))
	
}

func RowsToJSON(rows *sql.Rows) ([]byte, error) {
	colTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	count := len(colTypes)
	
	finalRows := []interface{}{};

	for rows.Next() {
		scanArgs := make([]interface{}, count)

		for i, v := range colTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4", "INT8", "INT":
				scanArgs[i] = new(sql.NullInt64)
			case "FLOAT", "FLOAT8", "FLOAT4":
				scanArgs[i] = new(sql.NullFloat64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		masterData := map[string]interface{}{}

		for i, v := range colTypes {
			if z, ok := scanArgs[i].(*sql.NullBool); ok {
				if z.Valid {
					masterData[v.Name()] = z.Bool
				} else {
					masterData[v.Name()] = nil
				}	
			continue
			}

			if z, ok := scanArgs[i].(*sql.NullString); ok {
				if z.Valid {
					masterData[v.Name()] = z.String
				} else {
					masterData[v.Name()] = nil
				}	
			continue
			}

			if z, ok := scanArgs[i].(*sql.NullInt64); ok {
				if z.Valid {
					masterData[v.Name()] = z.Int64
				} else {
					masterData[v.Name()] = nil
				}	
			continue
			}

			if z, ok := scanArgs[i].(*sql.NullFloat64); ok {
				if z.Valid {
					masterData[v.Name()] = z.Float64
				} else {
					masterData[v.Name()] = nil
				}	
			continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}

	js, err :=  json.Marshal(finalRows)
	if err != nil {
		return nil, err
	}

	return js, nil
}