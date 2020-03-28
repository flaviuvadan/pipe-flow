// sink package is responsible for holding logic associated with the sink - final result
package sink

import (
	"fmt"

	"github.com/flaviuvadan/pipe-flow/pipe"
)

// Sink struct represents the final state of the whole plumbing system
type Sink struct {
	filename string               // the name of the file the sink should dump data into
	pipes    []*pipe.Pipe         // the collection of pipes whose values are incoming to the sink
	data     map[string][]float64 // the data the sink collects from the pipes to output to a CSV
}

// New returns a new instance of a Sink
func NewSink(p []*pipe.Pipe) (*Sink, error) {
	if p == nil || len(p) == 0 {
		return nil, fmt.Errorf("cannot create a sink without pipes")
	}
	s := &Sink{
		pipes: p,
	}
	s.Collect()
	return s, nil
}

// Collect gets all the data from the pipes that are connected to this sink
func (s *Sink) Collect() {
	// there are many pipelines from which to get data from
	// have to merge all maps into a single one
	pipesData := []map[string][]float64{{}}
	for _, p := range s.pipes {
		pipesData = append(pipesData, p.GetOutput())
	}
	s.data = Merge(pipesData...)
}
