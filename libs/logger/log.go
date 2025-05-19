package logger

import (
	"github.com/curtis0505/bridge/libs/logger/v2"
)

type Logger struct {
	prefix string
	logger *logger.Logger
}

func (l *Logger) Trace(v ...interface{}) {
	//Trace(l.prefix, v...)
	l.logger.Debug(l.prefix, logger.BuildLogInput().WithData(v...))
}

func (l *Logger) Debug(v ...interface{}) {
	//Debug(l.prefix, v...)
	l.logger.Debug(l.prefix, logger.BuildLogInput().WithData(v...))
}

func (l *Logger) Info(v ...interface{}) {
	//Info(l.prefix, v...)
	l.logger.Info(l.prefix, logger.BuildLogInput().WithData(v...))
}

func (l *Logger) Warn(v ...interface{}) {
	//Warn(l.prefix, v...)
	l.logger.Warn(l.prefix, logger.BuildLogInput().WithData(v...))
}

func (l *Logger) Err(err error) error {
	//Error(l.prefix, err)
	return err
}

func (l *Logger) Error(v ...interface{}) {
	//Error(l.prefix, v...)
	l.logger.Error(l.prefix, logger.BuildLogInput().WithData(v...))
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		logger: logger.NewLogger(prefix),
		prefix: prefix,
	}
}
