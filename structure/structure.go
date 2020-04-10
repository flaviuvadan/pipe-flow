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
func NewStructure(dsc string) *Structure {
	return &Structure{
		description: dsc,
		source:      nil,
		sink:        nil,
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
	case *source.Source:
		s.source = v
	case *sink.Sink:
		s.sink = v
	default:
		return fmt.Errorf("provided interface cannot be cast to any known type")
	}
	return nil
}

// Flow launches the flow of all the pipelines that are registered with this structure
func (s *Structure) Flow() (string, error) {
	if s.source == nil {
		return "", fmt.Errorf("cannot flow with nil source")
	}
	if s.sink == nil {
		return "", fmt.Errorf("cannot flow with nil sink")
	}
	start := time.Now()

	q := NewQueue(len(s.junctions) + 1) // + 1 for source
	if err := q.Push(s.source); err != nil {
		return "", fmt.Errorf("failed to init queue by pushing source, err: %v", err)
	}
	for q.Size() != 0 {
		o := q.Pop()
		switch v := o.(type) {
		case *source.Source: // a source should always be the first thing that gets processed
			for _, p := range v.Pipes {
				if err := p.Flow(); err != nil {
					return "", fmt.Errorf("failed to perform pipe flow, err: %v", err)
				}
			}
		case *junction.Junction:
			// TODO: implement junction support after a bit more thinking, this seems too convoluted
		}
	}
	s.sink.Collect() // collect results
	if err := s.sink.Dump(); err != nil {
		return "", fmt.Errorf("failed to dump results from sink, err: %v", err)
	}
	duration := time.Now().Sub(start)
	return duration.String(), nil
}
