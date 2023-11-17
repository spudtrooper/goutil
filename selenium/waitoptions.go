// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package selenium

import "fmt"

type WaitOption struct {
	f func(*waitOptionImpl)
	s string
}

func (o WaitOption) String() string { return o.s }

type WaitOptions interface {
	Limit() int
	HasLimit() bool
	Message() string
	HasMessage() bool
}

func WaitLimit(limit int) WaitOption {
	return WaitOption{func(opts *waitOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}, fmt.Sprintf("selenium.WaitLimit(int %+v)", limit)}
}
func WaitLimitFlag(limit *int) WaitOption {
	return WaitOption{func(opts *waitOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}, fmt.Sprintf("selenium.WaitLimit(int %+v)", limit)}
}

func WaitMessage(message string) WaitOption {
	return WaitOption{func(opts *waitOptionImpl) {
		opts.has_message = true
		opts.message = message
	}, fmt.Sprintf("selenium.WaitMessage(string %+v)", message)}
}
func WaitMessageFlag(message *string) WaitOption {
	return WaitOption{func(opts *waitOptionImpl) {
		if message == nil {
			return
		}
		opts.has_message = true
		opts.message = *message
	}, fmt.Sprintf("selenium.WaitMessage(string %+v)", message)}
}

type waitOptionImpl struct {
	limit       int
	has_limit   bool
	message     string
	has_message bool
}

func (w *waitOptionImpl) Limit() int       { return w.limit }
func (w *waitOptionImpl) HasLimit() bool   { return w.has_limit }
func (w *waitOptionImpl) Message() string  { return w.message }
func (w *waitOptionImpl) HasMessage() bool { return w.has_message }

func makeWaitOptionImpl(opts ...WaitOption) *waitOptionImpl {
	res := &waitOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeWaitOptions(opts ...WaitOption) WaitOptions {
	return makeWaitOptionImpl(opts...)
}
