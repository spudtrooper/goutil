#!/bin/sh

set -e

go build main.go
rm -f ~/go/bin/goutil
cp main ~/go/bin/goutil