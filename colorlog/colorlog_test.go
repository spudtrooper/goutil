package colorlog

import (
	"reflect"
	"testing"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/sets"
)

func TestPrintf(t *testing.T) {
	globalLogger.specialStrings = sets.String([]string{"special"})
	var tests = []struct {
		name      string
		tmpl      string
		args      []interface{}
		wantTmpl  string
		wantArgs  []interface{}
		wantTrans []transform
		logger    *logger
	}{
		{
			tmpl:     "bool: %t",
			args:     []interface{}{true},
			wantTmpl: "bool: %s",
			wantArgs: []interface{}{"true"},
			wantTrans: []transform{
				{
					format: "%t",
					col:    globalLogger.boolean,
				},
			},
		},
		{
			tmpl:     "number: %d",
			args:     []interface{}{9},
			wantTmpl: "number: %s",
			wantArgs: []interface{}{"9"},
			wantTrans: []transform{
				{
					format: "%d",
					col:    globalLogger.number,
				},
			},
		},
		{
			tmpl:     "float: %f",
			args:     []interface{}{9.0},
			wantTmpl: "float: %s",
			wantArgs: []interface{}{"9.000000"},
			wantTrans: []transform{
				{
					format: "%f",
					col:    globalLogger.number,
				},
			},
		},
		{
			tmpl:     "float: %1.2f",
			args:     []interface{}{9.0},
			wantTmpl: "float: %s",
			wantArgs: []interface{}{"9.00"},
			wantTrans: []transform{
				{
					format: "%1.2f",
					col:    globalLogger.number,
				},
			},
		},
		{
			tmpl:     "string: %s",
			args:     []interface{}{"9"},
			wantTmpl: "string: %s",
			wantArgs: []interface{}{"9"},
			wantTrans: []transform{
				{
					format: "%s",
					col:    globalLogger.str,
				},
			},
		},
		{
			tmpl:     "special string: %s",
			args:     []interface{}{"special"},
			wantTmpl: "special string: %s",
			wantArgs: []interface{}{"special"},
			wantTrans: []transform{
				{
					format: "%s",
					col:    globalLogger.specialString,
				},
			},
		},
		{
			tmpl:     "special quoted: %q",
			args:     []interface{}{"special"},
			wantTmpl: "special quoted: %s",
			wantArgs: []interface{}{"\"special\""},
			wantTrans: []transform{
				{
					format: "%q",
					col:    globalLogger.specialString,
				},
			},
		},
		{
			tmpl:     "quoted: %q",
			args:     []interface{}{"9"},
			wantTmpl: "quoted: %s",
			wantArgs: []interface{}{"\"9\""},
			wantTrans: []transform{
				{
					format: "%q",
					col:    globalLogger.str,
				},
			},
		},
		{
			tmpl:     "uri: %s",
			args:     []interface{}{"http://foo.com"},
			wantTmpl: "uri: %s",
			wantArgs: []interface{}{"http://foo.com"},
			wantTrans: []transform{
				{
					format: "%s",
					col:    globalLogger.uri,
				},
			},
		},
	}
	for _, test := range tests {
		name := or.String(test.name, test.tmpl)
		t.Run(name, func(t *testing.T) {
			l := globalLogger
			if test.logger != nil {
				l = test.logger
			}
			l.Printf(test.tmpl, test.args...)

			gotTmpl, gotArgs, gotTrans := l.printf(test.tmpl, test.args...)
			if want, got := test.wantTmpl, gotTmpl; want != got {
				t.Errorf("args: want != got: %v %v", want, got)
			}
			if want, got := test.wantArgs, gotArgs; !reflect.DeepEqual(want, got) {
				t.Errorf("args: want != got: %v %v", want, got)
			}
			if want, got := test.wantTrans, gotTrans; !reflect.DeepEqual(want, got) {
				t.Errorf("args: want != got: %v %v", want, got)
			}
		})
	}
}
