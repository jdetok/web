package store

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jdetok/web/internal/errs"
	"github.com/jdetok/web/internal/mariadb"
)

type Player struct {
	PlayerId uint64
	Name string
	League string
}

type Season struct {
	SeasonId string
	Season string
	WSeason string
}

type Team struct {
	League string
	TeamId string
	TeamAbbr string
	CityTeam string 
}


// QUERY FOR PLAYER ID, PLAYER AND SAVE TO A LIST OF PLAYER STRUCTS
func GetPlayers(db *sql.DB) ([]Player, error) {
	fmt.Println("querying players & saving to struct")
	e := errs.ErrInfo{Prefix: "saving players to structs"}
	rows, err := db.Query(mariadb.Players.Q)
	if err != nil {
		e.Msg = "query failed"
		return nil, e.Error(err)
	}
	var players []Player
	for rows.Next() {
		var p Player
		rows.Scan(&p.PlayerId, &p.Name, &p.League)
		// convert to lowercase to match requests
		p.Name = strings.ToLower(p.Name) 
		p.League = strings.ToLower(p.League) 
		players = append(players, p)
	}
	return players, nil
} 

func SearchPlayers(players []Player, pSearch string) string {
	for _, p := range players {
		if p.Name == pSearch { // return match playerid (uint32) as string
			return strconv.FormatUint(p.PlayerId, 10)
		}
	}
	return ""
}
 // seasons
func GetSeasons(db *sql.DB) ([]Season, error) {
	fmt.Println("querying seasons & saving to struct")
	e := errs.ErrInfo{Prefix: "saving seasons to struct"}
	rows, err := db.Query(mariadb.Seasons.Q)
	if err != nil {
		e.Msg = "error querying db"
		e.Error(err)
	}

	var seasons []Season
	for rows.Next() {
		var szn Season
		rows.Scan(&szn.SeasonId, &szn.Season, &szn.WSeason)
		seasons = append(seasons, szn)
	}
	
	return seasons, nil
}

// teams
func GetTeams(db *sql.DB) ([]Team, error) {
	fmt.Println("querying teams & saving to struct")
	e := errs.ErrInfo{Prefix: "saving teams to struct"}
	rows, err := db.Query(mariadb.Teams.Q)
	if err != nil {
		e.Msg = "error querying db"
		e.Error(err)
	}

	var teams []Team
	for rows.Next() {
		var tm Team
		rows.Scan(&tm.League, &tm.TeamId, &tm.TeamAbbr, &tm.CityTeam)
		teams = append(teams, tm)
	}
	return teams, nil
}