package main

import (
	"fmt"

	"github.com/jdetok/web/external/get"
)

func main() {
	// get.GetRequest()
	url, err := get.BuildSeasTeamUrl(
		"2024",
		"REG", 
		"583ecae2-fb46-11e1-82cb-f4ce4684ea4c",
	)
	
	if err != nil {
		fmt.Printf("Error occured: %e", err)
	}

	res, err := get.Get(url)
	if err != nil {
		fmt.Printf("Error occured: %e", err)
	}
	fmt.Println(res)
	
}