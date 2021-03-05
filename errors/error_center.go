package errors

import (
	"fmt"
	"io"
	"strings"
)

// Center collects errors.
// This type can treat as error.
type Center struct {
	errors []error
}

// New creates a new instance of ErrorCenter
func New() *Center {
	return &Center{errors: []error{}}
}

// Push puts the given error into the receiver error center instance.
func (ec *Center) Push(err error) bool {
	if err != nil {
		switch err.(type) {
		case *Center:
			ec.errors = append(ec.errors, err.(*Center).errors...)
		default:
			ec.errors = append(ec.errors, err)
		}
	}
	return err != nil
}

// IsEmpty confirms the errors in the receiver error center instance is zero.
func (ec *Center) IsEmpty() bool {
	return len(ec.errors) == 0
}

// Error returns the error messages in the receiver error center instance.
func (ec *Center) Error() string {
	dest := new(strings.Builder)
	ec.Writeln(dest)
	return strings.TrimSpace(dest.String())
}

// Writeln prints the error messages in the receiver error center instance to the given destination.
func (ec *Center) Writeln(dest io.Writer) {
	for _, err := range ec.errors {
		fmt.Fprintln(dest, err.Error())
	}
}
