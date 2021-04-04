package errors

import "strings"

func New(code string, severity Severity, description ...string) *Error {
	return &Error{
		Code:        code,
		Severity:    severity,
		Description: description,
	}
}

func (e *Error) Error() string { return strings.Join(e.Description[:], ".") }

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

func Is(err error) (*Error, bool) {
	if err != nil {
		er, ok := err.(*Error)
		return er, ok
	}
	return nil, false
}
