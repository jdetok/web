package jsonops

import (
	"fmt"
	"os"
)

// save returned json to file (from HTTP response body)
func SaveJSON(path string, body []byte) {
	err := os.WriteFile(path, body, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON response to file at %s: %s\n", path, err)
		return
	}
	fmt.Printf("JSON response saved at %s\n", path)
} 
