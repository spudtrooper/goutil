package time

import (
	"reflect"
	"testing"
	"time"
)

func mustParse(t *testing.T, layout, value string) time.Time {
	res, err := time.Parse(layout, value)
	if err != nil {
		t.Fatalf("mustParse: %+v", err)
	}
	return res
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:  "simple dashes",
			input: "2018-12-31",
			want:  mustParse(t, "2006-01-02", "2018-12-31"),
		},
		{
			name:  "simple slashes",
			input: "2018/12/31",
			want:  mustParse(t, "2006/01/02", "2018/12/31"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
