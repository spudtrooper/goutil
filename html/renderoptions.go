// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package html

import "fmt"

//go:generate genopts --prefix=Render --outfile=renderoptions.go "noFormat:bool"

type RenderOption struct {
	f func(*renderOptionImpl)
	s string
}

func (o RenderOption) String() string { return o.s }

type RenderOptions interface {
	NoFormat() bool
	HasNoFormat() bool
}

func RenderNoFormat(noFormat bool) RenderOption {
	return RenderOption{func(opts *renderOptionImpl) {
		opts.has_noFormat = true
		opts.noFormat = noFormat
	}, fmt.Sprintf("html.RenderNoFormat(bool %+v)", noFormat)}
}
func RenderNoFormatFlag(noFormat *bool) RenderOption {
	return RenderOption{func(opts *renderOptionImpl) {
		if noFormat == nil {
			return
		}
		opts.has_noFormat = true
		opts.noFormat = *noFormat
	}, fmt.Sprintf("html.RenderNoFormat(bool %+v)", noFormat)}
}

type renderOptionImpl struct {
	noFormat     bool
	has_noFormat bool
}

func (r *renderOptionImpl) NoFormat() bool    { return r.noFormat }
func (r *renderOptionImpl) HasNoFormat() bool { return r.has_noFormat }

func makeRenderOptionImpl(opts ...RenderOption) *renderOptionImpl {
	res := &renderOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeRenderOptions(opts ...RenderOption) RenderOptions {
	return makeRenderOptionImpl(opts...)
}
