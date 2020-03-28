// pipe package is responsible for holding all the logic of pipes.
package pipe

// Pipe struct represents a pipeline through which data flows
type Pipe struct {
	input     map[string][]float64          // data that the pipe will apply the op to
	operation func(interface{}) interface{} // the operation that will be applied to the input data
	output    []float64                     // the output after applying the operation to the input
}

// New returns a new instance of Pipe
func NewPipe(op func(val interface{}) interface{}) *Pipe {
	return &Pipe{
		input:     nil,
		operation: op,
		output:    nil,
	}
}

// SetInput sets the inputs to the pipe, should only be accessed by a source
func (p *Pipe) SetInput(in map[string][]float64) {
	p.input = in
}

// GetOutput allows a consumer to get the output of this pipe
func (p *Pipe) GetOutput() []float64 {
	return p.output
}
