package check

import (
	"fmt"
	"log"
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

// CheckNonEmptyString fails if `s` is empty
func CheckNonEmptyString(s string, n Name) {
	Check(s != "", CheckMessage(fmt.Sprintf("%q expected to be non-empty", string(n))))
}
