package img

import (
	"image"
	"image/png"
	"os"
	"path"
	"testing"
)

func TestSimilarNaive(t *testing.T) {
	var tests = []struct {
		name   string
		af, bf string
		opts   []SimilarOptions
		want   bool
	}{
		{
			name: "same",
			af:   "a.png",
			bf:   "a.png",
			want: true,
		},
		{
			name: "similar",
			af:   "a.png",
			bf:   "b.png",
			want: true,
		},
		{
			name: "different",
			af:   "a.png",
			bf:   "c.png",
			want: false,
		},
		{
			name: "same/verbose",
			af:   "a.png",
			bf:   "a.png",
			opts: []SimilarOptions{SimilarVerboseDiffs(true)},
			want: true,
		},
		{
			name: "similar/low threshold",
			af:   "a.png",
			bf:   "b.png",
			opts: []SimilarOptions{SimilarDiffThreshold(0.0001)},
			want: false,
		},
		{
			name: "similar/factor of 1",
			af:   "a.png",
			bf:   "b.png",
			opts: []SimilarOptions{SimilarFactor(1)},
			want: true,
		},
		{
			name: "similar/sample of 1",
			af:   "a.png",
			bf:   "b.png",
			opts: []SimilarOptions{SimilarSample(1)},
			want: true,
		},
		{
			name: "different/factor of 1",
			af:   "a.png",
			bf:   "c.png",
			opts: []SimilarOptions{SimilarFactor(1)},
			want: false,
		},
		{
			name: "different/sample of 1",
			af:   "a.png",
			bf:   "c.png",
			opts: []SimilarOptions{SimilarSample(1)},
			want: false,
		},
		{
			name: "different/high max diffs",
			af:   "a.png",
			bf:   "c.png",
			opts: []SimilarOptions{SimilarMaxDiffs(10000000)},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			af := test.af
			bf := test.bf
			a, err := readImage(af)
			if err != nil {
				t.Fatalf("readImage(%q): %v", af, err)
			}
			b, err := readImage(bf)
			if err != nil {
				t.Fatalf("readImage(%q): %v", bf, err)
			}
			got, err := SimilarNaive(a, b, test.opts...)
			if err != nil {
				t.Fatalf("Similar(%q,%q): %v", af, bf, err)
			}
			if want := test.want; got != want {
				t.Errorf("Similar(%q,%q) want != got: %t, %t", af, bf, want, got)
			}
		})
	}
}

func TestSimilar(t *testing.T) {
	var tests = []struct {
		name   string
		af, bf string
		opts   []SimilarOptions
		want   bool
	}{
		{
			name: "same",
			af:   "a.png",
			bf:   "a.png",
			want: true,
		},
		{
			name: "similar",
			af:   "a.png",
			bf:   "b.png",
			want: true,
		},
		{
			name: "different",
			af:   "a.png",
			bf:   "c.png",
			want: false,
		},
		{
			name: "same/verbose",
			af:   "a.png",
			bf:   "a.png",
			opts: []SimilarOptions{SimilarVerboseDiffs(true)},
			want: true,
		},
		{
			name: "similar/low threshold",
			af:   "a.png",
			bf:   "b.png",
			opts: []SimilarOptions{SimilarDiffThreshold(0.0001)},
			want: false,
		},
		{
			name: "similar/factor of 1",
			af:   "a.png",
			bf:   "b.png",
			opts: []SimilarOptions{SimilarFactor(1)},
			want: true,
		},
		{
			name: "similar/sample of 1",
			af:   "a.png",
			bf:   "b.png",
			opts: []SimilarOptions{SimilarSample(1)},
			want: true,
		},
		{
			name: "different/factor of 1",
			af:   "a.png",
			bf:   "c.png",
			opts: []SimilarOptions{SimilarFactor(1)},
			want: false,
		},
		{
			name: "different/sample of 1",
			af:   "a.png",
			bf:   "c.png",
			opts: []SimilarOptions{SimilarSample(1)},
			want: false,
		},
		{
			name: "different/high max diffs",
			af:   "a.png",
			bf:   "c.png",
			opts: []SimilarOptions{SimilarMaxDiffs(10000000)},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			af := Png(path.Join("testdata", test.af))
			bf := Png(path.Join("testdata", test.bf))
			got, err := Similar(af, bf, test.opts...)
			if err != nil {
				t.Fatalf("Similar(%q,%q): %v", af, bf, err)
			}
			if want := test.want; got != want {
				t.Errorf("Similar(%q,%q) want != got: %t, %t", af, bf, want, got)
			}
		})
	}
}

func readImage(basename string) (image.Image, error) {
	f := path.Join("testdata", basename)
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
