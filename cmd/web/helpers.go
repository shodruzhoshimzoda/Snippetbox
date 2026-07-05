package main

import "net/http"


// this method will be used if we envountered eith any unexpected error in server side
func (app *application) serveError(w http.ResponseWriter, r *http.Request, error error) {

	var (
		method = r.Method
		url	= r.URL.RequestURI()
	)

	app.logger.Error(error.Error(), "method", method, "url", url)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


func (app *application) clientError(w http.ResponseWriter,   statusCode int) {

	http.Error(w, http.StatusText(statusCode), statusCode)
}