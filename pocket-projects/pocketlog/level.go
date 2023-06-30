package pocketlog

// Level represents available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly for debugging purpose
	LevelDebug Level = iota
	// LevelInfo represents informational level values in the application
	LevelInfo
	// LevelError represents highest logging level. To be used for error and traces.
	LevelError
)
