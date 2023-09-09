// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package sqlite3

import "fmt"

type PopulateSqlite3TableOption struct {
	f func(*populateSqlite3TableOptionImpl)
	s string
}

func (o PopulateSqlite3TableOption) String() string { return o.s }

type PopulateSqlite3TableOptions interface {
	CreateDBIfNotExists() bool
	HasCreateDBIfNotExists() bool
	DropIfExists() bool
	HasDropIfExists() bool
	LowerCaseColumnNames() bool
	HasLowerCaseColumnNames() bool
	PrimaryKey() string
	HasPrimaryKey() bool
	RemoveInvalidCharsFromColumnNames() bool
	HasRemoveInvalidCharsFromColumnNames() bool
	SnakeCaseColumnNames() bool
	HasSnakeCaseColumnNames() bool
}

func PopulateSqlite3TableCreateDBIfNotExists(createDBIfNotExists bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = createDBIfNotExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}
func PopulateSqlite3TableCreateDBIfNotExistsFlag(createDBIfNotExists *bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		if createDBIfNotExists == nil {
			return
		}
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = *createDBIfNotExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}

func PopulateSqlite3TableDropIfExists(dropIfExists bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		opts.has_dropIfExists = true
		opts.dropIfExists = dropIfExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableDropIfExists(bool %+v)", dropIfExists)}
}
func PopulateSqlite3TableDropIfExistsFlag(dropIfExists *bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		if dropIfExists == nil {
			return
		}
		opts.has_dropIfExists = true
		opts.dropIfExists = *dropIfExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableDropIfExists(bool %+v)", dropIfExists)}
}

func PopulateSqlite3TableLowerCaseColumnNames(lowerCaseColumnNames bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		opts.has_lowerCaseColumnNames = true
		opts.lowerCaseColumnNames = lowerCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableLowerCaseColumnNames(bool %+v)", lowerCaseColumnNames)}
}
func PopulateSqlite3TableLowerCaseColumnNamesFlag(lowerCaseColumnNames *bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		if lowerCaseColumnNames == nil {
			return
		}
		opts.has_lowerCaseColumnNames = true
		opts.lowerCaseColumnNames = *lowerCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableLowerCaseColumnNames(bool %+v)", lowerCaseColumnNames)}
}

func PopulateSqlite3TablePrimaryKey(primaryKey string) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		opts.has_primaryKey = true
		opts.primaryKey = primaryKey
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TablePrimaryKey(string %+v)", primaryKey)}
}
func PopulateSqlite3TablePrimaryKeyFlag(primaryKey *string) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		if primaryKey == nil {
			return
		}
		opts.has_primaryKey = true
		opts.primaryKey = *primaryKey
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TablePrimaryKey(string %+v)", primaryKey)}
}

func PopulateSqlite3TableRemoveInvalidCharsFromColumnNames(removeInvalidCharsFromColumnNames bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		opts.has_removeInvalidCharsFromColumnNames = true
		opts.removeInvalidCharsFromColumnNames = removeInvalidCharsFromColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableRemoveInvalidCharsFromColumnNames(bool %+v)", removeInvalidCharsFromColumnNames)}
}
func PopulateSqlite3TableRemoveInvalidCharsFromColumnNamesFlag(removeInvalidCharsFromColumnNames *bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		if removeInvalidCharsFromColumnNames == nil {
			return
		}
		opts.has_removeInvalidCharsFromColumnNames = true
		opts.removeInvalidCharsFromColumnNames = *removeInvalidCharsFromColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableRemoveInvalidCharsFromColumnNames(bool %+v)", removeInvalidCharsFromColumnNames)}
}

func PopulateSqlite3TableSnakeCaseColumnNames(snakeCaseColumnNames bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		opts.has_snakeCaseColumnNames = true
		opts.snakeCaseColumnNames = snakeCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableSnakeCaseColumnNames(bool %+v)", snakeCaseColumnNames)}
}
func PopulateSqlite3TableSnakeCaseColumnNamesFlag(snakeCaseColumnNames *bool) PopulateSqlite3TableOption {
	return PopulateSqlite3TableOption{func(opts *populateSqlite3TableOptionImpl) {
		if snakeCaseColumnNames == nil {
			return
		}
		opts.has_snakeCaseColumnNames = true
		opts.snakeCaseColumnNames = *snakeCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableSnakeCaseColumnNames(bool %+v)", snakeCaseColumnNames)}
}

type populateSqlite3TableOptionImpl struct {
	createDBIfNotExists                   bool
	has_createDBIfNotExists               bool
	dropIfExists                          bool
	has_dropIfExists                      bool
	lowerCaseColumnNames                  bool
	has_lowerCaseColumnNames              bool
	primaryKey                            string
	has_primaryKey                        bool
	removeInvalidCharsFromColumnNames     bool
	has_removeInvalidCharsFromColumnNames bool
	snakeCaseColumnNames                  bool
	has_snakeCaseColumnNames              bool
}

func (p *populateSqlite3TableOptionImpl) CreateDBIfNotExists() bool { return p.createDBIfNotExists }
func (p *populateSqlite3TableOptionImpl) HasCreateDBIfNotExists() bool {
	return p.has_createDBIfNotExists
}
func (p *populateSqlite3TableOptionImpl) DropIfExists() bool         { return p.dropIfExists }
func (p *populateSqlite3TableOptionImpl) HasDropIfExists() bool      { return p.has_dropIfExists }
func (p *populateSqlite3TableOptionImpl) LowerCaseColumnNames() bool { return p.lowerCaseColumnNames }
func (p *populateSqlite3TableOptionImpl) HasLowerCaseColumnNames() bool {
	return p.has_lowerCaseColumnNames
}
func (p *populateSqlite3TableOptionImpl) PrimaryKey() string  { return p.primaryKey }
func (p *populateSqlite3TableOptionImpl) HasPrimaryKey() bool { return p.has_primaryKey }
func (p *populateSqlite3TableOptionImpl) RemoveInvalidCharsFromColumnNames() bool {
	return p.removeInvalidCharsFromColumnNames
}
func (p *populateSqlite3TableOptionImpl) HasRemoveInvalidCharsFromColumnNames() bool {
	return p.has_removeInvalidCharsFromColumnNames
}
func (p *populateSqlite3TableOptionImpl) SnakeCaseColumnNames() bool { return p.snakeCaseColumnNames }
func (p *populateSqlite3TableOptionImpl) HasSnakeCaseColumnNames() bool {
	return p.has_snakeCaseColumnNames
}

func makePopulateSqlite3TableOptionImpl(opts ...PopulateSqlite3TableOption) *populateSqlite3TableOptionImpl {
	res := &populateSqlite3TableOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakePopulateSqlite3TableOptions(opts ...PopulateSqlite3TableOption) PopulateSqlite3TableOptions {
	return makePopulateSqlite3TableOptionImpl(opts...)
}
