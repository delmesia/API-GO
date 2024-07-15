package data

import (
	"time"

	"backend.delmesia/internal/validator"
)

// Movie represents a movie with its attributes.
//
// To control how keys appear in the JSON-encoded output,
// we can use struct tags to annotate the Movie struct like this:
// - "-" directive is used in struct tags to hide information that users don't need to see.
// - "omitempty" directive can hide fields if and only if they are empty.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Use the Check() method to execute the validation checks. This will add
	// the provided key and error message to the errors map
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	// Note that it's using the unique helper to check that all values in the input.Genres
	// slices are unique.
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
