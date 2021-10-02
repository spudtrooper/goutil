// Package img contains functions for iamges
package img

import (
	"image"
	"log"

	"github.com/go-errors/errors"
	"github.com/nfnt/resize"
	"github.com/rivo/duplo/haar"
)

// A SimilarOptions is an option to `Similar`
type SimilarOptions func(*similarOptionsImpl)

// SimilarVerboseDiffs enables logging of internals when `verboseDiffs` is true
// Defaults to false
func SimilarVerboseDiffs(verboseDiffs bool) SimilarOptions {
	return func(opts *similarOptionsImpl) {
		opts.verboseDiffs = verboseDiffs
	}
}

// SimilarDiffThreshold sets the maximum difference in coefficients used to define a difference
// Defaults to 0.01
func SimilarDiffThreshold(threshold float64) SimilarOptions {
	return func(opts *similarOptionsImpl) {
		opts.threshold = threshold
	}
}

// SimilarFactor sets the ratio between the original and diff matrix of coefficients
// Defaults to 5
func SimilarFactor(factor uint) SimilarOptions {
	return func(opts *similarOptionsImpl) {
		opts.factor = factor
	}
}

// SimilarSample sets the number of pixels to sample in each width and height
// Defaults to 5
func SimilarSample(sample uint) SimilarOptions {
	return func(opts *similarOptionsImpl) {
		opts.sample = sample
	}
}

// SimilarMaxDiffs sets the maximum number of differences allowed and still be considered similar
// Defaults to 0
func SimilarMaxDiffs(maxDiffs uint) SimilarOptions {
	return func(opts *similarOptionsImpl) {
		opts.maxDiffs = maxDiffs
	}
}

// Similar compares images and returns whether they are "similar"
func Similar(a, b image.Image, inputOpts ...SimilarOptions) (bool, error) {
	opts := &similarOptionsImpl{
		verboseDiffs: false,
		threshold:    .01,
		factor:       5,
		sample:       5,
		maxDiffs:     0,
	}
	for _, o := range inputOpts {
		o(opts)
	}
	ma := convolveMatrix(createMatrix(a), opts.factor, opts.sample)
	mb := convolveMatrix(createMatrix(b), opts.factor, opts.sample)
	if ma.Height != mb.Height {
		return false, errors.Errorf("heights must be equal: %d != %d", ma.Height, mb.Height)
	}
	if ma.Width != mb.Width {
		return false, errors.Errorf("widths must be equal: %d != %d", ma.Width, mb.Width)
	}
	diff := createDiffMatrix(ma, mb)
	var diffs uint
	for c := range diff.Coefs {
		for i := range diff.Coefs[c] {
			if diff.Coefs[c][i] > opts.threshold {
				diffs++
			}
		}
	}
	similar := diffs <= opts.maxDiffs
	if similar && opts.verboseDiffs {
		for i := range ma.Coefs {
			log.Printf("%d: %v\t%v\t%v", i, diff.Coefs[i], ma.Coefs[i], mb.Coefs[i])
		}
		log.Printf("diff: %v", diff)
		log.Printf("diffs: %v", diffs)
	}
	return similar, nil
}

func createMatrix(img image.Image) haar.Matrix {
	scaled := resize.Resize(128, 128, img, resize.Bicubic)
	matrix := haar.Transform(scaled)
	return matrix
}

func convolveMatrix(input haar.Matrix, factor, sample uint) haar.Matrix {
	width := input.Width / factor
	height := input.Height / factor
	matrix := haar.Matrix{
		Coefs:  make([]haar.Coef, width*height),
		Width:  uint(width),
		Height: uint(height),
	}
	for row := 0; row < int(matrix.Height); row++ {
		for col := 0; col < int(matrix.Width); col++ {
			var (
				startRow = uint(uint(row) * factor)
				endRow   = startRow + sample
				startCol = uint(uint(col) * factor)
				endCol   = startCol + sample
				coef     haar.Coef
			)
			for r := startRow; r < input.Height && r <= endRow; r++ {
				for c := startCol; c < input.Width && r <= endCol; c++ {
					coef.Add(input.Coefs[r*uint(input.Width)+c])
				}
			}
			coef.Divide(float64(sample) * float64(sample))
			matrix.Coefs[row*int(matrix.Width)+col] = coef
		}
	}
	return matrix
}

func createDiffMatrix(a, b haar.Matrix) haar.Matrix {
	width := uint(a.Width)
	height := uint(b.Height)
	diff := haar.Matrix{
		Coefs:  make([]haar.Coef, width*height),
		Width:  width,
		Height: height,
	}
	for r := uint(0); r < height; r++ {
		for c := uint(0); c < width; c++ {
			i := r*width + c
			ca := a.Coefs[i]
			cb := b.Coefs[i]
			diff.Coefs[i] = haar.Coef{
				ca[0] - cb[0],
				ca[1] - cb[1],
				ca[2] - cb[2],
			}
		}
	}
	return diff
}

type similarOptionsImpl struct {
	verboseDiffs bool
	threshold    float64
	factor       uint
	sample       uint
	maxDiffs     uint
}
