package server

import "errors"

var (
	ErrMalformedRequest      = errors.New("malformed request")
	ErrInvalidHttpStatusCode = errors.New("invalid http status code")
)
