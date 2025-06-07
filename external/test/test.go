package main

import (
	"encoding/json"
	"fmt"

	"github.com/jdetok/web/external/clean"
	"github.com/jdetok/web/external/get"
)

func main() {
	fmt.Println("TESTING EXTERNAL GET REQUEST PACKAGE")

	res, err := get.GetRequest("league/teams.json")
	if err != nil {
		fmt.Printf("Error occured: %s", err)
	}

// uncomment to print raw json response as string: 
	// fmt.Println(string(res))
	var league_resp clean.LeagueResp

	errl := json.Unmarshal(res, &league_resp)
	if errl != nil {
		fmt.Printf("Error occured: %s", errl)
	}

// pass LeagueResp struct to get.TeamIds to return a []string with each team's ID
	teamIds, err := get.TeamIds(league_resp)
	if err != nil {
		fmt.Printf("Error occured: %s", errl)
	}
// uncomment to print the team ids
	//fmt.Println(teamIds)

// pass the ids to get.TeamProfile to generate list of urls to get each team's profile
	urls, err := get.TeamProfileUrls(teamIds)
	if err != nil {
		fmt.Printf("Error occured: %s", errl)
	}
// uncomment to print the urls
	fmt.Println(urls)

	resps, respsstr, err := get.TeamProfiles(urls)
	if err != nil {
		fmt.Printf("Error occured: %s", errl)
	}

	fmt.Println(resps)
	fmt.Println(respsstr)
}