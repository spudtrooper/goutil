// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package hist

import "fmt"

//go:generate genopts --prefix=MakeHistogram --outfile=makehistoptions.go "sortAsc" "sortDesc"

type MakeHistogramOption struct {
	f func(*makeHistogramOptionImpl)
	s string
}

func (o MakeHistogramOption) String() string { return o.s }

type MakeHistogramOptions interface {
	SortAsc() bool
	HasSortAsc() bool
	SortDesc() bool
	HasSortDesc() bool
}

func MakeHistogramSortAsc(sortAsc bool) MakeHistogramOption {
	return MakeHistogramOption{func(opts *makeHistogramOptionImpl) {
		opts.has_sortAsc = true
		opts.sortAsc = sortAsc
	}, fmt.Sprintf("hist.MakeHistogramSortAsc(bool %+v)", sortAsc)}
}
func MakeHistogramSortAscFlag(sortAsc *bool) MakeHistogramOption {
	return MakeHistogramOption{func(opts *makeHistogramOptionImpl) {
		if sortAsc == nil {
			return
		}
		opts.has_sortAsc = true
		opts.sortAsc = *sortAsc
	}, fmt.Sprintf("hist.MakeHistogramSortAsc(bool %+v)", sortAsc)}
}

func MakeHistogramSortDesc(sortDesc bool) MakeHistogramOption {
	return MakeHistogramOption{func(opts *makeHistogramOptionImpl) {
		opts.has_sortDesc = true
		opts.sortDesc = sortDesc
	}, fmt.Sprintf("hist.MakeHistogramSortDesc(bool %+v)", sortDesc)}
}
func MakeHistogramSortDescFlag(sortDesc *bool) MakeHistogramOption {
	return MakeHistogramOption{func(opts *makeHistogramOptionImpl) {
		if sortDesc == nil {
			return
		}
		opts.has_sortDesc = true
		opts.sortDesc = *sortDesc
	}, fmt.Sprintf("hist.MakeHistogramSortDesc(bool %+v)", sortDesc)}
}

type makeHistogramOptionImpl struct {
	sortAsc      bool
	has_sortAsc  bool
	sortDesc     bool
	has_sortDesc bool
}

func (m *makeHistogramOptionImpl) SortAsc() bool     { return m.sortAsc }
func (m *makeHistogramOptionImpl) HasSortAsc() bool  { return m.has_sortAsc }
func (m *makeHistogramOptionImpl) SortDesc() bool    { return m.sortDesc }
func (m *makeHistogramOptionImpl) HasSortDesc() bool { return m.has_sortDesc }

func makeMakeHistogramOptionImpl(opts ...MakeHistogramOption) *makeHistogramOptionImpl {
	res := &makeHistogramOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeMakeHistogramOptions(opts ...MakeHistogramOption) MakeHistogramOptions {
	return makeMakeHistogramOptionImpl(opts...)
}
