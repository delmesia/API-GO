package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// ByName will get the value of ID paramater from the slice.
	// Since ByName's return value is a string, convert it base 10 integer
	// (with a bit of size 64). If the paramater couldn't be converted, or is less than 1,
	// the ID is invalid so return by using http.NotFound() function to return a 404 Not Found response.

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
