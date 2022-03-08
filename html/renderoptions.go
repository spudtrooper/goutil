package html

//go:generate genopts --opt_type=RenderOption --prefix=Render --outfile=renderoptions.go "noFormat:bool"

type RenderOption func(*renderOptionImpl)

type RenderOptions interface {
	NoFormat() bool
}

func RenderNoFormat(noFormat bool) RenderOption {
	return func(opts *renderOptionImpl) {
		opts.noFormat = noFormat
	}
}

type renderOptionImpl struct {
	noFormat bool
}

func (r *renderOptionImpl) NoFormat() bool { return r.noFormat }

func makeRenderOptionImpl(opts ...RenderOption) *renderOptionImpl {
	res := &renderOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRenderOptions(opts ...RenderOption) RenderOptions {
	return makeRenderOptionImpl(opts...)
}
