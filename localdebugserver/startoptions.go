// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package localdebugserver

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type StartOption struct {
	f func(*startOptionImpl)
	s string
}

func (o StartOption) String() string { return o.s }

type StartOptions interface {
	Port() int
	HasPort() bool
	Wait() bool
	HasWait() bool
}

func StartPort(port int) StartOption {
	return StartOption{func(opts *startOptionImpl) {
		opts.has_port = true
		opts.port = port
	}, fmt.Sprintf("localdebugserver.StartPort(int %+v)", port)}
}
func StartPortFlag(port *int) StartOption {
	return StartOption{func(opts *startOptionImpl) {
		if port == nil {
			return
		}
		opts.has_port = true
		opts.port = *port
	}, fmt.Sprintf("localdebugserver.StartPort(int %+v)", port)}
}

func StartWait(wait bool) StartOption {
	return StartOption{func(opts *startOptionImpl) {
		opts.has_wait = true
		opts.wait = wait
	}, fmt.Sprintf("localdebugserver.StartWait(bool %+v)", wait)}
}
func StartWaitFlag(wait *bool) StartOption {
	return StartOption{func(opts *startOptionImpl) {
		if wait == nil {
			return
		}
		opts.has_wait = true
		opts.wait = *wait
	}, fmt.Sprintf("localdebugserver.StartWait(bool %+v)", wait)}
}

type startOptionImpl struct {
	port     int
	has_port bool
	wait     bool
	has_wait bool
}

func (s *startOptionImpl) Port() int     { return or.Int(s.port, 8000) }
func (s *startOptionImpl) HasPort() bool { return s.has_port }
func (s *startOptionImpl) Wait() bool    { return s.wait }
func (s *startOptionImpl) HasWait() bool { return s.has_wait }

func makeStartOptionImpl(opts ...StartOption) *startOptionImpl {
	res := &startOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeStartOptions(opts ...StartOption) StartOptions {
	return makeStartOptionImpl(opts...)
}
