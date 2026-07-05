package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger 	*slog.Logger
}



func main() {
	
	addr := flag.String("addr",":4000", "HTTP network address")
	flag.Parse()
	

	// creating logger
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	
	app := &application{
		logger: log,
	}

	mux := http.NewServeMux()

	// serve static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))

	// routes registration
	mux.HandleFunc("/{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("POST /snippet/create",app. snippetCreate)

	


	log.Info("Server was runned ","addr", *addr)

	if err := http.ListenAndServe(*addr, mux);err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}


}