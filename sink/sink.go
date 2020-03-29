// sink package is responsible for holding logic associated with the sink - final result
package sink

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/flaviuvadan/pipe-flow/pipe"
)

const (
	Precision    = 3  // number of float decimals
	FloatBitSize = 64 // bit size of floats
)

// Sink struct represents the final state of the whole plumbing system
// if the filename was not specified, i.e it is "", results.csv is assumed
type Sink struct {
	filename string               // the name of the file the sink should dump data into
	pipes    []*pipe.Pipe         // the collection of pipes whose values are incoming to the sink
	data     map[string][]float64 // the data the sink collects from the pipes to output to a CSV
}

// New returns a new instance of a Sink
func NewSink(fn string, p []*pipe.Pipe) (*Sink, error) {
	if p == nil || len(p) == 0 {
		return nil, fmt.Errorf("cannot create a sink without pipes")
	}
	if fn == "" {
		fn = "results.csv"
	}
	s := &Sink{
		filename: fn,
		pipes:    p,
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

// Dump tries to create the CSV file named filename with the results of the sink
func (s *Sink) Dump() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get the current working directory")
	}

	f, err := os.Create(path.Join(cwd, s.filename))
	if err != nil {
		return fmt.Errorf("failed to create the dump CSV file")
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for k, v := range s.data {
		r := make([]string, 0, len(v)+1) // + 1 for the header
		r = append(r, k)
		for _, j := range v {
			r = append(r, strconv.FormatFloat(j, 'f', Precision, FloatBitSize))
		}
		if err := w.Write(r); err != nil {
			return fmt.Errorf("failed to write record to CSV file, err: %v", err)
		}
	}
	return nil
}
