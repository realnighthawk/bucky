package logger

import (
	"os"
	"time"

	"github.com/kumarabd/gokit/errors"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	Info(description ...interface{})
	Debug(description ...interface{})
	Error(err error)
}

type Logger struct {
	handler *logrus.Entry
}

func New(appname string, opts Options) (Handler, error) {

	log := logrus.New()

	switch opts.Format {
	case JsonLogFormat:
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	case SyslogLogFormat:
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		})
	}

	// log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	entry := log.WithFields(logrus.Fields{
		"app": appname,
	})

	return &Logger{
		handler: entry,
	}, nil
}

func (l *Logger) Error(err error) {
	l.handler.WithFields(logrus.Fields{
		"code":     errors.GetCode(err),
		"severity": errors.GetSeverity(err),
	}).Log(logrus.ErrorLevel, err.Error())
}

func (l *Logger) Info(description ...interface{}) {
	l.handler.Log(logrus.InfoLevel,
		description...,
	)
}

func (l *Logger) Debug(description ...interface{}) {
	l.handler.Log(logrus.DebugLevel,
		description...,
	)
}
