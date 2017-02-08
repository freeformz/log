package log_test

import (
	"bytes"
	"testing"

	"github.com/freeformz/log"
	"github.com/freeformz/log/level"
)

func TestBasicLog(t *testing.T) {
	var b bytes.Buffer
	out := `I am a message\nSo am I\n`
	l := log.New()
	l = l.WithWriter(&b)
	l.Log("I am a message")
	l.Log("So", "am", "I")
	if b.String() != out {
		t.Fatalf("%q != %q", b.String(), out)
	}
}

func TestFilter(t *testing.T) {
	var b bytes.Buffer
	l := log.New()
	l = l.WithWriter(&b)
	l.WithFilter(&level.Filter{At: level.Err})
	l.Log(level.Debug, "I am a debug message")
	l.Log(level.Err, "I am an error message")
	l.Log(level.Warn, "I am a warning message")
	t.Log(b.String())
}

func TestLevelWithoutFilter(t *testing.T) {
	var b bytes.Buffer
	l := log.New()
	l = l.WithWriter(&b)
	l.Log(level.Debug, "a", "b")
	l.Log(level.Err, "c", "d")
	l.Log(level.Warn, "e", "f")
	t.Log(b.String())
}
