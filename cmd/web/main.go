package main

import (
	"context"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/shodruzhoshimzoda/snippetbox/internal/models"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// creating logger
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// reading DSN
	if err := godotenv.Load(".env"); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	DSN := os.Getenv("DB_URL")
	if DSN == "" {
		log.Error("Missing DATABASE_DSN")
		os.Exit(1)
	}

	// connection to DB
	db, err := openDB(DSN)
	if err != nil {
		log.Error(err.Error())
	}
	defer db.Close(context.Background())

	templateCache, err := newTemplateCache()

	app := &application{
		logger:        log,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	app.logger.Info("connected to database")

	log.Info("Server was runed ", "addr", *addr)

	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

}

// openDB connection to Database

func openDB(dsn string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil

}
