package main

import (
	"fmt"

	"github.com/jdetok/go-api-jdeko.me/internal/mariadb"
)

func main() {
	db := mariadb.InitDB()
	resp, err := mariadb.DBJSONResposne(db, "select * from v_szn_totals")
	if err != nil {
		fmt.Printf("Error occured querying db: %v\n", err)
	}
	fmt.Println(string(resp))
}