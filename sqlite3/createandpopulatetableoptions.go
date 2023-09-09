// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package sqlite3

import "fmt"

type CreateAndPopulateTableOption struct {
	f func(*createAndPopulateTableOptionImpl)
	s string
}

func (o CreateAndPopulateTableOption) String() string { return o.s }

type CreateAndPopulateTableOptions interface {
	CreateDBIfNotExists() bool
	HasCreateDBIfNotExists() bool
	DropIfExists() bool
	HasDropIfExists() bool
	PrimaryKey() string
	HasPrimaryKey() bool
}

func CreateAndPopulateTableCreateDBIfNotExists(createDBIfNotExists bool) CreateAndPopulateTableOption {
	return CreateAndPopulateTableOption{func(opts *createAndPopulateTableOptionImpl) {
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = createDBIfNotExists
	}, fmt.Sprintf("sqlite3.CreateAndPopulateTableCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}
func CreateAndPopulateTableCreateDBIfNotExistsFlag(createDBIfNotExists *bool) CreateAndPopulateTableOption {
	return CreateAndPopulateTableOption{func(opts *createAndPopulateTableOptionImpl) {
		if createDBIfNotExists == nil {
			return
		}
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = *createDBIfNotExists
	}, fmt.Sprintf("sqlite3.CreateAndPopulateTableCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}

func CreateAndPopulateTableDropIfExists(dropIfExists bool) CreateAndPopulateTableOption {
	return CreateAndPopulateTableOption{func(opts *createAndPopulateTableOptionImpl) {
		opts.has_dropIfExists = true
		opts.dropIfExists = dropIfExists
	}, fmt.Sprintf("sqlite3.CreateAndPopulateTableDropIfExists(bool %+v)", dropIfExists)}
}
func CreateAndPopulateTableDropIfExistsFlag(dropIfExists *bool) CreateAndPopulateTableOption {
	return CreateAndPopulateTableOption{func(opts *createAndPopulateTableOptionImpl) {
		if dropIfExists == nil {
			return
		}
		opts.has_dropIfExists = true
		opts.dropIfExists = *dropIfExists
	}, fmt.Sprintf("sqlite3.CreateAndPopulateTableDropIfExists(bool %+v)", dropIfExists)}
}

func CreateAndPopulateTablePrimaryKey(primaryKey string) CreateAndPopulateTableOption {
	return CreateAndPopulateTableOption{func(opts *createAndPopulateTableOptionImpl) {
		opts.has_primaryKey = true
		opts.primaryKey = primaryKey
	}, fmt.Sprintf("sqlite3.CreateAndPopulateTablePrimaryKey(string %+v)", primaryKey)}
}
func CreateAndPopulateTablePrimaryKeyFlag(primaryKey *string) CreateAndPopulateTableOption {
	return CreateAndPopulateTableOption{func(opts *createAndPopulateTableOptionImpl) {
		if primaryKey == nil {
			return
		}
		opts.has_primaryKey = true
		opts.primaryKey = *primaryKey
	}, fmt.Sprintf("sqlite3.CreateAndPopulateTablePrimaryKey(string %+v)", primaryKey)}
}

type createAndPopulateTableOptionImpl struct {
	createDBIfNotExists     bool
	has_createDBIfNotExists bool
	dropIfExists            bool
	has_dropIfExists        bool
	primaryKey              string
	has_primaryKey          bool
}

func (c *createAndPopulateTableOptionImpl) CreateDBIfNotExists() bool { return c.createDBIfNotExists }
func (c *createAndPopulateTableOptionImpl) HasCreateDBIfNotExists() bool {
	return c.has_createDBIfNotExists
}
func (c *createAndPopulateTableOptionImpl) DropIfExists() bool    { return c.dropIfExists }
func (c *createAndPopulateTableOptionImpl) HasDropIfExists() bool { return c.has_dropIfExists }
func (c *createAndPopulateTableOptionImpl) PrimaryKey() string    { return c.primaryKey }
func (c *createAndPopulateTableOptionImpl) HasPrimaryKey() bool   { return c.has_primaryKey }

func makeCreateAndPopulateTableOptionImpl(opts ...CreateAndPopulateTableOption) *createAndPopulateTableOptionImpl {
	res := &createAndPopulateTableOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeCreateAndPopulateTableOptions(opts ...CreateAndPopulateTableOption) CreateAndPopulateTableOptions {
	return makeCreateAndPopulateTableOptionImpl(opts...)
}
