// // RUN THIS TO TEST DATABASE QUERIES
// // go run ./internal/db/test
package main

import (
	"github.com/jdetok/web/internal/db"
)

func main() {
	database := db.InitDB()
	db.SelectDB(database, `select team_name from team where lg = "NBA"`)
}

