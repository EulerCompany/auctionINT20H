package main

import "net/http"

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
	    app.notFound(w)
        return
    }   
	w.Write([]byte("index"))
}

func (app *application) createAuction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new auction..."))
}

func (app *application) showAuction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing particular auction"))
}
