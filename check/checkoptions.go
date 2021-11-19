package check

// ~/go/bin/genopts -opt_type=CheckOption --prefix=Check  message:string

type CheckOption func(*checkOptionImpl)

type CheckOptions interface {
	Message() string
}

func CheckMessage(message string) CheckOption {
	return func(opts *checkOptionImpl) {
		opts.message = message
	}
}

type checkOptionImpl struct {
	message string
}

func (c *checkOptionImpl) Message() string { return c.message }

func makeCheckOptionImpl(opts ...CheckOption) checkOptionImpl {
	var res checkOptionImpl
	for _, opt := range opts {
		opt(&res)
	}
	return res
}
