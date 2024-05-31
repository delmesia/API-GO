package main

import "net/http"

// logError method is a generic helper for logging an error message.:
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// errorResponse method is a generic helper for sending JSON-formatted error messages to the client
// with a given status code. using the type "any" instead of string can give more flexibility to what
// values we can include in the response
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	// Write the response using the writeJSON helper(). If it retursn an error, log it,
	// and fall back to sending  the client an empty response with a
	// 500 internal server error status code.
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse method will be used when the application encounters an unexpected problem at runtime.
// It will log the detailed error message, then uses the errorResponse() helper method to send a 500 Internal Server Error
// status code and JSON response to the client.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request."
	app.errorResponse(w, r, http.StatusInternalServerError, message)

}
