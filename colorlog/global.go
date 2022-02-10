package colorlog

import "github.com/fatih/color"

var globalLogger *logger

func init() {
	globalLogger = &logger{
		number:  *color.New(color.FgCyan),
		normal:  *color.New(color.FgWhite),
		str:     *color.New(color.FgMagenta),
		boolean: *color.New(color.FgYellow),
		uri:     *color.New(color.FgRed),
	}
}

// Println transforms the output be colorized according to the current rules
func Printf(tmpl string, args ...interface{}) {
	globalLogger.Printf(tmpl, args...)
}

// Println delegates straight to `log.Println`
func Println(s string) {
	globalLogger.Println(s)
}
