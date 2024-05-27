package main

import (
	"fmt"
	"net/http"
	"time"

	"backend.delmesia/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Time,
		Title:     "Casablanca",
		Runtime:   102,
		Genre:     []string{"drama", "romance", "war"},
		Version:   1,
	}

}
