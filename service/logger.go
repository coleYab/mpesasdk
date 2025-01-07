/*
Package service provides a custom logger with configurable log levels.

The Logger allows logging messages with different severity levels, which can be set and changed dynamically. The log levels supported are:
- DEBUG: Logs detailed debugging information.
- INFO: Logs general information messages.
- WARN: Logs warning messages indicating potential issues.
- ERROR: Logs error messages indicating failures or significant issues.

Logger methods include:
- Debug: Logs messages at the DEBUG level.
- Info: Logs messages at the INFO level.
- Warn: Logs messages at the WARN level.
- Error: Logs messages at the ERROR level.

Logger also provides:
- SetLevel: Allows changing the log level.
- ParseLevel: Converts a string to the corresponding LogLevel.
*/
package service

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// LogLevel represents the severity of log messages.
type LogLevel int

// Enum values for different log levels.
const (
	DEBUG LogLevel = iota // Logs detailed debugging information.
	INFO                  // Logs general information messages.
	WARN                  // Logs warning messages indicating potential issues.
	ERROR                 // Logs error messages indicating failures or significant issues.
)

// Logger is a custom logger with configurable log levels.
type Logger struct {
	level LogLevel  // The current log level of the logger.
	logger *log.Logger  // The underlying standard logger.
}

// NewLogger creates a new Logger with a specified log level.
// The default log output is os.Stdout, and the log format is the standard log format.
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SetLevel sets the logging level for the Logger.
// If set to a higher level, lower levels will be ignored.
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a message at the DEBUG level.
// This method logs if the current log level is DEBUG or lower (i.e., if level <= DEBUG).
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", format, args...)
	}
}

// Info logs a message at the INFO level.
// This method logs if the current log level is INFO or lower (i.e., if level <= INFO).
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", format, args...)
	}
}

// Warn logs a message at the WARN level.
// This method logs if the current log level is WARN or lower (i.e., if level <= WARN).
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= WARN {
		l.log("WARN", format, args...)
	}
}

// Error logs a message at the ERROR level.
// This method logs if the current log level is ERROR or lower (i.e., if level <= ERROR).
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", format, args...)
	}
}

// log is a helper function to format and print the log messages with the specified level.
func (l *Logger) log(level, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Printf("[%s] %s", level, message)
}

// ParseLevel converts a string to a LogLevel.
// It parses a string representation of a log level (e.g., "DEBUG", "INFO", "WARN", "ERROR")
// and returns the corresponding LogLevel value.
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

