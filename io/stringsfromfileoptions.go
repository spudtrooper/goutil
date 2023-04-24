// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package io

import "fmt"

//go:generate genopts --prefix=StringsFromFile --outfile=stringsfromfileoptions.go "skipEmpty:bool" "commentStart:string"

type StringsFromFileOption struct {
	f func(*stringsFromFileOptionImpl)
	s string
}

func (o StringsFromFileOption) String() string { return o.s }

type StringsFromFileOptions interface {
	CommentStart() string
	HasCommentStart() bool
	SkipEmpty() bool
	HasSkipEmpty() bool
}

func StringsFromFileCommentStart(commentStart string) StringsFromFileOption {
	return StringsFromFileOption{func(opts *stringsFromFileOptionImpl) {
		opts.has_commentStart = true
		opts.commentStart = commentStart
	}, fmt.Sprintf("io.StringsFromFileCommentStart(string %+v)", commentStart)}
}
func StringsFromFileCommentStartFlag(commentStart *string) StringsFromFileOption {
	return StringsFromFileOption{func(opts *stringsFromFileOptionImpl) {
		if commentStart == nil {
			return
		}
		opts.has_commentStart = true
		opts.commentStart = *commentStart
	}, fmt.Sprintf("io.StringsFromFileCommentStart(string %+v)", commentStart)}
}

func StringsFromFileSkipEmpty(skipEmpty bool) StringsFromFileOption {
	return StringsFromFileOption{func(opts *stringsFromFileOptionImpl) {
		opts.has_skipEmpty = true
		opts.skipEmpty = skipEmpty
	}, fmt.Sprintf("io.StringsFromFileSkipEmpty(bool %+v)", skipEmpty)}
}
func StringsFromFileSkipEmptyFlag(skipEmpty *bool) StringsFromFileOption {
	return StringsFromFileOption{func(opts *stringsFromFileOptionImpl) {
		if skipEmpty == nil {
			return
		}
		opts.has_skipEmpty = true
		opts.skipEmpty = *skipEmpty
	}, fmt.Sprintf("io.StringsFromFileSkipEmpty(bool %+v)", skipEmpty)}
}

type stringsFromFileOptionImpl struct {
	commentStart     string
	has_commentStart bool
	skipEmpty        bool
	has_skipEmpty    bool
}

func (s *stringsFromFileOptionImpl) CommentStart() string  { return s.commentStart }
func (s *stringsFromFileOptionImpl) HasCommentStart() bool { return s.has_commentStart }
func (s *stringsFromFileOptionImpl) SkipEmpty() bool       { return s.skipEmpty }
func (s *stringsFromFileOptionImpl) HasSkipEmpty() bool    { return s.has_skipEmpty }

func makeStringsFromFileOptionImpl(opts ...StringsFromFileOption) *stringsFromFileOptionImpl {
	res := &stringsFromFileOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeStringsFromFileOptions(opts ...StringsFromFileOption) StringsFromFileOptions {
	return makeStringsFromFileOptionImpl(opts...)
}
