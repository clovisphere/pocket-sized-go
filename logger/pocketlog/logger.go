package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// LogEntry is the JSON structure for each log message.
type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	//Time    string `json:"time"`
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it to your will.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf("debug", format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf("info", format, args...)
}

// Errorf formats and prints a message if the log message is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf("error", format, args...)
}

// logf prints the message to the output in JSON format.
func (l *Logger) logf(level, format string, args ...any) {
	entry := LogEntry{
		Level:   strings.ToLower(level),
		Message: fmt.Sprintf(format, args...),
		//Time:    time.Now().Format(time.RFC3339),
	}

	b, err := json.Marshal(entry)
	if err != nil {
		// fallback if JSON fails
		fmt.Fprintf(l.output, "[%-6s] %s\n", level, entry.Message)
		return
	}

	fmt.Fprintln(l.output, string(b))
}
