package csv

import (
	realCSV "encoding/csv"
	"strings"
)

type simple struct {
	buf strings.Builder
	w   *realCSV.Writer
}

func NewSimpleWriter() *simple {
	res := &simple{}
	w := realCSV.NewWriter(&res.buf)
	res.w = w
	return res
}

func (w *simple) Write(record []string) error       { return w.w.Write(record) }
func (w *simple) WriteAll(records [][]string) error { return w.w.WriteAll(records) }
func (w *simple) Flush()                            { w.w.Flush() }
func (w *simple) Error() error                      { return w.w.Error() }

func (w *simple) Done() string {
	w.w.Flush()
	return w.buf.String()
}
