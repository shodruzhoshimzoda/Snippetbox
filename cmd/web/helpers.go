package main

import (
	"fmt"
	"net/http"
)

// this method will be used if we encountered with any unexpected error in server side
func (app *application) serveError(w http.ResponseWriter, r *http.Request, error error) {

	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)

	app.logger.Error(error.Error(), "method", method, "url", url)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, statusCode int) {

	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template does not exist: %v", page)
		app.serveError(w, r, err)
		return
	}

	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serveError(w, r, err)
	}

}
