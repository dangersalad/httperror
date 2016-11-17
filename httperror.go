// Package httperror contains a struct type to represent an HTTP
// status error for developing http servers
package httperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPError is a struct type representing an http error response
type HTTPError struct {
	// The HTTP status code
	Code int `json:"code"`
	// The status string for the code
	Status string `json:"status"`
	// The error message
	Message string `json:"message"`
}

// New makes a new error given the code. The Status property is set
// from `http.StatusText()` and the remaining arguments are passed to
// `fmt.Sprintf`
func New(code int, messageArgs ...interface{}) *HTTPError {
	var message string
	status := http.StatusText(code)
	if len(messageArgs) == 0 {
		message = status
	} else {
		fmtStr := fmt.Sprintf("%v", messageArgs[0])
		message = fmt.Sprintf(fmtStr, messageArgs[1:]...)
	}
	return &HTTPError{code, status, message}
}

// IsHTTPError returns true if the provided error is an HTTPError
func IsHTTPError(err error) bool {
	_, ok := err.(HTTPError)
	if !ok {
		_, ok := err.(*HTTPError)
		return ok
	}
	return ok
}

// Error returns the string representation of this error
func (err HTTPError) Error() string {
	if err.Status == err.Message {
		return fmt.Sprintf("[%d] %s", err.Code, err.Status)
	} else {
		return fmt.Sprintf("[%d] %s: %s", err.Code, err.Status, err.Message)
	}
}

// Respond uses the supplied `http.ResponseWriter` to set the
// approprate headers and a json encoded representation of the
// HTTPError for the body
func (err *HTTPError) Respond(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(err.Code)
	json.NewEncoder(res).Encode(err)
}
