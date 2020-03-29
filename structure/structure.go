// structure package is responsible for holding the whole pipeline system (plumbing) and coordinating actions of
// starting and stopping
package structure

import (
	"fmt"
	"time"

	"github.com/flaviuvadan/pipe-flow/junction"
	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/flaviuvadan/pipe-flow/sink"
	"github.com/flaviuvadan/pipe-flow/source"
)

// Structure represents the state of the data processing system
type Structure struct {
	description string               // a description of the structure and what it does e.g the data it processes
	inform      bool                 // whether to inform users of the process of the pipelines as they are performing, state informs occur in junctions
	source      *source.Source       // data source
	sink        *sink.Sink           // data sink
	pipes       []*pipe.Pipe         // pipes that are coordinated by the structure
	junctions   []*junction.Junction // junctions that are intermediary steps between pipes
}

// New returns a new instance of a Structure
func NewStructure(dsc string, src *source.Source, snk *sink.Sink) *Structure {
	return &Structure{
		description: dsc,
		source:      src,
		sink:        snk,
		pipes:       []*pipe.Pipe{},
		junctions:   []*junction.Junction{},
	}
}

// Register adds pipes or junctions to the structure
func (s *Structure) Register(i interface{}) error {
	switch v := i.(type) {
	case *pipe.Pipe:
		s.pipes = append(s.pipes, v)
	case *junction.Junction:
		s.junctions = append(s.junctions, v)
	default:
		return fmt.Errorf("provided interface cannot be cast to any known type")
	}
	return nil
}

// Flow launches the flow of all the pipelines that are registered with this structure
func (s *Structure) Flow() (string, error) {
	start := time.Now()
	// TODO: this does not take into account the existence of junctions
	// TODO: implement junction support
	for _, p := range s.pipes {
		if err := p.Flow(); err != nil {
			return "", fmt.Errorf("pipe failed to flow data, err: %v", err)
		}
	}
	s.sink.Collect()
	if err := s.sink.Dump(); err != nil {
		return "", fmt.Errorf("sink failed to dump data, err: %v", err)
	}
	duration := time.Now().Sub(start)
	return duration.String(), nil
}
