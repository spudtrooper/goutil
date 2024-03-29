package log

import (
	"log"

	"github.com/spudtrooper/goutil/colorlog"
)

type Logger interface {
	Printf(tmpl string, args ...interface{})
	Println(s string)
	Fatalf(tmpl string, args ...interface{})
}

func MakeLog(prefix string, mOpts ...MakeLogOption) Logger {
	opts := MakeMakeLogOptions(mOpts...)
	if prefix != "" {
		prefix = "[" + prefix + "] "
	}
	if opts.Color() {
		return &colorLogger{base{prefix}}
	}
	return &logger{base{prefix}}
}

type defaultLogger struct{}

func (l *defaultLogger) Printf(tmpl string, args ...interface{}) { log.Printf(tmpl, args...) }
func (l *defaultLogger) Println(s string)                        { log.Println(s) }
func (l *defaultLogger) Fatalf(tmpl string, args ...interface{}) { log.Fatalf(tmpl, args...) }

func Must(logger Logger) Logger {
	if logger != nil {
		return logger
	}
	return &defaultLogger{}
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

func (l *logger) Fatalf(tmpl string, args ...interface{}) {
	log.Fatalf(l.prefix+tmpl, args...)
}

type colorLogger struct{ base }

func (l *colorLogger) Printf(tmpl string, args ...interface{}) {
	colorlog.Printf(l.prefix+tmpl, args...)
}

func (l *colorLogger) Println(s string) {
	colorlog.Println(l.prefix + s)
}

func (l *colorLogger) Fatalf(tmpl string, args ...interface{}) {
	colorlog.Fatalf(l.prefix+tmpl, args...)
}
