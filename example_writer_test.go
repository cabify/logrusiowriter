package logrusiowriter_test

import (
	"log"

	"github.com/cabify/logrusiowriter"
	"github.com/sirupsen/logrus"
)

func ExampleNew() {
	removeTimestampAndSetOutputToStdout(logrus.StandardLogger())

	log.SetOutput(logrusiowriter.New())
	log.SetFlags(0) // no date on standard logger

	log.Printf("Standard log")

	// Output:
	// level=info msg="Standard log\n"
}
