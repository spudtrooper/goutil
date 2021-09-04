package cond

import (
	"testing"
)

func TestInt(t *testing.T) {
	var tests = []struct {
		name string
		a, z int
		b    bool
		want int
	}{
		{
			name: "false",
			a:    1,
			z:    2,
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    1,
			z:    2,
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int(test.b, test.a, test.z); test.want != got {
				t.Fatalf("Int(%t, %d,%d): want %d, got %d", test.b, test.a, test.z, test.want, got)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	var tests = []struct {
		name string
		a, z int32
		b    bool
		want int32
	}{
		{
			name: "false",
			a:    1,
			z:    2,
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    1,
			z:    2,
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int32(test.b, test.a, test.z); test.want != got {
				t.Fatalf("Int32(%t, %d,%d): want %d, got %d", test.b, test.a, test.z, test.want, got)
			}
		})
	}
}
func TestInt64(t *testing.T) {
	var tests = []struct {
		name string
		a, z int64
		b    bool
		want int64
	}{
		{
			name: "false",
			a:    1,
			z:    2,
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    1,
			z:    2,
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int64(test.b, test.a, test.z); test.want != got {
				t.Fatalf("Int64(%t, %d,%d): want %d, got %d", test.b, test.a, test.z, test.want, got)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	var tests = []struct {
		name string
		a, z float32
		b    bool
		want float32
	}{
		{
			name: "false",
			a:    1,
			z:    2,
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    1,
			z:    2,
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float32(test.b, test.a, test.z); test.want != got {
				t.Fatalf("Float32(%t, %f,%f): want %f, got %f", test.b, test.a, test.z, test.want, got)
			}
		})
	}
}
func TestFloat64(t *testing.T) {
	var tests = []struct {
		name string
		a, z float64
		b    bool
		want float64
	}{
		{
			name: "false",
			a:    1,
			z:    2,
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    1,
			z:    2,
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float64(test.b, test.a, test.z); test.want != got {
				t.Fatalf("Float64(%t, %f,%f): want %f, got %f", test.b, test.a, test.z, test.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name string
		a, z string
		b    bool
		want string
	}{
		{
			name: "false",
			a:    "1",
			z:    "2",
			b:    false,
			want: "2",
		},
		{
			name: "true",
			a:    "1",
			z:    "2",
			b:    true,
			want: "1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := String(test.b, test.a, test.z); test.want != got {
				t.Fatalf("String(%t, %q,%q): want %q, got %q", test.b, test.a, test.z, test.want, got)
			}
		})
	}
}
