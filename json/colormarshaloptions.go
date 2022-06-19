// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package json

type ColorMarshalOption func(*colorMarshalOptionImpl)

type ColorMarshalOptions interface {
	Indent() int
}

func ColorMarshalIndent(indent int) ColorMarshalOption {
	return func(opts *colorMarshalOptionImpl) {
		opts.indent = indent
	}
}
func ColorMarshalIndentFlag(indent *int) ColorMarshalOption {
	return func(opts *colorMarshalOptionImpl) {
		opts.indent = *indent
	}
}

type colorMarshalOptionImpl struct {
	indent int
}

func (c *colorMarshalOptionImpl) Indent() int { return c.indent }

func makeColorMarshalOptionImpl(opts ...ColorMarshalOption) *colorMarshalOptionImpl {
	res := &colorMarshalOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeColorMarshalOptions(opts ...ColorMarshalOption) ColorMarshalOptions {
	return makeColorMarshalOptionImpl(opts...)
}
