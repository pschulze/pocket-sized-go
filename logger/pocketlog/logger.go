package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
	maxLen    int
}

// New returns you a logger, ready to log at the required threshold.
// If a log line's length exceeds maxLen, it will be truncated.
// The default output is os.Stdout.
// The default maximum log line length is 1000 runes.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout, maxLen: 1000}

	for _, opt := range opts {
		opt(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if LevelDebug < l.threshold {
		return
	}

	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if LevelInfo < l.threshold {
		return
	}

	l.logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if LevelError < l.threshold {
		return
	}

	l.logf(LevelError, format, args...)
}

// logf prints the message to the output.
func (l *Logger) logf(level Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	message = fmt.Sprintf("%s - %s", level, message)

	message = l.truncate(message)
	message = fmt.Sprintf("%s\n", message)

	_, _ = fmt.Fprintf(l.output, "%s", message)
}

func (l *Logger) truncate(message string) string {
	runes := []rune(message)
	if len(runes) <= l.maxLen-3 {
		return message
	}

	return string(runes[:l.maxLen-3]) + "..."
}
