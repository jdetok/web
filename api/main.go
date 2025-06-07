package main

import (
	"log"

	"github.com/jdetok/web/internal/db"
) 

func main() {

    // configs go here - 8080 for testing, will derive real vals from environment
    // cfg := config{
    //     addr: ":8080",
    // }

    // // initialize the app with the configs
    // app := &application{
    //     config: cfg,
    // }

    // // mount initializes mux (serves/routes HTTP) & handlers
    // mux := app.mount()

    // // run the server with the initialized mux 
    // log.Fatal(app.run(mux))

    // database testing
    database, err := db.Connect()
    if err != nil {
        log.Printf("An error occured: %s", err)
    }

    db.Select(database)

}
