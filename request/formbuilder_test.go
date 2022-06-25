package request

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"strings"
	"testing"
)

// ugh, whitespace is so annoying in these.
func norm(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func TestFormBuilder_Field(t *testing.T) {
	type add struct {
		key, val string
	}
	tests := []struct {
		name    string
		adds    []add
		wantErr bool
		want    func(contentType string) string
	}{
		{
			name: "empty",
			want: func(boundary string) string {
				return fmt.Sprintf(`
--%s--
`, boundary)
			},
		},
		{
			name: "one",
			adds: []add{
				{"key-1", "val-1"},
			},
			want: func(boundary string) string {
				return fmt.Sprintf(`
--%s
Content-Disposition: form-data; name="key-1"

val-1
--%s--
`, boundary, boundary)
			},
		},
		{
			name: "many",
			adds: []add{
				{"key-1", "val-1"},
				{"key-2", "val-2"},
				{"key-3", "val-3"},
			},
			want: func(boundary string) string {
				return fmt.Sprintf(`
--%s
Content-Disposition: form-data; name="key-1"

val-1
--%s
Content-Disposition: form-data; name="key-2"

val-2
--%s
Content-Disposition: form-data; name="key-3"

val-3
--%s--
`, boundary, boundary, boundary, boundary)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormBuilder()
			for _, a := range tt.adds {
				if err := f.Field(a.key, a.val); (err != nil) != tt.wantErr {
					t.Errorf("FormBuilder.Field() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if want, got := norm(tt.want(f.w.Boundary())), norm(f.String()); !reflect.DeepEqual(want, got) {
				log.Printf("got\n%s\n", f.String())
				t.Errorf("FormBuilder.String()\nwant <<<%s>>> != got <<<%s>>>", want, got)
			}
		})
	}
}

func TestFormBuilder_File(t *testing.T) {
	type add struct {
		key, file string
	}
	tests := []struct {
		name    string
		adds    []add
		wantErr bool
		want    func(contentType string) string
	}{
		{
			name: "one",
			adds: []add{
				{"key-1", "file-1.txt"},
			},
			want: func(boundary string) string {
				return fmt.Sprintf(`
--%s
Content-Disposition: form-data; name="key-1"; filename="file-1.txt"
Content-Type: application/octet-stream

file-1 content
--%s--
`, boundary, boundary)
			},
		},
		{
			name: "many",
			adds: []add{
				{"key-1", "file-1.txt"},
				{"key-2", "file-2.txt"},
				{"key-3", "file-3.txt"},
			},
			want: func(boundary string) string {
				return fmt.Sprintf(`
--%s
Content-Disposition: form-data; name="key-1"; filename="file-1.txt"
Content-Type: application/octet-stream

file-1 content
--%s
Content-Disposition: form-data; name="key-2"; filename="file-2.txt"
Content-Type: application/octet-stream

file-2 content
--%s
Content-Disposition: form-data; name="key-3"; filename="file-3.txt"
Content-Type: application/octet-stream

file-3 content
--%s--
`, boundary, boundary, boundary, boundary)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormBuilder()
			for _, a := range tt.adds {
				if err := f.File(a.key, path.Join("testdata", a.file)); (err != nil) != tt.wantErr {
					t.Errorf("FormBuilder.File() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if want, got := norm(tt.want(f.w.Boundary())), norm(f.String()); !reflect.DeepEqual(want, got) {
				log.Printf("got\n%s\n", f.String())
				t.Errorf("FormBuilder.String()\nwant <<<%s>>> != got <<<%s>>>", want, got)
			}
		})
	}
}
