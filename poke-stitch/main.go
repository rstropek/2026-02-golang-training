package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	gim "github.com/ozankasikci/go-image-merge"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	//logger := slog.New(&StackTraceHandler{Handler: slog.NewTextHandler(os.Stdout, nil)})

	httpClient := &http.Client{}
	handler := newStitchHandler(httpClient)

	mux := http.NewServeMux()
	mux.Handle("GET /poke-stitch", handler)

	address := ":3000" // port, in real life you might want to make this configurable
	logger.Info("Starting server", "address", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server error", "error", err)
			stop()
		}
	}()

	<-shutdownCtx.Done() // Wait for interrupt signal
	// We now know that an interrupt signal was received, so we can start the shutdown process.
	// E.g. cleanup something
	logger.Info("Shutting down server")
	gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(gracefulShutdownCtx); err != nil {
		logger.Error("Error during server shutdown", "error", err)
		os.Exit(1)
	}
}

func newStitchHandler(client *http.Client) http.Handler {
	return &stitchHandler{client: client}
}

type stitchHandler struct {
	client *http.Client
}

// We implement the http.Handler interface for stitchHandler
func (h *stitchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pokemonName := r.URL.Query().Get("pokemon")
	if pokemonName == "" {
		http.Error(w, "Missing 'pokemon' query parameter", http.StatusBadRequest)
		return
	}

	pokemon, err := h.fetchPokemon(pokemonName)
	if err != nil {
		http.Error(w, "Failed to fetch Pokemon data", http.StatusInternalServerError)
		return
	}

	spriteURLs := pokemon.SpriteURLs()
	images, err := h.fetchPokemonSprites(spriteURLs)
	if err != nil {
		http.Error(w, "Failed to fetch Pokemon sprites", http.StatusInternalServerError)
		return
	}

	if len(images) == 0 { // Just to be sure...
		http.Error(w, "No sprites found for the specified Pokemon", http.StatusNotFound)
		return
	}

	// Preallocate slice capacity to avoid repeated allocations during append.
	grids := make([]*gim.Grid, 0, len(images))
	for _, img := range images {
		// Keep nil filtering local; merge library expects concrete images.
		if img != nil {
			grids = append(grids, &gim.Grid{Image: img})
		}
	}
	if len(grids) == 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Integer math trick: ceil(len/2) for two columns.
	rows := (len(grids) + 1) / 2
	rgba, err := gim.New(grids, 2, rows).Merge()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, rgba); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *stitchHandler) fetchPokemon(pokemonName string) (Pokemon, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName), nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var pokemon Pokemon
	if err = json.NewDecoder(res.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (h *stitchHandler) fetchPokemonSprites(urls []string) ([]image.Image, error) {
	// Error handling: In case of an error, return the first error that occurs.
	images := make([]image.Image, len(urls))
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	var firstErr error
	var firstErrMu sync.Mutex // To protect access to firstErr

	for i, url := range urls {
		go func(index int, url string) {
			defer wg.Done()
			img, err := h.getImage(url)
			if err != nil {
				firstErrMu.Lock()
				if firstErr == nil {
					firstErr = err
				}
				firstErrMu.Unlock()
				return
			}
			images[i] = img
		}(i, url)
	}

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}

	return images, nil
}

/*
func (h *stitchHandler) getImageWithChannel(url string, ch chan<- image.Image) {
	// TODO: We are currently failing silently on errors here. We could send errors
	// through a separate channel or use a struct to encapsulate both image and error.

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	res, err := h.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return
	}

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return
	}

	ch <- img
}
*/

func (h *stitchHandler) getImage(url string) (image.Image, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
