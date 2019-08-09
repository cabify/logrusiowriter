# logrusiowriter
## `io.Writer` implementation using logrus

[![Travis CI build status](https://travis-ci.com/cabify/logrusiowriter.svg?branch=master)](https://travis-ci.com/cabify/logrusiowriter)
[![Coverage Status](https://coveralls.io/repos/github/cabify/logrusiowriter/badge.svg)](https://coveralls.io/github/cabify/logrusiowriter)
[![GoDoc](https://godoc.org/github.com/cabify/logrusiowriter?status.svg)](https://godoc.org/github.com/cabify/logrusiowriter)

# Motivation

Many golang libraries use the golang's `log` package to print their logs. This means that if your application
uses logrus to print structured logging, those packages will print a format that is (probably) incompatible with yours,
and you may end losing logs in your logs collector because they can't be parsed properly.

# Solution

Print the logs written using `log.Printf` through `logrus`, by setting `log.SetOutput` to an `io.Writer` implementation
that uses `logrus` as output, i.e.:

```go
	log.SetOutput(logrusiowriter.New())
```

See `example_*_test.go` files to find testable examples that serve as documentation.