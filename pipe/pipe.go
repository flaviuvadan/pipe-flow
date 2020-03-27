// pipe package is responsible for holding all the logic of pipes.
package pipe

// Pipe struct represents a pipeline through which data flows
type Pipe struct {
	input     map[string][]float64             // data that the pipe will apply the op to
	operation func(map[string][]float64) error // the operation that will be applied to the input data
	output    map[string][]float64             // the output after applying the operation to the input
}

// New returns a new instance of Pipe
func NewPipe(op func(map[string][]float64)) *Pipe {
	return &Pipe{}
}
