package jsonops

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jdetok/web/internal/errs"
)

// save returned json to file (from HTTP response body)
func SaveJSON(path string, body []byte) error {
	e := errs.ErrInfo{Prefix: "json write to file"}
	err := os.WriteFile(path, body, 0644)
	if err != nil {
		e.Msg = "json write to file failed"
		return e.Error(err)
	}
	return nil
}

func MapToJSON(path string, m map[string]any) []byte {
// marshal the map to return []byte
// two spaces of indent
	body, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return body
}
