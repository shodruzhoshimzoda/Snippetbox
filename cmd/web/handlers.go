package main

import (
	"fmt"
	"net/http"
	"strconv"
)



// handler for home-page
func home(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Server", "Go`")


	w.Write([]byte("hello from Snippetbox"))
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
 