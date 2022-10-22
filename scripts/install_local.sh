#!/bin/sh

set -e

go build main.go
# cp main ~/go/bin/goutil
rm -f ~/go/bin/goutil
ln -fns /Users/jeff/Projects/goutil/main ~/go/bin/goutil