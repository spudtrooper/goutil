package slice

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/sets"
)

//go:generate genopts --function Strings trimSpace
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

// StringDiff returns an array with all elements in `a` not in `b`.
func StringDiff(a, b []string) []string {
	res := []string{}
	bSet := sets.String(b)
	for _, x := range a {
		if !bSet[x] {
			res = append(res, x)
		}
	}
	return res
}

func Reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func NonEmptyStrings(ss []string) []string {
	res := []string{}
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}

//go:generate genopts --function Ints sep:string trimSpace
func Ints(s string, optss ...IntsOption) ([]int, error) {
	opts := MakeIntsOptions(optss...)
	sep := or.String(opts.Sep(), ",")
	res := []int{}
	for _, s := range NonEmptyStrings(strings.Split(s, sep)) {
		if opts.TrimSpace() {
			s = strings.TrimSpace(s)
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Errorf("strconv.Atoi(%q): %v", s, err)
		}
		res = append(res, i)
	}
	return res, nil
}
