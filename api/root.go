package main

import "net/http"

func (app *application) rootHandler(w http.ResponseWriter, r *http.Request) {
// THIS DOESN'T WORK
	http.FileServer(http.Dir("/home/jdeto/go/src/go-api/web/src"))
}