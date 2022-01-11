package hist

// genopts --opt_type=MakeHistogramOption --prefix=MakeHistogram --outfile=hist/makehistoptions.go 'sortAsc' 'sortDesc'

type MakeHistogramOption func(*makeHistogramOptionImpl)

type MakeHistogramOptions interface {
	SortAsc() bool
	SortDesc() bool
}

func MakeHistogramSortAsc(sortAsc bool) MakeHistogramOption {
	return func(opts *makeHistogramOptionImpl) {
		opts.sortAsc = sortAsc
	}
}

func MakeHistogramSortDesc(sortDesc bool) MakeHistogramOption {
	return func(opts *makeHistogramOptionImpl) {
		opts.sortDesc = sortDesc
	}
}

type makeHistogramOptionImpl struct {
	sortAsc  bool
	sortDesc bool
}

func (m *makeHistogramOptionImpl) SortAsc() bool  { return m.sortAsc }
func (m *makeHistogramOptionImpl) SortDesc() bool { return m.sortDesc }

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
