package formatstruct

import (
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
		o    interface{}
		want []string
		opts []Option
	}{
		{
			name: "empty",
			o:    struct{}{},
			want: []string{},
		},
		{
			name: "one",
			o: struct {
				A string
			}{
				A: "a",
			},
			want: []string{
				`A: a`,
			},
		},
		{
			name: "multi",
			o: struct {
				A, B, C string
			}{
				A: "a",
				B: "b",
				C: "c",
			},
			want: []string{
				`A: a`,
				`B: b`,
				`C: c`,
			},
		},
		{
			name: "multi unsorted",
			o: struct {
				B, C, A string
			}{
				A: "a",
				B: "b",
				C: "c",
			},
			want: []string{
				`A: a`,
				`B: b`,
				`C: c`,
			},
		},
		{
			name: "multi types",
			o: struct {
				A string
				B bool
				C int
				D float64
			}{
				A: "a",
				B: true,
				C: 3,
				D: 4.1,
			},
			want: []string{
				`A: a`,
				`B: true`,
				`C: 3`,
				`D: 4.1`,
			},
		},
		{
			name: "key transform",
			o: struct {
				A string
			}{
				A: "a",
			},
			want: []string{
				`A: aaa`,
			},
			opts: []Option{
				KeyTransform("A", func(val string) string {
					return val + val + val
				}),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Format(test.o, test.opts...)
			if err != nil {
				t.Fatalf("Format: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Format(%v: want %v, got %v", test.o, test.want, got)
			}
		})
	}
}
