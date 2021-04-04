package errors

type (
	Error struct {
		Code        string
		Severity    Severity
		Description []interface{}
	}
)

type Severity string

const (
	Emergency    Severity = "emergency"
	NoneSeverity Severity = "none"
	Warn         Severity = "warn"
	Alert        Severity = "alert"
	Critical     Severity = "critical"
	Fatal        Severity = "fatal"
)
