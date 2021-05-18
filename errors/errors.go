package errors

import (
	"fmt"
)

// New instantiates a new instance of error object
func New(code string, severity Severity, description ...interface{}) *Error {
	return &Error{
		Code:        code,
		Severity:    severity,
		Description: description,
	}
}

// Error returns the error description
func (e *Error) Error() string {
	return fmt.Sprint(e.Description...)
}

// GetCode returns the error code
func GetCode(err error) string {
	if obj := err.(*Error); obj != nil && obj.Code != " " {
		return obj.Code
	}
	return ""
}

// GetSeverity returns the severity level of the error
func GetSeverity(err error) Severity {
	if obj := err.(*Error); obj != nil {
		return obj.Severity
	}
	return NoneSeverity
}

// Is returns is the error object is valid
func Is(err error) bool {
	if err != nil {
		_, ok := err.(*Error)
		return ok
	}
	return false
}
