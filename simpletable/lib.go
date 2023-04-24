// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package simpletable

import "fmt"

type NewOption struct {
	f func(*newOptionImpl)
	s string
}

func (o NewOption) String() string { return o.s }

type NewOptions interface {
	Header() []string
	HasHeader() bool
	NoBorder() bool
	HasNoBorder() bool
}

func NewHeader(header []string) NewOption {
	return NewOption{func(opts *newOptionImpl) {
		opts.has_header = true
		opts.header = header
	}, fmt.Sprintf("simpletable.NewHeader([]string %+v)", header)}
}
func NewHeaderFlag(header *[]string) NewOption {
	return NewOption{func(opts *newOptionImpl) {
		if header == nil {
			return
		}
		opts.has_header = true
		opts.header = *header
	}, fmt.Sprintf("simpletable.NewHeader([]string %+v)", header)}
}

func NewNoBorder(noBorder bool) NewOption {
	return NewOption{func(opts *newOptionImpl) {
		opts.has_noBorder = true
		opts.noBorder = noBorder
	}, fmt.Sprintf("simpletable.NewNoBorder(bool %+v)", noBorder)}
}
func NewNoBorderFlag(noBorder *bool) NewOption {
	return NewOption{func(opts *newOptionImpl) {
		if noBorder == nil {
			return
		}
		opts.has_noBorder = true
		opts.noBorder = *noBorder
	}, fmt.Sprintf("simpletable.NewNoBorder(bool %+v)", noBorder)}
}

type newOptionImpl struct {
	header       []string
	has_header   bool
	noBorder     bool
	has_noBorder bool
}

func (n *newOptionImpl) Header() []string  { return n.header }
func (n *newOptionImpl) HasHeader() bool   { return n.has_header }
func (n *newOptionImpl) NoBorder() bool    { return n.noBorder }
func (n *newOptionImpl) HasNoBorder() bool { return n.has_noBorder }

func makeNewOptionImpl(opts ...NewOption) *newOptionImpl {
	res := &newOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeNewOptions(opts ...NewOption) NewOptions {
	return makeNewOptionImpl(opts...)
}
