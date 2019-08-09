package logrusiowriter_test

import (
	"os"

	"github.com/sirupsen/logrus"
)

var datelessFormatter = new(logrus.TextFormatter)

func init() {
	datelessFormatter.DisableTimestamp = true
}

// removeTimestampAndSetOutputToStdout removes date from logrus logs, and redirects them to os.Stdout
// this can't be done in init() because os.Stdout changes after calling init() in examples:
// see: https://unexpected-go.com/os-stdout-changes-after-init-in-examples.html
func removeTimestampAndSetOutputToStdout(logger *logrus.Logger) {
	logger.SetFormatter(datelessFormatter)
	logger.SetOutput(os.Stdout)
}
