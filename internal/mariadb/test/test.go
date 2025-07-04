package main

import (
	"encoding/json"
	"fmt"

	"github.com/jdetok/go-api-jdeko.me/internal/mariadb"
)


func main() {
	db := mariadb.InitDB()
	q := `
		select * from v_szn_avgs
		where season = ?
		and team = ?
	`
	resp, err := mariadb.DBJSONResposne(db, q, "2016-2017 Regular Season", "LAL")
	if err != nil {
		fmt.Printf("Error occured querying db: %v\n", err)
	}
	
	// fmt.Println(string(resp))

	var data any
	json.Unmarshal(resp, &data)
	fmt.Println(data)
}