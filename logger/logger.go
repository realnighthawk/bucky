package logger

import (
	"os"
	"time"

	"github.com/go-kit/log"
)

type LogHandler interface {
	Err(code string, des string)
	Debug(code string, des string)
	Info(code string, des string)
	Warn(code string, des string)
}

type Logger struct {
	logger *log.logger
}

func New(appname string) (*LogHandler, error) {
	l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	l = log.With(l, "ts", time.Now().String())
	l = log.With(l, "appname", appname)
	return &Logger{logger: l}
}

func (l *Logger) Err(location string, code string, des string) {
	l = log.With(l, "level", "ERROR")
	l.Log("reference", location, "code", code, "description", des)
}

func (l *Logger) Info(des string) {
	l = log.With(l, "level", "INFO")
	l.Log("description", des)
}

func (l *Logger) Warn(location string, des string) {
	l = log.With(l, "level", "WARN")
	l.Log("reference", location, "description", des)
}

func (l *Logger) Debug(des string) {
	l = log.With(l, "level", "DEBUG")
	l.Log("description", des)
}
