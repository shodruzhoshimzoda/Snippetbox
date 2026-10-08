package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	// "github.com/pingcap/log"

	"github.com/shodruzhoshimzoda/snippetbox/internal/models"
	"github.com/shodruzhoshimzoda/snippetbox/internal/validator"
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

type SnippetCreateForm struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = SnippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	/* if the request body size is more than 10M, we should return Bad Request	*/

	r.Body = http.MaxBytesReader(w, r.Body, 4096) // Limiting the request body size to 4096

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//title := r.PostFormValue("title")
	//content := r.PostFormValue("content")
	//
	//expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//
	//}
	//form := SnippetCreateForm{
	//	Title:   r.PostForm.Get("title"),
	//	Content: r.PostForm.Get("content"),
	//	Expires: expires,
	//}

	var form SnippetCreateForm

	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}
	form.FieldErrors = map[string]string{}

	form.CheckFieldError(validator.NotBlank(form.Title), "title", "This field can not be blank")
	form.CheckFieldError(validator.MaxChars(form.Title, 100), "title", "This field can not be more than 100 characters long")
	form.CheckFieldError(validator.NotBlank(form.Content), "content", "This field can not be blank")
	form.CheckFieldError(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field can not be less than 1 year")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}

func (app *application) snippetDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // convert id which is string to integer

	if err != nil || id < 1 {
		http.NotFound(w, r) // because we get invalid ID, the page with this id its not exist
		return
	}

	err = app.snippets.Delete(id)
	if err != nil {
		if errors.Is(err, models.ErrSnippetNotFound) {
			http.NotFound(w, r)
			return
		} else {
			app.serveError(w, r, err)
		}
		return

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
