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

type Op string
type Kind uint8

type Error struct {
	Op   Op
	Err  error
	Kind Kind
}

func (e *Error) Error() string {
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

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}

func Is(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Kind != 0 {
		return e.Kind == kind
	}

	return Is(kind, e.Err)
}
