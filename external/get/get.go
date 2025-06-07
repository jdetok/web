package get

import (
	"errors"
	"fmt"
	"io"
	"net/http"

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

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error occured: %e", err)
		return nil, err
	} 

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", env.GetString("API_KEY"))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error occured: %e", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error occured: %e", err)
		return nil, err
	}
	
	return body, nil
}

// general requests - just pass a string for everything after the root url (after en/)
func GetRequest(ext string) ([]byte, error){
	var genurl GeneralURL
	genurl.Root = env.GetString("API_ROOT")
	genurl.Ext = ext
	
	var url string = genurl.Root + genurl.Ext

	res, err := Get(url)
	if err != nil {
		fmt.Printf("Error requesting url: %e", err)
		return nil, err
	}
	
// return json []byte response
	return res, nil
}