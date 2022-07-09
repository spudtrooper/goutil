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

func TestIntDefault(t *testing.T) {
	var tests = []struct {
		name string
		a    int
		want int
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IntDefault(test.a); test.want != got {
				t.Errorf("IntDefault(%d): want %d, got %d", test.a, test.want, got)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	var tests = []struct {
		name string
		a, b int8
		want int8
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
			if got := Int8(test.a, test.b); test.want != got {
				t.Errorf("Int8(%d,%d): want %d, got %d", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestInt8Default(t *testing.T) {
	var tests = []struct {
		name string
		a    int8
		want int8
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int8Default(test.a); test.want != got {
				t.Errorf("Int8Default(%d): want %d, got %d", test.a, test.want, got)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	var tests = []struct {
		name string
		a, b int16
		want int16
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
			if got := Int16(test.a, test.b); test.want != got {
				t.Errorf("Int16(%d,%d): want %d, got %d", test.a, test.b, test.want, got)
			}
		})
	}
}

func TestInt16Default(t *testing.T) {
	var tests = []struct {
		name string
		a    int16
		want int16
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int16Default(test.a); test.want != got {
				t.Errorf("Int16Default(%d): want %d, got %d", test.a, test.want, got)
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

func TestInt32Default(t *testing.T) {
	var tests = []struct {
		name string
		a    int32
		want int32
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int32Default(test.a); test.want != got {
				t.Errorf("Int32Default(%d): want %d, got %d", test.a, test.want, got)
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

func TestInt64Default(t *testing.T) {
	var tests = []struct {
		name string
		a    int64
		want int64
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Int64Default(test.a); test.want != got {
				t.Errorf("Int64Default(%d): want %d, got %d", test.a, test.want, got)
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

func TestFloat32Default(t *testing.T) {
	var tests = []struct {
		name string
		a    float32
		want float32
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float32Default(test.a); test.want != got {
				t.Errorf("Float32Default(%f): want %f, got %f", test.a, test.want, got)
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

func TestFloat64Default(t *testing.T) {
	var tests = []struct {
		name string
		a    float64
		want float64
	}{
		{
			name: "defaults",
			a:    0,
			want: 0,
		},
		{
			name: "non-default default",
			a:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Float64Default(test.a); test.want != got {
				t.Errorf("Float64Default(%f): want %f, got %f", test.a, test.want, got)
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

func TestStringDefault(t *testing.T) {
	var tests = []struct {
		name string
		a    string
		want string
	}{
		{
			name: "defaults",
			a:    "",
			want: "",
		},
		{
			name: "non-default default",
			a:    "1",
			want: "1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := StringDefault(test.a); test.want != got {
				t.Errorf("StringDefault(%q): want %q, got %q", test.a, test.want, got)
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

func TestTimeDefault(t *testing.T) {
	var unset time.Time
	now := time.Now()
	var tests = []struct {
		name string
		a    time.Time
		want time.Time
	}{
		{
			name: "defaults",
			a:    unset,
			want: unset,
		},
		{
			name: "non-default default",
			a:    now,
			want: now,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := TimeDefault(test.a); test.want != got {
				t.Errorf("TimeDefault(%v): want %v, got %v", test.a, test.want, got)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	var unset time.Duration
	var tests = []struct {
		name string
		a, b time.Duration
		want time.Duration
	}{
		{
			name: "defaults",
			a:    unset,
			b:    unset,
			want: unset,
		},
		{
			name: "default non-default",
			a:    time.Second * 1,
			b:    unset,
			want: time.Second * 1,
		},
		{
			name: "non-default non-default",
			a:    time.Second * 1,
			b:    time.Second * 2,
			want: time.Second * 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if want, got := test.want, Duration(test.a, test.b); want != got {
				t.Errorf("Time(%v,%v): want %v, got %v", test.a, test.b, want, got)
			}
		})
	}
}

func TestDurationDefault(t *testing.T) {
	var unset time.Duration
	var tests = []struct {
		name string
		a    time.Duration
		want time.Duration
	}{
		{
			name: "defaults",
			a:    unset,
			want: unset,
		},
		{
			name: "non-default default",
			a:    time.Second * 1,
			want: time.Second * 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := DurationDefault(test.a); test.want != got {
				t.Errorf("DurationDefault(%v): want %v, got %v", test.a, test.want, got)
			}
		})
	}
}

func TestAny(t *testing.T) {
	if want, got := 0, Any(0, 0); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", 0, 0, want, got)
	}
	if want, got := 1, Any(1, 0); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", 1, 0, want, got)
	}
	if want, got := 1, Any(1, 2); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", 1, 2, want, got)
	}

	if want, got := "", Any("", ""); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", "", "", want, got)
	}
	if want, got := "1", Any("1", ""); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", "1", "", want, got)
	}
	if want, got := "1", Any("1", "2"); want != got {
		t.Errorf("Any(%v,%v): want %v, got %v", "1", "2", want, got)
	}
}
