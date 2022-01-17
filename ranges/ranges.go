package ranges

import "github.com/spudtrooper/goutil/check"

type Range interface {
	From(from int) Range
	To(to int) Range
	Loop(LoopFn)
}

type rng struct {
	from, to int
}

func MakeRange() Range {
	return &rng{}
}

type LoopFn func(int)

func Loop(from, to int, fn LoopFn) {
	MakeRange().From(from).To(to).Loop(fn)
}

func LoopTo(to int, fn LoopFn) {
	Loop(0, to, fn)
}

func (r *rng) From(from int) Range {
	r.from = from
	return r
}

func (r *rng) To(to int) Range {
	r.to = to
	return r
}

func (r *rng) Loop(fn LoopFn) {
	check.Check(r.from < r.to)
	for i := r.from; i < r.to; i++ {
		fn(i)
	}
}
