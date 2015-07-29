package httperror_test

import (
	"fmt"
	"github.com/dangersalad/httperror"
	"net/http"
	"testing"
)

// test creation of a simple http error
func TestNew(t *testing.T) {
	code := 404
	message := "Thing not found"
	err := httperror.New(code, message)
	if err.Code != code {
		t.Error("Code on error does not match")
	}
	if err.Message != message {
		t.Error("Message on error does not match")
	}
}

func ExampleNew_code() {
	err := httperror.New(418)
	fmt.Println(err)
	// Output: [418] I'm a teapot
}

func ExampleNew_message() {
	err := httperror.New(418, "Ah ah ah")
	fmt.Println(err)
	// Output: [418] I'm a teapot: Ah ah ah
}

func ExampleNew_formatted() {
	err := httperror.New(418, "Ah ah ah %s", "foo")
	fmt.Println(err)
	// Output: [418] I'm a teapot: Ah ah ah foo
}

// test allowing the error to fill in a message if one is not provided
func TestNewWithAutoMessage(t *testing.T) {
	code := 418
	message := "I'm a teapot"
	// leave out the message, let it use a standard one
	err := httperror.New(code)
	if err.Code != code {
		t.Error("Code on error does not match")
	}
	if err.Message != message {
		t.Error("Message on error does not match")
	}
}

// test using fmt.Sprintf style arguments to format an error message
func TestNewWithFormattedMessage(t *testing.T) {
	code := 418
	message := "Stop trying to %s"
	arg := "lick your own ass"
	err := httperror.New(code, message, arg)
	if err.Code != code {
		t.Error("Code on error does not match")
	}
	if err.Message != fmt.Sprintf(message, arg) {
		t.Error("Message on error does not match")
	}
}

// test the IsHTTPError function
func TestIsHTTPError(t *testing.T) {
	code := 418
	err := httperror.New(code)
	if !httperror.IsHTTPError(err) {
		t.Error("Not an HTTPError")
	}
	if !httperror.IsHTTPError(*err) {
		t.Error("Not an HTTPError")
	}
}

type MockResponseWriter struct {
	ResBody    []byte
	StatusCode int
}

func (res *MockResponseWriter) Header() http.Header {
	return http.Header(make(map[string][]string, 0))
}

func (res *MockResponseWriter) WriteHeader(val int) {
	res.StatusCode = val
}

func (res *MockResponseWriter) Write(val []byte) (int, error) {
	res.ResBody = val
	return len(val), nil
}

// test Respond
func TestRespond(t *testing.T) {
	code := 418
	err := httperror.New(code)
	res := new(MockResponseWriter)
	err.Respond(res)
	if res.StatusCode != code {
		t.Error("Response code not set")
	}
	if res.ResBody == nil {
		t.Error("Response body not set")
	}
}

func ExampleHTTPError_Respond() {
	var w http.ResponseWriter
	err := httperror.New(418)
	err.Respond(w)
}

// test Error
func TestErrorNoMessage(t *testing.T) {
	code := 418
	err := httperror.New(code)
	errStr := err.Error()
	expected := fmt.Sprintf("[%d] %s", err.Code, err.Status)
	if errStr != expected {
		t.Error("Error() not formatting strting output properly")
	}
}

func TestErrorMessage(t *testing.T) {
	code := 418
	err := httperror.New(code, "Ah ah ah")
	errStr := err.Error()
	expected := fmt.Sprintf("[%d] %s: %s", err.Code, err.Status, err.Message)
	if errStr != expected {
		t.Error("Error() not formatting strting output properly")
	}
}
