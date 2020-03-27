// source package is responsible for holding the logic of the start state of the plumbing system
package source

import "time"

// Source represents the beginning state of a pipeline
type Source struct {
	Name  string
	Start time.Time
}

// New returns a new instance of a Source
func NewSource() *Source {
	return &Source{}
}
