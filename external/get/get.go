package get

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdetok/web/internal/env"
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

func Get(url string, num int, save bool) ([]byte, int, error) {
// create get request with passed url	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error occured: %e\n", err)
		return nil, 0, err
	} 

// add headers to the request
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", env.GetString("API_KEY"))

// call the API
	res, err := http.DefaultClient.Do(req)

// ensure a 1 second delay occurs with every call (before error handling)
	time.Sleep(1 * time.Second)

	if err != nil {
		fmt.Printf("Error occured: %e\n", err)
		if res != nil {
			fmt.Printf("Error Status Code: %d", res.StatusCode)	
			return nil, res.StatusCode, err
		}
		return nil, 0, err
	}

	defer res.Body.Close()
	fmt.Printf("Status Code: %d\n", res.StatusCode)

// handle too many requests
	if res.StatusCode == 429 {
		var delay time.Duration = 30
		fmt.Printf("Too many requests (%d) - sleeping for %d seconds...\n", res.StatusCode, delay)
		time.Sleep(delay * time.Second)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error occured: %e\n", err)
		return nil, res.StatusCode, err
	}

// save the json file if 200 code was receieved (only if passed true for save)
// it's expected that many will return code 404
// novelty teams returned from teams endpoint (e.g. all star teams) don't have team profiles
	if save {
		root := "external/get/resp/"
		if res.StatusCode != 200 {
			fmt.Printf("Status %d - exiting without writing JSON\n", res.StatusCode)
			return body, res.StatusCode, nil
		}

		var jteam map[string]any
		if err:=  json.Unmarshal(body, &jteam); err != nil {
			fmt.Printf("Error parsing JSON: %s\n", err)
			return body, res.StatusCode, err
		}
// fall back name
		file := "response" + strconv.Itoa(num) + ".json"

// assign alias name (LAL, BOS, etc) from json as file name
		if tm, ok := jteam["alias"].(string); ok {
			file = strings.ToLower(tm) + ".json"
		}
		jsonpath := root + file

		SaveJSON(jsonpath, body)
		return body, res.StatusCode, nil
	}
	return body, res.StatusCode, nil
}

// save returned json to file
func SaveJSON(path string, body []byte) {
	err := os.WriteFile(path, body, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON response to file at %s: %s\n", path, err)
		return
	}
	fmt.Printf("JSON response saved at %s\n", path)
} 

// general requests - just pass a string for everything after the root url (after en/)
func GetRequest(ext string) ([]byte, error){

	var genurl GeneralURL
	genurl.Root = env.GetString("API_ROOT")
	genurl.Ext = ext
	
	var url string = genurl.Root + genurl.Ext

	// save = false prevents it from saving json file
	res, _, err := Get(url, 0, false)
	
	if err != nil {
		fmt.Printf("Error requesting url: %e\n", err)
		return nil, err
	}
	
// return json []byte response
	return res, nil
}