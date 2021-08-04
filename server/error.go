package server

import "github.com/realnighthawk/bucky/errors"

var (
	ErrInvalidKind    = errors.New("", errors.Alert, "Unknown server kind")
	ErrInvalidName    = errors.New("", errors.Alert, "Unknown server name")
	ErrInvalidVersion = errors.New("", errors.Alert, "Unknown server version")
)
