// sink package is responsible for holding logic associated with the sink - final result
package sink

// Sink struct represents the final state of the whole plumbing system
type Sink struct {
}

// New returns a new instance of a Sink
func New() *Sink {
	return &Sink{}
}
