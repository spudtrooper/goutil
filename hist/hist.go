package hist

import (
	"fmt"
	"sort"
	"sync"
)

type Histogram interface {
	Pairs() Pairs
	Add(key interface{}, val int)
}

func MakeHistogram(mOpts ...MakeHistogramOption) Histogram {
	opts := MakeMakeHistogramOptions(mOpts...)
	if opts.SortDesc() {
		return &desc{base{hist: map[string]int{}}}
	}
	if opts.SortAsc() {
		return &asc{base{hist: map[string]int{}}}
	}
	return &impl{base{hist: map[string]int{}}}
}

func Sync(h Histogram) Histogram {
	return &synced{h: h}
}

type base struct{ hist map[string]int }
type impl struct{ base }
type desc struct{ base }
type asc struct{ base }

type synced struct {
	h  Histogram
	mu sync.Mutex
}

func (h *synced) Add(key interface{}, val int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.h.Add(key, val)
}

func (h *synced) Pairs() Pairs {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.h.Pairs()
}

func (h *base) Add(key interface{}, val int) {
	h.hist[fmt.Sprintf("%s", key)] += val
}

func (h *impl) Pairs() Pairs {
	var res Pairs
	for k, v := range h.hist {
		res = append(res, Pair{k, v})
	}
	return res
}

func (h *desc) Pairs() Pairs {
	ps := makePairs(h.hist)
	sort.Sort(sort.Reverse(ps))
	return ps
}

func (h *asc) Pairs() Pairs {
	ps := makePairs(h.hist)
	sort.Sort(ps)
	return ps
}

// https://stackoverflow.com/questions/18695346/how-can-i-sort-a-mapstringint-by-its-values
func makePairs(wordFrequencies map[string]int) Pairs {
	ps := make(Pairs, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		ps[i] = Pair{k, v}
		i++
	}
	return ps
}

type Pair struct {
	Key   string
	Value int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
