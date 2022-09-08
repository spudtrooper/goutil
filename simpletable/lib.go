package simpletable

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type impl struct {
	table *tablewriter.Table
}

type SimpleTable interface {
	Append(row []string)
	Render()
}

//go:generate genopts --function New noBorder "header:[]string"
func New(writer io.Writer, optss ...NewOption) SimpleTable {
	opts := MakeNewOptions(optss...)

	table := tablewriter.NewWriter(writer)
	table.SetBorder(!opts.NoBorder())
	if len(opts.Header()) > 0 {
		table.SetHeader(opts.Header())
	}

	return &impl{
		table: table,
	}
}

func (s *impl) Append(row []string) { s.table.Append(row) }

func (s *impl) Render() { s.table.Render() }
