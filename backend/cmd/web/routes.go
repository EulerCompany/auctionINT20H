package main

import "net/http"


func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/", app.index)
    mux.HandleFunc("/auction/create", app.createAuction)
    mux.HandleFunc("/auction", app.showAuction)

    return mux
}
