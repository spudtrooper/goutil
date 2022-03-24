package fs

import "regexp"

//go:generate genopts --prefix=FindFiles --outfile=findfilesoptions.go "keep:*regexp.Regexp" "recursive:bool" "filter:FindFilesFilterFn"

type FindFilesOption func(*findFilesOptionImpl)

type FindFilesOptions interface {
	Keep() *regexp.Regexp
	Recursive() bool
	Filter() FindFilesFilterFn
}

func FindFilesKeep(keep *regexp.Regexp) FindFilesOption {
	return func(opts *findFilesOptionImpl) {
		opts.keep = keep
	}
}
func FindFilesKeepFlag(keep **regexp.Regexp) FindFilesOption {
	return func(opts *findFilesOptionImpl) {
		opts.keep = *keep
	}
}

func FindFilesRecursive(recursive bool) FindFilesOption {
	return func(opts *findFilesOptionImpl) {
		opts.recursive = recursive
	}
}
func FindFilesRecursiveFlag(recursive *bool) FindFilesOption {
	return func(opts *findFilesOptionImpl) {
		opts.recursive = *recursive
	}
}

func FindFilesFilter(filter FindFilesFilterFn) FindFilesOption {
	return func(opts *findFilesOptionImpl) {
		opts.filter = filter
	}
}
func FindFilesFilterFlag(filter *FindFilesFilterFn) FindFilesOption {
	return func(opts *findFilesOptionImpl) {
		opts.filter = *filter
	}
}

type findFilesOptionImpl struct {
	keep      *regexp.Regexp
	recursive bool
	filter    FindFilesFilterFn
}

func (f *findFilesOptionImpl) Keep() *regexp.Regexp      { return f.keep }
func (f *findFilesOptionImpl) Recursive() bool           { return f.recursive }
func (f *findFilesOptionImpl) Filter() FindFilesFilterFn { return f.filter }

func makeFindFilesOptionImpl(opts ...FindFilesOption) *findFilesOptionImpl {
	res := &findFilesOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeFindFilesOptions(opts ...FindFilesOption) FindFilesOptions {
	return makeFindFilesOptionImpl(opts...)
}
