// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package net

type ReadURLChanOption func(*readURLChanOptionImpl)

type ReadURLChanOptions interface {
	SkipEmpty() bool
	CommentStart() string
}

func ReadURLChanSkipEmpty(skipEmpty bool) ReadURLChanOption {
	return func(opts *readURLChanOptionImpl) {
		opts.skipEmpty = skipEmpty
	}
}
func ReadURLChanSkipEmptyFlag(skipEmpty *bool) ReadURLChanOption {
	return func(opts *readURLChanOptionImpl) {
		opts.skipEmpty = *skipEmpty
	}
}

func ReadURLChanCommentStart(commentStart string) ReadURLChanOption {
	return func(opts *readURLChanOptionImpl) {
		opts.commentStart = commentStart
	}
}
func ReadURLChanCommentStartFlag(commentStart *string) ReadURLChanOption {
	return func(opts *readURLChanOptionImpl) {
		opts.commentStart = *commentStart
	}
}

type readURLChanOptionImpl struct {
	skipEmpty    bool
	commentStart string
}

func (r *readURLChanOptionImpl) SkipEmpty() bool      { return r.skipEmpty }
func (r *readURLChanOptionImpl) CommentStart() string { return r.commentStart }

func makeReadURLChanOptionImpl(opts ...ReadURLChanOption) *readURLChanOptionImpl {
	res := &readURLChanOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeReadURLChanOptions(opts ...ReadURLChanOption) ReadURLChanOptions {
	return makeReadURLChanOptionImpl(opts...)
}
