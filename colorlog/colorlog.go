// package colorlog provides logging where entities are colored consistently, e.g. numbers are cyan
package colorlog

import (
	"log"
	"regexp"

	"github.com/fatih/color"
	"github.com/spudtrooper/goutil/sets"
)

var (
	formatRE = regexp.MustCompile(`(%[\-\+\d\.e]*[a-z])`)
	uriRE    = regexp.MustCompile(`^[a-z]+:\/\/.*`)
)

type logger struct {
	number         color.Color
	normal         color.Color
	str            color.Color
	boolean        color.Color
	uri            color.Color
	specialStrings sets.StringSet
	specialString  color.Color
}

func (l *logger) Number(c color.Color) {
	l.number = c
}
func (l *logger) Normal(c color.Color) {
	l.normal = c
}
func (l *logger) String(c color.Color) {
	l.str = c
}
func (l *logger) Bool(c color.Color) {
	l.boolean = c
}
func (l *logger) URI(c color.Color) {
	l.uri = c
}
func (l *logger) SpecialStrings(ss ...string) {
	l.specialStrings = sets.String(ss)
}
func (l *logger) SpecialString(c color.Color) {
	l.specialString = c
}

type transform struct {
	col    color.Color
	format string
}

// Println transforms the output be colorized according to the current rules
func (l *logger) Printf(tmpl string, args ...interface{}) {
	newTmpl, newArgs, _ := l.printf(tmpl, args...)
	log.Printf(newTmpl, newArgs...)
}

// Fatalf transforms the output be colorized according to the current rules
func (l *logger) Fatalf(tmpl string, args ...interface{}) {
	newTmpl, newArgs, _ := l.printf(tmpl, args...)
	log.Fatalf(newTmpl, newArgs...)
}

// Println delegates straight to `log.Println`
func (l *logger) Println(ss ...string) {
	if len(ss) == 0 {
		log.Println()
		return
	}
	if len(ss) != 1 {
		panic("Println() takes either 0 or 1 arg")
	}
	log.Println(ss[0])
}

func (l *logger) printf(tmpl string, args ...interface{}) (string, []interface{}, []transform) {
	newTmpl, trans := l.convert(tmpl, args)
	var newArgs []interface{}
	for i, arg := range args {
		t := trans[i]
		newArg := t.col.Sprintf(t.format, arg)
		newArgs = append(newArgs, newArg)
	}
	return newTmpl, newArgs, trans
}

func (l *logger) convert(tmpl string, args []interface{}) (string, []transform) {
	var trans []transform
	for i, m := range formatRE.FindAllStringSubmatch(tmpl, -1) {
		format := m[1]
		col := l.normal
		switch last := string(format[len(format)-1]); last {
		case "d", "f":
			col = l.number
		case "t":
			col = l.boolean
		case "q":
			if l.isSpecialString(args[i]) {
				col = l.specialString
			} else {
				col = l.str
			}
		case "s":
			if l.isSpecialString(args[i]) {
				col = l.specialString
			} else if isURI(args[i]) {
				col = l.uri
			}
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

func (l *logger) isSpecialString(s interface{}) bool {
	switch v := s.(type) {
	case string:
		if l.specialStrings[v] {
			return true
		}
	}
	return false
}

func isURI(s interface{}) bool {
	switch v := s.(type) {
	case string:
		return uriRE.MatchString(v)
	}
	return false
}
