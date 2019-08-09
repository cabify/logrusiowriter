package logrusiowriter_test

import (
	"fmt"

	"github.com/cabify/logrusiowriter"
	"github.com/sirupsen/logrus"
)

func ExampleWithConfig() {
	removeTimestampAndSetOutputToStdout(logrus.StandardLogger())

	config := logrusiowriter.Config{
		Level: "warning",
		Fields: map[string]interface{}{
			"config": "struct",
		},
	}

	writer := logrusiowriter.New(
		logrusiowriter.WithConfig(config),
	)

	_, _ = fmt.Fprint(writer, "Hello World!")
	// Output:
	// level=warning msg="Hello World!" config=struct
}

func ExampleWithFields() {
	removeTimestampAndSetOutputToStdout(logrus.StandardLogger())

	writer := logrusiowriter.New(
		logrusiowriter.WithFields(logrus.Fields{
			"config": "fields",
			"other":  288,
		}),
	)

	_, _ = fmt.Fprint(writer, "Hello World!")
	// Output:
	// level=info msg="Hello World!" config=fields other=288
}

func ExampleWithLevel() {
	removeTimestampAndSetOutputToStdout(logrus.StandardLogger())

	writer := logrusiowriter.New(
		logrusiowriter.WithLevel(logrus.ErrorLevel),
	)

	_, _ = fmt.Fprint(writer, "Hello World!")
	// Output:
	// level=error msg="Hello World!"
}

func ExampleWithLogger() {
	logger := logrus.New()
	removeTimestampAndSetOutputToStdout(logger)
	logger.SetLevel(logrus.TraceLevel)

	writer := logrusiowriter.New(
		logrusiowriter.WithLogger(logger),
	)

	_, _ = fmt.Fprint(writer, "Hello World!")
	// Output:
	// level=info msg="Hello World!"
}

func ExampleWithConfigInterface() {
	removeTimestampAndSetOutputToStdout(logrus.StandardLogger())

	writer := logrusiowriter.New(
		logrusiowriter.WithConfigInterface(configProvider{}),
	)

	_, _ = fmt.Fprint(writer, "Hello World!")
	// Output:
	// level=trace msg="Hello World!" config=interface
}

type configProvider struct{}

func (configProvider) Level() logrus.Level { return logrus.TraceLevel }

func (configProvider) Fields() logrus.Fields { return logrus.Fields{"config": "interface"} }

func (configProvider) Logger() logrus.FieldLogger {
	logger := logrus.New()
	removeTimestampAndSetOutputToStdout(logger)
	logger.SetLevel(logrus.TraceLevel)
	return logger
}
