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
	Info(v ...interface{})
	Debug(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

type logger struct {
	callDepth int
	logger    *log.Logger
}

func GetLogger() Logger {
	once.Do(func() {
		instance = &logger{
			callDepth: 3,
			logger:    log.New(os.Stderr, "", log.Lshortfile|log.Ldate|log.Ltime),
		}
	})
	return instance
}

func (l *logger) output(p string, v ...interface{}) {
	if err := l.logger.Output(l.callDepth, fmt.Sprintf("%s %s", p, fmt.Sprintln(v...))); err != nil {
		panic(err)
	}
}

func (l *logger) Info(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgGreen).Sprint("💡 [INFO]     "), v...)
}

func (l *logger) Debug(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgCyan).Sprint("👀 [DEBUG]    "), v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgYellow).Sprint("⚡  [WARNING] ️"), v...)
}

func (l *logger) Error(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgRed).Sprintf("💢 [ERROR]    "), v...)
}
