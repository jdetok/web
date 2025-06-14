package main

import (
	"fmt"

	"github.com/jdetok/web/internal/db"
	"github.com/jdetok/web/internal/jsonops"
)

func main() {
	fmt.Println("TESTING DATABASE PACKAGE")

	database, err := db.Connect()
    if err != nil {
        fmt.Printf("An error occured: %s", err)
    }
// pass true inse
    js, err := db.Select(database, db.CarrerStats, true)
	if err != nil {
		fmt.Printf("Error occured getting data from database: %s", err)
		return
	}

	// js is a []byte - print it as a string
	fmt.Println(string(js))

	// save to file
	jsonops.SaveJSON("json/db/player_career.json", js)
}

// recovery

// func main() {
// 	fmt.Println("TESTING DATABASE PACKAGE")

// 	database, err := db.Connect()
//     if err != nil {
//         fmt.Printf("An error occured: %s", err)
//     }

//     js, err := db.TestSelect(database)
// 	if err != nil {
// 		fmt.Printf("Error occured getting data from database: %s", err)
// 		return
// 	}

// 	// js is a []byte - print it as a string
// 	fmt.Println(string(js))
// }