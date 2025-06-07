package main

import (
	"log"

	"github.com/jdetok/web/internal/env"
	"github.com/joho/godotenv"
) 

func main() {

    err := godotenv.Load()
	if err != nil {
		 log.Println("dotenv didn't work")
	}

    // configs go here - 8080 for testing, will derive real vals from environment
    cfg := config{
        addr: env.GetString("SRV_IP"),
    }

    // initialize the app with the configs
    app := &application{
        config: cfg,
    }

    // mount initializes mux (serves/routes HTTP) & handlers
    mux := app.mount()

    // run the server with the initialized mux 
    log.Fatal(app.run(mux))

    

}
