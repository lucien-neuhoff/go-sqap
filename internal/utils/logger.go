package utils

import (
	"fmt"
	"go-sqap/internal/config"
	"log"
	"os"
	"runtime"
)

type Logger struct {
	cfg         *config.Config
	errorLogger *log.Logger
	infoLogger  *log.Logger
	debugLogger *log.Logger
}

func NewLogger(cfg *config.Config) *Logger {
	errorLog := log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.LstdFlags)
	infoLog := log.New(os.Stdout, "\033[36m[INFO]\033[0m ", log.LstdFlags)
	debugLog := log.New(os.Stdout, "\033[33m[DEBUG]\033[0m ", log.LstdFlags)
	return &Logger{
		cfg:         cfg,
		errorLogger: errorLog,
		infoLogger:  infoLog,
		debugLogger: debugLog,
	}
}

func (l *Logger) Error(args ...any) {
	_, file, line, _ := runtime.Caller(1)
	l.errorLogger.Printf("%s:%d %s", file, line, fmt.Sprintf("%s", args...))
}

func (l *Logger) Errorf(message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	l.errorLogger.Printf("%s:%d %s", file, line, fmt.Sprintf(message, args...))
}

func (l *Logger) Info(args ...any) {
	_, file, line, _ := runtime.Caller(1)
	l.infoLogger.Printf("%s:%d %s", file, line, fmt.Sprintf("%s", args...))
}

func (l *Logger) Infof(message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	l.infoLogger.Printf("%s:%d %s", file, line, fmt.Sprintf(message, args...))
}

func (l *Logger) Debug(args ...any) {
	if !l.cfg.DEBUG {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	l.debugLogger.Printf("%s:%d %s", file, line, fmt.Sprintf("%s", args...))
}

func (l *Logger) Debugf(message string, args ...any) {
	if !l.cfg.DEBUG {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	l.debugLogger.Printf("%s:%d %s", file, line, fmt.Sprintf(message, args...))
}
