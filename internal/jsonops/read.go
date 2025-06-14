package jsonops

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJSON(path string) []byte {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	
	byteVal, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(byteVal)
	 return byteVal
}

func MapJSONFile(path string) map[string]any {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteVal, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	
	var res map[string]any 
	json.Unmarshal(byteVal, &res)

	return res
}