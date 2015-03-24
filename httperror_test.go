package httperror

import (
	"./"
	"fmt"
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
