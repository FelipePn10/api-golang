package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
	l.info.Println(v...)
	l.warning.Println(v...)
	l.err.Println(v...)
}

func (l *Logger) Debugf(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Infof(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warningf(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Errorf(v ...interface{}) {
	l.err.Println(v...)
}
