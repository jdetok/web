package main

import (
	"fmt"

	"github.com/jdetok/web/internal/db"
)

func main() {
	fmt.Println("Hello db")

	database, err := db.Connect()
    if err != nil {
        fmt.Printf("An error occured: %s", err)
    }

    js, err := db.Select(database)
	if err != nil {
		fmt.Printf("Error occured getting data from database")
		return
	}

	fmt.Println(js)
}