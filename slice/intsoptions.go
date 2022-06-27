// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package slice

type IntsOption func(*intsOptionImpl)

type IntsOptions interface {
	Sep() string
	TrimSpace() bool
}

func IntsSep(sep string) IntsOption {
	return func(opts *intsOptionImpl) {
		opts.sep = sep
	}
}
func IntsSepFlag(sep *string) IntsOption {
	return func(opts *intsOptionImpl) {
		opts.sep = *sep
	}
}

func IntsTrimSpace(trimSpace bool) IntsOption {
	return func(opts *intsOptionImpl) {
		opts.trimSpace = trimSpace
	}
}
func IntsTrimSpaceFlag(trimSpace *bool) IntsOption {
	return func(opts *intsOptionImpl) {
		opts.trimSpace = *trimSpace
	}
}

type intsOptionImpl struct {
	sep       string
	trimSpace bool
}

func (i *intsOptionImpl) Sep() string     { return i.sep }
func (i *intsOptionImpl) TrimSpace() bool { return i.trimSpace }

func makeIntsOptionImpl(opts ...IntsOption) *intsOptionImpl {
	res := &intsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeIntsOptions(opts ...IntsOption) IntsOptions {
	return makeIntsOptionImpl(opts...)
}
