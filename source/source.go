// source package is responsible for holding the logic of the start state of the plumbing system
package source

// Source represents the beginning state of a pipeline
type Source struct {
}

// New returns a new instance of a Source 
func New() *Source {
	return &Source{}
}
