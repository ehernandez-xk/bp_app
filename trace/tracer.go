package trace

import (
	"fmt"
	"io"
)

// check if my struts implements the Tracer interface
var _ Tracer = (*tracer)(nil)
var _ Tracer = (*nilTracer)(nil)

// Tracer is the interface that describes an object capable of
// tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

//New returns a new tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off returns a nillTracer
func Off() Tracer {
	return &nilTracer{}
}
