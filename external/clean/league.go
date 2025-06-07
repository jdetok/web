package clean

type LeagueResp struct {
	League League `json:"league"`
	Teams []Team `json:"teams"`
}

type League struct {
	ID  string   `json:"id"`  
	Name string `json:"name"`  
}

type Team struct {
	ID     string  `json:"id"`              
	Name   string  `json:"name"`            
	Alias  string  `json:"alias"`           
	Market *string `json:"market,omitempty"`
}
