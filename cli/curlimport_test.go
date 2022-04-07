// curlimport will imort a curl command into the goutil/request framework.
package cli

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
		want    *curlCmd
		wantErr bool
	}{
		{
			name: "put",
			content: `curl 'https://community.hannity.com/u/tucker_carlson.json' \
			-X 'PUT' \
			-H 'authority: community.hannity.com' \
			--data-raw 'name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=' \
			--compressed`,
			want: &curlCmd{
				uri: "https://community.hannity.com/u/tucker_carlson.json",
				headers: []header{
					{key: "authority", val: "community.hannity.com"},
				},
				data:   "name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=",
				method: "PUT",
			},
		},
		{
			name: "post",
			content: `curl 'https://community.hannity.com/u/tucker_carlson.json' \
			-H 'authority: community.hannity.com' \
			--data-raw 'name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=' \
			--compressed`,
			want: &curlCmd{
				uri: "https://community.hannity.com/u/tucker_carlson.json",
				headers: []header{
					{key: "authority", val: "community.hannity.com"},
				},
				data:   "name=The+Tucker+Carlson&title=&primary_group_id=&flair_group_id=",
				method: "POST",
			},
		},
		{
			name: "get",
			content: `curl 'https://community.hannity.com/u/tucker_carlson.json' \
				-H 'authority: community.hannity.com' \
				--compressed`,
			want: &curlCmd{
				uri: "https://community.hannity.com/u/tucker_carlson.json",
				headers: []header{
					{key: "authority", val: "community.hannity.com"},
				},
				method: "GET",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseCurlCmd(test.content)
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
