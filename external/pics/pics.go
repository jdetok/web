package pics

// https://cdn.nba.com/headshots/nba/latest/1040x760/1630173.png
// https://cdn.wnba.com/headshots/wnba/latest/1040x760/1642777.png

func MakeUrl(lg, playerId string) string {
	return ("https://cdn." + lg + ".com/headshots/" + lg + "/latest/1040x760/" + playerId + ".png")
	// )s
}

func GetHeadshots(lg string) {
	
}