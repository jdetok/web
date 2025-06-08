package main

import (
	"fmt"

	"github.com/jdetok/web/internal/jsonops"
)

func main() {
	fmt.Println("TESTING INTERNAL JSON PACKAGE")
	
// reads indented file and returns single line json file
	// jsonops.SingleLine("json/teams.json", "json/teamsx.json")

	jsonops.IndentMany("json/teamprofiles", "json/teamprofiles/ind")
	jsonops.ShrinkMany("json/teamprofiles", "json/teamprofiles/mini")
	

// old testing: 
	// res := jsonops.MapJSONFile("json/teams.json")
	// // fmt.Println(res)

	// var body []byte = jsonops.MapToJSON("", res)

	// jsonops.SaveJSON("json/test.json", body)
}
