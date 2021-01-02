package movie

import "errors"

var (
	// ErrNotFound .
	ErrNotFound = errors.New("Scene not found")
	// ErrQueryDoesntMatch .
	ErrQueryDoesntMatch = errors.New("Query is invalid")
)
