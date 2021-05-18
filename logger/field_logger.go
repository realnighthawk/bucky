package logger

import (
	"github.com/gobuffalo/logger"
	"github.com/kumarabd/gokit/errors"
)

type fieldLogger struct {
	Logger Handler
}

// NewFieldLogger instantiates a new instance of field logger
func NewFieldLogger(log Handler) logger.FieldLogger {
	return &fieldLogger{
		Logger: log,
	}
}

// Debugf prints debug level log
func (l *fieldLogger) Debugf(s string, des ...interface{}) {
	l.Logger.Debug(s, des)
}

// Infof prints info level log
func (l *fieldLogger) Infof(s string, des ...interface{}) {
	l.Logger.Info(s, des)
}

// Infof prints log
func (l *fieldLogger) Printf(s string, des ...interface{}) {
	l.Logger.Info(s, des)
}

// Warnf prints warning level log
func (l *fieldLogger) Warnf(s string, des ...interface{}) {
	l.Logger.Warn(errors.New("", errors.Warn, s, des))
}

// Errorf prints error level log
func (l *fieldLogger) Errorf(s string, des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Alert, s, des))
}

// Fatalf prints fatal level log
func (l *fieldLogger) Fatalf(s string, des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Critical, s, des))
}

// Debug prints debug level log
func (l *fieldLogger) Debug(des ...interface{}) {
	l.Logger.Debug(des...)
}

// Info prints info level log
func (l *fieldLogger) Info(des ...interface{}) {
	l.Logger.Info(des...)
}

// Warn prints warning level log
func (l *fieldLogger) Warn(des ...interface{}) {
	l.Logger.Warn(errors.New("", errors.Warn, des))
}

// Error prints error level log
func (l *fieldLogger) Error(des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Alert, des))
}

// Fatal prints fatal level log
func (l *fieldLogger) Fatal(des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Critical, des))
}

// Panic prints panic level log
func (l *fieldLogger) Panic(des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Fatal, des))
}

// WithField logs with added field
func (l *fieldLogger) WithField(s string, val interface{}) logger.FieldLogger {
	l.Logger.WithField(s, val)
	return l
}

// WithFields logs with given added fields
func (l *fieldLogger) WithFields(val map[string]interface{}) logger.FieldLogger {
	l.Logger.WithFields(val)
	return l
}
