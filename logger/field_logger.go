package logger

import (
	"github.com/gobuffalo/logger"
	"github.com/kumarabd/gokit/errors"
)

type fieldLogger struct {
	Logger Handler
}

func NewFieldLogger(log Handler) logger.FieldLogger {
	return &fieldLogger{
		Logger: log,
	}
}

func (l *fieldLogger) Debugf(s string, des ...interface{}) {
	l.Logger.Debug(s, des)
}

func (l *fieldLogger) Infof(s string, des ...interface{}) {
	l.Logger.Info(s, des)
}

func (l *fieldLogger) Printf(s string, des ...interface{}) {
	l.Logger.Info(s, des)
}

func (l *fieldLogger) Warnf(s string, des ...interface{}) {
	l.Logger.Warn(errors.New("", errors.Warn, s, des))
}

func (l *fieldLogger) Errorf(s string, des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Alert, s, des))
}

func (l *fieldLogger) Fatalf(s string, des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Critical, s, des))
}

func (l *fieldLogger) Debug(des ...interface{}) {
	l.Logger.Debug(des...)
}

func (l *fieldLogger) Info(des ...interface{}) {
	l.Logger.Info(des...)
}

func (l *fieldLogger) Warn(des ...interface{}) {
	l.Logger.Warn(errors.New("", errors.Warn, des))
}

func (l *fieldLogger) Error(des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Alert, des))
}

func (l *fieldLogger) Fatal(des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Critical, des))
}

func (l *fieldLogger) Panic(des ...interface{}) {
	l.Logger.Error(errors.New("", errors.Fatal, des))
}

func (l *fieldLogger) WithField(s string, val interface{}) logger.FieldLogger {
	l.Logger.WithField(s, val)
	return l
}

func (l *fieldLogger) WithFields(val map[string]interface{}) logger.FieldLogger {
	l.Logger.WithFields(val)
	return l
}
