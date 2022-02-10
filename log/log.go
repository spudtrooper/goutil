package log

import (
	"log"

	"github.com/spudtrooper/goutil/colorlog"
)

type Logger interface {
	Printf(tmpl string, args ...interface{})
	Println(s string)
}

func MakeLog(prefix string, mOpts ...MakeLogOption) Logger {
	opts := MakeMakeLogOptions(mOpts...)
	if opts.Color() {
		return &colorLogger{base{"[" + prefix + "] "}}
	}
	return &logger{base{"[" + prefix + "] "}}
}

type base struct {
	prefix string
}

type logger struct{ base }

func (l *logger) Printf(tmpl string, args ...interface{}) {
	log.Printf(l.prefix+tmpl, args...)
}

func (l *logger) Println(s string) {
	log.Println(l.prefix + s)
}

type colorLogger struct{ base }

func (l *colorLogger) Printf(tmpl string, args ...interface{}) {
	colorlog.Printf(l.prefix+tmpl, args...)
}

func (l *colorLogger) Println(s string) {
	colorlog.Println(l.prefix + s)
}
