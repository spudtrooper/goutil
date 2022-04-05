package slice

import (
	"reflect"
	"testing"

	"github.com/spudtrooper/goutil/or"
)

func TestStrings(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		sep   string
		want  []string
		opts  []StringsOption
	}{
		{
			name:  "empty",
			input: "",
			sep:   ",",
			want:  []string{},
		},
		{
			name:  "one",
			input: "1",
			sep:   ",",
			want:  []string{"1"},
		},
		{
			name:  "uniques",
			input: "1,2,3",
			want:  []string{"1", "2", "3"},
			sep:   ",",
		},
		{
			name:  "dups",
			input: "1,2,3,1,2,3",
			sep:   ",",
			want:  []string{"1", "2", "3", "1", "2", "3"},
		},
		{
			name:  "uniques pipe",
			input: "1|2|3",
			sep:   "|",
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "uniques trimSpace",
			input: "1 , 2, 3 ",
			want:  []string{"1", "2", "3"},
			sep:   ",",
			opts:  []StringsOption{StringsTrimSpace(true)},
		},
	}
	for _, test := range tests {
		name := or.String(test.name, test.input)
		t.Run(name, func(t *testing.T) {
			if want, got := test.want, Strings(test.input, test.sep, test.opts...); !reflect.DeepEqual(want, got) {
				t.Errorf("Strings(%q,%q): want %v, got %v", test.input, test.sep, want, got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "one",
			input: []string{"1"},
			want:  []string{"1"},
		},
		{
			name:  "many",
			input: []string{"1", "2", "3"},
			want:  []string{"3", "2", "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Reverse(test.input)
			if want, got := test.want, test.input; !reflect.DeepEqual(want, got) {
				t.Errorf("Reverse(%v): want %v, got %v", test.input, want, got)
			}
		})
	}
}

func TestStringDiff(t *testing.T) {
	var tests = []struct {
		name string
		a, b []string
		want []string
	}{
		{
			name: "empty",
			a:    []string{},
			b:    []string{},
			want: []string{},
		},
		{
			name: "nil-a",
			a:    nil,
			b:    []string{},
			want: []string{},
		},
		{
			name: "nil-b",
			a:    []string{},
			b:    nil,
			want: []string{},
		},
		{
			name: "empty a, non-empty b",
			a:    []string{},
			b:    []string{"one"},
			want: []string{},
		},
		{
			name: "non-empty a, non-empty b",
			a:    []string{"one", "two", "three"},
			b:    []string{"one", "three"},
			want: []string{"two"},
		},
		{
			name: "non-empty a with dups, non-empty b",
			a:    []string{"one", "one", "two", "two", "three", "three"},
			b:    []string{"one", "three"},
			want: []string{"two", "two"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if want, got := test.want, StringDiff(test.a, test.b); !reflect.DeepEqual(want, got) {
				t.Errorf("StringDiff(%q,%q): want %v, got %v", test.a, test.b, want, got)
			}
		})
	}
}
