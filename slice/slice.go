// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package slice

import "fmt"

type StringsOption struct {
	f func(*stringsOptionImpl)
	s string
}

func (o StringsOption) String() string { return o.s }

type StringsOptions interface {
	TrimSpace() bool
	HasTrimSpace() bool
}

func StringsTrimSpace(trimSpace bool) StringsOption {
	return StringsOption{func(opts *stringsOptionImpl) {
		opts.has_trimSpace = true
		opts.trimSpace = trimSpace
	}, fmt.Sprintf("slice.StringsTrimSpace(bool %+v)", trimSpace)}
}
func StringsTrimSpaceFlag(trimSpace *bool) StringsOption {
	return StringsOption{func(opts *stringsOptionImpl) {
		if trimSpace == nil {
			return
		}
		opts.has_trimSpace = true
		opts.trimSpace = *trimSpace
	}, fmt.Sprintf("slice.StringsTrimSpace(bool %+v)", trimSpace)}
}

type stringsOptionImpl struct {
	trimSpace     bool
	has_trimSpace bool
}

func (s *stringsOptionImpl) TrimSpace() bool    { return s.trimSpace }
func (s *stringsOptionImpl) HasTrimSpace() bool { return s.has_trimSpace }

func makeStringsOptionImpl(opts ...StringsOption) *stringsOptionImpl {
	res := &stringsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeStringsOptions(opts ...StringsOption) StringsOptions {
	return makeStringsOptionImpl(opts...)
}
