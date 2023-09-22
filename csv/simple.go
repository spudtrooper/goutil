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
	var buf strings.Builder
	w := realCSV.NewWriter(&buf)
	return &simple{buf, w}
}

func (w *simple) Write(record []string) error    { return w.Write(record) }
func (w *simple) WriteAll(record []string) error { return w.WriteAll(record) }
func (w *simple) Flush()                         { w.Flush() }
func (w *simple) Error() error                   { return w.Error() }

func (w *simple) Done() string {
	w.w.Flush()
	return w.buf.String()
}
