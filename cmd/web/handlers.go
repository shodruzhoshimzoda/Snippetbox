package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	// "github.com/pingcap/log"
)

// handler for home-page
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go`")

	files := []string{
		"./ui/html/pages/home.html",
		"./ui/html/pages/partials/nav.html",
		"./ui/html/pages/base.html",
	}

	ts, err := template.ParseFiles(files...) // parse files from directory of templates
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	// w.Write([]byte("hello from Snippetbox"))
}

// handler for viewing snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) // convert id which is string to integer

	if err != nil || id < 1 {
		http.NotFound(w, r) // because we get invalid ID, the page with this id its not exist
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID: %v", id)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// chang http status code
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Display a form for creating new snippet"))
}
