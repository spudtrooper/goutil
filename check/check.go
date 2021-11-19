package check

import "fmt"

type Name string

// Check fails if `b` is false
func Check(b bool, checkOpts ...CheckOption) {
	if !b {
		opts := makeCheckOptionImpl(checkOpts...)
		msg := opts.message
		if msg == "" {
			msg = "check failure"
		}
		panic(msg)
	}
}

// Err fails if `err` is non-null
func Err(err error) {
	Check(err == nil, CheckMessage(err.Error()))
}

// CheckNonEmptyString fails if `s` is empty
func CheckNonEmptyString(s string, n Name) {
	Check(s != "", CheckMessage(fmt.Sprintf("%q expected to be non-empty", string(n))))
}
