// Package flags contains functions for creating flags with only defaults.
package flags

import (
	"flag"
	"time"
)

func String(name, desc string) *string {
	return flag.String(name, "", desc)
}

func Bool(name, desc string) *bool {
	return flag.Bool(name, false, desc)
}

func Int(name, desc string) *int {
	return flag.Int(name, 0, desc)
}

func Int64(name, desc string) *int64 {
	return flag.Int64(name, 0, desc)
}

func Uint64(name, desc string) *uint64 {
	return flag.Uint64(name, 0, desc)
}

func Uint(name, desc string) *uint {
	return flag.Uint(name, 0, desc)
}

func Float64(name, desc string) *float64 {
	return flag.Float64(name, 0, desc)
}

func Duration(name, desc string) *time.Duration {
	return flag.Duration(name, 0, desc)
}
