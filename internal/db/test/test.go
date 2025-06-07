package main

import (
	"fmt"

	"github.com/jdetok/web/internal/db"
)

func main() {
	fmt.Println("TESTING DATABASE PACKAGE")

	database, err := db.Connect()
    if err != nil {
        fmt.Printf("An error occured: %s", err)
    }

    js, err := db.Select(database)
	if err != nil {
		fmt.Printf("Error occured getting data from database: %s", err)
		return
	}

	// js is a []byte - print it as a string
	fmt.Println(string(js))
}