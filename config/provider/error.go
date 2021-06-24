package provider

import (
	"github.com/realnighthawk/bucky/errors"
)

// ErrGetObject returns error for get config
func ErrGetObject(err error) error {
	return errors.New("", errors.NoneSeverity, "Error getting object", err.Error())
}

// ErrSetObject returns error for set config
func ErrSetObject(err error) error {
	return errors.New("", errors.NoneSeverity, "Error setting object", err.Error())
}
