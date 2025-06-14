// when an endpoint like /select is requested, it should first check if a json string
// is stored in this store package
// any time a select is run on the database, it should return the data as json to this package
// then this package returns it to the handler

package store

import (
	"fmt"
	"time"

	"github.com/jdetok/web/internal/jsonops"
)

type CacheJSON struct {
	AllPlayers []byte
	lastUpdate time.Time
}

// var c = CacheJSON{
// 	lastUpdate: time.Now(),
// }

func CheckCache(j CacheJSON) {
	age := time.Since(j.lastUpdate)
	fmt.Println((age))	
}

// func main() {


// 	js, err := c.LoadPlayers(db.CarrerStatsByLg, "NBA")
// 	if err != nil {
// 		fmt.Printf("Error getting players: %s", err)
// 	}

// 	c.AllPlayers = js
// 	fmt.Println(c.AllPlayers)
// 	// time.Sleep(5 * time.Second)

// 	age := time.Since(c.lastUpdate)

// 	if (age > (5 * time.Second)) {
// 		fmt.Printf("Longer than 5! (%v)\n", age)
// 	} else {
// 		fmt.Printf("Less than 5 - not time yet! (%v)\n", age)
// 	}

// 	// CheckCache(c)	
// }


func (j CacheJSON) LoadPlayers(query, arg string) ([]byte, error) {

	// database, err := db.Connect()
    // if err != nil {
	// 	log.Printf("Error occured connecting to database: %s", err)
	// 	return nil, err
    // }

	// js, err := db.SelectArg(database, query, false, arg)
	// if err != nil {
	// 	log.Printf("Error occured selecting to database: %s", err)
	// 	return nil, err
	// }

	// jsonops.SaveJSON("./internal/store/json/players.json", js)

	js := jsonops.ReadJSON("./internal/store/json/players.json")
	
	// save the file with this time attached, then read time back in with .parse() and detemine 
	// whether to query database or just hit json
	fileTime := time.Now().Format("010206_150405")

	fmt.Println(fileTime)
	fmt.Printf("%T\n", fileTime)

	return js, nil
}



