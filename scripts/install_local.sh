#!/bin/sh

set -e

go build -o goutilctrl main.go
cp goutilctrl ~/go/bin

# Testing
goutilctrl