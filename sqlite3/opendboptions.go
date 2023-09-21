// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package sqlite3

import "fmt"

type OpenDBOption struct {
	f func(*openDBOptionImpl)
	s string
}

func (o OpenDBOption) String() string { return o.s }

type OpenDBOptions interface {
	CreateDBIfNotExists() bool
	HasCreateDBIfNotExists() bool
}

func OpenDBCreateDBIfNotExists(createDBIfNotExists bool) OpenDBOption {
	return OpenDBOption{func(opts *openDBOptionImpl) {
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = createDBIfNotExists
	}, fmt.Sprintf("sqlite3.OpenDBCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}
func OpenDBCreateDBIfNotExistsFlag(createDBIfNotExists *bool) OpenDBOption {
	return OpenDBOption{func(opts *openDBOptionImpl) {
		if createDBIfNotExists == nil {
			return
		}
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = *createDBIfNotExists
	}, fmt.Sprintf("sqlite3.OpenDBCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}

type openDBOptionImpl struct {
	createDBIfNotExists     bool
	has_createDBIfNotExists bool
}

func (o *openDBOptionImpl) CreateDBIfNotExists() bool    { return o.createDBIfNotExists }
func (o *openDBOptionImpl) HasCreateDBIfNotExists() bool { return o.has_createDBIfNotExists }

func makeOpenDBOptionImpl(opts ...OpenDBOption) *openDBOptionImpl {
	res := &openDBOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeOpenDBOptions(opts ...OpenDBOption) OpenDBOptions {
	return makeOpenDBOptionImpl(opts...)
}
