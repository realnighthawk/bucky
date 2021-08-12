package inmem

import "github.com/realnighthawk/bucky/errors"

var (
	ErrKeyNotExist = errors.New("", errors.Alert, "Key does not exist")
)
