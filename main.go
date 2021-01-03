package logging

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	instance *logger
	once     sync.Once
)

type Logger interface {
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type logger struct {
	callDepth int
	logger    *log.Logger
}

func GetLogger(prefix string) Logger {
	once.Do(func() {
		instance = &logger{
			callDepth: 3,
			logger:    log.New(os.Stderr, prefix, log.Lshortfile|log.Ldate|log.Ltime),
		}
	})
	return instance
}

func (l *logger) output(p, format string, v ...interface{}) {
	if err := l.logger.Output(l.callDepth, fmt.Sprintf(p+format, v...)); err != nil {
		panic(err)
	}
}

func (l *logger) Info(format string, v ...interface{}) {
	l.output(color.New(color.Bold, color.FgGreen).Sprint("[INFO] "), format, v...)
}

func (l *logger) Debug(format string, v ...interface{}) {
	l.output(color.New(color.Bold, color.FgCyan).Sprint("[DEBUG] "), format, v...)
}

func (l *logger) Warn(format string, v ...interface{}) {
	l.output(color.New(color.Bold, color.FgYellow).Sprint("[WARNING] "), format, v...)
}

func (l *logger) Error(format string, v ...interface{}) {
	l.output(color.New(color.Bold, color.FgRed).Sprintf("[ERROR] "), format, v...)
}
