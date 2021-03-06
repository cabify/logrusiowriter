package logrusiowriter

import (
	"io"

	"github.com/sirupsen/logrus"
)

// New creates a new io.Writer, and configures it with the provided configurers
// Provided Configurers overwrite previous values as they're applied
// If no Configurers provided, the writer will log with Info level and no fields using the logrus.StandardLogger
func New(cfg ...Configurer) io.Writer {
	w := &writer{
		logger:                  logrus.StandardLogger(),
		level:                   logrus.InfoLevel,
		fields:                  make(map[string]interface{}),
		trailingNewLineTrimming: true,
	}
	for _, c := range cfg {
		c(w)
	}

	return w
}

// Configurer configures the writer, use one of the With* functions to obtain one
type Configurer func(*writer)

// writer implements io.Writer
type writer struct {
	logger                  logrus.FieldLogger
	level                   logrus.Level
	fields                  map[string]interface{}
	trailingNewLineTrimming bool
}

// Write will write with the logger, level and fields set in the writer
func (w *writer) Write(bytes []byte) (int, error) {
	l := len(bytes)
	if w.trailingNewLineTrimming && l > 0 && bytes[l-1] == '\n' {
		bytes = bytes[:l-1]
	}
	w.logger.WithFields(w.fields).Log(w.level, string(bytes))
	return l, nil
}
