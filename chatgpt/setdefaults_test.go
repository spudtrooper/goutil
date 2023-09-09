package chatgpt

import (
	"reflect"
	"testing"
)

type Inner struct {
	Name  string
	Count int
}

type Test struct {
	Message string
	Number  float64
	Slice   []Inner
	Struct  Inner
}

func TestSetDefaults(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:  "Fill all fields",
			input: &Test{},
			expected: &Test{
				Message: "the Message",
				Number:  0,
				Slice:   []Inner{{Name: "the Name", Count: 0}},
				Struct:  Inner{Name: "the Name", Count: 0},
			},
		},
		{
			name: "Do not overwrite existing fields",
			input: &Test{
				Message: "Existing",
				Number:  42.0,
				Slice:   nil,
			},
			expected: &Test{
				Message: "Existing",
				Number:  42.0,
				Slice:   []Inner{{Name: "the Name", Count: 0}},
				Struct:  Inner{Name: "the Name", Count: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetDefaults(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("\n got: %+v\nwant: %+v", tt.input, tt.expected)
			}
		})
	}
}
