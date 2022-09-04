// Package flags contains functions for creating flags with only defaults.
package flags

import (
	"flag"
	"time"
)

// String creates a `string` flag with default value `""`
func String(name, desc string) *string {
	return flag.String(name, "", desc)
}

// Bool creates a `bool` flag with default value `false`
func Bool(name, desc string) *bool {
	return flag.Bool(name, false, desc)
}

// Int creates an `int` flag with default value `0`
func Int(name, desc string) *int {
	return flag.Int(name, 0, desc)
}

// Int64 creates an `int64` flag with default value `0`
func Int64(name, desc string) *int64 {
	return flag.Int64(name, 0, desc)
}

// Uint64 creates a `uint64` flag with default value `0`
func Uint64(name, desc string) *uint64 {
	return flag.Uint64(name, 0, desc)
}

// Uint creates a `uint` flag with default value `0`
func Uint(name, desc string) *uint {
	return flag.Uint(name, 0, desc)
}

// Float64 creates a `float64` flag with default value `0`
func Float64(name, desc string) *float64 {
	return flag.Float64(name, 0, desc)
}

// Duration creates a `duration` flag with default value `0`
func Duration(name, desc string) *time.Duration {
	return flag.Duration(name, 0, desc)
}
