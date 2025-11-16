package pocketlog_test

import (
	"testing"

	"github.com/pschulze/pocket-sized-go/logger/pocketlog"
)

const (
	debugMessage = "This is a debug message"
	infoMessage  = "This is an info message"
	errorMessage = "This is an error message"
)

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents += string(p)
	return len(p), nil
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "D - " + debugMessage + "\n" + "I - " + infoMessage + "\n" + "E - " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "I - " + infoMessage + "\n" + "E - " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "E - " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			lgr := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			lgr.Debugf(debugMessage)
			lgr.Infof(infoMessage)
			lgr.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_Truncation(t *testing.T) {
	tw := &testWriter{}
	maxLen := 10
	lgr := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(tw), pocketlog.WithMaxLen(maxLen))

	longMessage := "This message is definitely longer than out maxLen."
	lgr.Infof("%s", longMessage)

	expected := "I - Thi...\n"

	if tw.contents != expected {
		t.Errorf("invalid contents, expected %q, got %q", expected, tw.contents)
	}
}
