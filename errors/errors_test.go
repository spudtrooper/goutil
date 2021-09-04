package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrorCollector(t *testing.T) {
	var tests = []struct {
		name   string
		errors []error
		empty  bool
		err    error
	}{
		{
			name:   "no add",
			errors: []error{},
			empty:  true,
			err:    nil,
		},
		{
			name:   "one nil",
			errors: []error{nil},
			empty:  true,
			err:    nil,
		},
		{
			name:   "multiple nils",
			errors: []error{nil, nil, nil},
			empty:  true,
			err:    nil,
		},
		{
			name:   "one non-nil",
			errors: []error{fmt.Errorf("A")},
			empty:  false,
			err:    fmt.Errorf("A"),
		},
		{
			name:   "multiple non-nils",
			errors: []error{fmt.Errorf("A"), fmt.Errorf("B"), fmt.Errorf("C")},
			empty:  false,
			err: fmt.Errorf(strings.TrimSpace(`
1: A
2: B
3: C
`)),
		},
		{
			name:   "mixed",
			errors: []error{nil, fmt.Errorf("A"), nil, fmt.Errorf("B"), nil, fmt.Errorf("C"), nil},
			empty:  false,
			err: fmt.Errorf(strings.TrimSpace(`
1: A
2: B
3: C
`)),
		},
	}
	runTests := func(name string, ctor func() ErrorCollector, t *testing.T) {
		for _, test := range tests {
			t.Run(name+"."+test.name, func(t *testing.T) {
				c := ctor()
				for _, err := range test.errors {
					c.Add(err)
				}
				if empty := c.Empty(); test.empty != empty {
					t.Fatalf("Empty: want %t, got %t", test.empty, empty)
				}
				err := c.Build()
				if test.err != nil && err == nil {
					t.Fatalf("Build: want not nil, got nil")
				}
				if test.err == nil && err != nil {
					t.Fatalf("Build: want  nil, got %v", err)
				}
				if test.err != nil && err != nil && test.err.Error() != err.Error() {
					t.Fatalf("Build: want %v, got %v", test.err, err)
				}
			})
		}
	}
	runTests("MakeErrorCollector", MakeErrorCollector, t)
	runTests("MakeSyncErrorCollector", MakeSyncErrorCollector, t)
}
