package data

import (
	"database/sql"
	"errors"
)

// A custom ErrRecordNotFound error that will be used by the Get() method
// when looking up a movie that doesn't exist in the database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// This will wrap the MovieModel. This is optional, but as the build progresses,
// this can used to add models like UserModel and PermissionModel
type Models struct {
	Movies MovieModel
}

// For ease of use, NewModels() method will return a Models struct containing the
// initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
