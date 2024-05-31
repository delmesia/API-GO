package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// MarshalJSON method that satisfies the json.Marshaler interface.
// This should return a JSON-encoded value for the movie runtime.
func (r Runtime) MarshalJSON() ([]byte, error) {

	// Generate a string containing the movie runtime in the required format
	jsonValue := fmt.Sprintf("%d mins", r)
	// strconv.Quote() function is used to wrap the string in double quotes.
	// It needs to be surrounded by double quotes in order to be a valid *JSON string*
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert a value string to []byte slice and return
	return []byte(quotedJSONValue), nil

}
