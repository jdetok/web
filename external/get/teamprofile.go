package get

import (
	"fmt"

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

func TeamProfile(teamIds []string) ([]string, error) {
	root := env.GetString("API_ROOT")

	var urls []string

	for _, id := range teamIds {
		url := root + "teams/" + id + "/profile.json"
		urls = append(urls, url)
		fmt.Println(url)
	}
	return urls, nil
}