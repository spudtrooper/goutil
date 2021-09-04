package or

func Int(a, b int) int {
	if a != 0 {
		return a
	}
	return b
}

func Int32(a, b int32) int32 {
	if a != 0 {
		return a
	}
	return b
}

func Int64(a, b int64) int64 {
	if a != 0 {
		return a
	}
	return b
}

func String(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func Float32(a, b float32) float32 {
	if a != 0 {
		return a
	}
	return b
}

func Float64(a, b float64) float64 {
	if a != 0 {
		return a
	}
	return b
}
