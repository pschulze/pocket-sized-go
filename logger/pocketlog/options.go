package pocketlog

import "io"

// Option defines a functional option to our Logger.
// See https://golang.cafe/blog/golang-functional-options-pattern
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output writer of the logger.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

// WithMaxLen returns a configuration function that sets the maximum message length of the logger.
func WithMaxLen(length int) Option {
	return func(l *Logger) {
		l.maxLen = length
	}
}
