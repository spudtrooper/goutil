package io

//go:generate genopts --prefix=StringsFromFile --outfile=io/stringsfromfileoptions.go "skipEmpty:bool" "commentStart:string"

type StringsFromFileOption func(*stringsFromFileOptionImpl)

type StringsFromFileOptions interface {
	SkipEmpty() bool
	CommentStart() string
}

func StringsFromFileSkipEmpty(skipEmpty bool) StringsFromFileOption {
	return func(opts *stringsFromFileOptionImpl) {
		opts.skipEmpty = skipEmpty
	}
}
func StringsFromFileSkipEmptyFlag(skipEmpty *bool) StringsFromFileOption {
	return func(opts *stringsFromFileOptionImpl) {
		opts.skipEmpty = *skipEmpty
	}
}

func StringsFromFileCommentStart(commentStart string) StringsFromFileOption {
	return func(opts *stringsFromFileOptionImpl) {
		opts.commentStart = commentStart
	}
}
func StringsFromFileCommentStartFlag(commentStart *string) StringsFromFileOption {
	return func(opts *stringsFromFileOptionImpl) {
		opts.commentStart = *commentStart
	}
}

type stringsFromFileOptionImpl struct {
	skipEmpty    bool
	commentStart string
}

func (s *stringsFromFileOptionImpl) SkipEmpty() bool      { return s.skipEmpty }
func (s *stringsFromFileOptionImpl) CommentStart() string { return s.commentStart }

func makeStringsFromFileOptionImpl(opts ...StringsFromFileOption) *stringsFromFileOptionImpl {
	res := &stringsFromFileOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeStringsFromFileOptions(opts ...StringsFromFileOption) StringsFromFileOptions {
	return makeStringsFromFileOptionImpl(opts...)
}
