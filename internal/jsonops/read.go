package jsonops

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jdetok/web/internal/errs"
)

func ReadJSON(path string) ([]byte, error) {
	e := errs.ErrInfo{Prefix: "json read"}
	
	jsonFile, err := os.Open(path)
	if err != nil {
		e.Msg = "failed to open file"
		return nil, e.Error(err)
	}
	defer jsonFile.Close()
	
	js, err := io.ReadAll(jsonFile)
	if err != nil {
		e.Msg = "failed to read json file contents"
		return nil, e.Error(err)
	}
	return js, nil
}

func MapJSONFile(path string) (map[string]any, error) {
	e := errs.ErrInfo{Prefix: "json to Go map",}	
	js, err := ReadJSON(path)
	if err != nil {
		e.Msg = "failed to read json filed"
		return nil, e.Error(err)
	}
	
	var res map[string]any 
	if err = json.Unmarshal(js, &res); err != nil {
		e.Msg = "failed to unmarshal json file"
		return nil, e.Error(err)
	}

	return res, nil
}

// jsonFile, err := os.Open(path)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer jsonFile.Close()

	// byteVal, err := io.ReadAll(jsonFile)