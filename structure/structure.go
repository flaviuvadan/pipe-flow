// structure package is responsible for holding the whole pipeline system (plumbing) and coordinating actions of
// starting and stopping
package structure

import (
	"fmt"
	"time"

	"github.com/flaviuvadan/pipe-flow/pipe"
	"github.com/flaviuvadan/pipe-flow/sink"
	"github.com/flaviuvadan/pipe-flow/source"
)

// Structure represents the state of the data processing system
type Structure struct {
	Description string         // a Description of the structure and what it does e.g the data it processes
	Inform      bool           // whether to Inform users of the process of the pipelines as they are performing, state informs occur in junctions
	Source      *source.Source // data Source
	Sink        *sink.Sink     // data Sink
	Pipes       []*pipe.Pipe   // Pipes that are coordinated by the structure
}

// New returns a new instance of a Structure
func NewStructure(dsc string) *Structure {
	return &Structure{
		Description: dsc,
		Source:      nil,
		Sink:        nil,
		Pipes:       []*pipe.Pipe{},
	}
}

// Register adds Pipes or junctions to the structure
func (s *Structure) Register(i interface{}) error {
	switch v := i.(type) {
	case *pipe.Pipe:
		s.Pipes = append(s.Pipes, v)
	case *source.Source:
		s.Source = v
	case *sink.Sink:
		s.Sink = v
	default:
		return fmt.Errorf("provided interface cannot be cast to any known type")
	}
	return nil
}

// Flow launches the flow of all the pipelines that are registered with this structure
func (s *Structure) Flow() (string, error) {
	if s.Source == nil {
		return "", fmt.Errorf("cannot flow with nil Source")
	}
	if s.Sink == nil {
		return "", fmt.Errorf("cannot flow with nil Sink")
	}
	start := time.Now()
	// TODO: implement this
	duration := time.Now().Sub(start)
	return duration.String(), nil
}
