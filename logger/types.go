package logger

const (
	JsonLogFormat = iota
	SyslogLogFormat
)

// Format defines the logger format
type Format int

// Options supports different custom parameters for logger
type Options struct {
	Format     Format
	DebugLevel bool
}
