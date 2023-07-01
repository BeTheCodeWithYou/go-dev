package pocketlog

// Level represents available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly for debugging purpose
	LevelDebug Level = iota + 1
	// LevelInfo represents informational level values in the application
	LevelInfo
	// LevelWarn represents medium logging level, as name suggest, used for warning purpose.
	LevelWarn
	// LevelError is for reporting errors in the application.
	LevelError
	// LevelFatal represents application situation is catastrophic.
	LevelFatal
)
