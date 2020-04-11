// pipe package is responsible for holding all the logic of pipes.
package pipe

import (
	"fmt"
	"time"
)

// Pipe struct represents a pipeline through which data flows
type Pipe struct {
	Description string                           // a Description/name of the pipeline, used for monitoring
	input       map[string][]float64             // data that the pipe will apply the op to
	singleOps   []func(float64) (float64, error) // the singleOp that will be applied to independent input data points
	output      map[string][]float64             // the output after applying the singleOp to the input
	start       time.Time                        // start time of the pipeline
	end         time.Time                        // end time of the pipeline
}

// New returns a new instance of Pipe
func NewPipe(ds string, so []func(float64) (float64, error)) *Pipe {
	return &Pipe{
		Description: ds,
		input:       nil,
		singleOps:   so,
		output:      nil,
	}
}

// SetInput sets the inputs to the pipe, should only be accessed by a source
func (p *Pipe) SetInput(in map[string][]float64) {
	p.input = in
}

// GetInput returns the input that was specified to the pipe
func (p *Pipe) GetInput() map[string][]float64 {
	return p.input
}

// SetOutput sets the output of the pipe to a custom one that is not computed by Flow
// this is mostly implemented for testing purposes as we would otherwise have to set up
// CSV files for testing the Sink package
func (p *Pipe) SetOutput(ot map[string][]float64) {
	p.output = ot
}

// GetOutput allows a consumer to get the output of this pipe
func (p *Pipe) GetOutput() map[string][]float64 {
	return p.output
}

// GetFlowDuration tells how long the Flow operation needed to process the pipeline input
func (p *Pipe) GetFlowDuration() time.Duration {
	return p.end.Sub(p.start)
}

// Flow flows the specified input through the specified pipe singleOp and stores the output
func (p *Pipe) Flow() error {
	p.start = time.Now()
	if p.input == nil {
		return fmt.Errorf("cannot flow nil input through specified singleOps")
	}
	if p.singleOps == nil {
		return fmt.Errorf("cannot flow input through nil singleOps")
	}
	p.output = map[string][]float64{}
	for col, rows := range p.input {
		// TODO: use a sync/wait group to execute this in parallel
		p.output[col] = make([]float64, len(rows))
		for i, val := range rows {
			newVal, err := p.singleOps[0](val)
			if err != nil {
				return fmt.Errorf("failed to apply op to val %v on row %v with op msg: %v", val, i, err)
			}
			for _, op := range p.singleOps[1:] {
				newVal, err = op(val)
				if err != nil {
					return fmt.Errorf("failed to apply op to val %v on row %v with op msg: %v", val, i, err)
				}
			}
			p.output[col][i] = newVal
		}
	}
	p.end = time.Now()
	return nil
}
