// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package simpletable

type NewOption func(*newOptionImpl)

type NewOptions interface {
	NoBorder() bool
	Header() []string
}

func NewNoBorder(noBorder bool) NewOption {
	return func(opts *newOptionImpl) {
		opts.noBorder = noBorder
	}
}
func NewNoBorderFlag(noBorder *bool) NewOption {
	return func(opts *newOptionImpl) {
		opts.noBorder = *noBorder
	}
}

func NewHeader(header []string) NewOption {
	return func(opts *newOptionImpl) {
		opts.header = header
	}
}
func NewHeaderFlag(header *[]string) NewOption {
	return func(opts *newOptionImpl) {
		opts.header = *header
	}
}

type newOptionImpl struct {
	noBorder bool
	header   []string
}

func (n *newOptionImpl) NoBorder() bool   { return n.noBorder }
func (n *newOptionImpl) Header() []string { return n.header }

func makeNewOptionImpl(opts ...NewOption) *newOptionImpl {
	res := &newOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeNewOptions(opts ...NewOption) NewOptions {
	return makeNewOptionImpl(opts...)
}
