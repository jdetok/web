package jsonops

import (
	"encoding/json"
	"fmt"
)

// read existing unformatted file, unmarshal to a map than re marshal as indented
func Indent(oldFile, newFile string) {
	var m map[string]any = MapJSONFile(oldFile)

	body, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	SaveJSON(newFile, body)
}

// read an indented json file & create an unindented version
func SingleLine(oldFile, newFile string) {
	var m map[string]any = MapJSONFile(oldFile)

	body, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err.Error())
	}

	SaveJSON(newFile, body)
}