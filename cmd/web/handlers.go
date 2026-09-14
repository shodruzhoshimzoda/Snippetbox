package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	// "github.com/pingcap/log"

	"github.com/shodruzhoshimzoda/snippetbox/internal/models"
)

// handler for home-page
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.html", data)

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

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.html", data)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	expires, err := strconv.Atoi(r.PostFormValue("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
