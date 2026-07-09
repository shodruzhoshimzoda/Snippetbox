package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// creating logger
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app := &application{
		logger: log,
	}

	// reading DSN
	DSN := app.readDSN()

	// connection to DB
	db, err := openDB(DSN)
	if err != nil {
		app.logger.Error(err.Error())

		os.Exit(1)
	}

	defer db.Close()

	app.logger.Info("connected to database")

	log.Info("Server was runed ", "addr", *addr)

	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

}

func (app *application) readDSN() string {

	if err := godotenv.Load(".env"); err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}

	DSN := os.Getenv("DATABASE_DSN")
	if DSN == "" {
		app.logger.Error("Missing DATABASE_DSN")
		os.Exit(1)
	}
	return DSN

}

// openDB connection to Database

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
