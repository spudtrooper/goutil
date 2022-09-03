// Package or contains functions for eagerly choosing the first non-default value from two values.
// The name should probably be nondefault or something, but I embrace being the shittiest namer
// known to history.
package or

import (
	"time"
)

// SupportedType is the type of the supported types to generic versions of these functions.
type SupportedType interface {
	uint | uint16 | uint32 | uint64 | int | int16 | int32 | int64 | float32 | float64 | string | bool | time.Duration
}

func defaultValue[T SupportedType](v T) T {
	var a any = v
	var res any
	switch a.(type) {
	case uint:
		res = uint(0)
	case uint16:
		res = uint16(0)
	case uint32:
		res = uint32(0)
	case uint64:
		res = uint64(0)
	case int:
		res = int(0)
	case int16:
		res = int16(0)
	case int32:
		res = int32(0)
	case int64:
		res = int64(0)
	case float32:
		res = float32(0)
	case float64:
		res = float64(0)
	case string:
		res = ""
	case bool:
		res = false
	case time.Duration:
		res = time.Duration(0)
	default:
		panic("cannot reach this")
	}
	return res.(T)
}

// Any returns a if a isn't the default, otherwise b
func Any[T SupportedType](a, b T) T {
	if def := defaultValue(a); a != def {
		return a
	}
	return b
}

// Default returns a if a isn't the default otherwise 0
func Default[T SupportedType](a T) T { return Any(a, defaultValue(a)) }
