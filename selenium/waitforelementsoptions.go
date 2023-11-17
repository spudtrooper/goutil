// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package selenium

import "fmt"

type WaitForElementsOption struct {
	f func(*waitForElementsOptionImpl)
	s string
}

func (o WaitForElementsOption) String() string { return o.s }

type WaitForElementsOptions interface {
	Limit() int
	HasLimit() bool
	Message() string
	HasMessage() bool
}

func WaitForElementsLimit(limit int) WaitForElementsOption {
	return WaitForElementsOption{func(opts *waitForElementsOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}, fmt.Sprintf("selenium.WaitForElementsLimit(int %+v)", limit)}
}
func WaitForElementsLimitFlag(limit *int) WaitForElementsOption {
	return WaitForElementsOption{func(opts *waitForElementsOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}, fmt.Sprintf("selenium.WaitForElementsLimit(int %+v)", limit)}
}

func WaitForElementsMessage(message string) WaitForElementsOption {
	return WaitForElementsOption{func(opts *waitForElementsOptionImpl) {
		opts.has_message = true
		opts.message = message
	}, fmt.Sprintf("selenium.WaitForElementsMessage(string %+v)", message)}
}
func WaitForElementsMessageFlag(message *string) WaitForElementsOption {
	return WaitForElementsOption{func(opts *waitForElementsOptionImpl) {
		if message == nil {
			return
		}
		opts.has_message = true
		opts.message = *message
	}, fmt.Sprintf("selenium.WaitForElementsMessage(string %+v)", message)}
}

type waitForElementsOptionImpl struct {
	limit       int
	has_limit   bool
	message     string
	has_message bool
}

func (w *waitForElementsOptionImpl) Limit() int       { return w.limit }
func (w *waitForElementsOptionImpl) HasLimit() bool   { return w.has_limit }
func (w *waitForElementsOptionImpl) Message() string  { return w.message }
func (w *waitForElementsOptionImpl) HasMessage() bool { return w.has_message }

func makeWaitForElementsOptionImpl(opts ...WaitForElementsOption) *waitForElementsOptionImpl {
	res := &waitForElementsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeWaitForElementsOptions(opts ...WaitForElementsOption) WaitForElementsOptions {
	return makeWaitForElementsOptionImpl(opts...)
}
