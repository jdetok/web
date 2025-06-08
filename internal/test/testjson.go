package main

import (
	"fmt"

	"github.com/jdetok/web/internal/jsonops"
)

func main() {
	fmt.Println("TESTING INTERNAL JSON PACKAGE")
	res := jsonops.ReadJSON("json/teams.json")
	// fmt.Println(res)

	var body []byte = jsonops.MapToJSON("", res)

	jsonops.SaveJSON("json/test.json", body)
}
