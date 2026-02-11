package main

import (
	"flag"
	"log/slog"
	"os"
)

// Stores configuration settings for the application
type config struct {
	port int    // The port number the server will listen on
	env  string // The environment the application is running in (e.g., "dev", "prod")
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	// Use the flag package to read command-line flags for port and environment
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|prod)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
