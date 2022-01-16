package or

import (
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	var tests = []struct {
		name string
		a, b int
		want int
	}{
		{
			name: "defaults",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			b:    0,
			want: 1,
		},
		{
			name: "default non-default",
			a:    0,
			b:    2,
			want: 2,
		},
		{
			name: "non-default non-default",
			a:    1,
			b:    2,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int(test.a, test.b); test.want != got {
				t.Errorf("Int(%d,%d): want %d, got %d", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	var tests = []struct {
		name string
		a, b int32
		want int32
	}{
		{
			name: "defaults",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			b:    0,
			want: 1,
		},
		{
			name: "default non-default",
			a:    0,
			b:    2,
			want: 2,
		},
		{
			name: "non-default non-default",
			a:    1,
			b:    2,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int32(test.a, test.b); test.want != got {
				t.Errorf("Int32(%d,%d): want %d, got %d", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	var tests = []struct {
		name string
		a, b int64
		want int64
	}{
		{
			name: "defaults",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			b:    0,
			want: 1,
		},
		{
			name: "default non-default",
			a:    0,
			b:    2,
			want: 2,
		},
		{
			name: "non-default non-default",
			a:    1,
			b:    2,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int64(test.a, test.b); test.want != got {
				t.Errorf("Int64(%d,%d): want %d, got %d", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	var tests = []struct {
		name string
		a, b float32
		want float32
	}{
		{
			name: "defaults",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			b:    0,
			want: 1,
		},
		{
			name: "default non-default",
			a:    0,
			b:    2,
			want: 2,
		},
		{
			name: "non-default non-default",
			a:    1,
			b:    2,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float32(test.a, test.b); test.want != got {
				t.Errorf("Float32(%f,%f): want %f, got %f", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	var tests = []struct {
		name string
		a, b float64
		want float64
	}{
		{
			name: "defaults",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			b:    0,
			want: 1,
		},
		{
			name: "default non-default",
			a:    0,
			b:    2,
			want: 2,
		},
		{
			name: "non-default non-default",
			a:    1,
			b:    2,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float64(test.a, test.b); test.want != got {
				t.Errorf("Float64(%f,%f): want %f, got %f", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name string
		a, b string
		want string
	}{
		{
			name: "defaults",
			a:    "",
			b:    "",
			want: "",
		},
		{
			name: "non-default default",
			a:    "1",
			b:    "",
			want: "1",
		},
		{
			name: "default non-default",
			a:    "",
			b:    "2",
			want: "2",
		},
		{
			name: "non-default non-default",
			a:    "1",
			b:    "2",
			want: "1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := String(test.a, test.b); test.want != got {
				t.Errorf("String(%q,%q): want %q, got %q", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	var tests = []struct {
		name string
		a, b time.Time
		want time.Time
	}{
		{
			name: "defaults",
			a:    time.Time{},
			b:    time.Time{},
			want: time.Time{},
		},
		{
			name: "default non-default",
			a:    now,
			b:    time.Time{},
			want: now,
		},
		{
			name: "non-default non-default",
			a:    now,
			b:    now.Add(1 * time.Second),
			want: now,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if want, got := test.want, Time(test.a, test.b); want.Unix() != got.Unix() {
				t.Errorf("Time(%v,%v): want %v, got %v", test.a, test.b, want, got)
			}
		})
	}
}
