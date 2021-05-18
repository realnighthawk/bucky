package logger

const (
	// JsonLogFormat is json based logging format
	JsonLogFormat = iota
	// SyslogLogFormat is syslog based logging format
	SyslogLogFormat
)

// Format defines the logger format
type Format int

// Options supports different custom parameters for logger
type Options struct {
	Format     Format
	DebugLevel bool
}
