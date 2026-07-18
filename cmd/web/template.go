package main

import "github.com/shodruzhoshimzoda/snippetbox/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
