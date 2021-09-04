// Package cond contains functions for choosing between values.
package cond

// Int returns a if b, otherwise z
func Int(b bool, a, z int) int {
	if b {
		return a
	}
	return z
}

// Int32 returns a if b, otherwise z
func Int32(b bool, a, z int32) int32 {
	if b {
		return a
	}
	return z
}

// Int64 returns a if b, otherwise z
func Int64(b bool, a, z int64) int64 {
	if b {
		return a
	}
	return z
}

// String returns a if b, otherwise z
func String(b bool, a, z string) string {
	if b {
		return a
	}
	return z
}

// Float32 returns a if b, otherwise z
func Float32(b bool, a, z float32) float32 {
	if b {
		return a
	}
	return z
}

// Float64 returns a if b, otherwise z
func Float64(b bool, a, z float64) float64 {
	if b {
		return a
	}
	return z
}
