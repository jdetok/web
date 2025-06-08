package jsonops

import (
	"encoding/json"
	"fmt"
	"os"
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

// read directory with unindented json, indent all files
func IndentMany(oldPath, newPath string) {

	files, err := os.ReadDir(oldPath)
	if err != nil {
		fmt.Printf("error reading dir: %s\n", err)
	}

	for _, f := range files {
		if !f.IsDir() {
			oldf := oldPath + "/" + f.Name()
			indf := newPath + "/" + f.Name()
			Indent(oldf, indf)
		}
	} 
}

func ShrinkMany(oldPath, newPath string) {

	files, err := os.ReadDir(oldPath)
	if err != nil {
		fmt.Printf("error reading dir: %s\n", err)
	}

	for _, f := range files {
		if !f.IsDir(){
			oldf := oldPath + "/" + f.Name()
			minif := newPath + "/" + f.Name()
			SingleLine(oldf, minif)
		}
	} 
}