package sets

type IntSet map[int]bool
type Int32Set map[int32]bool
type Int64Set map[int64]bool
type StringSet map[string]bool
type Float32Set map[float32]bool
type Float64Set map[float64]bool

// Int creates a set from the list
func Int(lst []int) IntSet {
	res := map[int]bool{}
	for _, s := range lst {
		res[s] = true
	}
	return res
}

// Int32 creates a set from the list
func Int32(lst []int32) Int32Set {
	res := map[int32]bool{}
	for _, s := range lst {
		res[s] = true
	}
	return res
}

// Int64 creates a set from the list
func Int64(lst []int64) Int64Set {
	res := map[int64]bool{}
	for _, s := range lst {
		res[s] = true
	}
	return res
}

// String creates a set from the list
func String(lst []string) StringSet {
	res := map[string]bool{}
	for _, s := range lst {
		res[s] = true
	}
	return res
}

// Float32 creates a set from the list
func Float32(lst []float32) Float32Set {
	res := map[float32]bool{}
	for _, s := range lst {
		res[s] = true
	}
	return res
}

// Float64 creates a set from the list
func Float64(lst []float64) Float64Set {
	res := map[float64]bool{}
	for _, s := range lst {
		res[s] = true
	}
	return res
}
