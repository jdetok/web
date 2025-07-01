package pics

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jdetok/web/internal/env"
	"github.com/jdetok/web/internal/mariadb"
)

// https://cdn.nba.com/headshots/nba/latest/1040x760/1630173.png
// https://cdn.wnba.com/headshots/wnba/latest/1040x760/1642777.png

func MakeUrl(lg, playerId string) string {
	return ("https://cdn." + lg + ".com/headshots/" + lg + "/latest/1040x760/" + playerId + ".png")
	// )s
}

func GetHeadshots(lg string) {
	database := mariadb.InitDB()
	hShotPath := env.GetString("HS_PATH")
	q := `select player_id from player where active = 1 and lg = "` + strings.ToUpper(lg) + `"`
	playerIds, err := mariadb.SelectList(database, q)
	if err != nil {
		fmt.Println(err)
	}

	for _, pId := range playerIds {
		url := MakeUrl(lg, pId)

		resp, err := http.Get(url)
		if err != nil {
			if resp.StatusCode == 429 {
				fmt.Printf("too many requests error: %s\nsleeping for 30 seconds\n", err)
				time.Sleep(30 * time.Second)
			}
			fmt.Printf("error occured: %s\nsleeping for 10 seconds\n", err)
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		path := hShotPath + pId + ".png"
		os.WriteFile(path, data, 0666)
	}
}