// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package task

import (
	"fmt"

	"github.com/fatih/color"
)

//go:generate genopts --outfile=options.go "printDone" "color:*color.Color"

type Option struct {
	f func(*optionImpl)
	s string
}

func (o Option) String() string { return o.s }

type Options interface {
	Color() *color.Color
	HasColor() bool
	PrintDone() bool
	HasPrintDone() bool
}

func Color(color *color.Color) Option {
	return Option{func(opts *optionImpl) {
		opts.has_color = true
		opts.color = color
	}, fmt.Sprintf("task.Color(*color.Color %+v)", color)}
}
func ColorFlag(color **color.Color) Option {
	return Option{func(opts *optionImpl) {
		if color == nil {
			return
		}
		opts.has_color = true
		opts.color = *color
	}, fmt.Sprintf("task.Color(*color.Color %+v)", color)}
}

func PrintDone(printDone bool) Option {
	return Option{func(opts *optionImpl) {
		opts.has_printDone = true
		opts.printDone = printDone
	}, fmt.Sprintf("task.PrintDone(bool %+v)", printDone)}
}
func PrintDoneFlag(printDone *bool) Option {
	return Option{func(opts *optionImpl) {
		if printDone == nil {
			return
		}
		opts.has_printDone = true
		opts.printDone = *printDone
	}, fmt.Sprintf("task.PrintDone(bool %+v)", printDone)}
}

type optionImpl struct {
	color         *color.Color
	has_color     bool
	printDone     bool
	has_printDone bool
}

func (o *optionImpl) Color() *color.Color { return o.color }
func (o *optionImpl) HasColor() bool      { return o.has_color }
func (o *optionImpl) PrintDone() bool     { return o.printDone }
func (o *optionImpl) HasPrintDone() bool  { return o.has_printDone }

func makeOptionImpl(opts ...Option) *optionImpl {
	res := &optionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeOptions(opts ...Option) Options {
	return makeOptionImpl(opts...)
}
