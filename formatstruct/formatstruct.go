package formatstruct

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	lineRE = regexp.MustCompile(`\s+"([^"]+)"\: (.*),?`)
)

type Option interface {
	Transform(val string) string
	Key() string
}

type keyTransform struct {
	key string
	fn  func(string) string
}

func (k *keyTransform) Key() string {
	return k.key
}

func (k *keyTransform) Transform(val string) string {
	return k.fn(val)
}

func KeyTransform(key string, fn func(string) string) Option {
	return &keyTransform{key: key, fn: fn}
}

func Format(o interface{}, opts ...Option) ([]string, error) {
	/*
	 "ForceFollow": false,
	 "CacheDir": ".instagram_cache_dir",
	 "BadUserTries": 3,
	 "MaxLikesPerUser": 0,
	 "InstagramExportFile": ".instagram_export",
	 "Headless": false
	*/
	j, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		return nil, err
	}
	var longestKey int
	vals := map[string]string{}
	for _, line := range strings.Split(string(j), "\n") {
		m := lineRE.FindStringSubmatch(line)
		if len(m) == 3 {
			key, val := m[1], m[2]
			key = strings.ReplaceAll(key, `"`, "")
			val = strings.ReplaceAll(val, `"`, "")
			val = strings.ReplaceAll(val, `,`, "")
			for _, opt := range opts {
				if opt.Key() == key {
					val = opt.Transform(val)
				}
			}
			if longestKey < len(key) {
				longestKey = len(key)
			}
			vals[key] = val
		}
	}
	tmpl := fmt.Sprintf("%%-%ds: %%s", longestKey)
	var keys []string
	for k := range vals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	res := []string{}
	for _, k := range keys {
		v := vals[k]
		s := fmt.Sprintf(tmpl, k, v)
		res = append(res, s)
	}
	return res, nil
}
