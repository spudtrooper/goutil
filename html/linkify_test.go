package html

import (
	"testing"
)

func TestLinkify(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "no links",
			input: "blah blah blah",
			want:  "blah blah blah",
		},
		{
			name:  "http link",
			input: "http://foo-bar_baz.com",
			want:  `<a target="_" href="http://foo-bar_baz.com">foo-bar_baz.com</a>`,
		},
		{
			name:  "https link",
			input: "https://foo.com",
			want:  `<a target="_" href="https://foo.com">foo.com</a>`,
		},
		{
			name:  "http link ending in a period",
			input: "http://foo-bar_baz.com.",
			want:  `<a target="_" href="http://foo-bar_baz.com">foo-bar_baz.com</a>.`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Linkify(test.input); test.want != got {
				t.Errorf("Linkify(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}
