// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package sqlite3

import "fmt"

type DropTableIfExistsOption struct {
	f func(*dropTableIfExistsOptionImpl)
	s string
}

func (o DropTableIfExistsOption) String() string { return o.s }

type DropTableIfExistsOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func DropTableIfExistsVerbose(verbose bool) DropTableIfExistsOption {
	return DropTableIfExistsOption{func(opts *dropTableIfExistsOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("sqlite3.DropTableIfExistsVerbose(bool %+v)", verbose)}
}
func DropTableIfExistsVerboseFlag(verbose *bool) DropTableIfExistsOption {
	return DropTableIfExistsOption{func(opts *dropTableIfExistsOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("sqlite3.DropTableIfExistsVerbose(bool %+v)", verbose)}
}

type dropTableIfExistsOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (d *dropTableIfExistsOptionImpl) Verbose() bool    { return d.verbose }
func (d *dropTableIfExistsOptionImpl) HasVerbose() bool { return d.has_verbose }

func makeDropTableIfExistsOptionImpl(opts ...DropTableIfExistsOption) *dropTableIfExistsOptionImpl {
	res := &dropTableIfExistsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeDropTableIfExistsOptions(opts ...DropTableIfExistsOption) DropTableIfExistsOptions {
	return makeDropTableIfExistsOptionImpl(opts...)
}
