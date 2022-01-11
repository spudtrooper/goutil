package hist

// genopts --opt_type=MakeHistogramOption --prefix=MakeHistogram --outfile=hist/makehistoptions.go 'sorted'

type MakeHistogramOption func(*makeHistogramOptionImpl)

type MakeHistogramOptions interface {
	Sorted() bool
}

func MakeHistogramSorted(sorted bool) MakeHistogramOption {
	return func(opts *makeHistogramOptionImpl) {
		opts.sorted = sorted
	}
}

type makeHistogramOptionImpl struct {
	sorted bool
}

func (m *makeHistogramOptionImpl) Sorted() bool { return m.sorted }

func makeMakeHistogramOptionImpl(opts ...MakeHistogramOption) *makeHistogramOptionImpl {
	res := &makeHistogramOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeMakeHistogramOptions(opts ...MakeHistogramOption) MakeHistogramOptions {
	return makeMakeHistogramOptionImpl(opts...)
}
