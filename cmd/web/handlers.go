package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	// "github.com/pingcap/log"
)

// handler for home-page
func home(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Server", "Go`")

	files := []string{		
		"./ui/html/pages/home.html",
		"./ui/html/pages/partials/nav.html",
		"./ui/html/pages/base.html",
	}

	ts, err := template.ParseFiles(files...)   // parse files from directory of templates
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interranl Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil) 
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
 

	// w.Write([]byte("hello from Snippetbox"))
}

// handler for viewing snippet
func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) // convert id which is string to integer

	if err != nil || id < 1 {
		http.NotFound(w, r)			// because of we get error the page with this id its not exist
	}


	fmt.Fprintf(w, "Display a specific snippet with ID: %v", id)

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// chang http status code 
	w.WriteHeader(http.StatusCreated)


	w.Write([]byte("Display a form for creating new snippet"))
}
 