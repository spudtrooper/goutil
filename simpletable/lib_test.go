package simpletable

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name string
		opts []NewOption
		rows [][]string
		want string
	}{
		{
			name: "empty",
			want: `+
+`,
		},
		{
			name: "one",
			rows: [][]string{{"one"}},
			want: `+-----+
| one |
+-----+`,
		},
		{
			name: "multi",
			rows: [][]string{{"one", "two", "three"}, {"four", "five", "six"}},
			want: `+------+------+-------+
| one  | two  | three |
| four | five | six   |
+------+------+-------+`,
		},
		{
			name: "multi-header",
			opts: []NewOption{NewHeader([]string{"h1", "h2", "h3"})},
			rows: [][]string{{"one", "two", "three"}, {"four", "five", "six"}},
			want: `+------+------+-------+
|  H1  |  H2  |  H3   |
+------+------+-------+
| one  | two  | three |
| four | five | six   |
+------+------+-------+`,
		},
		{
			name: "multi-header-noborder",
			opts: []NewOption{
				NewHeader([]string{"h1", "h2", "h3"}),
				NewNoBorder(true),
			},
			rows: [][]string{{"one", "two", "three"}, {"four", "five", "six"}},
			want: `H1  |  H2  |  H3
-------+------+--------
  one  | two  | three
  four | five | six`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			s := New(&out, tt.opts...)
			for _, row := range tt.rows {
				s.Append(row)
			}
			s.Render()

			// Normalize the output because the terminal hides trailing whitespace
			var got string
			{
				var arr []string
				for _, line := range strings.Split(out.String(), "\n") {
					s := line
					s = strings.TrimRight(s, " ")
					arr = append(arr, s)
				}
				got = strings.TrimSpace(strings.Join(arr, "\n"))
			}
			if got != tt.want {
				log.Printf("\n<<%s>>", got)
				t.Errorf("Render() = \n<<%v>>, want \n<<%v>>", got, tt.want)
			}
		})
	}
}
