package check

import (
	"fmt"
	"log"

	"github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
)

type Name string

// Check fails if `b` is false
func Check(b bool, checkOpts ...CheckOption) {
	if !b {
		opts := makeCheckOptionImpl(checkOpts...)
		msg := opts.message
		if msg == "" {
			msg = "check failure"
		}
		log.Fatalf(msg)
	}
}

// Err fails if `err` is non-null
func Err(err error) {
	if err != nil {
		Check(false, CheckMessage(err.Error()))
	}
}

// ErrAll creates an aggregate error and fails on it, if it exists
func ErrAll(errs chan error) {
	parallel.WaitFor(func() {
		eb := errors.MakeErrorCollector()
		for e := range errs {
			eb.Add(e)
		}
		Err(eb.Build())
	})
}

// CheckNonEmptyString fails if `s` is empty
func CheckNonEmptyString(s string, n Name) {
	Check(s != "", CheckMessage(fmt.Sprintf("%q expected to be non-empty", string(n))))
}
