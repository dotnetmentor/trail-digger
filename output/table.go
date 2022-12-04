package output

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/dotnetmentor/trail-digger/trail"
)

type Table struct {
	tw            *tabwriter.Writer
	options       Options
	headerWritten bool
}

func NewTable(w io.Writer, o Options) *Table {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', tabwriter.Debug)

	return &Table{
		tw:            tw,
		options:       o,
		headerWritten: false,
	}
}

func (o *Table) Type() string {
	return TypeTable
}

func (o *Table) Write(w io.Writer, r *trail.Record) error {
	if !o.headerWritten {
		fmt.Fprintln(o.tw, "")
		fmt.Fprintln(o.tw, strings.Join(o.options.Fields, "\t  "))
		o.headerWritten = true
	}

	if r.ErrorCode == "" && o.options.ErrorsOnly {
		return nil
	}

	fr := NewRecordValueAccessor(r)
	values := make([]string, 0)
	for _, f := range o.options.Fields {
		value := fr.Value(f)
		values = append(values, value)
	}

	fmt.Fprintln(o.tw, strings.Join(values, "\t  "))
	return nil
}

func (o *Table) Flush() error {
	return o.tw.Flush()
}
