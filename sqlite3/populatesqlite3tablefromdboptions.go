// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package sqlite3

import "fmt"

type PopulateSqlite3TableFromDBOption struct {
	f func(*populateSqlite3TableFromDBOptionImpl)
	s string
}

func (o PopulateSqlite3TableFromDBOption) String() string { return o.s }

type PopulateSqlite3TableFromDBOptions interface {
	CreateDBIfNotExists() bool
	HasCreateDBIfNotExists() bool
	DeleteWhere() string
	HasDeleteWhere() bool
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
	Verbose() bool
	HasVerbose() bool
	ToDropTableIfExistsOptions() []DropTableIfExistsOption
	ToOpenDBOptions() []OpenDBOption
}

func PopulateSqlite3TableFromDBCreateDBIfNotExists(createDBIfNotExists bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = createDBIfNotExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}
func PopulateSqlite3TableFromDBCreateDBIfNotExistsFlag(createDBIfNotExists *bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if createDBIfNotExists == nil {
			return
		}
		opts.has_createDBIfNotExists = true
		opts.createDBIfNotExists = *createDBIfNotExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBCreateDBIfNotExists(bool %+v)", createDBIfNotExists)}
}

func PopulateSqlite3TableFromDBDeleteWhere(deleteWhere string) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_deleteWhere = true
		opts.deleteWhere = deleteWhere
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBDeleteWhere(string %+v)", deleteWhere)}
}
func PopulateSqlite3TableFromDBDeleteWhereFlag(deleteWhere *string) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if deleteWhere == nil {
			return
		}
		opts.has_deleteWhere = true
		opts.deleteWhere = *deleteWhere
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBDeleteWhere(string %+v)", deleteWhere)}
}

func PopulateSqlite3TableFromDBDropIfExists(dropIfExists bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_dropIfExists = true
		opts.dropIfExists = dropIfExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBDropIfExists(bool %+v)", dropIfExists)}
}
func PopulateSqlite3TableFromDBDropIfExistsFlag(dropIfExists *bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if dropIfExists == nil {
			return
		}
		opts.has_dropIfExists = true
		opts.dropIfExists = *dropIfExists
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBDropIfExists(bool %+v)", dropIfExists)}
}

func PopulateSqlite3TableFromDBLowerCaseColumnNames(lowerCaseColumnNames bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_lowerCaseColumnNames = true
		opts.lowerCaseColumnNames = lowerCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBLowerCaseColumnNames(bool %+v)", lowerCaseColumnNames)}
}
func PopulateSqlite3TableFromDBLowerCaseColumnNamesFlag(lowerCaseColumnNames *bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if lowerCaseColumnNames == nil {
			return
		}
		opts.has_lowerCaseColumnNames = true
		opts.lowerCaseColumnNames = *lowerCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBLowerCaseColumnNames(bool %+v)", lowerCaseColumnNames)}
}

func PopulateSqlite3TableFromDBPrimaryKey(primaryKey string) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_primaryKey = true
		opts.primaryKey = primaryKey
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBPrimaryKey(string %+v)", primaryKey)}
}
func PopulateSqlite3TableFromDBPrimaryKeyFlag(primaryKey *string) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if primaryKey == nil {
			return
		}
		opts.has_primaryKey = true
		opts.primaryKey = *primaryKey
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBPrimaryKey(string %+v)", primaryKey)}
}

func PopulateSqlite3TableFromDBRemoveInvalidCharsFromColumnNames(removeInvalidCharsFromColumnNames bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_removeInvalidCharsFromColumnNames = true
		opts.removeInvalidCharsFromColumnNames = removeInvalidCharsFromColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBRemoveInvalidCharsFromColumnNames(bool %+v)", removeInvalidCharsFromColumnNames)}
}
func PopulateSqlite3TableFromDBRemoveInvalidCharsFromColumnNamesFlag(removeInvalidCharsFromColumnNames *bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if removeInvalidCharsFromColumnNames == nil {
			return
		}
		opts.has_removeInvalidCharsFromColumnNames = true
		opts.removeInvalidCharsFromColumnNames = *removeInvalidCharsFromColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBRemoveInvalidCharsFromColumnNames(bool %+v)", removeInvalidCharsFromColumnNames)}
}

func PopulateSqlite3TableFromDBSnakeCaseColumnNames(snakeCaseColumnNames bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_snakeCaseColumnNames = true
		opts.snakeCaseColumnNames = snakeCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBSnakeCaseColumnNames(bool %+v)", snakeCaseColumnNames)}
}
func PopulateSqlite3TableFromDBSnakeCaseColumnNamesFlag(snakeCaseColumnNames *bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if snakeCaseColumnNames == nil {
			return
		}
		opts.has_snakeCaseColumnNames = true
		opts.snakeCaseColumnNames = *snakeCaseColumnNames
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBSnakeCaseColumnNames(bool %+v)", snakeCaseColumnNames)}
}

func PopulateSqlite3TableFromDBVerbose(verbose bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBVerbose(bool %+v)", verbose)}
}
func PopulateSqlite3TableFromDBVerboseFlag(verbose *bool) PopulateSqlite3TableFromDBOption {
	return PopulateSqlite3TableFromDBOption{func(opts *populateSqlite3TableFromDBOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("sqlite3.PopulateSqlite3TableFromDBVerbose(bool %+v)", verbose)}
}

type populateSqlite3TableFromDBOptionImpl struct {
	createDBIfNotExists                   bool
	has_createDBIfNotExists               bool
	deleteWhere                           string
	has_deleteWhere                       bool
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
	verbose                               bool
	has_verbose                           bool
}

func (p *populateSqlite3TableFromDBOptionImpl) CreateDBIfNotExists() bool {
	return p.createDBIfNotExists
}
func (p *populateSqlite3TableFromDBOptionImpl) HasCreateDBIfNotExists() bool {
	return p.has_createDBIfNotExists
}
func (p *populateSqlite3TableFromDBOptionImpl) DeleteWhere() string   { return p.deleteWhere }
func (p *populateSqlite3TableFromDBOptionImpl) HasDeleteWhere() bool  { return p.has_deleteWhere }
func (p *populateSqlite3TableFromDBOptionImpl) DropIfExists() bool    { return p.dropIfExists }
func (p *populateSqlite3TableFromDBOptionImpl) HasDropIfExists() bool { return p.has_dropIfExists }
func (p *populateSqlite3TableFromDBOptionImpl) LowerCaseColumnNames() bool {
	return p.lowerCaseColumnNames
}
func (p *populateSqlite3TableFromDBOptionImpl) HasLowerCaseColumnNames() bool {
	return p.has_lowerCaseColumnNames
}
func (p *populateSqlite3TableFromDBOptionImpl) PrimaryKey() string  { return p.primaryKey }
func (p *populateSqlite3TableFromDBOptionImpl) HasPrimaryKey() bool { return p.has_primaryKey }
func (p *populateSqlite3TableFromDBOptionImpl) RemoveInvalidCharsFromColumnNames() bool {
	return p.removeInvalidCharsFromColumnNames
}
func (p *populateSqlite3TableFromDBOptionImpl) HasRemoveInvalidCharsFromColumnNames() bool {
	return p.has_removeInvalidCharsFromColumnNames
}
func (p *populateSqlite3TableFromDBOptionImpl) SnakeCaseColumnNames() bool {
	return p.snakeCaseColumnNames
}
func (p *populateSqlite3TableFromDBOptionImpl) HasSnakeCaseColumnNames() bool {
	return p.has_snakeCaseColumnNames
}
func (p *populateSqlite3TableFromDBOptionImpl) Verbose() bool    { return p.verbose }
func (p *populateSqlite3TableFromDBOptionImpl) HasVerbose() bool { return p.has_verbose }

// ToDropTableIfExistsOptions converts PopulateSqlite3TableFromDBOption to an array of DropTableIfExistsOption
func (o *populateSqlite3TableFromDBOptionImpl) ToDropTableIfExistsOptions() []DropTableIfExistsOption {
	return []DropTableIfExistsOption{
		DropTableIfExistsVerbose(o.Verbose()),
	}
}

// ToOpenDBOptions converts PopulateSqlite3TableFromDBOption to an array of OpenDBOption
func (o *populateSqlite3TableFromDBOptionImpl) ToOpenDBOptions() []OpenDBOption {
	return []OpenDBOption{
		OpenDBCreateDBIfNotExists(o.CreateDBIfNotExists()),
	}
}

func makePopulateSqlite3TableFromDBOptionImpl(opts ...PopulateSqlite3TableFromDBOption) *populateSqlite3TableFromDBOptionImpl {
	res := &populateSqlite3TableFromDBOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakePopulateSqlite3TableFromDBOptions(opts ...PopulateSqlite3TableFromDBOption) PopulateSqlite3TableFromDBOptions {
	return makePopulateSqlite3TableFromDBOptionImpl(opts...)
}
