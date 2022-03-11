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

func Reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
