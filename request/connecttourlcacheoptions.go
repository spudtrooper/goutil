// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package request

import "fmt"

//go:generate genopts --prefix=ConnectToURLCache --outfile=connecttourlcacheoptions.go "port:int" "dbName:string"

type ConnectToURLCacheOption struct {
	f func(*connectToURLCacheOptionImpl)
	s string
}

func (o ConnectToURLCacheOption) String() string { return o.s }

type ConnectToURLCacheOptions interface {
	DbName() string
	HasDbName() bool
	Port() int
	HasPort() bool
}

func ConnectToURLCacheDbName(dbName string) ConnectToURLCacheOption {
	return ConnectToURLCacheOption{func(opts *connectToURLCacheOptionImpl) {
		opts.has_dbName = true
		opts.dbName = dbName
	}, fmt.Sprintf("request.ConnectToURLCacheDbName(string %+v)", dbName)}
}
func ConnectToURLCacheDbNameFlag(dbName *string) ConnectToURLCacheOption {
	return ConnectToURLCacheOption{func(opts *connectToURLCacheOptionImpl) {
		if dbName == nil {
			return
		}
		opts.has_dbName = true
		opts.dbName = *dbName
	}, fmt.Sprintf("request.ConnectToURLCacheDbName(string %+v)", dbName)}
}

func ConnectToURLCachePort(port int) ConnectToURLCacheOption {
	return ConnectToURLCacheOption{func(opts *connectToURLCacheOptionImpl) {
		opts.has_port = true
		opts.port = port
	}, fmt.Sprintf("request.ConnectToURLCachePort(int %+v)", port)}
}
func ConnectToURLCachePortFlag(port *int) ConnectToURLCacheOption {
	return ConnectToURLCacheOption{func(opts *connectToURLCacheOptionImpl) {
		if port == nil {
			return
		}
		opts.has_port = true
		opts.port = *port
	}, fmt.Sprintf("request.ConnectToURLCachePort(int %+v)", port)}
}

type connectToURLCacheOptionImpl struct {
	dbName     string
	has_dbName bool
	port       int
	has_port   bool
}

func (c *connectToURLCacheOptionImpl) DbName() string  { return c.dbName }
func (c *connectToURLCacheOptionImpl) HasDbName() bool { return c.has_dbName }
func (c *connectToURLCacheOptionImpl) Port() int       { return c.port }
func (c *connectToURLCacheOptionImpl) HasPort() bool   { return c.has_port }

func makeConnectToURLCacheOptionImpl(opts ...ConnectToURLCacheOption) *connectToURLCacheOptionImpl {
	res := &connectToURLCacheOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeConnectToURLCacheOptions(opts ...ConnectToURLCacheOption) ConnectToURLCacheOptions {
	return makeConnectToURLCacheOptionImpl(opts...)
}
