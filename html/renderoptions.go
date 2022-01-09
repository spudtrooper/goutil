package html

// genopts --opt_type=RenderOption --prefix=Render --outfile=html/renderoptions.go 'noformat:bool'

type RenderOption func(*renderOptionImpl)

type RenderOptions interface {
	Noformat() bool
}

func RenderNoformat(noformat bool) RenderOption {
	return func(opts *renderOptionImpl) {
		opts.noformat = noformat
	}
}

type renderOptionImpl struct {
	noformat bool
}

func (r *renderOptionImpl) Noformat() bool { return r.noformat }

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
