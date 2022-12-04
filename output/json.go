package output

import (
	"encoding/json"
	"io"

	"github.com/dotnetmentor/trail-digger/trail"
)

type Json struct {
	options Options
}

func NewJson(o Options) *Json {
	return &Json{
		options: o,
	}
}

func (o *Json) Type() string {
	return TypeJson
}

func (o *Json) Write(w io.Writer, r *trail.Record) error {
	if r.ErrorCode == "" && o.options.ErrorsOnly {
		return nil
	}

	return json.NewEncoder(w).Encode(r)
}

func (o Json) Flush() error {
	return nil
}
