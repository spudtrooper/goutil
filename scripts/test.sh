#!/bin/sh

set -e

go test ./cond ./errors ./formatstruct ./img ./lazycond ./or ./selenium
go run testing/term_test_main.go