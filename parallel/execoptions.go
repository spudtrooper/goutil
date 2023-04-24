// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package parallel

import "fmt"

//go:generate genopts --prefix=Exec --outfile=execoptions.go "exitOnError:bool"

type ExecOption struct {
	f func(*execOptionImpl)
	s string
}

func (o ExecOption) String() string { return o.s }

type ExecOptions interface {
	ExitOnError() bool
	HasExitOnError() bool
}

func ExecExitOnError(exitOnError bool) ExecOption {
	return ExecOption{func(opts *execOptionImpl) {
		opts.has_exitOnError = true
		opts.exitOnError = exitOnError
	}, fmt.Sprintf("parallel.ExecExitOnError(bool %+v)", exitOnError)}
}
func ExecExitOnErrorFlag(exitOnError *bool) ExecOption {
	return ExecOption{func(opts *execOptionImpl) {
		if exitOnError == nil {
			return
		}
		opts.has_exitOnError = true
		opts.exitOnError = *exitOnError
	}, fmt.Sprintf("parallel.ExecExitOnError(bool %+v)", exitOnError)}
}

type execOptionImpl struct {
	exitOnError     bool
	has_exitOnError bool
}

func (e *execOptionImpl) ExitOnError() bool    { return e.exitOnError }
func (e *execOptionImpl) HasExitOnError() bool { return e.has_exitOnError }

func makeExecOptionImpl(opts ...ExecOption) *execOptionImpl {
	res := &execOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeExecOptions(opts ...ExecOption) ExecOptions {
	return makeExecOptionImpl(opts...)
}
