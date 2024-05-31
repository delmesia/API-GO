package data

import "time"

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
	Genre     []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
