package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestInt(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  IntSet
	}{
		{
			name:  "empty",
			input: []int{},
			want:  map[int]bool{},
		},
		{
			name:  "one",
			input: []int{1},
			want: map[int]bool{
				1: true,
			},
		},
		{
			name:  "uniques",
			input: []int{1, 2, 3},
			want: map[int]bool{
				1: true,
				2: true,
				3: true,
			},
		},
		{
			name:  "dups",
			input: []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
			want: map[int]bool{
				1: true,
				2: true,
				3: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int(test.input); !reflect.DeepEqual(test.want, got) {
				t.Errorf("Int(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	var tests = []struct {
		name  string
		input []int32
		want  Int32Set
	}{
		{
			name:  "empty",
			input: []int32{},
			want:  map[int32]bool{},
		},
		{
			name:  "one",
			input: []int32{1},
			want: map[int32]bool{
				1: true,
			},
		},
		{
			name:  "uniques",
			input: []int32{1, 2, 3},
			want: map[int32]bool{
				1: true,
				2: true,
				3: true,
			},
		},
		{
			name:  "dups",
			input: []int32{1, 2, 3, 1, 2, 3, 1, 2, 3},
			want: map[int32]bool{
				1: true,
				2: true,
				3: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int32(test.input); !reflect.DeepEqual(test.want, got) {
				t.Errorf("Int32(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	var tests = []struct {
		name  string
		input []int64
		want  Int64Set
	}{
		{
			name:  "empty",
			input: []int64{},
			want:  map[int64]bool{},
		},
		{
			name:  "one",
			input: []int64{1},
			want: map[int64]bool{
				1: true,
			},
		},
		{
			name:  "uniques",
			input: []int64{1, 2, 3},
			want: map[int64]bool{
				1: true,
				2: true,
				3: true,
			},
		},
		{
			name:  "dups",
			input: []int64{1, 2, 3, 1, 2, 3, 1, 2, 3},
			want: map[int64]bool{
				1: true,
				2: true,
				3: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int64(test.input); !reflect.DeepEqual(test.want, got) {
				t.Errorf("Int64(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	var tests = []struct {
		name  string
		input []float32
		want  Float32Set
	}{
		{
			name:  "empty",
			input: []float32{},
			want:  map[float32]bool{},
		},
		{
			name:  "one",
			input: []float32{1},
			want: map[float32]bool{
				1: true,
			},
		},
		{
			name:  "uniques",
			input: []float32{1, 2, 3},
			want: map[float32]bool{
				1: true,
				2: true,
				3: true,
			},
		},
		{
			name:  "dups",
			input: []float32{1, 2, 3, 1, 2, 3, 1, 2, 3},
			want: map[float32]bool{
				1: true,
				2: true,
				3: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float32(test.input); !reflect.DeepEqual(test.want, got) {
				t.Errorf("Float32(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	var tests = []struct {
		name  string
		input []float64
		want  Float64Set
	}{
		{
			name:  "empty",
			input: []float64{},
			want:  map[float64]bool{},
		},
		{
			name:  "one",
			input: []float64{1},
			want: map[float64]bool{
				1: true,
			},
		},
		{
			name:  "uniques",
			input: []float64{1, 2, 3},
			want: map[float64]bool{
				1: true,
				2: true,
				3: true,
			},
		},
		{
			name:  "dups",
			input: []float64{1, 2, 3, 1, 2, 3, 1, 2, 3},
			want: map[float64]bool{
				1: true,
				2: true,
				3: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float64(test.input); !reflect.DeepEqual(test.want, got) {
				t.Errorf("Float64(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  StringSet
	}{
		{
			name:  "empty",
			input: []string{},
			want:  map[string]bool{},
		},
		{
			name:  "one",
			input: []string{"1"},
			want: map[string]bool{
				"1": true,
			},
		},
		{
			name:  "uniques",
			input: []string{"1", "2", "3"},
			want: map[string]bool{
				"1": true,
				"2": true,
				"3": true,
			},
		},
		{
			name:  "dups",
			input: []string{"1", "2", "3", "1", "2", "3", "1", "2", "3"},
			want: map[string]bool{
				"1": true,
				"2": true,
				"3": true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := String(test.input); !reflect.DeepEqual(test.want, got) {
				t.Errorf("String(%v): want %v, got %v", test.input, test.want, got)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	var tests = []struct {
		name  string
		input StringSet
		want  []string
	}{
		{
			name:  "empty",
			input: String([]string{}),
			want:  []string{},
		},
		{
			name:  "one",
			input: String([]string{"1"}),
			want:  []string{"1"},
		},
		{
			name:  "uniques",
			input: String([]string{"1", "2", "3"}),
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "dups",
			input: String([]string{"1", "2", "3", "1", "2", "3", "1", "2", "3"}),
			want:  []string{"1", "2", "3"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := Keys(test.input), test.want
			sort.Strings(got)
			sort.Strings(want)
			if !reflect.DeepEqual(want, got) {
				t.Errorf("Keys(%v): want %v, got %v", test.input, want, got)
			}
		})
	}
}

func TestSortedKeys(t *testing.T) {
	var tests = []struct {
		name  string
		input StringSet
		want  []string
	}{
		{
			name:  "empty",
			input: String([]string{}),
			want:  []string{},
		},
		{
			name:  "one",
			input: String([]string{"1"}),
			want:  []string{"1"},
		},
		{
			name:  "uniques",
			input: String([]string{"1", "2", "3"}),
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "dups",
			input: String([]string{"1", "2", "3", "1", "2", "3", "1", "2", "3"}),
			want:  []string{"1", "2", "3"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, want := SortedKeys(test.input), test.want
			if !reflect.DeepEqual(want, got) {
				t.Errorf("SortedKeys(%v): want %v, got %v", test.input, want, got)
			}
		})
	}
}

func TestIntSetAdd(t *testing.T) {
	set := Int([]int{1, 2})
	set.Add(3, 4, 5)

	if len(set) != 5 || !set[1] || !set[2] || !set[3] || !set[4] || !set[5] {
		t.Errorf("Add failed for IntSet, expected [1:2:3:4:5], got %v", set)
	}
}

func TestInt32SetAdd(t *testing.T) {
	set := Int32([]int32{1, 2})
	set.Add(3, 4, 5)

	if len(set) != 5 || !set[1] || !set[2] || !set[3] || !set[4] || !set[5] {
		t.Errorf("Add failed for Int32Set, expected [1:2:3:4:5], got %v", set)
	}
}

func TestInt64SetAdd(t *testing.T) {
	set := Int64([]int64{1, 2})
	set.Add(3, 4, 5)

	if len(set) != 5 || !set[1] || !set[2] || !set[3] || !set[4] || !set[5] {
		t.Errorf("Add failed for Int64Set, expected [1:2:3:4:5], got %v", set)
	}
}

func TestStringSetAdd(t *testing.T) {
	set := String([]string{"one", "two"})
	set.Add("three", "four")

	if len(set) != 4 || !set["one"] || !set["two"] || !set["three"] || !set["four"] {
		t.Errorf("Add failed for StringSet, expected [one:two:three:four], got %v", set)
	}
}

func TestFloat32SetAdd(t *testing.T) {
	set := Float32([]float32{1.1, 2.2})
	set.Add(3.3, 4.4)

	if len(set) != 4 || !set[1.1] || !set[2.2] || !set[3.3] || !set[4.4] {
		t.Errorf("Add failed for Float32Set, expected [1.1:2.2:3.3:4.4], got %v", set)
	}
}

func TestFloat64SetAdd(t *testing.T) {
	set := Float64([]float64{1.1, 2.2})
	set.Add(3.3, 4.4)

	if len(set) != 4 || !set[1.1] || !set[2.2] || !set[3.3] || !set[4.4] {
		t.Errorf("Add failed for Float64Set, expected [1.1:2.2:3.3:4.4], got %v", set)
	}
}

func TestIntSetSorted(t *testing.T) {
	set := Int([]int{3, 1, 2})
	set.Add(4, 5)

	slice := set.Slice()
	sorted := set.Sorted()
	sort.Ints(slice) // Ensure the original slice is sorted
	if !reflect.DeepEqual(sorted, slice) {
		t.Errorf("Sorted failed for IntSet, expected %v, got %v", slice, sorted)
	}
}

func TestStringSetSorted(t *testing.T) {
	set := String([]string{"three", "one", "two"})
	set.Add("four", "five")

	slice := set.Slice()

	sorted := set.Sorted()
	sort.Strings(slice) // Ensure the original slice is sorted
	if !reflect.DeepEqual(sorted, slice) {
		t.Errorf("Sorted failed for StringSet, expected %v, got %v", slice, sorted)
	}
}

func TestFloat64SetSorted(t *testing.T) {
	set := Float64([]float64{3.3, 1.1, 2.2})
	set.Add(4.4, 5.5)

	slice := set.Slice()
	sorted := set.Sorted()
	sort.Float64s(slice) // Ensure the original slice is sorted
	if !reflect.DeepEqual(sorted, slice) {
		t.Errorf("Sorted failed for Float64Set, expected %v, got %v", slice, sorted)
	}
}
