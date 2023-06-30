package pocketlog

// Logger for logging functionality
type Logger struct{
	threshold Level
}

func (l *Logger) DebugF(format string, args ... any){
		// implement me
}

func (l *Logger) InfoF(format string, args ...any){
	// implement me
}

func (l *Logger) ErrorF(format string, args ...any){
	// implement me
}

// New returns a logger. Ready to log at required thereshold level
func New(threshold Level) *Logger {
	return &Logger{threshold: threshold,}
}