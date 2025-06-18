package main

import (
	"fmt"
	"net/http"

	"github.com/jdetok/web/internal/env"
)

var fsPath string = env.GetString("STATIC_PATH")

func (app *application) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Printf("Referer: %s\n", r.Referer())
	http.FileServer(http.Dir(fsPath)).ServeHTTP(w ,r)
}

func (app *application) cssNoCache(w http.ResponseWriter, r *http.Request,) {
		w.Header().Set("Cache-Control", "no-store")
		http.StripPrefix("/css/", http.FileServer(http.Dir(fsPath + "/css"))).ServeHTTP(w, r)
	}

func (app *application) jsNoCache(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.StripPrefix("/js/", http.FileServer(http.Dir(fsPath + "/js"))).ServeHTTP(w, r)
	}
