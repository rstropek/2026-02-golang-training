package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"
)

// Stores configuration settings for the application
type config struct {
	port int    // The port number the server will listen on
	env  string // The environment the application is running in (e.g., "dev", "prod")
	db   struct {
		dsn string // Data Source Name for connecting to the database
	}
}

type application struct {
	config config
	logger *slog.Logger
	db     *sql.DB
}

func main() {
	var cfg config

	// Use the flag package to read command-line flags for port and environment
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|prod)")
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("POSTGRES_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("Database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
	}

	err = app.serve()
	if err != nil {
		logger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
