package aiot

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	ErrMissingOrInvalidCredentials = errors.New("missing or invalid credentials provided")
	ErrInvalidEmailOrPassword      = errors.New("invalid email or password")
)

type operation string
type kind uint8

type aiotError struct {
	Op   operation
	Err  error
	Kind kind
}

func (e *aiotError) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s -> ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	}
	return buf.String()
}

func makeE(args ...interface{}) error {
	e := &aiotError{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case operation:
			e.Op = arg
		case error:
			e.Err = arg
		case kind:
			e.Kind = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}
