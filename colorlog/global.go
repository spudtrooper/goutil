package colorlog

import "github.com/fatih/color"

var globalLogger *logger

func init() {
	globalLogger = &logger{
		number:        *color.New(color.FgCyan),
		normal:        *color.New(color.FgWhite),
		str:           *color.New(color.FgMagenta),
		quotedString:  *color.New(color.FgRed),
		boolean:       *color.New(color.FgYellow),
		uri:           *color.New(color.FgHiGreen).Add(color.Underline),
		specialString: *color.New(color.FgHiRed),
	}
}

// Printf transforms the output be colorized according to the current rules
func Printf(tmpl string, args ...interface{}) {
	globalLogger.Printf(tmpl, args...)
}

// Fatalf transforms the output be colorized according to the current rules
func Fatalf(tmpl string, args ...interface{}) {
	globalLogger.Fatalf(tmpl, args...)
}

// Println delegates straight to `log.Println`
func Println(ss ...string) {
	globalLogger.Println(ss...)
}

// Number sets the global number color
func Number(c color.Color) {
	globalLogger.Number(c)
}

// Normal sets the global normal color
func Normal(c color.Color) {
	globalLogger.Normal(c)
}

// String sets the global string color
func String(c color.Color) {
	globalLogger.String(c)
}

// Bool sets the global bool color
func Bool(c color.Color) {
	globalLogger.Bool(c)
}

// URI sets the global uri color
func URI(c color.Color) {
	globalLogger.URI(c)
}

// SpecialStrings sets the global specialstrings set
func SpecialStrings(ss ...string) {
	globalLogger.SpecialStrings(ss...)
}

// SpecialString sets the global specialstring color
func SpecialString(c color.Color) {
	globalLogger.SpecialString(c)
}
