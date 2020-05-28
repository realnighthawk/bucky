package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogHandler interface {
	Err(code string, des string)
	Debug(des string)
	Info(des string)
}

type Logger struct {
	handler logrus.Logger
}

func New(appname string) (LogHandler, error) {

	log = logrus.New()
	log.Out = os.Stdout

	host, _ := os.Hostname()
	log.WithFields(logrus.Fields{
		"host":    host,
		"app":     appname,
		"ts":      time.Now().String(),
		"country": "",
	})

	return &Logger{handler: log}, nil
}

func Log(description string, level string) {

	logrus.WithFields(logrus.Fields{
		"host":        host,
		"app":         appname,
		"ts":          time.Now().String(),
		"country":     "",
		"level":       level,
		"description": description,
	})
}

func (l *Logger) Err(code string, des string) {
	l.handler.SetLevel(logrus.ErrorLevel)
	l.handler.WithFields(logrus.Fields{
		"code":        code,
		"description": des,
	})
}

func (l *Logger) Info(des string) {
	l.handler.SetLevel(logrus.InfoLevel)
	l.handler.WithFields(logrus.Fields{
		"description": des,
	})
}

func (l *Logger) Debug(des string) {
	l.handler.SetLevel(logrus.DebugLevel)
	l.handler.WithFields(logrus.Fields{
		"description": des,
	})
}
