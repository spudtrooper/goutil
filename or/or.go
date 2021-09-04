package or

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
