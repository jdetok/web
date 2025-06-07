package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error occured getting env: %e", err)
	}
}

func GetString(key string) string {
	LoadEnv()
	val, found := os.LookupEnv(key)
	if !found {
		return ""
	}
	return val
}

func GetInt(key string) int {
	LoadEnv()
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