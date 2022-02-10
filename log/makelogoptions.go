package log

// genopts --opt_type=MakeLogOption --prefix=MakeLog --outfile=log/makelogoptions.go 'color'

type MakeLogOption func(*makeLogOptionImpl)

type MakeLogOptions interface {
	Color() bool
}

func MakeLogColor(color bool) MakeLogOption {
	return func(opts *makeLogOptionImpl) {
		opts.color = color
	}
}

type makeLogOptionImpl struct {
	color bool
}

func (m *makeLogOptionImpl) Color() bool { return m.color }

func makeMakeLogOptionImpl(opts ...MakeLogOption) *makeLogOptionImpl {
	res := &makeLogOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeMakeLogOptions(opts ...MakeLogOption) MakeLogOptions {
	return makeMakeLogOptionImpl(opts...)
}
