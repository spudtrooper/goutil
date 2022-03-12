package slice

//go:generate genopts --prefix=Strings --outfile=slice/stringsoptions.go "trimSpace"

type StringsOption func(*stringsOptionImpl)

type StringsOptions interface {
	TrimSpace() bool
}

func StringsTrimSpace(trimSpace bool) StringsOption {
	return func(opts *stringsOptionImpl) {
		opts.trimSpace = trimSpace
	}
}
func StringsTrimSpaceFlag(trimSpace *bool) StringsOption {
	return func(opts *stringsOptionImpl) {
		opts.trimSpace = *trimSpace
	}
}

type stringsOptionImpl struct {
	trimSpace bool
}

func (s *stringsOptionImpl) TrimSpace() bool { return s.trimSpace }

func makeStringsOptionImpl(opts ...StringsOption) *stringsOptionImpl {
	res := &stringsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeStringsOptions(opts ...StringsOption) StringsOptions {
	return makeStringsOptionImpl(opts...)
}
