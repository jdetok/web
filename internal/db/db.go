package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB {
	dbUser := "jdeto:"
	dbHost := "pi-jdeto"
	dbPort := "3307"
	database := "mysql"

	connStr := dbUser + "@tcp(" + dbHost + ":" + dbPort + ")/" + database

	// db, err := sql.Open("mysql", "jdeto:@tcp(pi-jdeto:3307)/mysql")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	} // ping to confirm connection02
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MariaDB!")
	return db
}

func dbUsers(db *sql.DB) {
	rows, err := db.Query("SELECT user, host FROM mysql.user")
	if err != nil {
		fmt.Printf("Error querying: %s", err)
		log.Fatal(err)
	}
	for rows.Next() {
		var user, host string
		err := rows.Scan(&user, &host)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User: %s | Host: %s\n", user, host)
	}
}