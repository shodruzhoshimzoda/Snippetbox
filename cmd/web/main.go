package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)





func main() {
	
	addr := flag.String("addr",":4000", "HTTP network address")
	flag.Parse()
	

	// creating logger
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	
	mux := http.NewServeMux()

	// serve static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))

	// routes registration
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)

	


	log.Info("Server was runned ","addr", *addr)

	if err := http.ListenAndServe(*addr, mux);err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}


}