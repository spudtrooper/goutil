// package colorlog provides logging where entities are colored consistently, e.g. numbers are cyan
package colorlog

import (
	"log"
	"regexp"

	"github.com/fatih/color"
)

type logger struct {
	number  color.Color
	normal  color.Color
	str     color.Color
	boolean color.Color
}

type transform struct {
	col    color.Color
	format string
}

// Println transforms the output be colorized according to the current rules
func (l *logger) Printf(tmpl string, args ...interface{}) {
	newTmpl, newArgs := l.printf(tmpl, args...)
	log.Printf(newTmpl, newArgs...)
}

func (l *logger) printf(tmpl string, args ...interface{}) (string, []interface{}) {
	newTmpl, trans := l.convert(tmpl)
	var newArgs []interface{}
	for i, arg := range args {
		t := trans[i]
		newArg := t.col.Sprintf(t.format, arg)
		newArgs = append(newArgs, newArg)
	}
	return newTmpl, newArgs
}

var formatRE = regexp.MustCompile(`(%[\-\+]?\d*\.?\d*[a-z])`)

func (l *logger) convert(tmpl string) (string, []transform) {
	var trans []transform
	for _, m := range formatRE.FindAllStringSubmatch(tmpl, -1) {
		format := m[1]
		col := l.normal
		switch last := string(format[len(format)-1]); last {
		case "d", "f":
			col = l.number
		case "t":
			col = l.boolean
		case "q":
			col = l.str
		}
		t := transform{
			format: format,
			col:    col,
		}
		trans = append(trans, t)
	}

	newTmpl := formatRE.ReplaceAllString(tmpl, "%s")

	return newTmpl, trans

}

// Println delegates straight to `log.Println`
func (l *logger) Println(s string) {
	log.Println(s)
}
