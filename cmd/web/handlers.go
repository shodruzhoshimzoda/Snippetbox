package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	// "github.com/pingcap/log"

	"github.com/shodruzhoshimzoda/snippetbox/internal/models"
)

// handler for home-page
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go`")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	//for _, snippet := range snippets {
	//	fmt.Fprintln(w, snippet)
	//}

	files := []string{
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
		"./ui/html/base.html",
	}

	data := templateData{
		Snippets: snippets,
	}

	ts, err := template.ParseFiles(files...) // parse files from directory of templates
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serveError(w, r, err)
		return
	}
	//
	//w.Write([]byte("hello from Snippetbox"))
}

// handler for viewing snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) // convert id which is string to integer

	if err != nil || id < 1 {
		http.NotFound(w, r) // because we get invalid ID, the page with this id its not exist
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrSnippetNotFound) {
			http.NotFound(w, r)
			return
		} else {
			app.serveError(w, r, err)
		}
		return

	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/view.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serveError(w, r, err)
		return
	}

	data := templateData{
		Snippet: snippet,
	}

	// execute template with data
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serveError(w, r, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
