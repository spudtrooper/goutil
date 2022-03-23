package parallel

//go:generate genopts --prefix=Exec --outfile=execoptions.go "exitOnError:bool"

type ExecOption func(*execOptionImpl)

type ExecOptions interface {
	ExitOnError() bool
}

func ExecExitOnError(exitOnError bool) ExecOption {
	return func(opts *execOptionImpl) {
		opts.exitOnError = exitOnError
	}
}
func ExecExitOnErrorFlag(exitOnError *bool) ExecOption {
	return func(opts *execOptionImpl) {
		opts.exitOnError = *exitOnError
	}
}

type execOptionImpl struct {
	exitOnError bool
}

func (e *execOptionImpl) ExitOnError() bool { return e.exitOnError }

func makeExecOptionImpl(opts ...ExecOption) *execOptionImpl {
	res := &execOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeExecOptions(opts ...ExecOption) ExecOptions {
	return makeExecOptionImpl(opts...)
}
