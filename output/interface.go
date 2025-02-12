package output

import (
	"io"

	"github.com/dotnetmentor/trail-digger/trail"
)

type Output interface {
	Type() string
	Write(w io.Writer, r *trail.Record) error
	Flush() error
}
