package pocketlog_test

import "pocket-projects/pocketlog"

func ExampleLogger_DebugF() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.DebugF("Hello, %s", "World")
	// Output: Hello, World
}
