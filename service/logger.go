package service;

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Logger is a custom logger with configurable log levels
type Logger struct {
	level LogLevel
	logger *log.Logger
}

// NewLogger creates a new Logger with a default level
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SetLevel sets the logging level for the Logger
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a message at the DEBUG level
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", format, args...)
	}
}

// Info logs a message at the INFO level
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", format, args...)
	}
}

// Warn logs a message at the WARN level
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= WARN {
		l.log("WARN", format, args...)
	}
}

// Error logs a message at the ERROR level
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", format, args...)
	}
}

// log is a helper function to format and print the log messages
func (l *Logger) log(level, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Printf("[%s] %s", level, message)
}

// ParseLevel converts a string to a LogLevel
func ParseLevel(level string) (LogLevel, error) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	default:
		return INFO, fmt.Errorf("unknown log level: %s", level)
	}
}

