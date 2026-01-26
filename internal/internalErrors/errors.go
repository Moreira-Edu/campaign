package internalerrors

import "errors"

var (
	ErrInernal error = errors.New("Internal server error")
)
