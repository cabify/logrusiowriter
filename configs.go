package logrusiowriter

import "github.com/sirupsen/logrus"

// Config holds the configuration to be used with WithConfig() configurer
// This struct is useful to embed into configuration structs parsed with libraries like envconfig
type Config struct {
	Level  string        `default:"info"`
	Fields logrus.Fields `default:"logger:stdlib"`
}

// WithLogger configures the logger with the one provided
func WithLogger(logger logrus.FieldLogger) Configurer {
	return func(w *writer) {
		w.logger = logger
	}
}

// WithLevel configures the level with the one provided
func WithLevel(lvl logrus.Level) Configurer {
	return func(w *writer) {
		w.level = lvl
	}
}

// WithFields configures the fields with the ones provided
func WithFields(fields logrus.Fields) Configurer {
	return func(w *writer) {
		w.fields = fields
	}
}

// WithConfig creates a configurer from the configuration provided as a struct
// If it's unable to parse the Level provided as a string, it will invoke the OnLevelParseError function and set the
// level returned by that function (a default value)
func WithConfig(cfg Config) Configurer {
	return func(w *writer) {
		lvl, err := logrus.ParseLevel(cfg.Level)
		if err != nil {
			lvl = OnLevelParseError(err)
		}
		w.level = lvl
		w.fields = cfg.Fields
	}
}

// WithConfigInterface creates a configurer from the configuration provided as an interface
func WithConfigInterface(cfg interface {
	Level() logrus.Level
	Fields() logrus.Fields
	Logger() logrus.FieldLogger
}) Configurer {
	return func(w *writer) {
		w.logger = cfg.Logger()
		w.level = cfg.Level()
		w.fields = cfg.Fields()
	}
}

// OnLevelParseError will be invoked if logrus is unable to parse the string level provided in the configuration
// The default behavior is to log it with logrus and return a default Info level,
// you can change this to log in some other system or to panic
// Changing this is not thread safe, so it might be a good idea to change it in a init() function
var OnLevelParseError = func(err error) logrus.Level {
	logrus.Errorf("Can't parse level: %s", err)
	return logrus.InfoLevel
}
