package pocketlog_test

import (
	"encoding/json"
	"learn-go-pockets/logger/pocketlog"
	"strings"
	"testing"
)

const (
	debugMessage = "What's in a word?"
	infoMessage  = "Oh 😞 so you think you are cute hein?"
	errorMessage = "To err is human, to forgive is divine."
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s!", "world")
	// Output:
	// {"level":"debug","message":"Hello, world!"}
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	tests := map[string]struct {
		level    pocketlog.Level
		expected []pocketlog.LogEntry
	}{
		"debug": {
			level: pocketlog.LevelDebug,
			expected: []pocketlog.LogEntry{
				{"debug", debugMessage},
				{"info", infoMessage},
				{"error", errorMessage},
			},
		},
		"info": {
			level: pocketlog.LevelInfo,
			expected: []pocketlog.LogEntry{
				{"info", infoMessage},
				{"error", errorMessage},
			},
		},
		"error": {
			level: pocketlog.LevelError,
			expected: []pocketlog.LogEntry{
				{"error", errorMessage},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			// Split log lines
			lines := splitLines(tw.contents)
			if len(lines) != len(tc.expected) {
				t.Fatalf("expected %d log lines, got %d", len(tc.expected), len(lines))
			}

			for i, line := range lines {
				var got pocketlog.LogEntry
				if err := json.Unmarshal([]byte(line), &got); err != nil {
					t.Fatalf("invalid JSON log: %v", err)
				}

				if got != tc.expected[i] {
					t.Errorf("expected %+v, got %+v", tc.expected[i], got)
				}
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implemnents the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

// splitLines splits output by newlines, trimming empties.
func splitLines(s string) []string {
	var lines []string
	for l := range strings.SplitSeq(s, "\n") {
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines
}
