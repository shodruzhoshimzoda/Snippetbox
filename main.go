package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)



// handler for home-page
func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello from Snippetbox"))
}

// handler for viewing snippet
func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) // convert id which is string to integer

	if err != nil || id < 1 {
		http.NotFound(w, r)			// because of we get error the page with this id its not exist
	}


	msg := fmt.Sprintf("Display a specific movie with id: %v", id)

	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating new snippet"))
}
 


func main() {
	
	mux := http.NewServeMux()


	// routes registration
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)

	


	log.Println("Server was runned on port: localhost:4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal("Error to run server: ", err)


}