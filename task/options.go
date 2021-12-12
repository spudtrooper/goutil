package task

import "github.com/fatih/color"

// genopts --outfile=task/options.go 'printDone' 'color:*color.Color'

type Option func(*optionImpl)

type Options interface {
	PrintDone() bool
	Color() *color.Color
}

func PrintDone(printDone bool) Option {
	return func(opts *optionImpl) {
		opts.printDone = printDone
	}
}

func Color(color *color.Color) Option {
	return func(opts *optionImpl) {
		opts.color = color
	}
}

type optionImpl struct {
	printDone bool
	color     *color.Color
}

func (o *optionImpl) PrintDone() bool     { return o.printDone }
func (o *optionImpl) Color() *color.Color { return o.color }

func makeOptionImpl(opts ...Option) *optionImpl {
	res := &optionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeOptions(opts ...Option) Options {
	return makeOptionImpl(opts...)
}
