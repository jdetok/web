package logs

import (
	"fmt"
	"net/http"
)

func LogHTTP(r *http.Request) {
	fmt.Printf("Received request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Printf("Referer: %s\n", r.Referer())
}