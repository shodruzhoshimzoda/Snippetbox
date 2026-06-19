package main

import (
	"log"
	"net/http"
)





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