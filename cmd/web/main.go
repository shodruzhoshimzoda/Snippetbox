package main

import (
	"flag"
	"log"
	"net/http"
)





func main() {
	
	addr := flag.String("addr",":4000", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()

	

	// serve static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))

	// routes registration
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)

	


	log.Println("Server was runned on port: ", *addr)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal("Error to run server: ", err)


}