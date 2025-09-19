package pocketlog_test

import "learn-go-pockets/logger/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, nil)
	debugLogger.Debugf("Hello, %s!", "world")
	// Output:
	// Hello, world!
}
