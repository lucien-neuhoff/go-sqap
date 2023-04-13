package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type Logger struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

func NewLogger() *Logger {
	errorLog := log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	infoLog := log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	return &Logger{
		errorLogger: errorLog,
		infoLogger:  infoLog,
	}
}

func (l *Logger) Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	l.errorLogger.Printf("%s:%d %s", file, line, message)
}

func (l *Logger) Info(message string) {
	_, file, line, _ := runtime.Caller(1)
	l.infoLogger.Printf("%s:%d %s", file, line, message)
}

func (l *Logger) Errorf(message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	l.errorLogger.Printf("%s:%d %s", file, line, fmt.Sprintf(message, args...))
}
