package types

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidEncoding = errors.New("invalid data encoding")
)

const (
	EventHupReceived    = "Event hup received from poller"
	InvalidEncoding     = "Invalid Data Encoding Found"
	InternalServerError = "Something went wrong"
)
