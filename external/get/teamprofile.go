package get

import (
	"fmt"
	"time"

	"github.com/jdetok/web/external/clean"
	"github.com/jdetok/web/internal/env"
)

func TeamIds(league clean.LeagueResp) ([]string, error) {
	var teamIds []string

	for _, team := range league.Teams {
		teamIds = append(teamIds, team.ID)
	}

	return teamIds, nil
}

// takes list of team ids and generates list of urls to get their team profile
func TeamProfileUrls(teamIds []string) ([]string, error) {
	root := env.GetString("API_ROOT")

	var urls []string

	for _, id := range teamIds {
		url := root + "teams/" + id + "/profile.json"
		urls = append(urls, url)
		fmt.Println(url)
	}
	return urls, nil
}

// takes the list of urls and requests each
func TeamProfiles(urls []string) ([][]byte, []string, error) {
	var resps [][]byte
	var respsstr []string
	
	for i, url := range urls {
		res, _, err := Get(url, i, true)
		if err != nil {
			fmt.Printf("Error getting team profile at %s", url)
		}
		resps = append(resps, res)
		respsstr = append(respsstr, string(res))
		
// only one call per second
		time.Sleep(1 * time.Second)
	}
	return resps, respsstr, nil
}