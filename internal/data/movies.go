package data

import (
	"database/sql"
	"time"

	"backend.delmesia/internal/validator"

	"github.com/lib/pq"
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

// A struct type which wraps a sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

// The Insert() method accepts a pointer to a movie struct, which should contain the data
// for the new record.
func (m MovieModel) Insert(movie *Movie) error {
	// The SQL query for inserting a new record in the movies table and returning
	// the system-generated data.
	query := `
		INSERT INTO movies (title, year, runtime, genres)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version`
	// args will contain the values for the placeholder parameters from the movie struct.
	// Declaring slice immediately next to the SQL query helps to make it nice and clear in the query.
	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil

}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}
