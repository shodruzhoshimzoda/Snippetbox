package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/form/v4"
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

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) templateData {

	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) decodePostForm(r *http.Request, dst any) error {

	if err := r.ParseForm(); err != nil {
		return err
	}

	err := app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {

		var InvalidDecodeError *form.InvalidDecoderError
		if errors.As(err, &InvalidDecodeError) {
			panic(err)
		}
		return err

	}

	return nil
}
