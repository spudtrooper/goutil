package log

import "log"

type logger struct {
	prefix string
}

func MakeLog(prefix string) *logger {
	return &logger{"[" + prefix + "] "}
}

func (l *logger) Printf(tmpl string, args ...interface{}) {
	log.Printf(l.prefix+tmpl, args...)
}

func (l *logger) Println(s string) {
	log.Println(l.prefix + s)
}
