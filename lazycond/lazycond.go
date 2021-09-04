// Package lazycond contains functions for choosing between lazy values.
package lazycond

// Int returns a if b, otherwise z
func Int(b bool, a, z func() int) int {
	if b {
		return a()
	}
	return z()
}

// Int32 returns a if b, otherwise z
func Int32(b bool, a, z func() int32) int32 {
	if b {
		return a()
	}
	return z()
}

// Int64 returns a if b, otherwise z
func Int64(b bool, a, z func() int64) int64 {
	if b {
		return a()
	}
	return z()
}

// String returns a if b, otherwise z
func String(b bool, a, z func() string) string {
	if b {
		return a()
	}
	return z()
}

// Float32 returns a if b, otherwise z
func Float32(b bool, a, z func() float32) float32 {
	if b {
		return a()
	}
	return z()
}

// Float64 returns a if b, otherwise z
func Float64(b bool, a, z func() float64) float64 {
	if b {
		return a()
	}
	return z()
}
