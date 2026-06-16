package main

import (
	"log"
	"net/http"
)

// handler for home-page
func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello from Snippetbox"))
}





func main() {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/home", home)



	log.Println("Server was runned on port: localhost:4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal("Error to run server: ", err)




}