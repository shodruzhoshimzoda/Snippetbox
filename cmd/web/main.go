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

	

	log.Info("Server was runned ","addr", *addr)

	if err := http.ListenAndServe(*addr, app.routes());err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}


}