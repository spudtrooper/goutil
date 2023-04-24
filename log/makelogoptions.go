// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package log

import "fmt"

//go:generate genopts --prefix=MakeLog --outfile=makelogoptions.go "color"

type MakeLogOption struct {
	f func(*makeLogOptionImpl)
	s string
}

func (o MakeLogOption) String() string { return o.s }

type MakeLogOptions interface {
	Color() bool
	HasColor() bool
}

func MakeLogColor(color bool) MakeLogOption {
	return MakeLogOption{func(opts *makeLogOptionImpl) {
		opts.has_color = true
		opts.color = color
	}, fmt.Sprintf("log.MakeLogColor(bool %+v)", color)}
}
func MakeLogColorFlag(color *bool) MakeLogOption {
	return MakeLogOption{func(opts *makeLogOptionImpl) {
		if color == nil {
			return
		}
		opts.has_color = true
		opts.color = *color
	}, fmt.Sprintf("log.MakeLogColor(bool %+v)", color)}
}

type makeLogOptionImpl struct {
	color     bool
	has_color bool
}

func (m *makeLogOptionImpl) Color() bool    { return m.color }
func (m *makeLogOptionImpl) HasColor() bool { return m.has_color }

func makeMakeLogOptionImpl(opts ...MakeLogOption) *makeLogOptionImpl {
	res := &makeLogOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeMakeLogOptions(opts ...MakeLogOption) MakeLogOptions {
	return makeMakeLogOptionImpl(opts...)
}
