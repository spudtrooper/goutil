// Package or contains functions for eagerly choosing the first non-default value from two values.
// The name should probably be nondefault or something, but I embrace being the shittiest namer
// known to history.
package or

import (
	"time"
)

// Ints returns a if len(a) != 0, otherwise b
func Ints(a, b []int) []int {
	if len(a) != 0 {
		return a
	}
	return b
}

// Int returns a if a != 0, otherwise b
func Int(a, b int) int {
	if a != 0 {
		return a
	}
	return b
}

// IntDefault returns a if a != 0, otherwise 0
func IntDefault(a int) int { return Int(a, 0) }

// Int8 returns a if a != 0, otherwise b
func Int8(a, b int8) int8 {
	if a != 0 {
		return a
	}
	return b
}

// Int8Default returns a  if it's non-zero, otherwise 0.
func Int8Default(a int8) int8 { return Int8(a, int8(0)) }

// Int16 returns a if a != 0, otherwise b
func Int16(a, b int16) int16 {
	if a != 0 {
		return a
	}
	return b
}

// Int16Default returns a  if it's non-zero, otherwise 0.
func Int16Default(a int16) int16 { return Int16(a, int16(0)) }

// Int32 returns a if a != 0, otherwise b
func Int32(a, b int32) int32 {
	if a != 0 {
		return a
	}
	return b
}

// Int32Default returns a  if it's non-zero, otherwise 0.
func Int32Default(a int32) int32 { return Int32(a, int32(0)) }

// Int64 returns a if a != 0, otherwise b
func Int64(a, b int64) int64 {
	if a != 0 {
		return a
	}
	return b
}

// Int64Default returns a  if it's non-zero, otherwise 0.
func Int64Default(a int64) int64 { return Int64(a, int64(0)) }

// String returns a if a != "", otherwise b
func String(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// StringDefault returns a  if it's non-zero, otherwise 0.
func StringDefault(a string) string { return String(a, "") }

// Float32 returns a if a != 0, otherwise b
func Float32(a, b float32) float32 {
	if a != 0 {
		return a
	}
	return b
}

// Float32Default returns a  if it's non-zero, otherwise 0.
func Float32Default(a float32) float32 { return Float32(a, float32(0)) }

// Float64 returns a if a != 0, otherwise b
func Float64(a, b float64) float64 {
	if a != 0 {
		return a
	}
	return b
}

// Float64Default returns a  if it's non-zero, otherwise 0.
func Float64Default(a float64) float64 { return Float64(a, float64(0)) }

// Time returns a if a.IsZero(), otherwise b
func Time(a, b time.Time) time.Time {
	if !a.IsZero() {
		return a
	}
	return b
}

// TimeDefault returns a  if it's non-zero, otherwise 0.
func TimeDefault(a time.Time) time.Time { return Time(a, time.Time{}) }

// Duration returns a if a != 0, otherwise  b
func Duration(a, b time.Duration) time.Duration {
	if a != 0 {
		return a
	}
	return b
}

// DurationDefault returns a  if it's non-zero, otherwise 0.
func DurationDefault(a time.Duration) time.Duration { return Duration(a, time.Duration(0)) }
