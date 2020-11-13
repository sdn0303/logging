package logging

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	logger LoggerIFace = &Logger{}
	once   sync.Once
)

type LoggerIFace interface {
	Info(v ...interface{})
	Debug(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

type Logger struct {
	CallDepth int
	logger    *log.Logger
}

func GetLogger() LoggerIFace {
	once.Do(func() {
		logger = &Logger{
			CallDepth: 3,
			logger:    log.New(os.Stderr, "", log.Lshortfile|log.Ldate|log.Ltime),
		}
	})
	return logger
}

func (l *Logger) output(p string, v ...interface{}) {
	if err := l.logger.Output(l.CallDepth, fmt.Sprintf("%s %s", p, fmt.Sprintln(v...))); err != nil {
		panic(err)
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgGreen).Sprint("💡 [INFO]     "), v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgCyan).Sprint("👀 [DEBUG]    "), v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgYellow).Sprint("⚡  [WARNING] ️"), v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.output(color.New(color.Bold, color.FgRed).Sprintf("💢 [ERROR]    "), v...)
}
