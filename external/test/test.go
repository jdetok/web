package main

import (
	"encoding/json"
	"fmt"

	"github.com/jdetok/web/external/clean"
	"github.com/jdetok/web/external/get"
)

func main() {
	res, err := get.GetRequest("league/teams.json")
	if err != nil {
		fmt.Printf("Error occured: %e", err)
	}

	// uncomment to print raw json response: 
	// fmt.Println(string(res))
	var league_resp clean.LeagueResp

	errl := json.Unmarshal(res, &league_resp)
	if errl != nil {
		fmt.Printf("Error occured: %e", errl)
	}

	teamIds, err := get.TeamIds(league_resp)
	
	if err != nil {
		fmt.Printf("Error occured: %e", errl)
	}

	fmt.Println(teamIds)

	urls, err := get.TeamProfile(teamIds)
	if err != nil {
		fmt.Printf("Error occured: %e", errl)
	}
	fmt.Println(len(urls))


}