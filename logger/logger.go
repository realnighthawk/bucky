package logger

import (
	"os"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

type LogHandler interface {
	Err(code string, des string)
	Debug(des string)
	Info(des string)
	EnableDebug(b bool)
}

type Logger struct {
	handler *logrus.Logger
	debug   bool
}

func New(appname string) (LogHandler, error) {

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
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

func Log(appname string, description string, level string) {
	host, _ := os.Hostname()
	logrus.WithFields(logrus.Fields{
		"host":    host,
		"app":     appname,
		"ts":      time.Now().String(),
		"country": "",
		"level":   level,
	}).Info(description)
}

func (l *Logger) EnableDebug(b bool) {
	l.debug = b
}

func (l *Logger) Err(code string, description string) {
	l.handler.SetLevel(logrus.ErrorLevel)
	l.handler.WithFields(logrus.Fields{
		"code": code,
	}).Error(description)
}

func (l *Logger) Info(description string) {
	l.handler.SetLevel(logrus.InfoLevel)
	l.Info(description)
}

func (l *Logger) Debug(description string) {
	if l.debug {
		l.handler.SetLevel(logrus.DebugLevel)
		l.Debug(description)
	}
}
