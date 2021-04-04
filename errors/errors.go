package errors

import (
	"fmt"
)

func New(code string, severity Severity, description ...interface{}) *Error {
	return &Error{
		Code:        code,
		Severity:    severity,
		Description: description,
	}
}

func (e *Error) Error() string {
	return fmt.Sprint(e.Description...)
}

func GetCode(err error) string {
	if obj := err.(*Error); obj != nil && obj.Code != " " {
		return obj.Code
	}
	return ""
}

func GetSeverity(err error) Severity {
	if obj := err.(*Error); obj != nil {
		return obj.Severity
	}
	return NoneSeverity
}

func Is(err error) bool {
	if err != nil {
		_, ok := err.(*Error)
		return ok
	}
	return false
}
