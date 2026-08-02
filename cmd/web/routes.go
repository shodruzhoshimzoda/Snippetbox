package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	// serve static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// routes registration
	mux.HandleFunc("/{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)

	return app.recoverPanic(app.logRequestMiddleware(commonHeaders(mux)))

}
