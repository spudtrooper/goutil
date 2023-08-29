package curlimport

import (
	"reflect"
	"testing"
)

func TestParseCurlCmd(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		content string
		want    *CurlCmd
		wantErr bool
	}{
		{
			name: "put",
			content: `curl 'https://community.hannity.com/u/tucker_carlson.json' \
			-X 'PUT' \
			-H 'authority: community.hannity.com' \
			--data-raw 'name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=' \
			--compressed`,
			want: &CurlCmd{
				URI: "https://community.hannity.com/u/tucker_carlson.json",
				Headers: []Header{
					{Key: "authority", Val: "community.hannity.com"},
				},
				Data:   "name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=",
				Method: "PUT",
			},
		},
		{
			name: "post",
			content: `curl 'https://community.hannity.com/u/tucker_carlson.json' \
			-H 'authority: community.hannity.com' \
			--data-raw 'name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=' \
			--compressed`,
			want: &CurlCmd{
				URI: "https://community.hannity.com/u/tucker_carlson.json",
				Headers: []Header{
					{Key: "authority", Val: "community.hannity.com"},
				},
				Data:   "name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=",
				Method: "POST",
			},
		},
		{
			name: "get",
			content: `curl 'https://community.hannity.com/u/tucker_carlson.json' \
				-H 'authority: community.hannity.com' \
				--compressed`,
			want: &CurlCmd{
				URI: "https://community.hannity.com/u/tucker_carlson.json",
				Headers: []Header{
					{Key: "authority", Val: "community.hannity.com"},
				},
				Method: "GET",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Parse(test.content)
			if (err != nil) != test.wantErr {
				t.Errorf("parseCurlCmd() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("parseCurlCmd() = %+v, want %+v", got, test.want)
			}
		})
	}
}
