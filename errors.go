package aiot

import "errors"

var (
	ErrMissingOrInvalidCredentials = errors.New("missing or invalid credentials provided")
	ErrInvalidEmailOrPassword      = errors.New("invalid email or password")
)
