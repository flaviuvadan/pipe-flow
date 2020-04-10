// junction package is responsible for representing junctions in the plumbing infrastructure.
package junction

import (
	"fmt"
	"github.com/flaviuvadan/pipe-flow/pipe"
)

// Junction represents the structure of the junction
type Junction struct {
	in  map[string][]float64 // results from the previous pipe
	out *pipe.Pipe           // future pipe that runs on in
}

// New returns a new instance of a junction struct
func NewJunction(out *pipe.Pipe) *Junction {
	return &Junction{
		out: out,
	}
}

// Continue makes the pipelines continue operating on the originally specified input
func (j *Junction) Continue() (*Junction, error) {
	if err := j.out.Flow(); err != nil {
		return nil, fmt.Errorf("failed to continue flow, err: %v", err)
	}
	return j.out.GetJunction(), nil
}

// SetNextInput takes the given input and sets it as input to the future pipe
func (j *Junction) SetNextInput(i map[string][]float64) {
	j.out.SetInput(i)
}
