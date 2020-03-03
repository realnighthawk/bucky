package logger

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

type LogHandler interface {
	Err(location string, code string, des string)
	Debug(des string)
	Info(des string)
	Warn(location string, des string)
}

type Logger struct {
	logger log.Logger
}

func New(appname string) (LogHandler, error) {
	l := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	l = log.With(l, "ts", time.Now().String())
	l = log.With(l, "appname", appname)
	return &Logger{logger: l}, nil
}

func (l *Logger) Err(location string, code string, des string) {
	logger := log.With(l.logger, "level", "ERROR")
	logger.Log("reference", location, "code", code, "description", des)
}

func (l *Logger) Info(des string) {
	logger := log.With(l.logger, "level", "INFO")
	logger.Log("description", des)
}

func (l *Logger) Warn(location string, des string) {
	logger := log.With(l.logger, "level", "WARN")
	logger.Log("reference", location, "description", des)
}

func (l *Logger) Debug(des string) {
	logger := log.With(l.logger, "level", "DEBUG")
	logger.Log("description", des)
}

func (l *Logger) With(key string, val string) {
	return log.With(l.logger.key, value)
}
