// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package request

//go:generate genopts --prefix=ConnectToURLCache --outfile=connecttourlcacheoptions.go "port:int" "dbName:string"

type ConnectToURLCacheOption func(*connectToURLCacheOptionImpl)

type ConnectToURLCacheOptions interface {
	Port() int
	DbName() string
}

func ConnectToURLCachePort(port int) ConnectToURLCacheOption {
	return func(opts *connectToURLCacheOptionImpl) {
		opts.port = port
	}
}
func ConnectToURLCachePortFlag(port *int) ConnectToURLCacheOption {
	return func(opts *connectToURLCacheOptionImpl) {
		opts.port = *port
	}
}

func ConnectToURLCacheDbName(dbName string) ConnectToURLCacheOption {
	return func(opts *connectToURLCacheOptionImpl) {
		opts.dbName = dbName
	}
}
func ConnectToURLCacheDbNameFlag(dbName *string) ConnectToURLCacheOption {
	return func(opts *connectToURLCacheOptionImpl) {
		opts.dbName = *dbName
	}
}

type connectToURLCacheOptionImpl struct {
	port   int
	dbName string
}

func (c *connectToURLCacheOptionImpl) Port() int      { return c.port }
func (c *connectToURLCacheOptionImpl) DbName() string { return c.dbName }

func makeConnectToURLCacheOptionImpl(opts ...ConnectToURLCacheOption) *connectToURLCacheOptionImpl {
	res := &connectToURLCacheOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeConnectToURLCacheOptions(opts ...ConnectToURLCacheOption) ConnectToURLCacheOptions {
	return makeConnectToURLCacheOptionImpl(opts...)
}
