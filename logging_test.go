package logging

import (
	"testing"
)

var l = GetLogger()

func TestLogger_Info(t *testing.T) {

	t.Run("info", func(t *testing.T) {
		l.Info("Hello, World")
	})
}

func TestLogger_Debug(t *testing.T) {
	t.Run("debug", func(t *testing.T) {
		l.Debug("Hello, World")
	})
}

func TestLogger_Warning(t *testing.T) {
	t.Run("warning", func(t *testing.T) {
		l.Warn("Hello, World")
	})
}

func TestLogger_Error(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		l.Error("Hello, World")
	})
}
