package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) serve() error {
	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", app.config.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
		// Add additional server settings as you like,
		// please consult the documention for details.
	}

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		app.logger.Info("Starting server", "port", app.config.port, "env", app.config.env)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			app.logger.Error("Server error", "error", err)
			stop()
		}
	}()

	<-shutdownCtx.Done() // Wait for interrupt signal
	// We now know that an interrupt signal was received, so we can start the shutdown process.
	// E.g. cleanup something
	app.logger.Info("Shutting down server")
	gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return server.Shutdown(gracefulShutdownCtx)
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/heroes", app.listHeroesHandler)
	router.HandlerFunc(http.MethodGet, "/heroes/:id", app.showHeroHandler)
	router.HandlerFunc(http.MethodPost, "/heroes", app.createHeroHandler)
	router.HandlerFunc(http.MethodDelete, "/heroes/:id", app.deleteHeroHandler)

	return router
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Here we use a map to setup the response object.
	// As an alternative, we could also create a struct
	// type and use that instead.
	data := map[string]string{
		"status": "available",
		"env":    app.config.env,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
