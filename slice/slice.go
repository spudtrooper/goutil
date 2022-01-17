package slice

import "strings"

func Strings(input, sep string, sOpts ...StringsOption) []string {
	opts := MakeStringsOptions(sOpts...)
	if input == "" {
		return []string{}
	}
	var res []string
	for _, s := range strings.Split(input, sep) {
		if opts.TrimSpace() {
			s = strings.TrimSpace(s)
		}
		res = append(res, s)
	}
	return res
}
