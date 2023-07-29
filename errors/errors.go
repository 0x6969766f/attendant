package errors

import "errors"

var (
	As = errors.As
	Is = errors.Is

	ErrNotFound = errors.New("models: resource not found")
)
