package hist

type HistogramChannelBuilder interface {
	Build(mOpts ...MakeHistogramOption) Histogram
	Add(key interface{})
	Close()
}

type builder struct {
	keys chan interface{}
}

func MakeHistogramChannelBuilder() HistogramChannelBuilder {
	return &builder{keys: make(chan interface{})}
}

func (b *builder) Build(mOpts ...MakeHistogramOption) Histogram {
	h := MakeHistogram(mOpts...)
	for k := range b.keys {
		h.Add(k, 1)
	}
	return h
}

func (b *builder) Add(key interface{}) {
	b.keys <- key
}

func (b *builder) Close() {
	close(b.keys)
}
