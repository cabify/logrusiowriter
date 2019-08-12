package logrusiowriter

import (
	"bytes"
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

	t.Run("trimming", func(t *testing.T) {
		const newLineAppendedByLogrus = "\n"
		for _, tc := range []struct {
			desc     string
			trimming bool
			input    string
			expected string
		}{
			{
				desc:     "zerolength with trimming",
				trimming: true,
				input:    "",
				expected: "level=info" + newLineAppendedByLogrus,
			},
			{
				desc:     "zerolength without trimming",
				trimming: false,
				input:    "",
				expected: "level=info" + newLineAppendedByLogrus,
			},
			{
				desc:     "only newline with trimming",
				trimming: true,
				input:    "\n",
				expected: "level=info" + newLineAppendedByLogrus,
			},
			{
				desc:     "only newline without trimming",
				trimming: false,
				input:    "\n",
				expected: `level=info msg="\n"` + newLineAppendedByLogrus,
			},
			{
				desc:     "with trailing newline and trimming",
				trimming: true,
				input:    "message\n",
				expected: "level=info msg=message" + newLineAppendedByLogrus,
			},
			{
				desc:     "with trailing newline and no trimming",
				trimming: false,
				input:    "message\n",
				expected: `level=info msg="message\n"` + newLineAppendedByLogrus,
			},
			{
				desc:     "no newline with trimming",
				trimming: true,
				input:    "message",
				expected: "level=info msg=message" + newLineAppendedByLogrus,
			},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				buf := &bytes.Buffer{}

				bufLogger := logrus.New()
				bufLogger.SetOutput(buf)
				bufLogger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true

				writer := New(WithLogger(bufLogger), WithTrailingNewLineTrimming(tc.trimming))

				_, _ = writer.Write([]byte(tc.input))

				if buf.String() != tc.expected {
					t.Errorf("Unexpected output\nExpected: '%s'\nGot:      '%s'", tc.expected, buf.String())
				}
			})
		}
	})
}
