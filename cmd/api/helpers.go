package main

import (
	"encoding/json"
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

// A writeJSON() helper for sending response.
// Destination: w http.ResponseWriter
// HTTP status code to send: status int
// Data to encode to JSON: data any
// Header map containing additional HTTP headers to be included in the response: headers http.Headerj
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, return an error if there was one.
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in the terminal
	js = append(js, '\n')

	// At this point, we're sure that we wont encounter any more errors before writing
	// the response, it's now safe to add any headers we want to include.
	// Loop through the header map and add each header to http.ResponseWriter header map.
	// Note: Go doesnt throw an error if you try to range over or read over a nil map
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code and
	// json response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
