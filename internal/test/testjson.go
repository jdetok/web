package main

import (
	"fmt"

	"github.com/jdetok/web/internal/jsonops"
)

func main() {
	fmt.Println("TESTING INTERNAL JSON PACKAGE")
	res := jsonops.ReadJSON("external/get/resp/teams.json")
	fmt.Println(res)
}	
