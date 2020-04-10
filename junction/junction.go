// junction package is responsible for representing junctions in the plumbing infrastructure.
package junction

import (
	"fmt"

	"github.com/flaviuvadan/pipe-flow/pipe"
)

// Junction represents the structure of the junction
type Junction struct {
	in  *pipe.Pipe
	out *pipe.Pipe
}

// New returns a new instance of a junction struct
func NewJunction(in *pipe.Pipe, out *pipe.Pipe) *Junction {
	return &Junction{
		in:  in,
		out: out,
	}
}

// GetState returns a formatted string that describes how long the in pipeline operated for
func (j *Junction) GetState() string {
	return fmt.Sprintf("IN pipeline description: %v | duration: %v\n"+
		"OUT pipeline description: %v\n", j.in.Description, j.in.GetFlowDuration(), j.out.Description)
}

// Continue makes the pipelines continue operating on the originally specified input
func (j *Junction) Continue() error {
	j.out.SetInput(j.in.GetOutput())
	return j.out.Flow()
}
