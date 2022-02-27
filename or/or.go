// Package or contains functions for eagerly choosing the first non-default value from two values.
// The name should probably be nondefault or something, but I embrace being the shittiest namer
// known to history.
package or

import "time"

// Int returns a if a != 0, otherwise b
func Int(a, b int) int {
	if a != 0 {
		return a
	}
	return b
}

// Int32 returns a if a != 0, otherwise b
func Int32(a, b int32) int32 {
	if a != 0 {
		return a
	}
	return b
}

// Int64 returns a if a != 0, otherwise b
func Int64(a, b int64) int64 {
	if a != 0 {
		return a
	}
	return b
}

// String returns a if a != "", otherwise b
func String(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// Float32 returns a if a != 0, otherwise b
func Float32(a, b float32) float32 {
	if a != 0 {
		return a
	}
	return b
}

// Float64 returns a if a != 0, otherwise b
func Float64(a, b float64) float64 {
	if a != 0 {
		return a
	}
	return b
}

// Time returns a if a.IsZero(), otherwise b
func Time(a, b time.Time) time.Time {
	if !a.IsZero() {
		return a
	}
	return b
}

// Duration returns a if a != 0, otherwise  b
func Duration(a, b time.Duration) time.Duration {
	if a != 0 {
		return a
	}
	return b
}
