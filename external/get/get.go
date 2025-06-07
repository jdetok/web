package get

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jdetok/web/internal/env"
	"github.com/joho/godotenv"
)

type GeneralURL struct {
	Root string
	Ext string
}

func (comps GeneralURL) ReqGenURL() (string, error) {
	if comps.Root == "" || comps.Ext == "" {
		return "", errors.New("no part of url can be blank")
	}
	
	return comps.Root + comps.Ext, nil
}

func Get(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error occured: %e", err)
		return "", err
	} 

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", env.GetString("API_KEY"))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error occured: %e", err)
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error occured: %e", err)
		return "", err
	}
	// fmt.Println(string(body))
	return string(body), nil
}


func GetRequest(){
	// load from .env
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error occured getting env: %e", err)
		return
	}

	var genurl GeneralURL
	genurl.Root = env.GetString("API_ROOT")
	genurl.Ext = "seasons/2024/REG/teams/583eca2f-fb46-11e1-82cb-f4ce4684ea4c/statistics.json"

	url, err := genurl.ReqGenURL()
	if err != nil {
		fmt.Printf("Error occured building url: %e", err)
	}

	res, err := Get(url)
	if err != nil {
		fmt.Printf("Error requesting url: %e", err)
	}

	fmt.Println(res)
	
}