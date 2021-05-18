package errors

type (
	// Error defines the attributes of an error
	Error struct {
		Code        string
		Severity    Severity
		Description []interface{}
	}
)

// Severity describes the severity level of the error
type Severity string

const (
	// Emergency level severity
	Emergency Severity = "emergency"
	// NoneSeverity level severity
	NoneSeverity Severity = "none"
	// Warn level severity
	Warn Severity = "warn"
	// Alert level severity
	Alert Severity = "alert"
	// Critical level severity
	Critical Severity = "critical"
	// Fatal level severity
	Fatal Severity = "fatal"
)
