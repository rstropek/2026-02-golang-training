package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rstropek/hero-manager/internal/data"
)

func (app *application) listHeroesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list heroes")
}

func (app *application) showHeroHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	// Simulate fetching the hero by ID and return a response.
	hero := data.Hero{
		ID:        int64(id),
		FirstSeen: time.Now(),
		Name:      "Superman",
		CanFly:    true,
		RealName:  "Clark Kent",
		Abilities: []string{"Flight", "Super Strength", "Heat Vision"},
		Version:   1,
	}

	js, err := json.Marshal(hero)
	if err != nil {
		http.Error(w, "failed to marshal hero", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) createHeroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create hero")
}

func (app *application) deleteHeroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete hero")
}
