package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

	fmt.Fprintf(w, "show hero %d", id)
}

func (app *application) createHeroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create hero")
}

func (app *application) deleteHeroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete hero")
}
