package logrusiowriter

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestWithConfig(t *testing.T) {
	t.Run("can't parse level, configures info level by default", func(t *testing.T) {
		expectedLevel := logrus.InfoLevel

		cfg := Config{
			Level:  "none",
			Fields: logrus.Fields{},
		}

		w := New(WithConfig(cfg))

		configuredLevel := w.(*writer).level
		if configuredLevel != expectedLevel {
			t.Errorf("Configured level should be %s, but it was %s", expectedLevel, configuredLevel)
		}
	})

	t.Run("custom OnLevelParseError", func(t *testing.T) {
		originalOnLevelParseError := OnLevelParseError
		defer func() { OnLevelParseError = originalOnLevelParseError }()

		expectedLevel := logrus.WarnLevel

		cfg := Config{
			Level:  "none",
			Fields: logrus.Fields{},
		}

		var providedErr error
		OnLevelParseError = func(err error) logrus.Level {
			providedErr = err
			return expectedLevel
		}

		w := New(WithConfig(cfg))

		configuredLevel := w.(*writer).level
		if configuredLevel != expectedLevel {
			t.Errorf("Configured level should be %s, but it was %s", expectedLevel, configuredLevel)
		}

		if providedErr == nil {
			t.Errorf("Error provided to OnLevelParseError should not be nil")
		}
	})
}
