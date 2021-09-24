package lazycond

import (
	"testing"
)

func TestInt(t *testing.T) {
	var tests = []struct {
		name string
		a, z func() int
		b    bool
		want int
	}{
		{
			name: "false",
			a:    func() int { return 1 },
			z:    func() int { return 2 },
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    func() int { return 1 },
			z:    func() int { return 2 },
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int(test.b, test.a, test.z); test.want != got {
				t.Errorf("Int(%t, %d,%d): want %d, got %d", test.b, test.a(), test.z(), test.want, got)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	var tests = []struct {
		name string
		a, z func() int32
		b    bool
		want int32
	}{
		{
			name: "false",
			a:    func() int32 { return 1 },
			z:    func() int32 { return 2 },
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    func() int32 { return 1 },
			z:    func() int32 { return 2 },
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int32(test.b, test.a, test.z); test.want != got {
				t.Errorf("Int32(%t, %d,%d): want %d, got %d", test.b, test.a(), test.z(), test.want, got)
			}
		})
	}
}
func TestInt64(t *testing.T) {
	var tests = []struct {
		name string
		a, z func() int64
		b    bool
		want int64
	}{
		{
			name: "false",
			a:    func() int64 { return 1 },
			z:    func() int64 { return 2 },
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    func() int64 { return 1 },
			z:    func() int64 { return 2 },
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int64(test.b, test.a, test.z); test.want != got {
				t.Errorf("Int64(%t, %d,%d): want %d, got %d", test.b, test.a(), test.z(), test.want, got)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	var tests = []struct {
		name string
		a, z func() float32
		b    bool
		want float32
	}{
		{
			name: "false",
			a:    func() float32 { return 1 },
			z:    func() float32 { return 2 },
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    func() float32 { return 1 },
			z:    func() float32 { return 2 },
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float32(test.b, test.a, test.z); test.want != got {
				t.Errorf("Float32(%t, %f,%f): want %f, got %f", test.b, test.a(), test.z(), test.want, got)
			}
		})
	}
}
func TestFloat64(t *testing.T) {
	var tests = []struct {
		name string
		a, z func() float64
		b    bool
		want float64
	}{
		{
			name: "false",
			a:    func() float64 { return 1 },
			z:    func() float64 { return 2 },
			b:    false,
			want: 2,
		},
		{
			name: "true",
			a:    func() float64 { return 1 },
			z:    func() float64 { return 2 },
			b:    true,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float64(test.b, test.a, test.z); test.want != got {
				t.Errorf("Float64(%t, %f,%f): want %f, got %f", test.b, test.a(), test.z(), test.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name string
		a, z func() string
		b    bool
		want string
	}{
		{
			name: "false",
			a:    func() string { return "1" },
			z:    func() string { return "2" },
			b:    false,
			want: "2",
		},
		{
			name: "true",
			a:    func() string { return "1" },
			z:    func() string { return "2" },
			b:    true,
			want: "1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := String(test.b, test.a, test.z); test.want != got {
				t.Errorf("String(%t, %q,%q): want %q, got %q", test.b, test.a(), test.z(), test.want, got)
			}
		})
	}
}
