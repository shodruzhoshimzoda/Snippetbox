package main

import (
	"html/template"
	"path/filepath"

	"github.com/shodruzhoshimzoda/snippetbox/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {

	// initialize new map to act cache
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-on
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err  := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}	

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}	

		cache[name] = ts
	}
	return cache, nil
}
