package errors

import "errors"

var (
	ErrCannotPopulateNil = errors.New("cannot populate from nil value")
	ErrNotImplemented    = errors.New("not implemented")
	ErrJobIsInvalid      = errors.New("job is invalid")
)
