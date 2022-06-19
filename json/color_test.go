package json

import "testing"

func TestColorMarshal(t *testing.T) {
	type foo struct {
		Bar int
	}
	tests := []struct {
		name    string
		input   interface{}
		indent  int
		want    string
		wantErr bool
	}{
		{
			name:  "empty",
			input: nil,
			want:  "{}",
		},
		{
			name:  "simple",
			input: foo{Bar: 1},
			want:  "{\n  \"Bar\": 1\n}",
		},
		{
			name:   "simple",
			input:  foo{Bar: 1},
			indent: 4,
			want:   "{\n    \"Bar\": 1\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ColorMarshal(tt.input, ColorMarshalIndent(tt.indent))
			if (err != nil) != tt.wantErr {
				t.Errorf("ColorMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if want := tt.want; got != want {
				t.Errorf("ColorMarshal(), want %q, got %q", want, got)

			}
		})
	}
}
