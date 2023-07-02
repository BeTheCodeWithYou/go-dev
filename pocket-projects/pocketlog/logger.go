package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger for logging functionality
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a logger. Ready to log at required thereshold level
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

func (l *Logger) DebugF(format string, args ...any) {

	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)

}

func (l *Logger) InfoF(format string, args ...any) {

	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelInfo {
		return
	}
	l.logf(format, args...)
}

func (l *Logger) WarnF(format string, args ...any) {

	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelWarn {
		return
	}

	l.logf(format, args...)
}

func (l *Logger) ErrorF(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelError {
		return
	}
	l.logf(format, args...)
}

func (l *Logger) FatalF(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelFatal {
		return
	}

	l.logf(format, args...)
}

// logf prints the message to the output and does formatting of the message.
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
