package errors

import "errors"

var (
	ErrNotFound    = errors.New("resource not found")
	ErrMissingData = errors.New("insufficient data")
	ErrParse       = errors.New("parse error")
)
