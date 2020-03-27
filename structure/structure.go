// structure package is responsible for holding the whole pipeline system (plumbing) and coordinating actions of
// starting and stopping
package structure

import (
	"github.com/flaviuvadan/pipe-flow/source"
	"github.com/flaviuvadan/pipe-flow/sink"
	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/flaviuvadan/pipe-flow/junction"
)

// Structure represents the state of the data processing system
type Structure struct {
	source    *source.Source
	sink      *sink.Sink
	pipes     []*pipe.Pipe
	junctions []*junction.Junction
}

// New returns a new instance of a Structure
func NewStructure(src *source.Source, snk *sink.Sink, pps []*pipe.Pipe, jnc []*junction.Junction) *Structure {
	return &Structure{
		source:    src,
		sink:      snk,
		pipes:     pps,
		junctions: jnc,
	}
}
