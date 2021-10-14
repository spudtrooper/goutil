package net

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	tests := []struct {
		name string
		url  string
		res  []*regexp.Regexp
	}{
		{
			name: "jeffpalm.com",
			url:  "http://jeffpalm.com",
			res: []*regexp.Regexp{
				regexp.MustCompile(`Jeff Palm`),
			},
		},
		{
			name: "google.com",
			url:  "https://google.com",
			res: []*regexp.Regexp{
				regexp.MustCompile(`<title>Google</title>`),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := test.name + ".tmp"
			os.Remove(f)
			defer os.Remove(f)

			err := DownloadFile(f, test.url)
			if err != nil {
				t.Fatalf("DownloadFile: %v", err)
			}
			b, err := ioutil.ReadFile(f)
			if err != nil {
				t.Fatalf("ReadFile: %v", err)
			}
			s := string(b)
			for _, re := range test.res {
				if !re.MatchString(s) {
					t.Errorf("re=%v doesn't match: %s", re, s)
				}
			}
		})
	}
}

func TestReadURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		res  []*regexp.Regexp
	}{
		{
			name: "jeffpalm.com",
			url:  "http://jeffpalm.com",
			res: []*regexp.Regexp{
				regexp.MustCompile(`Jeff Palm`),
			},
		},
		{
			name: "google.com",
			url:  "https://google.com",
			res: []*regexp.Regexp{
				regexp.MustCompile(`<title>Google</title>`),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := ReadURL(test.url)
			if err != nil {
				t.Fatalf("DownloadFile: %v", err)
			}
			s := string(b)
			for _, re := range test.res {
				if !re.MatchString(s) {
					t.Errorf("re=%v doesn't match: %s", re, s)
				}
			}
		})
	}
}
