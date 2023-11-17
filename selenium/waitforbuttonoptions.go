// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package selenium

import "fmt"

type WaitForButtonOption struct {
	f func(*waitForButtonOptionImpl)
	s string
}

func (o WaitForButtonOption) String() string { return o.s }

type WaitForButtonOptions interface {
	Limit() int
	HasLimit() bool
	Message() string
	HasMessage() bool
	ToWaitOptions() []WaitOption
}

func WaitForButtonLimit(limit int) WaitForButtonOption {
	return WaitForButtonOption{func(opts *waitForButtonOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}, fmt.Sprintf("selenium.WaitForButtonLimit(int %+v)", limit)}
}
func WaitForButtonLimitFlag(limit *int) WaitForButtonOption {
	return WaitForButtonOption{func(opts *waitForButtonOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}, fmt.Sprintf("selenium.WaitForButtonLimit(int %+v)", limit)}
}

func WaitForButtonMessage(message string) WaitForButtonOption {
	return WaitForButtonOption{func(opts *waitForButtonOptionImpl) {
		opts.has_message = true
		opts.message = message
	}, fmt.Sprintf("selenium.WaitForButtonMessage(string %+v)", message)}
}
func WaitForButtonMessageFlag(message *string) WaitForButtonOption {
	return WaitForButtonOption{func(opts *waitForButtonOptionImpl) {
		if message == nil {
			return
		}
		opts.has_message = true
		opts.message = *message
	}, fmt.Sprintf("selenium.WaitForButtonMessage(string %+v)", message)}
}

type waitForButtonOptionImpl struct {
	limit       int
	has_limit   bool
	message     string
	has_message bool
}

func (w *waitForButtonOptionImpl) Limit() int       { return w.limit }
func (w *waitForButtonOptionImpl) HasLimit() bool   { return w.has_limit }
func (w *waitForButtonOptionImpl) Message() string  { return w.message }
func (w *waitForButtonOptionImpl) HasMessage() bool { return w.has_message }

// ToWaitOptions converts WaitForButtonOption to an array of WaitOption
func (o *waitForButtonOptionImpl) ToWaitOptions() []WaitOption {
	return []WaitOption{
		WaitLimit(o.Limit()),
		WaitMessage(o.Message()),
	}
}

func makeWaitForButtonOptionImpl(opts ...WaitForButtonOption) *waitForButtonOptionImpl {
	res := &waitForButtonOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeWaitForButtonOptions(opts ...WaitForButtonOption) WaitForButtonOptions {
	return makeWaitForButtonOptionImpl(opts...)
}
