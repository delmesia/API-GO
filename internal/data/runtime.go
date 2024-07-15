package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("Invalid runtime format")

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

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// The incoming JSON value will be in the string format "<runtime> mins",
	// and we need to remove the surrounding double-quotes from string. If
	// unable to unquote, return ErrInvalidRuntimeFormat
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// Split the string to isolate the part containing the number
	parts := strings.Split(unquotedJSONValue, " ")

	// Sanity check the parts to make sure it's in the expected format.
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// Otherwise, parse the string containing the number into an int32
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}
