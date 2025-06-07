package get

import (
	"fmt"

	"github.com/jdetok/web/internal/env"
	"github.com/joho/godotenv"
)

type SeasonStatUrl struct {
	Root string
	Endpoint string
	Season string
	SeasonType string
	Resource string
	Team string
	Object string
}

func BuildSeasTeamUrl(szn, szn_type, team string) (string, error) {
	fmt.Println("Trying to build season team url")
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error occured getting env: %e", err)
	}
	var u SeasonStatUrl
	u.Root = env.GetString("API_ROOT")
	u.Endpoint = "seasons/"
	u.Season = szn + "/"
	u.SeasonType = szn_type + "/"
	u.Resource = "teams/"
	u.Team = team + "/"
	u.Object = "statistics.json"

	url := u.Root + u.Endpoint + u.Season + u.SeasonType + u.Resource + u.Team + u.Object
	fmt.Println(url)
	return url, nil
}