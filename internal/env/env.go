package env

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return ""
	}
	return val
}

func GetInt(key string) int {
	val, found := os.LookupEnv(key)
	if !found {
		return 0
	}

// convert key from string to int
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}

	return valAsInt
}