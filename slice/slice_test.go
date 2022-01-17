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
