// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package fs

import (
	"fmt"
	"regexp"
)

//go:generate genopts --prefix=FindFiles --outfile=findfilesoptions.go "keep:*regexp.Regexp" "recursive:bool" "filter:FindFilesFilterFn"

type FindFilesOption struct {
	f func(*findFilesOptionImpl)
	s string
}

func (o FindFilesOption) String() string { return o.s }

type FindFilesOptions interface {
	Filter() FindFilesFilterFn
	HasFilter() bool
	Keep() *regexp.Regexp
	HasKeep() bool
	Recursive() bool
	HasRecursive() bool
}

func FindFilesFilter(filter FindFilesFilterFn) FindFilesOption {
	return FindFilesOption{func(opts *findFilesOptionImpl) {
		opts.has_filter = true
		opts.filter = filter
	}, fmt.Sprintf("fs.FindFilesFilter(FindFilesFilterFn %+v)", filter)}
}
func FindFilesFilterFlag(filter *FindFilesFilterFn) FindFilesOption {
	return FindFilesOption{func(opts *findFilesOptionImpl) {
		if filter == nil {
			return
		}
		opts.has_filter = true
		opts.filter = *filter
	}, fmt.Sprintf("fs.FindFilesFilter(FindFilesFilterFn %+v)", filter)}
}

func FindFilesKeep(keep *regexp.Regexp) FindFilesOption {
	return FindFilesOption{func(opts *findFilesOptionImpl) {
		opts.has_keep = true
		opts.keep = keep
	}, fmt.Sprintf("fs.FindFilesKeep(*regexp.Regexp %+v)", keep)}
}
func FindFilesKeepFlag(keep **regexp.Regexp) FindFilesOption {
	return FindFilesOption{func(opts *findFilesOptionImpl) {
		if keep == nil {
			return
		}
		opts.has_keep = true
		opts.keep = *keep
	}, fmt.Sprintf("fs.FindFilesKeep(*regexp.Regexp %+v)", keep)}
}

func FindFilesRecursive(recursive bool) FindFilesOption {
	return FindFilesOption{func(opts *findFilesOptionImpl) {
		opts.has_recursive = true
		opts.recursive = recursive
	}, fmt.Sprintf("fs.FindFilesRecursive(bool %+v)", recursive)}
}
func FindFilesRecursiveFlag(recursive *bool) FindFilesOption {
	return FindFilesOption{func(opts *findFilesOptionImpl) {
		if recursive == nil {
			return
		}
		opts.has_recursive = true
		opts.recursive = *recursive
	}, fmt.Sprintf("fs.FindFilesRecursive(bool %+v)", recursive)}
}

type findFilesOptionImpl struct {
	filter        FindFilesFilterFn
	has_filter    bool
	keep          *regexp.Regexp
	has_keep      bool
	recursive     bool
	has_recursive bool
}

func (f *findFilesOptionImpl) Filter() FindFilesFilterFn { return f.filter }
func (f *findFilesOptionImpl) HasFilter() bool           { return f.has_filter }
func (f *findFilesOptionImpl) Keep() *regexp.Regexp      { return f.keep }
func (f *findFilesOptionImpl) HasKeep() bool             { return f.has_keep }
func (f *findFilesOptionImpl) Recursive() bool           { return f.recursive }
func (f *findFilesOptionImpl) HasRecursive() bool        { return f.has_recursive }

func makeFindFilesOptionImpl(opts ...FindFilesOption) *findFilesOptionImpl {
	res := &findFilesOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeFindFilesOptions(opts ...FindFilesOption) FindFilesOptions {
	return makeFindFilesOptionImpl(opts...)
}
