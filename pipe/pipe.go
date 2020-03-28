// pipe package is responsible for holding all the logic of pipes.
package pipe

import "fmt"

// Pipe struct represents a pipeline through which data flows
type Pipe struct {
	input    map[string][]float64           // data that the pipe will apply the op to
	singleOp func(float64) (float64, error) // the singleOp that will be applied to independent input data points
	output   map[string][]float64           // the output after applying the singleOp to the input
}

// New returns a new instance of Pipe
func NewPipe(so func(float64) (float64, error)) *Pipe {
	return &Pipe{
		input:    nil,
		singleOp: so,
		output:   nil,
	}
}

// SetInput sets the inputs to the pipe, should only be accessed by a source
func (p *Pipe) SetInput(in map[string][]float64) {
	p.input = in
}

// GetOutput allows a consumer to get the output of this pipe
func (p *Pipe) GetOutput() map[string][]float64 {
	return p.output
}

// Flow flows the specified input through the specified pipe singleOp and stores the output
func (p *Pipe) Flow() error {
	if p.input == nil {
		return fmt.Errorf("cannot flow nil input through specified singleOp")
	}
	if p.singleOp == nil {
		return fmt.Errorf("cannot flow input through nil singleOp")
	}
	p.output = map[string][]float64{}
	for c, r := range p.input {
		// TODO: use a sync/wait group to execute this in parallel
		p.output[c] = make([]float64, len(r))
		for i, v := range r {
			newV, err := p.singleOp(v)
			if err != nil {
				return fmt.Errorf("failed to apply singleOp to val %v on row %v with op msg: %v", v, i, err)
			}
			p.output[c][i] = newV
		}
	}
	return nil
}
