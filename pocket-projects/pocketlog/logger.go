package pocketlog

import "fmt"

// Logger for logging functionality
type Logger struct {
	threshold Level
}

func (l *Logger) DebugF(format string, args ...any) {

	if l.threshold > LevelDebug {
		return
	}

	fmt.Printf(format+"\n", args...)
}

func (l *Logger) InfoF(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	fmt.Printf(format+"\n ", args...)

}

func (l *Logger) WarnF(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	fmt.Printf(format+"\n", args...)
}

func (l *Logger) ErrorF(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	fmt.Printf(format+"\n", args...)
}

func (l *Logger) FatalF(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}

	fmt.Printf(format+"\n", args...)
}

// New returns a logger. Ready to log at required thereshold level
func New(threshold Level) *Logger {
	return &Logger{threshold: threshold}
}
