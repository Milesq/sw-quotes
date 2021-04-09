package getscene

import "errors"

var (
	// ErrNotFound .
	ErrNotFound = errors.New("scene not found")
	// ErrQueryDoesntMatch .
	ErrQueryDoesntMatch = errors.New("query is invalid")
)
