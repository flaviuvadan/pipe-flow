// junction package is responsible for representing junctions in the plumbing infrastructure.
package junction

import "github.com/flaviuvadan/pipe-flow/pipe"

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
