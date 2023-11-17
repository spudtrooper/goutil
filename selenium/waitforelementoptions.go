// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package selenium

import "fmt"

type WaitForElementOption struct {
	f func(*waitForElementOptionImpl)
	s string
}

func (o WaitForElementOption) String() string { return o.s }

type WaitForElementOptions interface {
	Limit() int
	HasLimit() bool
	Message() string
	HasMessage() bool
}

func WaitForElementLimit(limit int) WaitForElementOption {
	return WaitForElementOption{func(opts *waitForElementOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}, fmt.Sprintf("selenium.WaitForElementLimit(int %+v)", limit)}
}
func WaitForElementLimitFlag(limit *int) WaitForElementOption {
	return WaitForElementOption{func(opts *waitForElementOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}, fmt.Sprintf("selenium.WaitForElementLimit(int %+v)", limit)}
}

func WaitForElementMessage(message string) WaitForElementOption {
	return WaitForElementOption{func(opts *waitForElementOptionImpl) {
		opts.has_message = true
		opts.message = message
	}, fmt.Sprintf("selenium.WaitForElementMessage(string %+v)", message)}
}
func WaitForElementMessageFlag(message *string) WaitForElementOption {
	return WaitForElementOption{func(opts *waitForElementOptionImpl) {
		if message == nil {
			return
		}
		opts.has_message = true
		opts.message = *message
	}, fmt.Sprintf("selenium.WaitForElementMessage(string %+v)", message)}
}

type waitForElementOptionImpl struct {
	limit       int
	has_limit   bool
	message     string
	has_message bool
}

func (w *waitForElementOptionImpl) Limit() int       { return w.limit }
func (w *waitForElementOptionImpl) HasLimit() bool   { return w.has_limit }
func (w *waitForElementOptionImpl) Message() string  { return w.message }
func (w *waitForElementOptionImpl) HasMessage() bool { return w.has_message }

func makeWaitForElementOptionImpl(opts ...WaitForElementOption) *waitForElementOptionImpl {
	res := &waitForElementOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeWaitForElementOptions(opts ...WaitForElementOption) WaitForElementOptions {
	return makeWaitForElementOptionImpl(opts...)
}
